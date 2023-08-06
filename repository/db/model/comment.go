package model

type Comment struct {
	Id         int64  // 视频评论id
	UserId     int64  // 评论用户信息
	Content    string // 评论内容
	CreateDate string // 评论发布日期，格式 mm-dd
}
