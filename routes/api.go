package routes

import (
	"admin-manage-shop/app/http/controllers/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {
	api := r.Group("api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "Hello World",
			})
		})

		authGroup := api.Group("/auth")
		vcc := new(auth.VerifyCodeController)
		authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
	}
}
