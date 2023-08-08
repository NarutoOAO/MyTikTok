package dao

import (
	"MyTikTok/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// NewUserDao this file is to search videos and create videos
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{NewDBClient(ctx)}
}

func (dao *UserDao) CreateUser(user *model.User) error {
	err := dao.DB.Model(&model.User{}).Create(&user).Error
	return err
}

func (dao *UserDao) GetUser(id int64) (user *model.User, err error) {
	err = dao.DB.Model(&model.User{}).Where("id=?", id).First(&user).Error
	return
}

func (dao *UserDao) IfExistOrNot(name string) (user *model.User, exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.User{}).Where("name=?", name).Count(&count).Error
	if err != nil {
		return nil, false, err
	}
	if count == 0 {
		return nil, false, nil
	}
	err = dao.DB.Model(&model.User{}).Where("name=?", name).First(&user).Error
	if err != nil {
		return nil, true, err
	}
	return user, true, nil
}
