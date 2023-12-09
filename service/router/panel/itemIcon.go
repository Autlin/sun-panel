package panel

import (
	"Mi-Panel/api/api_v1"
	"Mi-Panel/api/api_v1/middleware"

	"github.com/gin-gonic/gin"
)

func InitItemIcon(router *gin.RouterGroup) {
	itemIcon := api_v1.ApiGroupApp.ApiPanel.ItemIcon
	r := router.Group("", middleware.LoginInterceptor)
	{
		r.POST("/panel/itemIcon/edit", itemIcon.Edit)
		r.POST("/panel/itemIcon/getListByGroupId", itemIcon.GetListByGroupId)
		r.POST("/panel/itemIcon/deletes", itemIcon.Deletes)
		r.POST("/panel/itemIcon/saveSort", itemIcon.SaveSort)
	}
}
