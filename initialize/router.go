package initialize

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"rebuildServer/global"
	"rebuildServer/router"
)

// 初始化总路由

func Routers() *gin.Engine {
	Router := gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	// Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	// Router.Static("/favicon.ico", "./dist/favicon.ico")
	// Router.Static("/static", "./dist/assets")   // dist里面的静态资源
	// Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	//Router.Use(middleware.LoadTls())	// 打开就可以是用https
	global.GVA_LOG.Info("use middleware logger")
	// 跨域，如需跨域可以打开下面的注释
	//Router.Use(middleware.Cors()) // 直接放行全部跨域请求
	//Router.Use(middleware.CorsByRules()) // 按照配置的规则放行跨域请求
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")

	// 方便统一添加路由器前缀 多服务上线使用

	// 获取路由器组示例
	systemRouter := router.RouterGroupApp.System
	exampleRouter := router.RouterGroupApp.Example
	autocodeRouter := router.RouterGroupApp.Autocode
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(context *gin.Context) {
			context.JSON(200, "ok")
		})
	}
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由，不做鉴权
		systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	}
}
