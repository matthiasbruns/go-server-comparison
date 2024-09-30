package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type putModel struct {
	Name string `json:"name" binding:"required,min=2,max=100,alphanum" swaggertype:"string" example:"John"`
}

func main() {
	r := gin.Default()

	r.GET("/helloworld", getHelloWorld)
	r.PUT("/helloworld", putHelloWorld)

	err := r.Run()

	if err != nil {
		panic(err)
	}
}

func getHelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

func putHelloWorld(c *gin.Context) {
	var model putModel
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("hello world %s", model.Name),
	})
}
