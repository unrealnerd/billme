package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	
	inventoryAPI := InitInventoryAPI()

	r := gin.Default()
	r.GET("/", inventoryAPI.Ping)
	r.GET("/get/:id", inventoryAPI.Find)
	r.POST("/add", inventoryAPI.Add)
	
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
