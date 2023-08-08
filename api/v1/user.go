package v1

import (
	util "MyTikTok/pkg/utils"
	service2 "MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	service := service2.UserService{}
	service.UserName = c.Query("username")
	service.PassWord = c.Query("password")
	response := service.Register(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func Login(c *gin.Context) {
	service := service2.UserService{}
	service.UserName = c.Query("username")
	service.PassWord = c.Query("password")
	response := service.Login(c.Request.Context())
	c.JSON(http.StatusOK, response)
}

func GetUserInfo(c *gin.Context) {
	service := service2.UserService{}
	service.Token = c.Query("token")
	claim, _ := util.ParseToken(service.Token)
	response := service.GetUserInfo(c.Request.Context(), int64(claim.ID))
	c.JSON(http.StatusOK, response)
}
