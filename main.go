// package main is copied from 
// https://github.com/PacktPublishing/Building-Distributed-Applications-in-Gin/blob/main/chapter01/main.go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	http.ListenAndServe(":8080", router)
}