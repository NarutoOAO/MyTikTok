package v1

import (
	service2 "MyTikTok/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Feed(c *gin.Context) {
	service := service2.FeedService{}
	//code := e.SUCCESS
	//if err := c.ShouldBind(service); err == nil {
	//	code = e.ERROR
	//	c.JSON(http.StatusBadRequest, types.Response{
	//		StatusCode: int32(code),
	//		StatusMsg:  e.GetMsg(code),
	//	})
	//	util.LogrusObj.Info(err)
	//	return
	//}
	response := service.Feed(c.Request.Context())
	c.JSON(http.StatusOK, response)
}
