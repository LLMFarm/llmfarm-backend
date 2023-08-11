package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	v1 "llmfarm/api/bot/v1"
	"llmfarm/internal/model/entity"
	chainService "llmfarm/internal/service/chain"

	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type BotService struct {
}

func New() *BotService {
	return &BotService{}
}

func (s *BotService) List(word string, userId uint) ([]v1.Bot, error) {
	//查询关联表 chainBotUserRelation
	// var chainBotUserRelation []entity.ChainBotUserRelation
	// chainBotUserRelationR, err := g.Model("chain_bot_user_relation").Where("userId = ? and relationType = ?", userId, "bot").All()
	// if err != nil {
	// 	return nil, errors.New("查询失败")
	// }
	// chainBotUserRelationR.Structs(&chainBotUserRelation)

	var botIds []int
	// for _, relation := range chainBotUserRelation {
	// 	botIds = append(botIds, relation.RelationId)
	// }

	var userBot []v1.Bot
	fmt.Println("botIds", botIds)
	//SELECT bot.*,chain.`name` as `chainName` FROM bot,chain WHERE chain.id = bot.chainId AND bot.deleteTime is null and bot.userType = "user" and bot.userId= 13

	// mBot := g.Model("bot").Where("deleteTime is null and userType = ? and userId=?", "user", userId)
	mBot := g.Model("bot").
		InnerJoin("chain", "chain.id = bot.chainId").
		Fields("bot.*, chain.chainName as chainName,chain.useKnowledge as useKnowledge").
		Where("bot.deleteTime IS NULL and bot.botName like ?", "%"+word+"%").
		Where("bot.userType = ?", "user")
	if len(botIds) > 0 {
		mBot = mBot.Where("(bot.userId = ? OR bot.id in (?) )", userId, botIds)
	} else {
		mBot = mBot.Where("bot.userId = ?", userId)
	}

	userBotRes, err := mBot.OrderDesc("updateTime").All()
	if err != nil {
		return userBot, errors.New("查询失败")
	}
	userBotRes.Structs(&userBot)
	var myBotIds []int
	for i, bot := range userBot {
		//判断userId 如果是创建者，返回creator
		if bot.UserId == userId {
			userBot[i].UserPermission = "creator"
			myBotIds = append(myBotIds, int(bot.Id))
		} else {
			// 判断 Id 在不在 chainIds 里面
			for _, botId := range botIds {
				if int(bot.Id) == botId {
					userBot[i].UserPermission = "share"
					break
				}
			}
		}
	}

	//查询一下我的chain 看看那些开启分享了
	// if len(myBotIds) > 0 {
	// 	//如果chain_bot_user_relation 中有chainId 则说明开启分享了
	// 	var shareBotIds []int
	// 	shareBotIdsR, err := g.Model("chain_bot_user_relation").Where("relationId in (?) and relationType = ? and deleteTime is null", myBotIds, "bot").All()
	// 	if err != nil {
	// 		return nil, errors.New("查询失败")
	// 	}
	// 	var shareChainBotUserRelation []entity.ChainBotUserRelation
	// 	shareBotIdsR.Structs(&shareChainBotUserRelation)
	// 	for _, relation := range shareChainBotUserRelation {
	// 		shareBotIds = append(shareBotIds, relation.RelationId)
	// 	}
	// 	for i, bot := range userBot {
	// 		for _, shareBotId := range shareBotIds {
	// 			if int(bot.Id) == shareBotId {
	// 				userBot[i].ShareOpen = true
	// 				break
	// 			}
	// 		}
	// 	}
	// }
	return userBot, nil
}

func (s *BotService) GetPublicBotList(word string, userId uint, limit, pageNumber, promptType, scenarioId int) ([]*v1.Bot, int, error) {
	return nil, 0, errors.New("创建文件夹失败")

	// where := g.Map{}
	// botWhere := g.Map{}
	// if word != "" {
	// 	where["name like ?"] = "%" + word + "%"
	// 	botWhere["bot.name like ?"] = "%" + word + "%"
	// }
	// if scenarioId > 0 {
	// 	where["scenarioId"] = scenarioId
	// 	botWhere["bot.scenarioId"] = scenarioId
	// }
	// if promptType > 0 {
	// 	where["promptType"] = promptType
	// 	botWhere["bot.promptType"] = promptType
	// }
	// botWhere["bot.isPublic = ?"] = 1
	// where["isPublic"] = 1
	// totalCount, err := g.Model("bot").Where("deleteTime is null").Count(where)
	// if err != nil {
	// 	return nil, 0, err
	// }
	// pageCount := totalCount/limit + 1

	// offset := limit * (pageNumber - 1)
	// var systemBot []*v1.Bot
	// systemBotSql := g.Model("bot").
	// 	InnerJoin("chain", "chain.id = bot.chainId").
	// 	Fields("bot.*, chain.name as chainName,chain.useKnowledge as useKnowledge").
	// 	Where("bot.deleteTime IS NULL").
	// 	Where(botWhere).
	// 	OrderDesc("updateTime").Limit(offset, limit)
	// systemBotRes, err := systemBotSql.OrderDesc("createTime").All()
	// if err != nil {
	// 	return nil, 0, err
	// }

	// systemBotRes.Structs(&systemBot)
	// botUsedUserCountMap, _ := s.GetBotUsedUserCount(userId)
	// botUsedCountMap, _ := s.GetBotUsedCount(userId)

	// for _, bot := range systemBot {
	// 	fmt.Println(bot.Id)
	// 	userCount := botUsedUserCountMap[gconv.String(bot.Id)]
	// 	bot.UseUserCount = userCount
	// 	useCount := botUsedCountMap[gconv.String(bot.Id)]
	// 	bot.UseCount = useCount
	// 	if bot.UserType == "system" {
	// 		bot.ShowName = "官方"
	// 	} else {
	// 		var user *entity.User
	// 		fmt.Println("=====", bot.UserId)
	// 		g.Model("user").Where("id = ?", bot.UserId).Scan(&user)
	// 		if user != nil {
	// 			bot.ShowName = user.Name
	// 		}
	// 	}
	// }
	// return systemBot, pageCount, nil
}

// todo 改成下载chain，然后删掉
// 下载bot
// func (s *BotService) DownloadBot(botId, userId uint) (*entity.Bot, error) {
// 	var systemBot *entity.Bot
// 	r, err := g.Model("bot").Where("id=?", botId).One()
// 	if err != nil {
// 		return nil, errors.New("系统bot查询失败")
// 	}
// 	r.Struct(&systemBot)

// 	g.Model("bot").Data(g.Map{"downloadCount": systemBot.DownloadCount + 1}).Where("id=?", botId).Update()

// 	//获取系统bot的chain信息
// 	instantiationChain, err := chainService.New().Detail(systemBot.ChainId)
// 	if err != nil {
// 		return nil, errors.New("chain模板查询失败")
// 	}
// 	var chainTemplate *entity.Chain
// 	gconv.Struct(instantiationChain, &chainTemplate)
// 	chainTemplate.Type = "template"
// 	//复制chain模板
// 	templateChainNew, err := chainService.New().CopyChain(chainTemplate, userId, 0)
// 	//复制chain实例化
// 	instantiationChainNew, err := chainService.New().CopyChain(instantiationChain, userId, templateChainNew.Id)
// 	//复制bot
// 	var bot *entity.Bot = new(entity.Bot)
// 	bot.UserId = userId
// 	bot.Name = systemBot.Name
// 	bot.ChainId = instantiationChainNew.Id
// 	bot.UserType = "user"
// 	bot.Uuid = uuid.New()
// 	bot.CreateTime = gtime.Now()
// 	bot.UpdateTime = gtime.Now()
// 	bot.IsPublic = 0
// 	bot.ScenarioId = systemBot.ScenarioId
// 	bot.Icon = systemBot.Icon
// 	res, err := g.Model("bot").Insert(&bot)
// 	if err != nil {
// 		return nil, errors.New("bot复制失败")
// 	}

// 	newBotId, _ := res.LastInsertId()
// 	bot.Id = uint(newBotId)
// 	return bot, nil
// }

type CheckUpdateBot struct {
	InstanceId         int
	TemplateId         int
	InstanceUpdateTime *gtime.Time
	TemplateUpdateTime *gtime.Time
}

func (s *BotService) Create(req *v1.CreateBotReq, userId uint) (*entity.Bot, error) {
	var bot *entity.Bot = new(entity.Bot)
	bot.UserId = userId
	bot.BotName = req.Name
	bot.ChainId = req.ChainId
	bot.UserType = "user"
	bot.CreateTime = gtime.Now()
	bot.UpdateTime = gtime.Now()

	res, err := g.Model("bot").Insert(&bot)
	if err != nil {
		return nil, errors.New("创建失败")
	}
	botId, err := res.LastInsertId()
	if err != nil {
		return nil, errors.New("创建失败")
	}
	bot.Id = uint(botId)

	return bot, err
}

func (s *BotService) Modify(req *v1.ModifyBotReq, userId uint) (*entity.Bot, error) {
	var bot *entity.Bot
	r, err := g.Model("bot").Where("id=?  and userId=?", req.BotId, userId).One()
	if err != nil {
		return nil, errors.New("bot查询失败")
	}
	r.Struct(&bot)
	if bot.CreateTime == nil {
		return nil, errors.New("bot不存在")
	}
	if bot.UserType == "system" {
		return nil, errors.New("系统bot不能编辑")
	}

	updateData := g.Map{"botName": req.Name}

	_, err = g.Model("bot").Data(updateData).Where("id=? and userId=?", req.BotId, userId).Update()
	if err != nil {
		return nil, errors.New("修改失败")
	}
	bot.BotName = req.Name
	return bot, err
}

func (s *BotService) Delete(req *v1.DeleteBotReq, userId uint) error {
	var bot *entity.Bot
	r, err := g.Model("bot").Where("id=?  and userId=?", req.BotId, userId).One()
	if err != nil {
		return errors.New("bot查询失败")
	}
	r.Struct(&bot)
	if bot.CreateTime == nil {
		return errors.New("bot不存在")
	}
	if bot.UserType == "system" {
		return errors.New("系统bot不能删除")
	}
	_, err = g.Model("bot").Data(g.Map{"deleteTime": gtime.Now()}).Where("id=? and userId=?", req.BotId, userId).Update()
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}

func (s *BotService) Detail(botId uint, userId uint) (*v1.BotDetail, error) {
	var bot *v1.BotDetail
	//r, err := g.Model("bot").Where(" deleteTime is null and id=? ", botId).One()
	r, err := g.Model("bot").
		InnerJoin("chain", "chain.id = bot.chainId").
		InnerJoin("user", "user.id = bot.userId").
		Fields("bot.*, user.name as userName,user.phone as userPhone ,chain.chainName as chainName,chain.useKnowledge as useKnowledge").
		Where("bot.deleteTime IS NULL").
		Where("bot.id = ?", botId).
		One()
	if err != nil {
		return nil, errors.New("bot不存在")
	}
	r.Struct(&bot)
	if bot == nil {
		return nil, errors.New("bot不存在")
	}
	//根据botId 获取Chain信息
	chain, _ := chainService.New().Detail(bot.ChainId, userId)
	// 定义 一维数组 serviceList
	var serviceList []string
	Nodes := chain.Chain.Nodes
	//循环 Nodes 获取节点信息
	//VectorSearch
	bot.UseKnowledge = 0
	for _, node := range Nodes {
		//根据节点类型获取节点信息
		fmt.Println("node=====", node)

		data := node["data"].(map[string]interface{})
		if data["service"] != nil {
			service := data["service"].(string)
			fmt.Println("service=====", service)
			serviceList = append(serviceList, service)
		}
		if data["type"] != nil {
			fmt.Println("type=====", data["type"])
			nodeType := data["type"].(string)
			if nodeType == "VectorSearch" {
				fmt.Println("nodeType=====", nodeType)
				bot.UseKnowledge = 1
			}
		}

	}

	if bot.UserId == userId {
		bot.UserPermission = "creator"
	} else {
		//去关联表里查询
		var chainBotUserRelation = new(entity.ChainBotUserRelation)
		r, err := g.Model("chain_bot_user_relation").Where("relationId = ? and userId = ? and relationType = ? and deleteTime is null", botId, userId, "bot").One()
		if err != nil {
			return nil, errors.New("查询失败")
		}
		r.Struct(&chainBotUserRelation)
		bot.UserPermission = chainBotUserRelation.UserPermission
	}
	return bot, err
}

func (s *BotService) GetLibs(botId uint) ([]*entity.FolderRes, error) {
	var botFolders []*entity.BotFolder
	err := g.Model("botFolder").Where("botId=? ", botId).Scan(&botFolders)
	if err != nil {
		return nil, errors.New("查询关联表失败")
	}

	libIds := g.Slice{}
	for _, botLib := range botFolders {
		libIds = append(libIds, botLib.FolderId)
	}

	var folderReses []*entity.FolderRes
	if err := g.Model("folder").Where("id IN(?)", libIds).
		ScanList(&folderReses, "Folder"); err != nil {
		return nil, err
	}
	if err := g.Model("user").Where("id", gdb.ListItemValuesUnique(folderReses, "Folder", "UserId")).
		ScanList(&folderReses, "UserInfo", "Folder", "id:userId"); err != nil {
		return nil, err
	}

	return folderReses, nil
}

func (s *BotService) AddLibs(botId uint, libIds []int) ([]*entity.FolderRes, error) {
	folders := g.Slice{}
	updateData := g.List{}
	for _, id := range libIds {
		folders = append(folders, id)
		updateData = append(updateData, g.Map{"botId": botId, "folderId": id})
	}

	_, err := g.Model("botFolder").Data(updateData).Save()
	var folderFiles []*entity.FolderFile
	err = g.Model("folderFile").Where("folderId IN(?)", folders).Scan(&folderFiles)

	filesData := g.List{}
	if len(folderFiles) > 0 {
		for _, file := range folderFiles {
			filesData = append(filesData, g.Map{"botId": botId, "fileId": file.FileId})
		}

		_, err = g.Model("botUserFile").Data(filesData).Save()
		if err != nil {
			return nil, err
		}
	}
	libs, err := s.GetLibs(botId)
	if err != nil {
		return nil, err
	}
	return libs, nil
}

func (s *BotService) SubLibs(botId uint, libIds []int) ([]*entity.FolderRes, error) {
	_, err := g.Model("botFolder").Where("botId=? and folderId IN(?)", botId, libIds).Delete()
	if err != nil {
		return nil, errors.New("去除文件夹失败")
	}

	var folderFiles []*entity.FolderFile
	g.Model("folderFile").Where("folderId IN (?)", libIds).Scan(&folderFiles)

	fileIds := []int{}
	for _, folderFile := range folderFiles {
		fileIds = append(fileIds, folderFile.FileId)
	}
	if len(fileIds) > 0 {
		// 如果除了libIds还有其他的folder与该bot关联且folder中有file，则不删除，否则删除botfile的关联
		var otherBotFolders []*entity.BotFolder
		g.Model("botFolder").Where("botId=? and folderId NOT IN(?)", botId, libIds).Scan(&otherBotFolders)
		if otherBotFolders != nil && len(otherBotFolders) > 0 {
			otherFolderIds := []int{}
			for _, otherBotFolder := range otherBotFolders {
				otherFolderIds = append(otherFolderIds, otherBotFolder.FolderId)
			}
			noDeleteFileIds := []int{}
			var noDeleteFolderFiles []*entity.FolderFile
			g.Model("folderFile").Where("folderId IN (?) and fileId in(?)", otherFolderIds, fileIds).Scan(&noDeleteFolderFiles)
			for _, noDeleteFolderFile := range noDeleteFolderFiles {
				noDeleteFileIds = append(noDeleteFileIds, noDeleteFolderFile.FileId)
			}

			fileSet := gset.NewFrom(fileIds)
			noDeleteFileSet := gset.NewFrom(noDeleteFileIds)
			deleteIds := fileSet.Diff(noDeleteFileSet).Slice()
			g.Model("botUserFile").Where("botId=? and fileId IN (?)", botId, deleteIds).Delete()

		} else {
			g.Model("botUserFile").Where("botId=? and fileId IN (?)", botId, fileIds).Delete()
		}
	}
	libs, err := s.GetLibs(botId)
	if err != nil {
		return nil, err
	}
	return libs, nil
}

func (s *BotService) GetFileIds(botId uint, userId uint) ([]int, error) {
	var botUserFiles []*entity.BotUserFile
	err := g.Model("botUserFile").Where("botId=? ", botId).Scan(&botUserFiles)
	if err != nil {
		return nil, errors.New("查询关联表失败")
	}

	userFileIds := g.Slice{}
	for _, botUserFile := range botUserFiles {
		userFileIds = append(userFileIds, botUserFile.FileId)
	}

	var userFiles []*entity.UserFile
	m := g.Model("userFile").Where("id IN(?)", userFileIds)
	err = m.Scan(&userFiles)
	if err != nil {
		return nil, err
	}

	fileIds := []int{}
	for _, rf := range userFiles {
		fileIds = append(fileIds, rf.FileId)
	}
	return fileIds, nil
}

// 获取使用过的bot
func (s *BotService) UsedBots(userId uint) ([]*entity.Bot, error) {
	res, err := g.Model("message").Fields("botId,createTime").Where("userId = ?", userId).Group("botId,createTime").OrderDesc("createTime").All()
	if err != nil {
		return nil, err
	}
	bIds := g.Slice{}
	var bots []*entity.Bot

	if len(res) == 0 {
		g.Model("bot").Where("id IN(?) and userType = 'user'  and deleteTime is null", bIds).Scan(&bots)
	} else {
		orderString := "FIELD(id,"
		for i, botId := range res {
			bIds = append(bIds, botId["botId"])
			orderString += "'" + botId["botId"].String() + "'"
			if i != len(res)-1 {
				orderString += ", "
			}
		}
		orderString += ")"
		fmt.Println("orderString====", orderString)
		g.Model("bot").Where("id IN(?) and userType = 'user'  and deleteTime is null", bIds).Order(orderString).Scan(&bots)
	}

	return bots, nil
}

// 获取bot对话人数
func (s *BotService) GetBotUsedUserCount(userId uint) (map[string]int, error) {
	botUserCountMap := map[string]int{}

	res, err := g.Model("message").Fields("count(distinct userId) count,botId").Group("botId").All()
	if err != nil {
		return botUserCountMap, errors.New("去除文件夹失败")
	}
	//botUserCountMap := gmap.New()
	for _, re := range res {
		botUserCountMap[re["botId"].String()] = re["count"].Int()
	}

	return botUserCountMap, nil
}

// 获取bot对话次数
func (s *BotService) GetBotUsedCount(userId uint) (map[string]int, error) {
	botUseCountMap := map[string]int{}

	res, err := g.Model("message").Fields("count(userId) count,botId").Group("botId").All()
	if err != nil {
		return nil, errors.New("去除文件夹失败")
	}
	for _, re := range res {
		botUseCountMap[re["botId"].String()] = re["count"].Int()
	}

	return botUseCountMap, nil
}

// 机器人增加文件
func (s *BotService) AddFiles(fileIds []int, botId int, userId uint) ([]*entity.UserFileRes, error) {
	updateData := g.List{}
	for _, id := range fileIds {
		updateData = append(updateData, g.Map{"botId": botId, "fileId": id})
	}

	_, err := g.Model("botUserFile").Data(updateData).Save()
	if err != nil {
		return nil, errors.New("插入失败")
	}

	files, err := s.GetFilesInBot(botId, userId)
	if err != nil {
		return nil, err
	}
	return files, err
}

// 机器人去除文件
func (s *BotService) SubFiles(fileIds []int, botId int, userId uint) ([]*entity.UserFileRes, error) {
	var bot *entity.Bot
	if err := g.Model("bot").Where("id = ?", botId).Scan(&bot); err != nil {
		return nil, err
	}
	if bot.UserId != userId {
		return nil, errors.New("您不是该知识库创建者")
	}

	_, err := g.Model("botUserFile").Where("botId=? and fileId IN(?)", botId, fileIds).Delete()
	if err != nil {
		return nil, errors.New("去除文件失败")
	}

	//同时去掉bot关联文件夹里的文件
	var botFolders []*entity.BotFolder
	g.Model("botFolder").Where("botId = ?", botId).Scan(&botFolders)
	if botFolders != nil && len(botFolders) > 0 {
		folderIds := []int{}
		for _, botFolder := range botFolders {
			folderIds = append(folderIds, botFolder.FolderId)
		}
		g.Model("folderFile").Where("folderId IN (?) and fileId IN (?)", folderIds, fileIds).Delete()
	}

	files, err := s.GetFilesInBot(botId, userId)
	if err != nil {
		return nil, err
	}
	return files, err
}

// 获取机器人下的文件列表
func (s *BotService) GetFilesInBot(botId int, userId uint) ([]*entity.UserFileRes, error) {
	var botUserFiles []*entity.BotUserFile
	g.Model("botUserFile").Where("botId=?", botId).Scan(&botUserFiles)

	newFileIds := g.Slice{}
	for _, file := range botUserFiles {
		newFileIds = append(newFileIds, file.FileId)
	}

	var botFolders []*entity.BotFolder
	folderIds := []int{}
	g.Model("botFolder").Where("botId = ?", botId).Scan(&botFolders)
	for _, botFolder := range botFolders {
		folderIds = append(folderIds, botFolder.FolderId)
	}
	where := g.Map{}
	if len(folderIds) > 0 {
		where["f.id IN(?)"] = folderIds
	}

	var userFileInfos []*entity.UserFileInfo
	g.Model("userFile").LeftJoin("file r", "userFile.fileId = r.id").Fields("userFile.*,r.url as url,r.parseStatus as parseStatus").Where("userFile.deleteTime is null and userFile.id IN(?)", newFileIds).Scan(&userFileInfos)

	var fileReses []*entity.UserFileRes
	// if err := g.Model("userFile").Where("id IN(?)", newFileIds).
	// 	ScanList(&fileReses, "File"); err != nil {
	// 	return nil, err
	// }
	m := g.Model("userFile,file r,folderFile ff,folder f").
		Where("f.id=ff.folderId and ff.fileId=userFile.id and userFile.fileId = r.id and userFile.deleteTime is null and userFile.id IN(?)", newFileIds).Where(where).
		Fields("userFile.*,r.url as url,r.parseStatus as parseStatus,f.folderName as folderName")
	m = m.OrderDesc("userFile.createTime")
	m.Scan(&fileReses)
	err := m.ScanList(&fileReses, "File")
	fileMap := map[int][]string{}
	fileFlagMap := map[int]int{}
	for _, fileRes := range fileReses {
		if fileMap[int(fileRes.File.Id)] == nil {
			fileMap[int(fileRes.File.Id)] = []string{fileRes.FolderName}
		} else {
			fileMap[int(fileRes.File.Id)] = append(fileMap[int(fileRes.File.Id)], fileRes.FolderName)
		}
	}

	for index, fileRes := range fileReses {
		if fileFlagMap[int(fileRes.File.Id)] == 0 {
			a, _ := json.Marshal(fileMap[int(fileRes.File.Id)])
			fileRes.FolderName = string(a)
			fileFlagMap[int(fileRes.File.Id)] = 1
		} else {
			fileReses = append(fileReses[:index], fileReses[index+1:]...)
		}
	}

	for _, userFileInfo := range userFileInfos {
		if fileFlagMap[int(userFileInfo.Id)] == 0 {
			fileRes := new(entity.UserFileRes)
			fileRes.File = userFileInfo
			fileRes.FolderName = "[]"
			fileReses = append(fileReses, fileRes)
		}
	}

	if err != nil {
		return nil, errors.New("查询失败")
	}

	if err := g.Model("user").Where("id", gdb.ListItemValuesUnique(fileReses, "File", "UserId")).
		ScanList(&fileReses, "UserInfo", "File", "id:userId"); err != nil {
		return nil, err
	}

	return fileReses, nil
}

func (s *BotService) ChainBotList(userId, chainId int) ([]*entity.Bot, error) {
	var bots []*entity.Bot
	g.Model("bot").Where("chainId = ? and deleteTime is null and userId = ?", chainId, userId).OrderDesc("createTime").Scan(&bots)

	return bots, nil
}

func (s *BotService) BotNoSetting(botId int) ([]string, error) {
	var bot *entity.Bot
	g.Model("bot").Where("id = ?", botId).Scan(&bot)

	var chain *entity.Chain
	g.Model("chain").Where("id = ?", bot.ChainId).Scan(&chain)

	serviceSet := gset.New(true)
	for _, node := range chain.Nodes {
		data := node["data"].(map[string]interface{})
		if data["service"] != nil {
			serviceSet.Add(data["service"].(string))
		}
	}

	var settings []*entity.Setting
	g.Model("setting").Where("chainId = ? and deleteTime is null", chain.Id).Scan(&settings)
	settingSet := gset.New(true)
	settingValueSet := gset.New(true)
	for _, setting := range settings {
		settingData, _ := json.Marshal(setting)
		fmt.Println("settingData========", string(settingData))
		data, _ := json.Marshal(setting.Value)
		fmt.Println("data================", string(data))
		if string(data) == "{}" {
			settingValueSet.Add(setting.ServiceType)
		} else {
			//循环变量 setting.Value
			for _, value := range setting.Value {
				if value == nil {
					settingValueSet.Add(setting.ServiceType)
				}
			}
		}
		settingSet.Add(setting.ServiceType)
	}

	noset := serviceSet.Diff(settingSet)

	noset = noset.Union(settingValueSet)

	nosetStr := []string{}
	for _, nos := range noset.Slice() {
		nosetStr = append(nosetStr, nos.(string))
	}

	if chain.UseKnowledge == 1 {
		// var botUserFiles []*entity.BotUserFile
		// g.Model("botUserFile").Where("botId = ?", botId).Scan(&botUserFiles)
		// if len(botUserFiles) == 0 {
		// 	nosetStr = append(nosetStr, "UseKnowledge")
		// }
		var BotDataSetData []*entity.BotDataSet
		g.Model("bot_data_set").Where("botId = ?", botId).Scan(&BotDataSetData)
		if len(BotDataSetData) == 0 {
			nosetStr = append(nosetStr, "UseKnowledge")
		}
	}

	return nosetStr, nil
}

// BotAddDataSetReq
func (s *BotService) BotAddDataSet(req *v1.BotAddDataSetReq, userId uint) error {
	BotDataSetData := g.List{}
	for _, id := range req.DataSetIds {
		BotDataSetData = append(BotDataSetData, g.Map{"botId": req.BotId, "dataSetId": id})
	}

	_, err := g.Model("bot_data_set").Data(BotDataSetData).Save()
	if err != nil {
		return errors.New("添加失败")
	}
	return nil
}

// 删除botDataSet
func (s *BotService) BotDeleteDataSet(req *v1.BotDeleteDataSetReq, userId uint) error {
	_, err := g.Model("bot_data_set").Where("botId=? and dataSetId=?", req.BotId, req.DataSetId).Delete()
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}

// BotAddBotSetFolderRes
func (s *BotService) BotAddSetFolder(req *v1.BotAddBotSetFolderReq, userId uint) error {
	//批量添加 bot_set_folder
	folders := g.Slice{}
	updateData := g.List{}
	for _, id := range req.SetFolderIds {
		folders = append(folders, id)
		updateData = append(updateData, g.Map{"botId": req.BotId, "setFolderId": id})
	}
	_, err := g.Model("bot_set_folder").Data(updateData).Save()
	if err != nil {
		return errors.New("添加失败")
	}

	//获取 folder_data_set 中的 setFolderId
	var folderDataSets []entity.FolderDataSet
	m := g.Model("folder_data_set").Where("setFolderId in (?)", req.SetFolderIds)
	err = m.Scan(&folderDataSets)

	if err != nil {
		fmt.Println("=====", err)
		return errors.New("查询失败")
	}
	var dataSetIds []int
	for _, folderDataSet := range folderDataSets {
		dataSetIds = append(dataSetIds, int(folderDataSet.DataSetId))
	}

	fmt.Println("dataSetIds========================", dataSetIds)
	dataSet := g.Slice{}
	dataSetData := g.List{}
	for _, id := range dataSetIds {
		dataSet = append(dataSet, id)
		dataSetData = append(dataSetData, g.Map{"botId": req.BotId, "dataSetId": id})
	}
	_, err = g.Model("bot_data_set").Data(dataSetData).Save()
	return nil

}

// 删除botSetFolder
func (s *BotService) BotDeleteSetFolder(req *v1.BotDeleteBotSetFolderReq, userId uint) error {
	_, err := g.Model("bot_set_folder").Where("botId=? and setFolderId=?", req.BotId, req.SetFolderId).Delete()
	if err != nil {
		return errors.New("folder删除失败")
	}
	//删除 botsetFolder下的 data_set 文件 关联表是 folder_data_set 和 bot_data_set
	//先获取 botsetFolder下的 data_set
	var FolderDataSets []entity.FolderDataSet
	err = g.Model("folder_data_set").Where("setFolderId=?", req.SetFolderId).Scan(&FolderDataSets)
	if err != nil {
		return errors.New("查询失败")
	}
	var dataSetIds []int
	for _, FolderDataSet := range FolderDataSets {
		//判断数据集 有没有和bot关联的其他folder有关联
		var folderDataSets []entity.FolderDataSet
		err = g.Model("folder_data_set").Where("dataSetId=? and setFolderId!=?", FolderDataSet.DataSetId, req.SetFolderId).Scan(&folderDataSets)
		if err != nil {
			return errors.New("查询失败")
		}
		var folderIds []int
		for _, folderDataSet := range folderDataSets {
			folderIds = append(folderIds, int(folderDataSet.SetFolderId))
		}

		if len(folderIds) > 0 {
			//查询bot和folder 有没有关联
			// select count(*) from bot_set_folder where botId=1 and setFolderId in(1,2,3)
			c, err := g.Model("bot_set_folder").Where("botId=? and setFolderId in(?)", req.BotId, folderIds).Count()
			if err != nil {
				return errors.New("查询失败")
			}
			if c == 0 {
				dataSetIds = append(dataSetIds, int(FolderDataSet.DataSetId))
			}
		} else {
			dataSetIds = append(dataSetIds, int(FolderDataSet.DataSetId))
		}
	}

	fmt.Println("dataSetIds===========", dataSetIds)
	//删除 botsetFolder下的 data_set 文件 关联表是 bot_data_set
	_, err = g.Model("bot_data_set").Where("botId=? and dataSetId in(?)", req.BotId, dataSetIds).Delete()
	if err != nil {
		return errors.New("dataSet删除失败")
	}

	return nil
}

// BotGetBotSetFolderRes
func (s *BotService) BotGetSetFolder(req *v1.BotGetBotSetFolderReq, userId uint) ([]*v1.SetFolderInfo, error) {
	//根据botID 获取 bot_set_folder 中的setFolderId
	var botSetFolders []entity.BotSetFolder
	err := g.Model("bot_set_folder").Fields("setFolderId").Where("botId=?", req.BotId).Scan(&botSetFolders)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	var setFolderIds []int
	for _, botSetFolder := range botSetFolders {
		setFolderIds = append(setFolderIds, int(botSetFolder.SetFolderId))
	}

	if len(setFolderIds) > 0 {
		//根据setFolderId 获取 setFolder
		var setFolders []*v1.SetFolderInfo
		err = g.Model("set_folder s").Fields("s.*,u.name as `userName`").Where("s.deleteTime is null and s.id in(?)", setFolderIds).
			InnerJoin("user u", "u.id = s.userId").
			Scan(&setFolders)
		if err != nil {
			fmt.Println("查询失败2 ========", err)
			return nil, errors.New("查询失败")
		}
		return setFolders, nil
	}
	return nil, nil
}

// GetBotDataSet
func (s *BotService) GetBotDataSetIds(botId uint) ([]int, error) {
	//根据botID 获取 bot_data_set 中的dataSetId
	// var dataSetIds []int
	// m := g.Model("bot_data_set").Fields("dataSetId").Where("botId=?", botId)
	// r, err := m.All()
	// if r.Len() > 0 {
	// 	m.Scan(&dataSetIds)
	// 	fmt.Println("dataSetIds=====", dataSetIds)
	// }
	// if err != nil {
	// 	return nil, errors.New("查询失败")
	// }

	var botDataSet []*entity.BotDataSet
	err := g.Model("bot_data_set").Where("botId=?", botId).Scan(&botDataSet)
	if err != nil {
		return nil, errors.New("查询失败")
	}

	var dataSetIds []int
	for _, botDataSet := range botDataSet {
		dataSetIds = append(dataSetIds, int(botDataSet.DataSetId))
	}

	return dataSetIds, nil
}
