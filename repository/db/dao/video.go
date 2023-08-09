package dao

import (
	"MyTikTok/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

// NewVideoDao this file is to search videos and create videos
func NewVideoDao(ctx context.Context) *VideoDao {
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) GetVideos() (videos []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Order("created_at desc").Limit(30).Find(&videos).Error
	return
}

func (dao *VideoDao) CreateVideo(video *model.Video) (err error) {
	err = dao.DB.Model(&model.Video{}).Create(video).Error
	return
}

func (dao *VideoDao) GetVideosById(userId int64) (videos []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Where("user_id=?", userId).Find(&videos).Error
	return
}
