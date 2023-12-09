package system

import (
	"Mi-Panel/api/api_v1"
	"Mi-Panel/api/api_v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitModuleConfigRouter(router *gin.RouterGroup) {
	api := api_v1.ApiGroupApp.ApiSystem.ModuleConfigApi
	r := router.Group("", middleware.LoginInterceptor)
	r.POST("/system/moduleConfig/getByName", api.GetByName)
	r.POST("/system/moduleConfig/save", api.Save)

}
