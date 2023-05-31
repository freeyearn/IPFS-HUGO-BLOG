package middlewares

import (
	"IPFS-Blog-Hugo/utils/currentUser"
	"IPFS-Blog-Hugo/utils/parser"
	"github.com/gin-gonic/gin"
)

// AuthUserType 垂直鉴权
func AuthUserType(allowType ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}
		// 验证权限
		if !user.AuthType(allowType...) {
			parser.JsonAccessDenied(c, "您无权访问！")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}

// AuthUserID 水平鉴权
func AuthUserID(allowUserID ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		//查询账号信息
		user, err := currentUser.GetUser(c)
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}
		// 验证权限
		var f bool
		for _, userID := range allowUserID {
			if user.UserID == userID {
				f = true
				break
			}
		}
		if !f {
			parser.JsonAccessDenied(c, "您无权访问！")
			c.Abort()
			return
		}
		c.Next()
		return
	}
}
