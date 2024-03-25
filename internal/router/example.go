package router

import (
	"net/http"
)

func NewExampleRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handleGetExample)

	return router
}

func handleGetExample(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get example"))
}
