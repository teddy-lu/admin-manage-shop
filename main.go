package main

import (
	"admin-manage-shop/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	bootstrap.SetupRoute(r)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println(err.Error())
	}
}
