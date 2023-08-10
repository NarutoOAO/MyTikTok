package v1

import (
	service2 "MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentOperation(c *gin.Context) {
	var response interface{}
	service := service2.CommentService{}
	service.Token = c.Query("token")
	id := c.Query("video_id")
	vId, _ := strconv.Atoi(id)
	service.VideoId = int64(vId)
	operation := c.Query("action_type")
	if operation == "1" {
		service.CommentText = c.Query("comment_text")
		response = service.CreateComment(c.Request.Context())
	} else if operation == "2" {
		id = c.Query("comment_id")
		cId, _ := strconv.Atoi(id)
		service.CommentId = int64(cId)
		response = service.DeleteComment(c.Request.Context())
	}
	c.JSON(http.StatusOK, response)
}

func GetComments(c *gin.Context) {
	service := service2.CommentService{}
	//service.Token = c.Query("token")
	Id := c.Query("video_id")
	vId, _ := strconv.Atoi(Id)
	service.VideoId = int64(vId)
	response := service.GetComments(c.Request.Context())
	c.JSON(http.StatusOK, response)
}
