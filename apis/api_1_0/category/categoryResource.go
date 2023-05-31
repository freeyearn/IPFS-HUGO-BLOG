package category

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/parser"
	"IPFS-Blog-Hugo/utils/structs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func CategoryGetBasicHandler(c *gin.Context) {
	var err error

	category_id := c.Param("category_id")

	var CategoryService services.CategoryService

	err = CategoryService.Get(map[string]any{"category_id": category_id})

	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", CategoryService)
}

func CategoryGetHandler(c *gin.Context) {
	var err error

	var Parser struct {
		AutoId       int64     `json:"auto_id" form:"auto_id"`
		CategoryId   string    `json:"category_id" form:"category_id"`
		UserId       string    `json:"user_id" form:"user_id"`
		CategoryName string    `json:"category_name" form:"category_name"`
		IsDeleted    int       `json:"is_deleted" form:"is_deleted"`
		CreateTime   time.Time `json:"create_time" form:"create_time"`
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

	var CategoryService services.CategoryService
	CategoryService.Assign(Parser)
	err = CategoryService.Get(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", CategoryService)
}

func CategoryPostHandler(c *gin.Context) {
	var err error

	var Parser struct {
		AutoId       int64     `json:"auto_id" form:"auto_id"`
		CategoryId   string    `json:"category_id" form:"category_id"`
		UserId       string    `json:"user_id" form:"user_id" binding:"required"`
		CategoryName string    `json:"category_name" form:"category_name" binding:"required"`
		IsDeleted    int       `json:"is_deleted" form:"is_deleted"`
		CreateTime   time.Time `json:"create_time" form:"create_time"`
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

	var CategoryService services.CategoryService
	CategoryService.Assign(Parser)

	categoryId := utils.CreateRandomId(viper.GetInt("security.IDSuffixNum"))
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	CategoryService.SetCategoryId(categoryId)

	err = CategoryService.Add(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", CategoryService)
}

func CategoryPutHandler(c *gin.Context) {
	var err error

	categoryId := c.Param("category_id")
	var Parser struct {
		UserId       string `json:"user_id" form:"user_id"`
		CategoryName string `json:"category_name" form:"category_name"`
		IsDeleted    int    `json:"is_deleted" form:"is_deleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var CategoryService services.CategoryService

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	//不能修改业务主键
	delete(args, "user_id")

	err = CategoryService.Update(map[string]any{"category_id": categoryId}, args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", CategoryService)
}

func CategoryDeleteHandler(c *gin.Context) {
	var err error

	var Parser struct {
		CategoryId string `json:"category_id" form:"category_id" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var CategoryService services.CategoryService
	CategoryService.Assign(Parser)

	err = CategoryService.Update(map[string]any{"category_id": Parser.CategoryId},
		map[string]any{"is_deleted": 1})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", CategoryService)
}

// CategoryGetListHandler 获取列表
func CategoryGetListHandler(c *gin.Context) {
	var err error

	var Parser struct {
		UserId string `json:"user_id" form:"user_id"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var CategoryService services.CategoryService
	CategoryService.Assign(Parser)

	results, err := CategoryService.GetCategories()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}

// CategoryGetListByPage 获取列表（分页）
func CategoryGetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		services.CategoryService
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
