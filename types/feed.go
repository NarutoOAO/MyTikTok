package types

import "MyTikTok/repository/db/model"

type FeedResponse struct {
	Response
	VideoList []*model.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}
