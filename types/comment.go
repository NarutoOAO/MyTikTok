package types

type CommentResponse struct {
	Response
	*CommentInfo `json:"comment"`
}

type CommentInfo struct {
	Id         int64     `json:"id,omitempty"`
	User       *UserInfo `json:"user"`
	Content    string    `json:"content,omitempty"`
	CreateDate string    `json:"create_date,omitempty"`
}

type CommentsResponse struct {
	Response
	CommentInfos []*CommentInfo `json:"comment_list"`
}
