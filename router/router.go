package router

import (
	api "MyTikTok/api/v1"
	"MyTikTok/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	//r.POST("chat", api.Chat)
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "pong"})
		})
		v1.GET("/douyin/feed/", api.Feed)
		authed := v1.Group("/")

		authed.Use(middleware.JWT())
		{
		}
	}
	return r
}
