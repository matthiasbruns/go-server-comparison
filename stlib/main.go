package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/matthiasbruns/go-server-comparison/stlib/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// NameModel model info
// @Description Model for put request - concatenates hello world with name
type nameModel struct {
	Name string `json:"name" form:"name" validate:"required,min=2,max=100,alphanum" swaggertype:"string" example:"John"`
} // @name NameModel

// messageResponse model info
// @Description Response that contains a message
type messageResponse struct {
	Message string `json:"message" validate:"required"`
} // @name Message

type errorResponse struct {
	Error string `json:"error" validate:"required"`
} // @name ErrorResponse

var validate = validator.New(validator.WithRequiredStructEnabled())

// @BasePath /api/v1
func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
	))

	mux.HandleFunc("GET /api/v1/helloworld", getHelloWorld)
	mux.HandleFunc("PUT /api/v1/helloworld", putHelloWorld)
	mux.HandleFunc("POST /api/v1/helloworld", postHelloWorld)

	err := http.ListenAndServe(":8082", mux)
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
func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	writeMessage(w, "Hello World")
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
func putHelloWorld(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	model := nameModel{
		Name: q.Get("name"),
	}

	// validate
	if err := validate.Struct(&model); err != nil {
		writeError(w, err)
		return
	}

	writeMessage(w, "Hello World "+model.Name)
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
func postHelloWorld(w http.ResponseWriter, r *http.Request) {
	var model nameModel
	if err := json.NewDecoder(r.Body).Decode(&model); err != nil {
		writeError(w, err)
		return
	}

	// validate
	if err := validate.Struct(&model); err != nil {
		writeError(w, err)
		return
	}

	writeMessage(w, "Hello World "+model.Name)
}

func writeMessage(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	msg := messageResponse{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		writeError(w, err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	errResp := errorResponse{
		Error: err.Error(),
	}

	if err := json.NewEncoder(w).Encode(errResp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
