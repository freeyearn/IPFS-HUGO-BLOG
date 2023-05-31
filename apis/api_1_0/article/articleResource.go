package admin

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils/parser"
	"IPFS-Blog-Hugo/utils/structs"
	"github.com/gin-gonic/gin"
	"time"
)

func ArticleGetBasicHandler(c *gin.Context) {
	var err error

	articleId := c.Param("id")

	var ArticleService services.ArticleService

	err = ArticleService.Get(map[string]any{"article_id": articleId})

	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticleGetHandler(c *gin.Context) {
	var err error

	var Parser struct {
		ArticleId   string    `json:"article_id" form:"article_id"`
		CID         string    `json:"cid" form:"cid"`
		AuthorId    string    `json:"author_id" form:"author_id"`
		Title       string    `json:"title" form:"title" binding:"required"`
		Description string    `json:"description" form:"description" binding:"required"`
		Tags        string    `json:"tags" form:"tags"`
		Category    string    `json:"category" form:"category"`
		Keyword     string    `json:"keyword" form:"keyword"`
		Next        string    `json:"next" form:"next"`
		Prev        string    `json:"prev" form:"prev"`
		PublishTime time.Time `json:"publish_time" form:"publish_time"`
		IsDeleted   int32     `json:"is_deleted" form:"is_deleted"`
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

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)
	err = ArticleService.Get(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticlePostHandler(c *gin.Context) {
	var err error

	var Parser struct {
		ArticleId   string    `json:"article_id" form:"article_id"`
		CID         string    `json:"cid" form:"cid"`
		AuthorId    string    `json:"author_id" form:"author_id"`
		Title       string    `json:"title" form:"title" binding:"required"`
		Description string    `json:"description" form:"description" binding:"required"`
		Tags        string    `json:"tags" form:"tags"`
		Category    string    `json:"category" form:"category"`
		Keyword     string    `json:"keyword" form:"keyword"`
		Next        string    `json:"next" form:"next"`
		Prev        string    `json:"prev" form:"prev"`
		PublishTime time.Time `json:"publish_time" form:"publish_time"`
		IsDeleted   int32     `json:"is_deleted" form:"is_deleted"`
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

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)

	err = ArticleService.Add(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticlePutHandler(c *gin.Context) {
	var err error

	article_id := c.Param("article_id")
	var Parser struct {
		ArticleId   string    `json:"article_id" form:"article_id"`
		CID         string    `json:"cid" form:"cid"`
		AuthorId    string    `json:"author_id" form:"author_id"`
		Title       string    `json:"title" form:"title" binding:"required"`
		Description string    `json:"description" form:"description" binding:"required"`
		Tags        string    `json:"tags" form:"tags"`
		Category    string    `json:"category" form:"category"`
		Keyword     string    `json:"keyword" form:"keyword"`
		Next        string    `json:"next" form:"next"`
		Prev        string    `json:"prev" form:"prev"`
		PublishTime time.Time `json:"publish_time" form:"publish_time"`
		IsDeleted   int32     `json:"is_deleted" form:"is_deleted"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ArticleService services.ArticleService

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	//不能修改业务主键
	delete(args, "article_id")

	err = ArticleService.Update(map[string]any{"article_id": article_id}, args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticleDeleteHandler(c *gin.Context) {
	var err error

	var Parser struct {
		ArticleId int `json:"article_id" form:"article_id" binding:"required"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)

	args, err := structs.StructToMap(Parser, "json")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	err = ArticleService.Delete(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "", ArticleService)
}

// ArticleGetListHandler get list of articles.
func ArticleGetListHandler(c *gin.Context) {
	var err error
	var ArticleService services.ArticleService

	err = c.ShouldBind(&ArticleService)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	results, err := ArticleService.GetList()
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results)
}

// ArticleGetListByPage get list divided by parameter: page and size.
func ArticleGetListByPage(c *gin.Context) {
	var err error

	var Parser struct {
		AuthorId string `json:"author_id" form:"author_id" binding:"required"`
		Title    string `json:"title" form:"title"`
		Category string `json:"category" form:"category"`
		parser.ListParser
	}

	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	var articleService services.ArticleService
	articleService.Assign(Parser)
	results, count, err := articleService.GetListByPage(Parser.ListParser)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	parser.JsonOK(c, "", results, map[string]any{
		"totalCount": count,
		"totalPage":  int(count)/Parser.Size + 1,
	})
}
