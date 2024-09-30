package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matthiasbruns/go-server-comparison/gin/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// putModel model info
// @Description Model for put request - concatenates hello world with name
type putModel struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=100,alphanum" swaggertype:"string" example:"John"`
} // @name PutModel

// messageResponse model info
// @Description Response that contains a message
type messageResponse struct {
	Message string `json:"message" binding:"required"`
} // @name Message

type errorResponse struct {
	Error string `json:"error" binding:"required"`
} // @name ErrorResponse

func main() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/helloworld", getHelloWorld)
		v1.PUT("/helloworld", putHelloWorld)
		v1.POST("/helloworld", postHelloWorld)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := r.Run()

	if err != nil {
		panic(err)
	}
}

// getHelloWorld godoc
// @Summary getHelloWorld returns hello world
// @Description getHelloWorld returns hello world
// @Tags helloworld
// @Produce json
// @Success      200  {object}   Message
// @Failure      500  {object}   errorResponse
// @Router /helloworld [get]
func getHelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}

// putHelloWorld godoc
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name
// @Tags helloworld
// @Accept json
// @Produce json
// @Param        query.name    query     PutModel  true  "what name" PutModel
// @Success      200  {object}   Message
// @Failure      400  {object}   errorResponse
// @Failure      500  {object}   errorResponse
// @Router /helloworld [put]
func putHelloWorld(c *gin.Context) {
	var model putModel
	if err := c.ShouldBindQuery(&model); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("hello world %s", model.Name),
	})
}

// postHelloWorld godoc
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name
// @Tags helloworld
// @Accept json
// @Produce json
// @Param        query    body     PutModel  true  "what name" PutModel
// @Success      200  {object}   Message
// @Failure      400  {object}   errorResponse
// @Failure      500  {object}   errorResponse
// @Router /helloworld [post]
func postHelloWorld(c *gin.Context) {
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
