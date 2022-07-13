package bootstrap

import (
	"admin-manage-shop/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func SetupRoute(router *gin.Engine) {
	registerGlobalMiddleware(router)

	routes.RegisterAPIRoutes(router)

	setup404Handler(router)
}

func setup404Handler(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "no route 404")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"error_code": 404,
				"error_msg":  "no route 404",
			})
		}
	})
}

func registerGlobalMiddleware(router *gin.Engine) {
	router.Use(gin.Logger(), gin.Recovery())
}
