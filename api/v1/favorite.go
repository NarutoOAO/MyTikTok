package v1

import (
	service2 "MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	service := service2.FavoriteService{}
	service.Token = c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	vId, _ := strconv.Atoi(videoId)
	aT, _ := strconv.Atoi(actionType)
	service.VideoId = int64(vId)
	service.ActionType = int32(aT)
	response := service.Action(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func GetFavoriteList(c *gin.Context) {
	service := service2.FavoriteService{}
	service.Token = c.Query("token")
	response := service.ShowFavoriteList(c.Request.Context())
	c.JSON(http.StatusOK, response)
}
