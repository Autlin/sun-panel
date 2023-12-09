package panel

import (
	"Mi-Panel/api/api_v1"
	"Mi-Panel/api/api_v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserConfig(router *gin.RouterGroup) {
	api := api_v1.ApiGroupApp.ApiPanel.UserConfig
	r := router.Group("", middleware.LoginInterceptor)
	{
		r.POST("/panel/userConfig/set", api.Set)
		r.POST("/panel/userConfig/get", api.Get)
	}
}
