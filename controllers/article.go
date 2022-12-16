package controllers

import (
	"net/http"
	"rein/final/database"
	"rein/final/repository"
	"rein/final/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllArticle(c *gin.Context) {
	var (
		result gin.H
	)

	articles, err := repository.GetAllArticle(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": articles,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetArticleById(c *gin.Context) {
	var (
		result gin.H
	)

	var article structs.Article
	id, _ := strconv.Atoi(c.Param("id"))

	article.ID = int64(id)
	articles, err := repository.GetArticleById(database.DbConnection, article)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": articles,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertArticle(c *gin.Context) {
	var article structs.Article

	err := c.ShouldBindJSON(&article)
	if err != nil {
		panic(err)
	}

	err = repository.InsertArticle(database.DbConnection, article)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully inserted",
	})
}

func UpdateArticle(c *gin.Context) {
	var article structs.Article
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&article)
	if err != nil {
		panic(err)
	}

	article.ID = int64(id)

	err = repository.UpdateArticle(database.DbConnection, article)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully updated",
	})
}

func DeleteArticle(c *gin.Context) {
	var article structs.Article
	id, err := strconv.Atoi(c.Param("id"))

	article.ID = int64(id)

	err = repository.DeleteArticle(database.DbConnection, article)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully deleted",
	})
}
