package router

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jacobmeredith/boilerplate-go-api/internal/service"
)

func NewExampleRouter(exampleService *service.ExampleService) *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("POST /", handleCreateExample(exampleService))
	router.HandleFunc("GET /{id}", handleGetExample(exampleService))

	return router
}

func handleGetExample(exampleStore *service.ExampleService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		example, err := exampleStore.GetExample(uint(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.Write([]byte(example.Message))
	}
}

func handleCreateExample(exampleStore *service.ExampleService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		type requestBody struct {
			Message string `json:"message"`
		}

		var body requestBody

		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = exampleStore.CreateExample(body.Message)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
