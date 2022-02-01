package main

import (
	"packform/models"
	controllers "packform/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	models.ConnectDataBase()

	r.GET("/orders", controllers.GetOrders)
	// r.GET("/ordersquery", controllers.QueryStrings()) //query?name=ivan&prodcut=xxx
	r.Run(":5000") // listen and serve on 0.0.0.0:5000 (for windows "localhost:5000")
}
