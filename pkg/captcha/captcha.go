package captcha

import (
	"admin-manage-shop/pkg/app"
	"admin-manage-shop/pkg/config"
	"admin-manage-shop/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {
	once.Do(func() {
		internalCaptcha = &Captcha{}

		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix:   config.GetString("app.name") + ":captcha",
		}

		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.maxskew"),
			config.GetInt("captcha.dotcount"),
		)

		//driver := base64Captcha.NewDriverMath(
		//	config.GetInt("captcha.height"),
		//	config.GetInt("captcha.width"),
		//	15,
		//	0,
		//	nil,
		//	nil,
		//	[]string{"3Dumb.ttf"},
		//)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)

	})
	return internalCaptcha
}

func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	if !app.IsProduction() && id == config.GetString("captcha.testing_key") {
		return true
	}
	return c.Base64Captcha.Verify(id, answer, false)
}
