package main

import (
	"github.com/gin-gonic/gin"
	//"fmt"
)

func pingHandle (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}		

func main() {
	InitializeFile()

	r := gin.Default()

	r.GET("/ping", pingHandle)

	v1API := r.Group("/api/v1")
	
	v1API.GET("/articles", GetAllArticles)
	v1API.GET("/articles/:articleId", GetArticle)
	v1API.POST("/articles", CreateArticle)
	v1API.POST("/articles/:articleId/update", UpdateArticle)

	r.Run(":5000") 
}