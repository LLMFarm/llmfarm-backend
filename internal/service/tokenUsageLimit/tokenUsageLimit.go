package tokenUsageLimit

import (
	"context"
	"errors"
	"fmt"
	v1 "llmfarm/api/tokenUsageLimit/v1"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type TokenUsageLimitService struct{}

func New() *TokenUsageLimitService {
	return &TokenUsageLimitService{}
}

func (s *TokenUsageLimitService) GetUserById(userId uint) (*entity.User, error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and id =  ?", userId).One()
	r.Struct(&user)
	return user, err
}

// list
func (s *TokenUsageLimitService) List(ctx context.Context, versionType string, req *v1.TokenUsageLimitListReq, userId uint) ([]v1.TokenUsageLimit, error) {
	// 获取用户信息
	// language, _ := g.Cfg().Get(r.Context(), "language")
	// i18n.SetLanguage(language.String())
	userInfo, err := s.GetUserById(userId)
	if err != nil {
		return nil, errors.New("查询失败")
	}
	var tokenUsageLimits []v1.TokenUsageLimit
	// 获取环境变量 versionType == "enterprise"
	if userInfo.Is_membership == 1 && versionType != "enterprise" {
		// 查询meberToken表
		tokenUsageLimitList := g.Model("memberToken m").Fields(" m.tokenTotal as `DailyTokens`, m.tokenUse AS `UsageToken`").Where("m.userId = ? ", userId)
		res, err := tokenUsageLimitList.All()
		if err != nil {
			return nil, errors.New("查询失败")
		}
		res.Structs(&tokenUsageLimits)
		tokenUsageLimits[0].Title = g.I18n().T(ctx, "总用量")
		return tokenUsageLimits, err

	} else {
		tokenUsageLimitList := g.Model("tokenUsageLimit t").Fields(" t.dailyTokens as `DailyTokens`, SUM(mes.usageTotalTokens) AS `UsageToken`").LeftJoin("messageModelUsageTokens mes", "mes.messageId IN (SELECT id FROM message WHERE userId = t.userId  ) ").Where("t.userId = ? ", userId)
		//Group("m.`name`,m.id,t.dailyTokens")
		res, err := tokenUsageLimitList.All()
		if err != nil {
			return nil, errors.New("查询失败")
		}
		res.Structs(&tokenUsageLimits)
		tokenUsageLimits[0].Title = g.I18n().T(ctx, "总用量")
		return tokenUsageLimits, err
	}
}

// 根据用户Id 和 模型Id 获取 DailyTokens
func (s *TokenUsageLimitService) GetUserdailyTokens(userId uint) (uint, error) {
	var tokenUsageLimit = new(entity.TokenUsageLimit)
	r, err := g.Model("tokenUsageLimit").Where(" deleteTime is null and userId = ? ", userId).One()
	if err != nil {
		return 0, errors.New("未查询到")
	}
	r.Struct(&tokenUsageLimit)
	var total uint
	total = tokenUsageLimit.DailyTokens
	return total, err
}

// 创建用户的tokenUsageLimit
func (s *TokenUsageLimitService) Create(userId uint) error {
	var TokenUsageLimit = new(entity.TokenUsageLimit)
	TokenUsageLimit.UserId = userId
	DailyTokens, _ := g.Cfg().Get(context.TODO(), "tokenUsageLimit.dailyTokens")
	fmt.Println("DailyTokens", DailyTokens)
	TokenUsageLimit.DailyTokens = DailyTokens.Uint()
	TokenUsageLimit.DefaultCreateTime()
	TokenUsageLimit.DefaultUpdateTime()
	res, err := g.Model("tokenUsageLimit").Insert(&TokenUsageLimit)
	if err != nil {
		fmt.Println("err", err, res)
		return errors.New("创建失败")
	}
	productId, err := res.LastInsertId()
	TokenUsageLimit.Id = uint(productId)
	return err
}
