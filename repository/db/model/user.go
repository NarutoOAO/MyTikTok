package model

import "github.com/jinzhu/gorm"

type User struct {
	*gorm.Model
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
