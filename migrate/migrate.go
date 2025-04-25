package main

import (
	"github.com/InspectorGadget/gorm-crud/initializers"
	"github.com/InspectorGadget/gorm-crud/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Post{})
	if err != nil {
		panic(err)
	}
}
