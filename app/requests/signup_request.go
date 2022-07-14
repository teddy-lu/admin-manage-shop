package requests

import (
	"admin-manage-shop/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginRequest struct {
	Account  string `json:"account,omitempty" valid:"account"`
	Password string `json:"password,omitempty" valid:"password"`

	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
}

func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"account":        []string{"required"},
		"password":       []string{"required", "min:6"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"account": []string{
			"required: 登陆帐号必须",
		},
		"password": []string{
			"required: 密码必须",
			"min: 密码长度最少大于6",
		},
		"captcha_id": []string{
			"required:图片验证码id必须",
		},
		"captcha_answer": []string{
			"required:验证码必须",
			"digits:验证码长度不正确",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*LoginRequest)
	errs = validators.ValidateVerifyCode(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}
