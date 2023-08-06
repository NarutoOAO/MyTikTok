package router

import (
	"MyTikTok/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
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

		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
		}
	}
	return r
}
