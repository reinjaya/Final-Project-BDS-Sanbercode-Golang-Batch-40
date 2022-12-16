package controllers

import (
	"net/http"
	"rein/final/database"
	"rein/final/repository"
	"rein/final/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllLike(c *gin.Context) {
	var (
		result gin.H
	)

	likes, err := repository.GetAllLike(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": likes,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetLikeById(c *gin.Context) {
	var (
		result gin.H
	)

	var like structs.Like
	id, _ := strconv.Atoi(c.Param("id"))

	like.ID = int64(id)
	likes, err := repository.GetLikeByIdArticle(database.DbConnection, like)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": likes,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetLikeByIdArticle(c *gin.Context) {
	var (
		result gin.H
	)

	var like structs.Like
	id, _ := strconv.Atoi(c.Param("id"))

	like.ID = int64(id)
	likes, err := repository.GetLikeByIdArticle(database.DbConnection, like)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": likes,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertLike(c *gin.Context) {
	var like structs.Like

	err := c.ShouldBindJSON(&like)
	if err != nil {
		panic(err)
	}

	err = repository.InsertLike(database.DbConnection, like)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully inserted",
	})
}

func UpdateLike(c *gin.Context) {
	var like structs.Like
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&like)
	if err != nil {
		panic(err)
	}

	like.ID = int64(id)

	err = repository.UpdateLike(database.DbConnection, like)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully updated",
	})
}

func DeleteLike(c *gin.Context) {
	var like structs.Like
	id, err := strconv.Atoi(c.Param("id"))

	like.ID = int64(id)

	err = repository.DeleteLike(database.DbConnection, like)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully deleted",
	})
}
