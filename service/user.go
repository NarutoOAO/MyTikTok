package service

import (
	"MyTikTok/pkg/e"
	util "MyTikTok/pkg/utils"
	dao2 "MyTikTok/repository/db/dao"
	"MyTikTok/repository/db/model"
	"MyTikTok/types"
	"context"
)

type UserService struct {
	UserName string
	PassWord string
	Token    string
}

func (service *UserService) Register(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewUserDao(ctx)
	var user *model.User
	var err error
	_, exist, err := dao.IfExistOrNot(service.UserName)
	if exist {
		code = e.ErrorExistUser
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	user = &model.User{
		Name:            service.UserName,
		FollowCount:     0,
		FollowerCount:   0,
		IsFollow:        false,
		Avatar:          "",
		BackGroundImage: "",
		Signature:       "这个人很懒，还没有个人简介～",
		TotalFavorite:   0,
		WorkCount:       0,
		FavoriteCount:   0,
	}
	err = user.SetPassword(service.PassWord)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	err = dao.CreateUser(user)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.Id, user.Name)
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return types.UserLoginResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		UserId: user.Id,
		Token:  token,
	}
}

func (service *UserService) Login(ctx context.Context) interface{} {
	var token string
	var user *model.User
	var err error
	code := e.SUCCESS
	dao := dao2.NewUserDao(ctx)
	user, exist, err := dao.IfExistOrNot(service.UserName)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if !exist {
		code = e.ErrorNotExistUser
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if !user.CheckPassword(service.PassWord) {
		code = e.ErrorNotComparePassword
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	token, err = util.GenerateToken(user.Id, user.Name)
	return types.UserLoginResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		UserId: user.Id,
		Token:  token,
	}
}

func (service *UserService) GetUserInfo(ctx context.Context, id int64) interface{} {
	var user *model.User
	var err error
	code := e.SUCCESS
	dao := dao2.NewUserDao(ctx)
	user, err = dao.GetUser(id)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return types.UserInfoResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		UserInfo: types.BuildUserInfo(user),
	}
}
