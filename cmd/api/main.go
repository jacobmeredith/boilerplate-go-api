package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/jacobmeredith/boilerplate-go-api/internal/database"
	"github.com/jacobmeredith/boilerplate-go-api/internal/middleware"
	"github.com/jacobmeredith/boilerplate-go-api/internal/router"
	"github.com/jacobmeredith/boilerplate-go-api/internal/service"
	"github.com/jacobmeredith/boilerplate-go-api/internal/store"
)

func main() {
	port := flag.String("port", "8080", "The port you want the api server to run on")
	flag.Parse()

	db, err := database.GetDatabase("local.db")
	if err != nil {
		panic("unable to connect to database")
	}

	// Stores
	exampleStore := store.NewExampleStore(db)

	// Services
	exampleService := service.NewExampleService(exampleStore)

	// Routers
	exampleRouter := router.NewExampleRouter(exampleService)

	// Main app router
	router := http.NewServeMux()
	router.Handle("/example/", http.StripPrefix("/example", exampleRouter))

	// Middleware that is applied app wide
	middlewareStack := middleware.CreateMiddlewareStack(middleware.Logging)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", *port),
		Handler: middlewareStack(router),
	}

	server.ListenAndServe()
}
