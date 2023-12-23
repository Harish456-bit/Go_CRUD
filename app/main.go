package main

import (
	controllers "Models/Controllers"
	models "Models/Models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDB()

	router.POST("/Posts", controllers.CreatePost)
	router.GET("/Posts", controllers.ReadPost)
	router.GET("/Posts/:tdn", controllers.FindPost)
	router.PUT("/Posts/:tdn", controllers.UpdatePost)
	router.DELETE("/Posts/:tdn", controllers.DeletePost)

	router.Run("localhost:8080")
}
