package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/InspectorGadget/gorm-crud/initializers"
	"github.com/InspectorGadget/gorm-crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// Get data off post body
	var body models.Post
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Parameters are required",
			},
		)
		return
	}

	// Create post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)
	if result.Error != nil {
		c.Status(400)
	}

	// Return post
	c.JSON(
		http.StatusOK,
		gin.H{
			"post": post,
		},
	)
}
func PostsIndex(c *gin.Context) {
	// Get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return posts
	c.JSON(
		http.StatusOK,
		gin.H{
			"posts": posts,
		},
	)
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Parameter must be in int()",
			},
		)
		return
	}

	var post = models.Post{}
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Post not found!",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"post": post,
		},
	)
}

func PostUpdate(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Parameter must be in int()",
			},
		)
		return
	}

	// Get Data off request body
	var body models.Post
	err = c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Parameters are required",
			},
		)
		return
	}

	// Find the post we're updating
	var post = models.Post{}
	result := initializers.DB.First(&post, id)
	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Post not found!",
			},
		)
		return
	}

	// Update it
	result = initializers.DB.Model(&post).Updates(
		models.Post{
			Title: body.Title,
			Body:  body.Body,
		},
	)
	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Unable to update post!",
				"error":   result.Error.Error(),
			},
		)
		return
	}

	// Respond it
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"message": "Post updated!",
			"post":    post,
		},
	)
}

func PostDelete(c *gin.Context) {
	// Get ID from param
	id := c.Param("id")
	_, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Parameter must be in int()",
			},
		)
		return
	}

	// Find post
	var post = models.Post{}
	result := initializers.DB.Find(&post, id)
	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Unable to update post!",
				"error":   result.Error.Error(),
			},
		)
		return
	}

	// Delete post
	result = initializers.DB.Delete(&post)
	if result.Error != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"message": "Unable to delete post!",
				"error":   result.Error.Error(),
			},
		)
		return
	}

	// Return status
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"message": fmt.Sprintf("Post with ID (%v) deleted!", post.ID),
		},
	)
}
