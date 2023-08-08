package dao

import (
	"MyTikTok/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

// NewVideoDaoDao this file is to search videos and create videos
func NewVideoDao(ctx context.Context) *VideoDao {
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) GetVideos() (videos []*model.Video, err error) {
	err = dao.DB.Model(&model.Video{}).Order("created_at desc").Limit(30).Find(&videos).Error
	return
}
