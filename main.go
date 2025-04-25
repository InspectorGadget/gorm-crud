package main

import (
	"github.com/InspectorGadget/gorm-crud/controllers"
	"github.com/InspectorGadget/gorm-crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/posts", controllers.PostsIndex)
	r.POST("/posts", controllers.PostCreate)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run(":3000")
}
