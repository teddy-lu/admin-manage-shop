package config

import "admin-manage-shop/pkg/config"

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			"name": config.Env("APP_NAME", "AdminShop"),

			"env": config.Env("APP_ENV", "production"),

			"debug": config.Env("APP_DEBUG", false),

			"port": config.Env("APP_PORT", "3000"),

			// 加密会话、JWT 加密
			"key": config.Env("APP_KEY", ""),

			"url": config.Env("APP_URL", "http://localhost:3000"),

			"timezone": config.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
