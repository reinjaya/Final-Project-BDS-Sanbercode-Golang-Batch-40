package controllers

import (
	"net/http"
	"rein/final/database"
	"rein/final/repository"
	"rein/final/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllComment(c *gin.Context) {
	var (
		result gin.H
	)

	comments, err := repository.GetAllComment(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comments,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCommentById(c *gin.Context) {
	var (
		result gin.H
	)

	var comment structs.Comment
	id, _ := strconv.Atoi(c.Param("id"))

	comment.ID = int64(id)
	comments, err := repository.GetCommentById(database.DbConnection, comment)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comments,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetCommentByIdArticle(c *gin.Context) {
	var (
		result gin.H
	)

	var comment structs.Comment
	id, _ := strconv.Atoi(c.Param("id"))

	comment.ID = int64(id)
	comments, err := repository.GetCommentByIdArticle(database.DbConnection, comment)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": comments,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertComment(c *gin.Context) {
	var comment structs.Comment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}

	err = repository.InsertComment(database.DbConnection, comment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully inserted",
	})
}

func UpdateComment(c *gin.Context) {
	var comment structs.Comment
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		panic(err)
	}

	comment.ID = int64(id)

	err = repository.UpdateComment(database.DbConnection, comment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully updated",
	})
}

func DeleteComment(c *gin.Context) {
	var comment structs.Comment
	id, err := strconv.Atoi(c.Param("id"))

	comment.ID = int64(id)

	err = repository.DeleteComment(database.DbConnection, comment)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully deleted",
	})
}
