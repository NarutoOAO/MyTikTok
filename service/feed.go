package service

import (
	"MyTikTok/conf"
	"MyTikTok/pkg/e"
	util "MyTikTok/pkg/utils"
	dao2 "MyTikTok/repository/db/dao"
	"MyTikTok/repository/db/model"
	"MyTikTok/types"
	"context"
	"mime/multipart"
	"path/filepath"
	"time"
)

type FeedService struct {
	Token  string
	Title  string
	UserId int64
}

func (service *FeedService) Feed(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewVideoDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	videoList, err := dao.GetVideos()
	var videoListResp []*types.Video
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	for _, v := range videoList {
		user, _ := dao1.GetUser(v.UserId)
		video := &types.Video{
			Id:            v.Id,
			Author:        types.BuildUserInfo(user),
			PlayUrl:       conf.Host + conf.HttpPort + conf.VideoPath + v.PlayUrl,
			CoverUrl:      conf.Host + conf.HttpPort + conf.VideoCoverImagePath + v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		}
		videoListResp = append(videoListResp, video)
	}
	return types.FeedResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		VideoList: videoListResp,
		NextTime:  time.Now().Unix(),
	}
}

func (service *FeedService) Publish(ctx context.Context, userId int64, data multipart.File, fileHeader *multipart.FileHeader) interface{} {
	code := e.SUCCESS
	dao := dao2.NewVideoDao(ctx)
	var video *model.Video
	var err error
	filename := filepath.Base(fileHeader.Filename)
	path, err := util.UploadVideoToLocalStatic(data, userId, filename)
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	video = &model.Video{
		UserId:        userId,
		PlayUrl:       path,
		CoverUrl:      "cover.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         service.Title,
	}
	err = dao.CreateVideo(video)
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

func (service *FeedService) GetVideoList(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewVideoDao(ctx)
	dao1 := dao2.NewUserDao(ctx)
	videoList, err := dao.GetVideosById(service.UserId)
	var videoListResp []*types.Video
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	for _, v := range videoList {
		user, _ := dao1.GetUser(v.UserId)
		video := &types.Video{
			Id:            v.Id,
			Author:        types.BuildUserInfo(user),
			PlayUrl:       conf.Host + conf.HttpPort + conf.VideoPath + v.PlayUrl,
			CoverUrl:      conf.Host + conf.HttpPort + conf.VideoCoverImagePath + v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
		}
		videoListResp = append(videoListResp, video)
	}
	return types.FeedResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		VideoList: videoListResp,
		NextTime:  time.Now().Unix(),
	}
}
