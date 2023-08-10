package service

import (
	"MyTikTok/pkg/e"
	util "MyTikTok/pkg/utils"
	dao2 "MyTikTok/repository/db/dao"
	"MyTikTok/repository/db/model"
	"MyTikTok/types"
	"context"
	"time"
)

type CommentService struct {
	Token       string
	VideoId     int64
	UserId      int64
	CommentText string
	CommentId   int64
}

func (service *CommentService) CreateComment(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewCommentDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	dao3 := dao2.NewVideoDao(ctx)
	var err error
	claim, err := util.ParseToken(service.Token)
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  "token解析失败",
		}
	}
	comment := &model.Comment{
		VideoId:    service.VideoId,
		UserId:     claim.ID,
		Content:    service.CommentText,
		CreateDate: time.Now().Format("01-02"),
	}
	err = dao.CreateComment(comment)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	user, err := dao1.GetUser(claim.ID)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	video, err := dao3.GetVideoById(service.VideoId)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	video.CommentCount++
	err = dao3.UpdateVideo(video.Id, video)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return types.CommentResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		CommentInfo: &types.CommentInfo{
			Id:         comment.Id,
			User:       types.BuildUserInfo(user),
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		},
	}
}

func (service *CommentService) DeleteComment(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewCommentDao(ctx)
	dao3 := dao2.NewVideoDao(ctx)
	var err error
	comment, err := dao.GetComment(service.CommentId)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	claim, _ := util.ParseToken(service.Token)
	if comment.UserId != claim.ID {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  "你没有权限删除评论",
		}
	}
	err = dao.DeleteComment(service.CommentId)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	video, err := dao3.GetVideoById(service.VideoId)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	video.CommentCount--
	err = dao3.UpdateVideo(video.Id, video)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return types.Response{
		StatusCode: code,
		StatusMsg:  e.GetMsg(code),
	}
}

func (service *CommentService) GetComments(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewCommentDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	var err error
	var commentList []*types.CommentInfo
	//claim, err := util.ParseToken(service.Token)
	//if err != nil {
	//	code = e.ERROR
	//	return types.Response{
	//		StatusCode: code,
	//		StatusMsg:  "token解析失败",
	//	}
	//}
	comments, err := dao.GetComments(service.VideoId)
	if err != nil {
		code = e.ErrorDatabase
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	for _, v := range comments {
		user, _ := dao1.GetUser(v.UserId)
		commentInfo := &types.CommentInfo{
			Id:         v.Id,
			User:       types.BuildUserInfo(user),
			Content:    v.Content,
			CreateDate: v.CreateDate,
		}
		commentList = append(commentList, commentInfo)
	}
	return types.CommentsResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		CommentInfos: commentList,
	}
}
