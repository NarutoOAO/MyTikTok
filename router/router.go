package router

import (
	api "MyTikTok/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Cors())
	r.StaticFS("/static", http.Dir("./static"))
	v1 := r.Group("/douyin")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "pong"})
		})
		v1.GET("/feed/", api.Feed)
		v1.POST("/user/register/", api.Register)
		v1.POST("/user/login/", api.Login)
		v1.GET("/user/", api.GetUserInfo)
		v1.POST("/publish/action/", api.Publish)
		v1.GET("/publish/list/", api.PublishList)

		//评论操作
		v1.POST("/comment/action/", api.CommentOperation)
		v1.GET("/comment/list/", api.GetComments)
		v1.POST("/favorite/action/", api.FavoriteAction)
		v1.GET("/favorite/list/", api.GetFavoriteList)

	}
	return r
}
