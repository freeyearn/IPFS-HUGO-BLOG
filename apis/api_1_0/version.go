package api_1_0

import (
	"IPFS-Blog-Hugo/utils/codes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetVersion(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    codes.OK,
		"message": "当前接口版本信息。",
		"data": gin.H{
			"version": "1.0",
		},
	})
	return
}
