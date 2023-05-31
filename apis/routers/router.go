package routers

import (
	"IPFS-Blog-Hugo/apis/api_1_0"
	middlewares2 "IPFS-Blog-Hugo/apis/middlewares"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
	"strings"
)

func InitRouter(engine *gin.Engine) {
	// 加载全局中间件
	engine.Use(middlewares2.CorsMiddleware())
	engine.Use(middlewares2.LogMiddleware())
	engine.Use(middlewares2.LoadUser())

	//配置静态html
	engine.Static("/web/static", "static")
	engine.HTMLRender = loadTemplates("/web/template")
	engine.StaticFile("/favicon.ico", "/web/static/favicon.ico")

	//部署vue
	//engine.LoadHTMLGlob("/dist/*.html")             // 添加入口index.html
	//engine.LoadHTMLFiles("/web/static/*/*")         // 添加资源路径
	//engine.Static("/static", "./dist/static")       // 添加资源路径
	//engine.StaticFile("/hello/", "dist/index.html") //前端接口

	//初始化路由
	api_1_0.InitAPIRouter(engine)
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	// 非模板嵌套
	htmls, err := filepath.Glob(templatesDir + "/htmls/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, html := range htmls {
		r.AddFromGlob(filepath.Base(html), html)
	}

	// 布局模板
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	// 嵌套的内容模板
	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// template自定义函数
	funcMap := template.FuncMap{
		"StringToLower": func(str string) string {
			return strings.ToLower(str)
		},
	}

	// 将主模板，include页面，layout子模板组合成一个完整的html页面
	for _, include := range includes {
		files := []string{}
		files = append(files, templatesDir+"/frame.html", include)
		files = append(files, layouts...)
		r.AddFromFilesFuncs(filepath.Base(include), funcMap, files...)
	}

	return r
}

func loadTemplates1(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	basesLogin, err := filepath.Glob(templatesDir + "/bases/base_login.html")
	if err != nil {
		panic(err.Error())
	}
	includesLogin, err := filepath.Glob(templatesDir + "/login/*.html")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 template map
	for _, include := range includesLogin {
		layoutCopy := make([]string, len(basesLogin))
		copy(layoutCopy, basesLogin)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	bases, err := filepath.Glob(templatesDir + "/bases/base.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/others/*.html")
	if err != nil {
		panic(err.Error())
	}
	// 为layouts/和includes/目录生成 template map
	for _, include := range includes {
		layoutCopy := make([]string, len(bases))
		copy(layoutCopy, bases)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}
