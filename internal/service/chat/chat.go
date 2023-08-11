package service

import (
	"errors"
	"fmt"
	v1 "llmfarm/api/chat/v1"
	"llmfarm/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type ChatService struct{}

func New() *ChatService {
	return &ChatService{}
}

// 创建会话列表
func (s *ChatService) Create(req *v1.CreateChatReq, userId uint) (*entity.Chat, error) {
	var chat *entity.Chat = new(entity.Chat)
	chat.ChatName = req.Name
	chat.Type = req.Type
	chat.UserId = userId
	chat.ContextType = req.ContextType
	chat.IsTemp = req.IsTemp
	chat.CreateTime = gtime.Now()
	chat.UpdateTime = gtime.Now()
	res, err := g.Model("chat").Insert(&chat)
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("创建失败")
	}
	chatId, err := res.LastInsertId()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("创建失败")
	}
	chat.Id = uint(chatId)
	return chat, err
}

func (s *ChatService) List(req *v1.ChatListReq, userId uint) ([]entity.Chat, error) {
	var chats []entity.Chat
	mChat := g.Model("chat").Where("deleteTime is null and isTemp=0 and userId=?", userId)
	if req.Keyword != "" {
		mChat.Where("name like ?", "%"+req.Keyword+"%")
	}
	if req.Type != "" {
		mChat.Where("type = ?", req.Type)
	}
	res, err := mChat.OrderDesc("createTime").All()
	if err != nil {
		return chats, err
	}
	res.Structs(&chats)
	return chats, err
}

func (s *ChatService) ModifyChatContentType(req *v1.ModifyChatContextTypeReq, userId uint) (*entity.Chat, error) {
	var chat *entity.Chat
	r, err := g.Model("chat").Where("id=? and userId=?", req.ChatId, userId).One()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("会话不存在")
	}
	r.Struct(&chat)
	if chat == nil {
		return nil, errors.New("会话不存在")
	}
	_, err = g.Model("chat").Data(g.Map{"contextType": req.ContextType}).Where("id=? and userId=?", req.ChatId, userId).Update()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("修改失败")
	}
	chat.ContextType = req.ContextType
	return chat, nil
}

func (s *ChatService) Modify(req *v1.ModifyChatReq, userId uint) (*entity.Chat, error) {
	var chat *entity.Chat
	r, err := g.Model("chat").Where("id=? and userId=?", req.ChatId, userId).One()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("会话不存在")
	}
	r.Struct(&chat)
	if chat == nil {
		return nil, errors.New("会话不存在")
	}
	_, err = g.Model("chat").Data(g.Map{"chatName": req.Name}).Where("id=? and userId=?", req.ChatId, userId).Update()
	if err != nil {
		fmt.Println("err", err)
		return nil, errors.New("修改失败")
	}
	chat.ChatName = req.Name
	return chat, nil
}

func (s *ChatService) Delete(req *v1.DeleteChatReq, userId uint) error {
	_, err := g.Model("chat").Data(g.Map{"deleteTime": gtime.Now()}).Where("id=? and userId=?", req.ChatId, userId).Update()
	if err != nil {
		fmt.Println("err", err)
		return errors.New("删除失败")
	}
	return nil
}

func (s *ChatService) ClearChat(req *v1.ClearChatReq, userId uint) error {
	_, err := g.Model("chat").Data(g.Map{"deleteTime": gtime.Now()}).Where("userId=?", userId).Update()
	if err != nil {
		fmt.Println("err", err)
		return errors.New("清除失败")
	}
	return nil
}
