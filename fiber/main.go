package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/matthiasbruns/go-server-comparison/fiber/docs"
)

// nameModel model info
// @Description Model for put request - concatenates hello world with name
type nameModel struct {
	Name string `json:"name" form:"name" binding:"required,min=2,max=100,alphanum" swaggertype:"string" example:"John"`
} // @name NameModel

// messageResponse model info
// @Description Response that contains a message
type messageResponse struct {
	Message string `json:"message" binding:"required"`
} // @name Message

type errorResponse struct {
	Error string `json:"error" binding:"required"`
} // @name ErrorResponse

// @BasePath /api/v1
func main() {
	// Initialize a new Fiber app
	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	v1 := app.Group("/api/v1")
	{
		v1.Get("/helloworld", getHelloWorld)
		v1.Put("/helloworld", putHelloWorld)
		v1.Post("/helloworld", postHelloWorld)
	}

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}

// getHelloWorld godoc
// @Summary getHelloWorld returns hello world
// @Description getHelloWorld returns hello world
// @Tags helloworld
// @Produce json
// @Success      200  {object}   Message
// @Failure      500  {object}   errorResponse
// @Router /helloworld [get]
func getHelloWorld(c *fiber.Ctx) error {
	return c.JSON(messageResponse{Message: "hello world"})
}

// putHelloWorld godoc
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name
// @Tags helloworld
// @Accept json
// @Produce json
// @Param        query.name    query     NameModel  true  "what name" NameModel
// @Success      200  {object}   Message
// @Failure      400  {object}   errorResponse
// @Failure      500  {object}   errorResponse
// @Router /helloworld [put]
func putHelloWorld(c *fiber.Ctx) error {
	var query nameModel
	if err := c.QueryParser(&query); err != nil {
		return c.Status(400).JSON(errorResponse{Error: err.Error()})
	}
	return c.JSON(messageResponse{Message: fmt.Sprintf("hello world %s", query.Name)})
}

// postHelloWorld godoc
// @Summary getHelloWorld returns hello world with name
// @Description getHelloWorld returns hello world with name
// @Tags helloworld
// @Accept json
// @Produce json
// @Param        query    body     NameModel  true  "what name" NameModel
// @Success      200  {object}   Message
// @Failure      400  {object}   errorResponse
// @Failure      500  {object}   errorResponse
// @Router /helloworld [post]
func postHelloWorld(c *fiber.Ctx) error {
	var query nameModel
	if err := c.BodyParser(&query); err != nil {
		return c.Status(400).JSON(errorResponse{Error: err.Error()})
	}
	return c.JSON(messageResponse{Message: fmt.Sprintf("hello world %s", query.Name)})
}
