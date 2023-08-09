package types

import (
	"MyTikTok/conf"
	"MyTikTok/repository/db/model"
)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	*UserInfo `json:"user"`
}

type UserInfo struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackGroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorite   int64  `json:"total_favorite"`   //获赞数量
	WorkCount       int64  `json:"work_count"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
}

func BuildUserInfo(user *model.User) *UserInfo {
	return &UserInfo{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackGroundImage: conf.Host + conf.HttpPort + conf.BackGroundImagePath + user.BackGroundImage,
		Signature:       user.Signature,
		TotalFavorite:   user.TotalFavorite,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}
