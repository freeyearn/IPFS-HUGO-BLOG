package user

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitUserRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("user")
	Api.GET("", UserGetHandler)
	Api.GET(":user_id", UserGetBasicHandler)
	Api.POST("", UserPostHandler)
	Api.PUT(":user_id", UserPutHandler)
	Api.DELETE("", UserDeleteHandler)
	Api.GET("list", UserGetListHandler)
	Api.GET("list/page", UserGetListByPage)
	Api.POST("register", Register)
	Api.POST("login", Login)
	Api.GET("user_info", GetUserInfoByToken)
}
