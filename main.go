package main

import (
	"packform/controllers"
	"packform/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/orders", controllers.GetOrders())
	r.GET("/ordersquery", controllers.QueryStrings()) //query?name=ivan&prodcut=xxx
	r.Run()                                           // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//Close connection to database when the main function finishes
	// defer db.
}
