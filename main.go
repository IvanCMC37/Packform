package main

import (
	"packform/models"
	controllers "packform/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/orders", controllers.GetOrders)
	// r.GET("/ordersquery", controllers.QueryStrings()) //query?name=ivan&prodcut=xxx
	r.Run(":5000") // listen and serve on 0.0.0.0:5000 (for windows "localhost:5000")
}
