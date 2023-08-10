package dao

import (
	"MyTikTok/repository/db/model"
	"context"

	"gorm.io/gorm"
)

type CommentDao struct {
	*gorm.DB
}

// NewCommentDao this file is to search videos and create videos
func NewCommentDao(ctx context.Context) *CommentDao {
	return &CommentDao{NewDBClient(ctx)}
}

func (dao *CommentDao) CreateComment(comment *model.Comment) error {
	err := dao.DB.Model(&model.Comment{}).Create(&comment).Error
	return err
}

func (dao *CommentDao) DeleteComment(cId int64) (err error) {
	err = dao.DB.Where("id=?", cId).Delete(&model.Comment{}).Error
	return
}

func (dao *CommentDao) GetComment(id int64) (comment *model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("id=?", id).Find(&comment).Error
	return
}

func (dao *CommentDao) GetComments(vId int64) (comments []*model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("video_id=?", vId).Find(&comments).Error
	return
}
