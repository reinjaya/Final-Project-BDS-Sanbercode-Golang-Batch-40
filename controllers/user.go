package controllers

import (
	"net/http"
	"rein/final/database"
	"rein/final/repository"
	"rein/final/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var (
		result gin.H
	)

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetUserById(c *gin.Context) {
	var (
		result gin.H
	)

	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = int64(id)
	users, err := repository.GetUserById(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetUserByUsername(c *gin.Context) {
	var (
		result gin.H
	)

	var user structs.User
	user.UserName = c.Param("username")

	users, err := repository.GetUserByUsername(database.DbConnection, user)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully inserted",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}

	user.ID = int64(id)

	err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully updated",
	})
}

func DeleteUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("id"))

	user.ID = int64(id)

	err = repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Successfully deleted",
	})
}
