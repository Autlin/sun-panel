package system

import (
	"Mi-Panel/api/api_v1"

	"github.com/gin-gonic/gin"
)

func InitNoticeRouter(router *gin.RouterGroup) {
	api := api_v1.ApiGroupApp.ApiSystem.NoticeApi

	router.POST("/notice/getListByDisplayType", api.GetListByDisplayType)
}
