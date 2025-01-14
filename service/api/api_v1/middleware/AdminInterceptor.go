package middleware

import (
	// "calendar-note-gin/api/v1/common/apiReturn"
	// . "calendar-note-gin/api/v1/common/base"

	"Mi-Panel/api/api_v1/common/apiReturn"
	"Mi-Panel/api/api_v1/common/base"

	"github.com/gin-gonic/gin"
)

func AdminInterceptor(c *gin.Context) {
	currentUser, _ := base.GetCurrentUserInfo(c)
	if currentUser.Role != 1 {
		apiReturn.ErrorNoAccess(c)
		c.Abort()
		return
	}
}
