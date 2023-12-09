package system

import (
	"Mi-Panel/api/api_v1"
	"Mi-Panel/api/api_v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitLogin(router *gin.RouterGroup) {
	loginApi := api_v1.ApiGroupApp.ApiSystem.LoginApi

	router.POST("/login", loginApi.Login)
	router.POST("/logout", middleware.LoginInterceptor, loginApi.Logout)
	router.POST("/login/sendResetPasswordVCode", loginApi.SendResetPasswordVCode)
	router.POST("/login/resetPasswordByVCode", loginApi.ResetPasswordByVCode)

}
