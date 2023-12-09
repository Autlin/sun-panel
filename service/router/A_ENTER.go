package router

import (
	"Mi-Panel/global"
	// "Mi-Panel/router/admin"
	"Mi-Panel/router/openness"
	"Mi-Panel/router/panel"
	"Mi-Panel/router/system"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func InitRouters(addr string) error {
	router := gin.Default()
	rootRouter := router.Group("/")
	routerGroup := rootRouter.Group("api")

	// 管理员接口

	// 接口
	system.Init(routerGroup)
	// admin.Init(routerGroup)
	panel.Init(routerGroup)
	openness.Init(routerGroup)

	// WEB文件服务
	{
		webPath := "./web"
		router.StaticFile("/", webPath+"/index.html")
		router.StaticFile("/favicon.ico", webPath+"/favicon.ico")
		router.StaticFile("/favicon.svg", webPath+"/favicon.svg")
		router.Static("/assets", webPath+"/assets")
	}

	// 上传的文件
	router.Static("/uploads", "./uploads")

	global.Logger.Info("Mi-Panel is Started.  Listening and serving HTTP on ", addr)
	return router.Run(addr)
}
