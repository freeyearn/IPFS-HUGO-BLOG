// coding: utf-8
// @Author : lryself
// @Software: GoLand

package api_1_0

import (
	article "IPFS-Blog-Hugo/apis/api_1_0/article"
	"IPFS-Blog-Hugo/apis/api_1_0/category"
	user "IPFS-Blog-Hugo/apis/api_1_0/user"
	"github.com/gin-gonic/gin"
)

var (
	Api *gin.RouterGroup
)

func InitAPIRouter(engine *gin.Engine) {
	Api = engine.Group("api")
	Api.Any("version", GetVersion)

	user.InitUserRouterGroup(Api)
	article.InitArticleRouterGroup(Api)
	category.InitCategoryRouterGroup(Api)
}
