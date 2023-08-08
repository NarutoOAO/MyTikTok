package types

import "MyTikTok/repository/db/model"

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
	Id              int64  // 用户id
	Name            string // 用户名称
	FollowCount     int64  // 关注总数
	FollowerCount   int64  // 粉丝总数
	IsFollow        bool   // true-已关注，false-未关注
	Avatar          string //用户头像
	BackGroundImage string //用户个人页顶部大图
	Signature       string //个人简介
	TotalFavorite   int64  //获赞数量
	WorkCount       int64  //作品数量
	FavoriteCount   int64  //点赞数量
}

func BuildUserInfo(user *model.User) *UserInfo {
	return &UserInfo{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackGroundImage: user.BackGroundImage,
		Signature:       user.Signature,
		TotalFavorite:   user.TotalFavorite,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}
}
