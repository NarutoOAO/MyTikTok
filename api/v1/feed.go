package v1

import (
	util "MyTikTok/pkg/utils"
	service2 "MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Feed(c *gin.Context) {
	service := service2.FeedService{}
	response := service.Feed(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func PublishList(c *gin.Context) {
	service := service2.FeedService{}
	//service.Token = c.Query("token")
	Id := c.Query("user_id")
	uId, _ := strconv.Atoi(Id)
	service.UserId = int64(uId)
	response := service.GetVideoList(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func Publish(c *gin.Context) {
	service := service2.FeedService{}
	service.Token = c.PostForm("token")
	service.Title = c.PostForm("title")
	data, fileHeader, _ := c.Request.FormFile("data")
	claim, _ := util.ParseToken(service.Token)
	response := service.Publish(c.Request.Context(), claim.ID, data, fileHeader)
	c.JSON(http.StatusOK, response)
}
