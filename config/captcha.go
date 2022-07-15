package config

import "admin-manage-shop/pkg/config"

func init() {
	config.Add("captcha", func() map[string]interface{} {
		return map[string]interface{}{
			"height": 60,

			"width": 240,

			"length": 6,
			// 数字的最大倾斜角度
			"maxskew": 0.7,
			// 图片背景里的混淆点数量
			"dotcount": 80,
			// 分钟
			"expire_time": 15,

			"debug_expire_time": 10080,

			"testing_key": "captcha_skip_test",
		}
	})
}
