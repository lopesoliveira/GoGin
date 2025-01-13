package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lopesoliveira/GoGin/db"
	"net/http"
)

func main() {

	db.SqlServerDb()

	router := gin.Default()

	//router.GET("/", func(c *gin.Context) {})
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
