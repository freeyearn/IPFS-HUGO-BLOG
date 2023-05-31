package category

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitCategoryRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("category")
	Api.GET("", CategoryGetHandler)
	Api.GET(":category_id", CategoryGetBasicHandler)
	Api.POST("", CategoryPostHandler)
	Api.PUT(":category_id", CategoryPutHandler)
	Api.DELETE("", CategoryDeleteHandler)
	Api.GET("list", CategoryGetListHandler)
	Api.GET("list/page", CategoryGetListByPage)
}
