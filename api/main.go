package main

import (
	"github.com/gin-gonic/gin"
)
func main() {
	
	inventoryAPI := InitInventoryAPI()

	r := gin.Default()
	r.GET("/", inventoryAPI.Ping)
	r.POST("/add", inventoryAPI.Add)
	// r.GET("/qh/id/:id", quoteheaderAPI.Find)
	r.Run(":8000") // listen and serve on 0.0.0.0:8000
}
