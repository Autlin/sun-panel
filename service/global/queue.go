package global

import (
	"Mi-Panel/lib/queue"
	"Mi-Panel/lib/queue/queueMemory"
	"Mi-Panel/lib/queue/queueRedis"
	"Mi-Panel/structs"
)

// 缓存驱动
const (
	QUEUE_DRIVE_REDIS  = "redis"
	QUEUE_DRIVE_MEMORY = "memory"
)

// 创建一个队列
// name:缓存名称
func NewQueuer(name string) queue.Queuer {
	drive := Config.GetValueString("base", "queue_drive")
	if drive == "" {
		drive = CACHE_DRIVE_MEMORY
	}
	var queuer queue.Queuer
	Logger.Debugln("队列驱动:", drive)
	switch drive {
	case CACHE_DRIVE_MEMORY:
		queuer = queueMemory.New()
	case CACHE_DRIVE_REDIS:
		redisConfig := structs.IniConfigRedis{}
		if err := Config.GetSection("redis", &redisConfig); err != nil {
			redisConfig.Prefix = ""
		}
		queuer = queueRedis.New(RedisDb, redisConfig.Prefix+name)
	}

	return queuer
}
