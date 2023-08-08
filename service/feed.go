package service

import (
	"MyTikTok/pkg/e"
	dao2 "MyTikTok/repository/db/dao"
	"MyTikTok/types"
	"context"
	"time"
)

type FeedService struct {
}

func (service *FeedService) Feed(ctx context.Context) interface{} {
	code := e.SUCCESS
	dao := dao2.NewVideoDao(ctx)
	videoList, err := dao.GetVideos()
	if err != nil {
		code = e.ERROR
		return types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		}
	}
	return types.FeedResponse{
		Response: types.Response{
			StatusCode: code,
			StatusMsg:  e.GetMsg(code),
		},
		VideoList: videoList,
		NextTime:  time.Now().Unix(),
	}
}
