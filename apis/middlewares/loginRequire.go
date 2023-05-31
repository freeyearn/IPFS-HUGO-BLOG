package middlewares

import (
	"IPFS-Blog-Hugo/utils/currentUser"
	"IPFS-Blog-Hugo/utils/parser"
	"github.com/gin-gonic/gin"
)

func LoginRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取当前用户
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}

		//查询用户信息
		//var userInfoService services.UserInfoService
		//userInfoService.SetUserID(user.UserID)
		//err = userInfoService.Get()
		//if err != nil {
		//	parser.JsonDBError(c, "", err)
		//	c.Abort()
		//	return
		//}
		//user.UserType = userInfoService.UserType
		c.Set("user", user)
		c.Next()
	}
}
