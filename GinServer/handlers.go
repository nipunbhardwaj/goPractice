package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"encoding/json"
	"log"
	"io/ioutil"
	"math/rand"
)

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Intro string `json:"intro"`
	Content string `json:"content"`
	Author_id string  `json:"author_id"`
}

var articles = []Article{}

func InitializeFile () {
	fileData, err := ioutil.ReadFile("data.json")
	
	if err != nil {
		log.Fatal(err)
	}
	log.Println("File Opened!")

	marshallErr := json.Unmarshal(fileData, &articles)
	if marshallErr != nil {
		log.Fatal(err)
	}
}

func GetAllArticles (c *gin.Context) {
	c.JSON(200, gin.H{
		"data": articles,
	})
	return
}

func GetArticle (c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleId"))

	if err != nil {
		log.Fatal(err)
	}
	
	for _, article := range articles {
		if article.Id == articleId {
			c.JSON(200, gin.H{
				"message": "Article Found!" ,
				"data": article,
			})
			return
		}
	}
	c.JSON(200, gin.H{
		"message": "Article not found!",
	})
	return
}

func CreateArticle (c *gin.Context) {
	var newArticle Article
	if c.ShouldBind(&newArticle) == nil {
		newArticle.Id = rand.Int()
		articles = append(articles, newArticle)

		//save article to disk
		saveArticleToDisk()

		c.JSON(200, gin.H{
			"message": "Article Saved!",
		})
		return
	}
	c.JSON(400, gin.H{
		"message": "Article Data Not Found!",
	})
	return
}

func UpdateArticle (c *gin.Context) {
	var updateArticle Article
	articleId, err := strconv.Atoi(c.Param("articleId"))

	if err != nil {
		log.Fatal(err)
	}

	if c.ShouldBind(&updateArticle) == nil {
		updateArticle.Id = articleId

		for index, article := range articles {
		if article.Id == articleId {
				articles[index] = updateArticle
				break
			}
		}
	
		//save article to disk
		saveArticleToDisk()

		c.JSON(200, gin.H{
			"message": "Article Updated!",
		})
		return
	}
	c.JSON(400, gin.H{
		"message": "Article Data Not Found!",
	})
	return
}

func saveArticleToDisk () {
	done := make(chan bool)
	go func () {
		b, err := json.Marshal(articles)
	    if err != nil {
	        log.Fatal(err)
	        return
	    }
		ioutil.WriteFile("data.json", b, 0644)
		done <- true
	}()
	<- done
}