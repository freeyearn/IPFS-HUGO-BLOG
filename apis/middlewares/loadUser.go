package middlewares

import (
	"github.com/gin-gonic/gin"
	//"go_hugo_ipfs_api/utils/currentUser"
	//"go_hugo_ipfs_api/utils/parser"
)

func LoadUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		if token == "" {
			c.Next()
			return
		}

		// todo token verify
		//res, err := rpcReq.VerifyToken(token)
		//if err != nil {
		//	c.Next()
		//	return
		//}

		//加载用户信息到上下文
		//User, err := currentUser.NewUser(res.UserID, currentUser.UnKnown)
		//defer currentUser.Release(User)
		//if err != nil {
		//	parser.JsonDBError(c, "用户信息未找到", err)
		//	c.Abort()
		//	return
		//}

		//c.Set("user", User)
		//c.Next()
	}
}
