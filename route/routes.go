package route

import (
	"go_bbs/controllers"
	"go_bbs/logger"
	"net/http"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	}
	r := gin.New()
	// 最重要的就是这个日志库
	r.Use(logger.GinLogger())

	//swagger 接口文档
	// http://localhost:8080/swagger/index.html 可以看到接口文档
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	// 注册

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	// 注册
	r.POST("/register", controllers.RegisterHandler)
	r.LoadHTMLFiles("./dist/index.html")
	r.Static("/css", "./dist/css")
	r.Static("/js", "./dist/js")
	r.Static("/img", "./dist/img")

	r.GET("/vue", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	// 注册pprof 相关路由
	pprof.Register(r)

	return r
}
