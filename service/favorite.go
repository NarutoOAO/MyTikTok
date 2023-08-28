package service

import (
	"MyTikTok/conf"
	"MyTikTok/pkg/e"
	util "MyTikTok/pkg/utils"
	dao2 "MyTikTok/repository/db/dao"
	"MyTikTok/repository/db/model"
	"MyTikTok/types"
	"context"
)

type FavoriteService struct {
	VideoId    int64
	ActionType int32
	Token      string
}

func (service *FavoriteService) Action(ctx context.Context) interface{} {
	code := e.SUCCESS
	var favorite *model.Favorite
	var err error
	dao := dao2.NewFavoriteDao(ctx)
	claim, _ := util.ParseToken(service.Token)
	count, err := dao.CountFavorite(claim.ID, service.VideoId)
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	if service.ActionType == 1 {
		if count != 0 {
			code = e.ErrorExistFavorite
			return types.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
		favorite = &model.Favorite{
			UserId:  claim.ID,
			VideoId: service.VideoId,
		}
		err = dao.CreateFavorite(favorite)
		if err != nil {
			code = e.ErrorExistFavorite
			return types.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
	} else if service.ActionType == 2 {
		favorite, err = dao.GetFavorite(claim.ID, service.VideoId)
		if err != nil {
			code = e.ERROR
			return types.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
		err = dao.DeleteFavorite(favorite)
		if err != nil {
			code = e.ERROR
			return types.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
	}
	return types.Response{
		StatusCode: code,
		StatusMsg:  e.GetMsg(code),
	}
}

func (service *FavoriteService) ShowFavoriteList(ctx context.Context) interface{} {
	code := e.SUCCESS
	var favorites []*model.Favorite
	var videos []*model.Video
	var err error
	var videoListResp []*types.Video
	dao := dao2.NewFavoriteDao(ctx)
	dao1 := dao2.NewVideoDao(ctx)
	dao3 := dao2.NewUserDao(ctx)
	claim, _ := util.ParseToken(service.Token)
	favorites, err = dao.GetFavorites(claim.ID)
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	for _, v := range favorites {
		video, err := dao1.GetVideoById(v.VideoId)
		if err != nil {
			code = e.ERROR
			return types.Response{
				StatusCode: code,
				StatusMsg:  e.GetMsg(code),
			}
		}
		videos = append(videos, video)
	}
	for _, v := range videos {
		user, _ := dao3.GetUser(v.UserId)
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
	return types.FavoritesResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		VideoList: videoListResp,
	}
}
