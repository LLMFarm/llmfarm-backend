package user

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	v1 "llmfarm/api/user/v1"
	"llmfarm/internal/model/entity"
	"llmfarm/internal/service/session"
	tokenUsageLimitService "llmfarm/internal/service/tokenUsageLimit"
	"llmfarm/library/response"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type UserService struct{}

func New() *UserService {
	return &UserService{}
}

func (s *UserService) List() ([]entity.User, error) {
	a := 2
	m := g.Model("user")
	if a == 2 {
		m.Where("id > ?", 1)
	}

	r, err := m.All()

	var users []entity.User
	r.Structs(&users)
	return users, err
}

func (s *UserService) Login(ctx context.Context, req *v1.UserLoginReq) (*v1.UserLoginResult, error) {
	// 登录后将用户信息存储到session
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and phone = ? or email = ?", req.Username, req.Username).One()
	r.Struct(&user)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("登录失败")
	}
	if user == nil {
		return nil, errors.New("账户不存在")
	}

	if user.User_status != 1 {
		return nil, errors.New("账户已被禁用")
	}
	encryptPassword, _ := gmd5.EncryptString(req.Password + user.Salt)
	fmt.Println(encryptPassword)
	if encryptPassword != user.Password {
		return nil, errors.New("密码错误")
	}
	session.Session.SetUser(ctx, user)
	userLogin := new(v1.UserLoginResult)
	userLogin.Name = user.Name
	userLogin.Email = user.Email
	userLogin.Phone = user.Phone
	userLogin.UserId = user.Id
	userLogin.Is_admin = user.Is_admin
	return userLogin, nil
}

func (s *UserService) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	exists, err := s.CheckPhone(req.Phone)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	} else {
		if exists {
			return nil, errors.New("用户已存在")
		}
	}

	salt := grand.S(6, false)
	encryptPassword, _ := gmd5.EncryptString(req.Password + salt)
	var user *entity.User = new(entity.User)
	user.Phone = req.Phone
	user.Password = encryptPassword
	user.Salt = salt
	user.Is_admin = 0    //是否后端管理员
	user.User_status = 1 //用户状态 0 禁用 1 正常
	//用户名称 用户+手机号后4位
	userName := g.I18n().T(ctx, "用户")
	user.Name = userName + req.Phone[len(req.Phone)-4:]
	user.DefaultCreateTime()
	user.DefaultUpdateTime()
	var r sql.Result
	r, err = g.Model("user").Insert(user)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	userId, _ := r.LastInsertId()
	// 生成默认bot
	CreateDefaultChainBot(uint(userId))
	//生成 TokenUsageLimit 默认数据
	SetTokenUsageLimit(uint(userId))
	user.Id = uint(userId)
	session.Session.SetUser(ctx, user)
	// //更新邀请码状态
	// g.Model("invitationCode").Where("code = ? and status = ?", req.InvitationCode, "可用").Update(g.Map{
	// 	"status":           "不可用",
	// 	"userId":           userId,
	// 	"verificationTime": gtime.Now(),
	// })
	return
}

// 注册谷歌用户 createGoogleUser
func (s *UserService) CreateGoogleUser(ctx context.Context, email string, googleId string, name string) (err error) {
	salt := grand.S(6, false)
	//encryptPassword, _ := gmd5.EncryptString(req.Password + salt)
	var user *entity.User = new(entity.User)
	user.Phone = ""
	user.GoogleId = googleId
	user.Email = email
	user.Salt = salt
	user.Is_admin = 0    //是否后端管理员
	user.User_status = 1 //用户状态 0 禁用 1 正常
	//用户名称 用户+手机号后4位
	user.Name = name
	user.DefaultCreateTime()
	user.DefaultUpdateTime()
	var r sql.Result
	r, err = g.Model("user").Insert(user)
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}
	userId, _ := r.LastInsertId()
	// 生成默认bot
	CreateDefaultChainBot(uint(userId))
	//生成 TokenUsageLimit 默认数据
	SetTokenUsageLimit(uint(userId))
	user.Id = uint(userId)
	session.Session.SetUser(ctx, user)

	return
}

func (s *UserService) CheckPhone(Phone string) (bool, error) {
	count, err := g.Model("user").Where("deleteTime is null and phone = ?", Phone).Count()
	if err != nil {
		return true, err
	}
	return count > 0, nil
}

func (s *UserService) SignOut(ctx context.Context, req *v1.UserSignOutReq) error {

	return nil
}

func (s *UserService) LoginWithVerifyCode(ctx context.Context, req *v1.UserLoginWithVerifyCodeReq) (*v1.UserLoginResult, error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and phone = ? or email = ?", req.Phone, req.Phone).One()
	r.Struct(&user)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("登录异常")
	}
	var verificationCode *entity.VerificationCode
	r, err = g.Model("verificationCode").Where("code=? and phone=? and isVerified=0 and expireTime>?", req.Code, req.Phone, gtime.Now()).One()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("验证码错误")
	}
	r.Struct(&verificationCode)
	if verificationCode == nil {
		return nil, errors.New("验证码错误")
	}
	_, err = g.Model("verificationCode").Data(g.Map{"isVerified": 1}).Where("id=?", verificationCode.Id).Update()
	if err != nil {
		return nil, err
	}
	if user != nil && user.User_status != 1 {
		return nil, errors.New("账户已被禁用")

	}
	if user == nil {
		return nil, errors.New("抱歉，您的账号不存在，请点击注册，开始创建您的新账号。")
		// salt := grand.S(6, false)
		// user = new(entity.User)
		// user.DefaultCreateTime()
		// user.DefaultUpdateTime()
		// user.Password = "123456"
		// user.Phone = req.Phone
		// user.Salt = salt
		// user.Is_admin = 0    //是否后端管理员
		// user.User_status = 1 //用户状态 0 禁用 1 正常
		// //用户名称 用户+手机号后4位
		// user.Name = "用户" + req.Phone[len(req.Phone)-4:]
		// r, err := g.Model("user").Insert(&user)
		// if err != nil {
		// 	return nil, err
		// }
		// userId, _ := r.LastInsertId()
		// user.Id = uint(userId)
		// //生成 TokenUsageLimit 默认数据
		// SetTokenUsageLimit(uint(userId))
	}
	session.Session.SetUser(ctx, user)
	userLogin := new(v1.UserLoginResult)
	userLogin.Name = user.Name
	userLogin.Email = user.Email
	userLogin.Phone = user.Phone
	userLogin.UserId = user.Id
	userLogin.Is_admin = user.Is_admin
	return userLogin, nil
}

func CreateDefaultChainBot(userId uint) {
	chain := new(entity.Chain)
	chain.CreateTime = gtime.Now()
	chain.UpdateTime = gtime.Now()
	chain.UserId = userId
	chain.ChainName = "GPT-3.5"
	nodesStr := `[{"id": "6caa6d97-b44f-4689-9204-c9a64cb956e6", "data": {"desc": "用户输入内容", "type": "UserInputText", "title": "UserInputText", "value": {}}, "size": {"width": 300, "height": 0}, "ports": {"items": [{"id": "text", "attrs": {"circle": {"r": 6, "style": {"stroke": "#52c41a"}, "magnet": true}}, "group": "text"}], "groups": {"text": {"position": {"args": {"dy": 140}, "name": "right"}}}}, "shape": "flow-node", "position": {"x": -540, "y": -260}}, {"id": "52231773-e824-4bc0-a1cb-5fb9c87fd797", "data": {"desc": "Call OpenAI LLM", "type": "ChatCompletion", "title": "ChatCompletion", "value": {"modelName": "gpt-3.5-turbo"}}, "size": {"width": 300, "height": 0}, "ports": {"items": [{"id": "prompt", "attrs": {"circle": {"r": 6, "style": {"stroke": "#52c41a"}, "magnet": true}}, "group": "prompt"}, {"id": "ChatOpenAI", "attrs": {"circle": {"r": 6, "style": {"stroke": "#52c41a"}, "magnet": true}}, "group": "ChatOpenAI"}], "groups": {"prompt": {"position": {"args": {"dy": 320}, "name": "left"}}, "ChatOpenAI": {"position": {"args": {"dy": 380}, "name": "right"}}}}, "shape": "flow-node", "position": {"x": 30, "y": -310}}]`
	var nodes []map[string]interface{}
	json.Unmarshal([]byte(nodesStr), &nodes)
	chain.Nodes = nodes
	edgesStr := `[{"id": "631f31bd-0d6d-44fb-a56d-58826185083b", "attrs": {"line": {"stroke": "#c2c8d5", "targetMarker": {"name": "block", "size": 8}}}, "shape": "next", "router": {"name": "normal"}, "source": {"cell": "6caa6d97-b44f-4689-9204-c9a64cb956e6", "port": "text"}, "target": {"cell": "52231773-e824-4bc0-a1cb-5fb9c87fd797", "port": "prompt"}, "zIndex": 0, "connector": {"args": {"radius": 4}, "name": "smooth"}}]`
	var edges []map[string]interface{}
	json.Unmarshal([]byte(edgesStr), &edges)
	chain.Edges = edges
	chain.UserType = "user"
	chain.ChainLevel = 1

	chainRes, err := g.Model("chain").Insert(&chain)
	if err != nil {
		fmt.Println("初始化创建chain失败")
	}
	chainId, _ := chainRes.LastInsertId()
	chain.Id = uint(chainId)

	bot := new(entity.Bot)
	bot.CreateTime = gtime.Now()
	bot.UpdateTime = gtime.Now()
	bot.BotName = "GPT-3.5"
	bot.ChainId = chain.Id
	bot.UserType = "user"
	bot.UserId = userId
	bot.Icon = ""
	bot.IsDefault = 1

	if _, err := g.Model("bot").Insert(&bot); err != nil {
		fmt.Println("初始化创建bot失败")
	}
}

func SetTokenUsageLimit(userId uint) {
	// models, err := modelService.New().List(nil)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// for i := 0; i < len(models); i++ {
	tokenUsageLimitService.New().Create(userId)
	//}
}

func (s *UserService) ForgetPassword(ctx context.Context, req *v1.UserForgetPasswordReq) (res *v1.UserForgetPasswordRes, err error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and phone = ? or email = ?", req.Phone, req.Phone).One()
	r.Struct(&user)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("用户不存在")
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	var verificationCode *entity.VerificationCode
	r, err = g.Model("verificationCode").Where("code=? and phone=? and isVerified=0 and expireTime>?", req.Code, req.Phone, gtime.Now()).One()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("验证码错误")
	}
	r.Struct(&verificationCode)
	if verificationCode == nil {
		return nil, errors.New("验证码错误")
	}
	_, err = g.Model("verificationCode").Data(g.Map{"isVerified": 1}).Where("id=?", verificationCode.Id).Update()
	newEncryptedPassword, _ := gmd5.EncryptString(req.Password + user.Salt)
	_, err = g.Model("user").
		Data("password", string(newEncryptedPassword)).
		Where("phone", user.Phone).
		Update()
	if err != nil {
		response.JsonFailExit(ctx, err.Error())
	}

	return
}

func (s *UserService) Detail(req *v1.FindUserReq) (*entity.User, error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and phone = ? or email = ?", req.Phone, req.Email).One()
	r.Struct(&user)
	return user, err
}

func (s *UserService) GetUserById(userId uint) (*entity.User, error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and id =  ?", userId).One()
	r.Struct(&user)
	return user, err
}

//更新用户信息

func (s *UserService) UpdateUserById(userId uint) (*entity.User, error) {
	var user *entity.User
	r, err := g.Model("user").Where("deleteTime is null and id =  ?", userId).One()
	if err != nil {
		return nil, err
	}
	r.Struct(&user)
	//判断会员截止时间
	if user.Is_membership == 1 && user.Member_deadline.After(gtime.Now()) == false {
		//更新用户状态
		_, err = g.Model("user").Data(g.Map{"is_membership": 0}).Where("deleteTime is null and id = ?", user.Id).Update()
		user.Is_membership = 0
	}
	return user, nil
}
