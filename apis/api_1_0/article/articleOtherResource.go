package admin

import (
	"IPFS-Blog-Hugo/internal/services"
	"IPFS-Blog-Hugo/utils"
	"IPFS-Blog-Hugo/utils/jwt"
	"IPFS-Blog-Hugo/utils/message"
	"IPFS-Blog-Hugo/utils/parser"
	"IPFS-Blog-Hugo/utils/structs"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func ArticleUpload(c *gin.Context) {
	var err error

	var Parser struct {
		AuthorId    string `json:"author_id" form:"author_id"  binding:"required"`
		AuthorName  string `json:"author_name" form:"author_name"  binding:"required"`
		Title       string `json:"title" form:"title" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
		Tags        string `json:"tags" form:"tags"`
		Category    string `json:"category" form:"category"`
		Keyword     string `json:"keyword" form:"keyword"`
		Next        string `json:"next" form:"next"`
		Prev        string `json:"prev" form:"prev"`
		Status      int    `json:"status" form:"status"`
		Content     string `json:"content" form:"content"`
		Date        string `json:"-"`
		LastMod     string `json:"-"`
		Categories  string `json:"-"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	args, err := structs.StructToMap(Parser, "json")
	delete(args, "author_name")
	delete(args, "content")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)

	// if the data exists
	err = ArticleService.Get(args)
	if err == nil {
		parser.JsonDBError(c, "data exists", errors.New("data exists error"))
		return
	}

	ArticleService.ArticleId = utils.CreateRandomId(viper.GetInt("security.IDSuffixNum"))

	// handle file
	path := filepath.Join(
		utils.CreateFileUploadDir(viper.GetString("blog.ContentDir")),
		utils.CreateFileName(viper.GetString("blog.Suffix")))

	filePath := strings.Split(path, "post")
	fileUrl := strings.Replace(filePath[len(filePath)-1], "\\", "/", -1)
	fileUrl = strings.Replace(fileUrl, ".md", "/index.html", 1)
	ArticleService.Path = fileUrl

	Parser.Date = utils.GetCurrentDay()
	Parser.LastMod = Parser.Date

	// get categories category are json encoded array of Model Category's ID
	// need to change them to CategoryName, saving them to template
	if Parser.Category != "" {
		var category []string
		err := json.Unmarshal([]byte(Parser.Category), &category)
		if err != nil {
			parser.JsonInternalError(c, "category parse error", err)
			return
		}
		categories := make([]string, len(category))
		for idx, value := range category {
			categoryService := services.CategoryService{}
			err := categoryService.Get(map[string]any{
				"category_id": value,
			})
			if err != nil {
				message.PrintErr("category parse err:", err)
			}
			categories[idx] = categoryService.GetCategoryName()
		}
		Parser.Categories = "["
		for _, v := range categories {
			Parser.Categories = Parser.Categories + fmt.Sprintf("\"%s\", ", v)
		}

		Parser.Categories = Parser.Categories + "]"
	}

	// generate blog
	err = utils.GenerateBlog(Parser, viper.GetString("template.Blog"), path)
	if err != nil {
		parser.JsonInternalError(c, "blog save error", err)
		return
	}

	err = ArticleService.Add(args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// hugo generate new blogs
	// status marks the blog is a draft or isn't.
	if Parser.Status == 1 {
		go func() {
			services.CompileAndUpload()
		}()
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticleUpdate(c *gin.Context) {
	var err error
	articleId := c.Param("id")
	var Parser struct {
		ArticleId   string `json:"article_id" form:"article_id"`
		AuthorName  string `json:"author_name" form:"author_name"  binding:"required"`
		Title       string `json:"title" form:"title" binding:"required"`
		Description string `json:"description" form:"description" binding:"required"`
		Tags        string `json:"tags" form:"tags"`
		Category    string `json:"category" form:"category"`
		Keyword     string `json:"keyword" form:"keyword"`
		Next        string `json:"next" form:"next"`
		Prev        string `json:"prev" form:"prev"`
		Status      int    `json:"status" form:"status"`
		Content     string `json:"content" form:"content"`
		Date        string `json:"-"`
		LastMod     string `json:"-"`
		Categories  string `json:"-"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}
	Parser.ArticleId = articleId

	args, err := structs.StructToMap(Parser, "json")

	delete(args, "content")
	delete(args, "author_name")
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)

	// if the data exists
	err = ArticleService.Get(map[string]any{"article_id": articleId})
	if err != nil {
		parser.JsonDBError(c, "data not exists", err)
		return
	}

	// handle file
	path := strings.Replace(ArticleService.Path, "/index.html", "."+viper.GetString("blog.Suffix"), 1)
	savePath, _ := os.Getwd()
	path = filepath.Join(savePath, viper.GetString("blog.Dir"), viper.GetString("blog.ContentDir"), path)
	//fmt.Println("path:", path)
	Parser.Date = utils.TransferToDay(ArticleService.PublishTime)
	Parser.LastMod = utils.GetCurrentDay()

	// get categories category are json encoded array of Model Category's ID
	// need to change them to CategoryName, saving them to template
	if Parser.Category != "" && Parser.Category != ArticleService.Category {
		var category []string
		err := json.Unmarshal([]byte(Parser.Category), &category)
		if err != nil {
			parser.JsonInternalError(c, "category parse error", err)
		}
		categories := make([]string, len(category))
		for idx, value := range category {
			categoryService := services.CategoryService{}
			err := categoryService.Get(map[string]any{
				"category_id": value,
			})
			if err != nil {
				message.PrintErr("category query err: id:", value, " err:", err)
			}
			categories[idx] = categoryService.GetCategoryName()
		}
		Parser.Categories = "["
		for _, v := range categories {
			Parser.Categories = Parser.Categories + fmt.Sprintf("\"%s\", ", v)
		}

		Parser.Categories = Parser.Categories + "]"
	}

	// generate blog
	err = utils.GenerateBlog(Parser, viper.GetString("template.Blog"), path)
	if err != nil {
		parser.JsonInternalError(c, "blog save error", err)
		return
	}

	err = ArticleService.Update(map[string]any{"article_id": articleId}, args)
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// hugo generate new blogs
	// status marks the blog is a draft or isn't.
	if Parser.Status == 1 {
		go func() {
			services.CompileAndUpload()
		}()
	}
	parser.JsonOK(c, "", ArticleService)
}

func ArticleDelete(c *gin.Context) {
	var err error
	articleId := c.Param("id")

	var ArticleService services.ArticleService
	ArticleService.ArticleId = articleId

	// if the data exists
	err = ArticleService.Get(map[string]any{"article_id": articleId})
	if err != nil {
		parser.JsonDBError(c, "data not exists", err)
		return
	}

	// handle file
	path := strings.Replace(ArticleService.Path, "/index.html", "."+viper.GetString("blog.Suffix"), 1)
	savePath, _ := os.Getwd()
	path = filepath.Join(savePath, viper.GetString("blog.Dir"), viper.GetString("blog.ContentDir"), path)

	err = ArticleService.Update(map[string]any{"article_id": articleId}, map[string]any{"is_deleted": 1})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// remove blog
	err = os.Remove(path)
	if err != nil {
		parser.JsonInternalError(c, "delete failed", err)
		return
	}

	go func() {
		services.CompileAndUpload()
	}()

	parser.JsonOK(c, "", ArticleService)
}

func ArticleCount(c *gin.Context) {
	var err error
	var Parser struct {
		AuthorId string `json:"author_id" form:"author_id"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)

	count := ArticleService.Count()

	parser.JsonOK(c, "", map[string]any{
		"count": count,
	})
}

func ArticleFetch(c *gin.Context) {
	var err error

	var Parser struct {
		ArticleId string `json:"article_id" form:"article_id"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	token := c.GetHeader("token")
	jwtClaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
	if err != nil {
		parser.JsonAccessDenied(c, "please login")
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)
	err = ArticleService.Get(map[string]any{
		"article_id": Parser.ArticleId,
		"author_id":  jwtClaim.UserId,
	})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// handle file
	path := strings.Replace(ArticleService.Path, "/index.html", "."+viper.GetString("blog.Suffix"), 1)
	savePath, _ := os.Getwd()
	path = filepath.Join(savePath, viper.GetString("blog.Dir"), viper.GetString("blog.ContentDir"), path)

	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	contents := string(file)
	contentList := strings.Split(contents, "---")
	content := contentList[len(contentList)-1]
	content = content[2 : len(content)-1]

	type Result struct {
		services.ArticleService
		Content string
	}
	parser.JsonOK(c, "", Result{
		ArticleService,
		content,
	})
}

func ArticleStatusUpdate(c *gin.Context) {
	var err error

	var Parser struct {
		ArticleId string `json:"article_id" form:"article_id"`
		Status    int    `json:"status" form:"status"`
	}
	err = c.ShouldBind(&Parser)
	if err != nil {
		parser.JsonParameterIllegal(c, "", err)
		return
	}

	token := c.GetHeader("token")
	jwtClaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
	if err != nil {
		parser.JsonAccessDenied(c, "please login")
		return
	}

	var ArticleService services.ArticleService
	ArticleService.Assign(Parser)
	err = ArticleService.Get(map[string]any{
		"article_id": Parser.ArticleId,
		"author_id":  jwtClaim.UserId,
	})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}

	// handle file
	path := strings.Replace(ArticleService.Path, "/index.html", "."+viper.GetString("blog.Suffix"), 1)
	savePath, _ := os.Getwd()
	path = filepath.Join(savePath, viper.GetString("blog.Dir"), viper.GetString("blog.ContentDir"), path)

	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	contents := string(file)
	contentList := strings.Split(contents, "---")
	content := contentList[len(contentList)-1]
	content = content[2 : len(content)-1]

	var Blog struct {
		services.ArticleService
		AuthorName string `json:"author_name" form:"author_name"`
		Content    string `json:"content" form:"content"`
		Date       string `json:"-"`
		LastMod    string `json:"-"`
		Categories string `json:"-"`
	}
	Blog.ArticleService = ArticleService
	Blog.AuthorName = jwtClaim.Username
	Blog.Content = content
	Blog.Date = utils.GetCurrentDay()
	Blog.LastMod = Blog.Date
	Blog.Status = Parser.Status
	// get categories category are json encoded array of Model Category's ID
	// need to change them to CategoryName, saving them to template
	if Blog.Category != "" {
		var category []string
		err := json.Unmarshal([]byte(Blog.Category), &category)
		if err != nil {
			parser.JsonInternalError(c, "category parse error", err)
			return
		}
		categories := make([]string, len(category))
		for idx, value := range category {
			categoryService := services.CategoryService{}
			err := categoryService.Get(map[string]any{
				"category_id": value,
			})
			if err != nil {
				message.PrintErr("category parse err:", err)
			}
			categories[idx] = categoryService.GetCategoryName()
		}
		Blog.Categories = "["
		for _, v := range categories {
			Blog.Categories = Blog.Categories + fmt.Sprintf("\"%s\", ", v)
		}

		Blog.Categories = Blog.Categories + "]"
	}

	// generate blog
	err = utils.CoverBlog(Blog, viper.GetString("template.Blog"), path)
	if err != nil {
		parser.JsonInternalError(c, "blog save error", err)
		return
	}

	// update status
	err = ArticleService.Update(
		map[string]any{
			"article_id": Parser.ArticleId,
			"author_id":  jwtClaim.UserId,
		},
		map[string]any{
			"status": Parser.Status,
		})
	if err != nil {
		parser.JsonDBError(c, "", err)
		return
	}
	parser.JsonOK(c, "status update success", ArticleService)
}
