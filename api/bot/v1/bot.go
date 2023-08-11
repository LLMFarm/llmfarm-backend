package v1

import (
	"llmfarm/api"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetPublicBotListReq struct {
	g.Meta     `path:"/bot/getPublicBotList" tags:"机器人相关接口" method:"get" summary:"查询公开机器人列表"`
	Word       string `json:"word" dc:"关键词" `
	Limit      int    `json:"limit" dc:"每页数量" v:"required"`
	PageNumber int    `json:"pageNumber" dc:"当前页码 从1开始" v:"required"`
	PromptType int    `json:"PromptType" dc:" 1:提示词 2:提示词工程 3:提示词流程图 如果是全部传0" `
	ScenarioId int    `json:"scenarioId" dc:"场景Id 如果是全部传0"`
}
type Bot struct {
	entity.Bot
	ChainName      string `json:"chainName" dc:"chain名称"`
	UseKnowledge   uint   `json:"useKnowledge" dc:"是否使用知识库" `
	UseUserCount   int    `json:"useUserCount" dc:"使用人数" `
	UseCount       int    `json:"useCount" dc:"使用次数" `
	g.Meta         `mime:"application/json" example:"string"`
	ShowName       string `json:"showName" dc:"显示名称，如果是系统的就展示官方，如果是个人的展示用户名"`
	ShareOpen      bool   `json:"shareOpen" default:"false" `
	UserPermission string `json:"userPermission" dc:"用户权限"`
}

type BotDetail struct {
	entity.Bot
	ChainName         string `json:"chainName" dc:"chain名称"`
	UseKnowledge      uint   `json:"useKnowledge" dc:"是否使用知识库" `
	g.Meta            `mime:"application/json" example:"string"`
	BotSettingConfigs []*SettingConfig `json:"configs" dc:"配置列表"`
	UserPermission    string           `json:"userPermission" dc:"用户权限"`
	UserName          string           `json:"userName" dc:"用户姓名"`
	UserPhone         string           `json:"userPhone" dc:"用户手机号"`
}

type SettingConfig struct {
	*entity.SettingConfig
	SettingId uint `json:"settingId" dc:"settingId"`
	g.Meta    `mime:"application/json" example:"string"`
}

type UpdateBot struct {
	entity.Bot
	NeedUpdate bool `json:"needUpdate" dc:"需要更新"`
	g.Meta     `mime:"application/json" example:"string"`
}
type ListBotReq struct {
	g.Meta `path:"/bot/list" tags:"机器人相关接口" method:"get" summary:"查询机器人列表"`
	Word   string `json:"word" dc:"关键词" `
}

type GetPublicBotListRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data BotListData `json:"data" dc:"数据体"`
}
type BotListData struct {
	BotList   []*Bot `json:"botList" dc:"bot列表"`
	PageCount int    `json:"pageCount" dc:"总页数"`
}
type ListBotRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []Bot `json:"data" dc:"数据体"`
}

type DetailBotReq struct {
	g.Meta `path:"/bot/detail" tags:"机器人相关接口" method:"get" summary:"查询机器人详情"`
	BotId  uint `json:"botId" dc:"机器人id" v:"required"`
}

type DetailBotRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Bot `json:"data" dc:"数据体"`
}

type DeleteBotReq struct {
	g.Meta `path:"/bot" tags:"机器人相关接口" method:"delete" summary:"删除机器人"`
	BotId  uint `json:"botId" dc:"机器人id" v:"required"`
}

type DeleteBotRes struct {
	g.Meta `mime:"application/json"`
	api.Common
}

type CreateBotReq struct {
	g.Meta     `path:"/bot" tags:"机器人相关接口" method:"post" summary:"创建机器人"`
	Name       string `json:"name" dc:"机器人名称" v:"required"`
	ChainId    uint   `json:"chainId" dc:"提示链id" v:"required"`
	IsPublic   bool   `json:"isPublic" dc:"是否公开" `
	Icon       string `json:"icon" dc:"图标" v:"required-if:isPublic,true"`
	ScenarioId uint   `json:"scenarioId" dc:"场景Id" v:"required-if:isPublic,true"`
	Desc       string `json:"desc" dc:"机器人描述" v:"required-if:isPublic,true"`
}

type CreateBotRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Bot `json:"data" dc:"数据体"`
}

type ModifyBotReq struct {
	g.Meta `path:"/bot" tags:"机器人相关接口" method:"put" summary:"修改机器人"`
	BotId  uint   `json:"botId" dc:"机器人id" v:"required"`
	Name   string `json:"name" dc:"机器人名称" v:"required"`
	Icon   string `json:"icon" dc:"图标" v:"required-if:isPublic,true"`
}

type ModifyBotRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.Bot `json:"data" dc:"数据体"`
}

type AddLibsReq struct {
	g.Meta `path:"/bot/addLibs" tags:"机器人相关接口" method:"post" summary:"机器人添加文件夹"`
	LibIds []int `json:"libIds" dc:"文件夹id数组" v:"required"`
	BotId  int   `json:"botId" dc:"机器人id" v:"required"`
}

type AddLibsRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FolderRes `json:"data" dc:"数据体"`
}

type SubLibsReq struct {
	g.Meta `path:"/bot/subLibs" tags:"机器人相关接口" method:"post" summary:"机器人去除文件夹"`
	LibIds []int `json:"libIds" dc:"文件夹id数组" v:"required"`
	BotId  int   `json:"botId" dc:"机器人id" v:"required"`
}

type SubLibsRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FolderRes `json:"data" dc:"数据体"`
}

type GetLibsReq struct {
	g.Meta `path:"/bot/getLibs" tags:"机器人相关接口" method:"post" summary:"获取文件夹列表"`
	BotId  int `json:"botId" dc:"机器人id" v:"required"`
}

type GetLibsRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FolderRes `json:"data" dc:"数据体"`
}

type GetUsedBotListReq struct {
	g.Meta `path:"/bot/usedBotList" tags:"机器人相关接口" method:"get" summary:"获取最近使用过的机器人"`
}

type GetUsedBotListRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.Bot `json:"data" dc:"数据体"`
}

type AddFilesReq struct {
	g.Meta  `path:"/bot/addFiles" tags:"机器人相关接口" method:"post" summary:"机器人添加文件"`
	FileIds []int `json:"fileIds" dc:"资源文件id数组" v:"required"`
	BotId   int   `json:"botId" dc:"机器人id" v:"required"`
}

type AddFilesRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FileRes `json:"data" dc:"数据体"`
}

type SubFilesReq struct {
	g.Meta  `path:"/bot/subFiles" tags:"机器人相关接口" method:"post" summary:"机器人去除文件"`
	FileIds []int `json:"fileIds" dc:"资源文件id数组" v:"required"`
	BotId   int   `json:"botId" dc:"机器人id" v:"required"`
}

type SubFilesRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FileRes `json:"data" dc:"数据体"`
}

type GetFilesReq struct {
	g.Meta `path:"/bot/getFiles" tags:"机器人相关接口" method:"post" summary:"获取机器人文件列表"`
	BotId  int `json:"botId" dc:"机器人id" v:"required"`
}

type GetFilesRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.FileRes `json:"data" dc:"数据体"`
}

type BotNoSettingReq struct {
	g.Meta `path:"/bot/botNoSetting" tags:"机器人相关接口" method:"post" summary:"获取机器人未配置的类型数组"`
	BotId  int `json:"botId" dc:"机器人id" v:"required"`
}

type BotNoSettingRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []string `json:"data" dc:"数据体"`
}

// bot 添加 dataSet
type BotAddDataSetReq struct {
	g.Meta     `path:"/bot/dataSet" tags:"数据集相关接口" method:"post" summary:"bot关联数据集"`
	DataSetIds []int `json:"dataSetIds" dc:"数据集Id" v:"required"`
	BotId      int   `json:"botId" dc:"机器人Id" v:"required"`
}

type BotAddDataSetRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.DataSet `json:"data" dc:"数据体"`
}

// bot 删除 dataSet
type BotDeleteDataSetReq struct {
	g.Meta    `path:"/bot/dataSet" tags:"数据集相关接口" method:"delete" summary:"bot删除数据集"`
	DataSetId int `json:"dataSetId" dc:"数据集Id" v:"required"`
	BotId     int `json:"botId" dc:"机器人Id" v:"required"`
}

type BotDeleteDataSetRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.DataSet `json:"data" dc:"数据体"`
}

// bot 添加 botSetFolder
type BotAddBotSetFolderReq struct {
	g.Meta       `path:"/bot/botSetFolder" tags:"botSetFolder相关接口" method:"post" summary:"添加botSetFolder"`
	SetFolderIds []int `json:"setFolderIds" dc:"文件夹Id" v:"required"`
	BotId        int   `json:"botId" dc:"机器人Id" v:"required"`
}

type BotAddBotSetFolderRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.BotSetFolder `json:"data" dc:"数据体"`
}

// bot 删除 botSetFolder
type BotDeleteBotSetFolderReq struct {
	g.Meta      `path:"/bot/botSetFolder" tags:"botSetFolder相关接口" method:"delete" summary:"删除botSetFolder"`
	SetFolderId int `json:"setFolderId" dc:"文件夹Id" v:"required"`
	BotId       int `json:"botId" dc:"机器人Id" v:"required"`
}

type BotDeleteBotSetFolderRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data entity.BotSetFolder `json:"data" dc:"数据体"`
}

// 根据bot 获取bot_set_folder
type BotGetBotSetFolderReq struct {
	g.Meta `path:"/bot/botSetFolder" tags:"botSetFolder相关接口" method:"get" summary:"获取botSetFolder"`
	BotId  int `json:"botId" dc:"机器人Id" v:"required"`
}

type BotGetBotSetFolderRes struct {
	g.Meta `mime:"application/json"`
	api.Common
	Data []entity.BotSetFolder `json:"data" dc:"数据体"`
}

type SetFolderInfo struct {
	entity.SetFolder
	UserName string `json:"userName"`
}
