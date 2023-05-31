package admin

import (
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitArticleRouterGroup(engine *gin.RouterGroup) {
	Api = engine.Group("article")
	Api.GET("", ArticleGetHandler)
	Api.GET(":id", ArticleGetBasicHandler)
	Api.POST("", ArticleUpload)
	Api.PUT(":id", ArticleUpdate)
	Api.DELETE(":id", ArticleDelete)
	Api.GET("list", ArticleGetListHandler)
	Api.GET("list/page", ArticleGetListByPage)
	Api.GET("count", ArticleCount)
	Api.GET("fetch", ArticleFetch)
	Api.PUT("status", ArticleStatusUpdate)
}
