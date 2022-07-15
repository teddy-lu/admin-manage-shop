package auth

import (
	"admin-manage-shop/app/http/controllers"
	"admin-manage-shop/pkg/captcha"
	"admin-manage-shop/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VerifyCodeController struct {
	controllers.BaseApiController
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
