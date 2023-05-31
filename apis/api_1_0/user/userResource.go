package user

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils/parser"
	"IPFS-Blog-Hugo/utils/structs"
	"github.com/gin-gonic/gin"
	"time"
)

func UserGetBasicHandler(c *gin.Context) {
	var err error

	userId := c.Param("user_id")

	var UserService services.UserService

	err = UserService.Get(map[string]any{"user_id": userId})

	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

func UserGetHandler(c *gin.Context) {
	var err error

	var Parser struct {
		UserId     string    `json:"user_id" form:"user_id"`
		Wallet     string    `json:"wallet" form:"wallet"`
		Name       string    `json:"name" form:"name"`
		Password   string    `json:"password" form:"password"`
		Age        int       `json:"age" form:"age"`
		Status     int       `json:"status" form:"status"`
		CreateTime time.Time `json:"create_time" form:"create_time"`
		UpdateTime time.Time `json:"update_time" form:"update_time"`
		IsDeleted  int       `json:"is_deleted" form:"is_deleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService
	UserService.Assign(Parser)
	err = UserService.Get(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

func UserPostHandler(c *gin.Context) {
	var err error

	var Parser struct {
		UserId     string    `json:"user_id" form:"user_id"`
		Wallet     string    `json:"wallet" form:"wallet"`
		Name       string    `json:"name" form:"name"`
		Password   string    `json:"password" form:"password"`
		Age        int       `json:"age" form:"age"`
		Status     int       `json:"status" form:"status"`
		CreateTime time.Time `json:"create_time" form:"create_time"`
		UpdateTime time.Time `json:"update_time" form:"update_time"`
		IsDeleted  int       `json:"is_deleted" form:"is_deleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService
	UserService.Assign(Parser)

	err = UserService.Add(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

func UserPutHandler(c *gin.Context) {
	var err error

	userId := c.Param("user_id")
	var Parser struct {
		UserId     string    `json:"user_id" form:"user_id"`
		Wallet     string    `json:"wallet" form:"wallet"`
		Name       string    `json:"name" form:"name"`
		Password   string    `json:"password" form:"password"`
		Age        int       `json:"age" form:"age"`
		Status     int       `json:"status" form:"status"`
		CreateTime time.Time `json:"create_time" form:"create_time"`
		UpdateTime time.Time `json:"update_time" form:"update_time"`
		IsDeleted  int       `json:"is_deleted" form:"is_deleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	//不能修改业务主键
	delete(args, "user_id")

	err = UserService.Update(map[string]any{"user_id": userId}, args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

func UserDeleteHandler(c *gin.Context) {
	var err error

	var Parser struct {
		Id int `json:"id" form:"id" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var UserService services.UserService
	UserService.Assign(Parser)

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	err = UserService.Delete(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", UserService)
}

// UserGetListHandler 获取列表
func UserGetListHandler(c *gin.Context) {
	var err error
	var UserService services.UserService

	err = c.ShouldBind(&UserService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := UserService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}

// UserGetListByPage 获取列表（分页）
func UserGetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		services.UserService
		parser.ListParser
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, count, err := Parser.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results, map[string]any{
		"totalCount": count,
		"totalPage":  int(count)/Parser.Size + 1,
	})
}
