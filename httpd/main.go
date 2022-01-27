package main

import (
	"packform/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/order", handler.OrderGet())

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
