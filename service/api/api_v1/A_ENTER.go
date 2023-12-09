package api_v1

import (
	"Mi-Panel/api/api_v1/openness"
	"Mi-Panel/api/api_v1/panel"
	"Mi-Panel/api/api_v1/system"
)

type ApiGroup struct {
	ApiSystem system.ApiSystem // 系统功能api
	ApiOpen   openness.ApiPpenness
	ApiPanel  panel.ApiPanel
}

var ApiGroupApp = new(ApiGroup)
