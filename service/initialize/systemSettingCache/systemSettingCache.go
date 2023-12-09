package systemSettingCache

import (
	"Mi-Panel/global"
	"Mi-Panel/lib/cmn/systemSetting"
	"time"
)

func InItSystemSettingCache() *systemSetting.SystemSettingCache {
	return &systemSetting.SystemSettingCache{
		Cache: global.NewCache[interface{}](5*time.Hour, -1, "systemSettingCache"),
	}
}
