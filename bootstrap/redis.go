package bootstrap

import (
	"admin-manage-shop/pkg/config"
	"admin-manage-shop/pkg/redis"
	"fmt"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
