package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

const (
	PassWordCost        = 12       //password
	Active       string = "active" //active users
)

type User struct {
	*gorm.Model
	Id              int64  // 用户id
	Name            string `gorm:"unique"` // 用户名称
	PassWord        string
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

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PassWord = string(bytes)
	return nil
}

// CheckPassword check your password
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password))
	return err == nil
}

// AvatarUrl set avatar
func (user *User) AvatarURL() string {
	signedGetURL := user.Avatar
	return signedGetURL
}
