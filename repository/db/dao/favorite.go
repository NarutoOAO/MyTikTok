package dao

import (
	"MyTikTok/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	err := dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
	return err
}

func (dao *FavoriteDao) CountFavorite(uId int64, vId int64) (count int64, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=? and video_id=?", uId, vId).Count(&count).Error
	return
}

func (dao *FavoriteDao) GetFavorite(uId int64, vId int64) (favorite *model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=? and video_id=?", uId, vId).First(&favorite).Error
	return
}

func (dao *FavoriteDao) GetFavorites(uId int64) (favorites []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id=?", uId).Find(&favorites).Error
	return
}

func (dao *FavoriteDao) DeleteFavorite(favorite *model.Favorite) error {
	err := dao.DB.Model(&model.Favorite{}).Delete(&favorite).Error
	return err
}
