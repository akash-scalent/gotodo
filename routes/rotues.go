package routes

import (
	"net/http"

	"github.com/akash-scalent/gotodo/controllers"
	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Router struct {
	Handler *http.Handler
	logger *zerolog.Logger
}

var routerLogger zerolog.Logger

func NewRouter(logger *zerolog.Logger) *Router {
	routerLogger = logger.With().Str("component", "router").Logger()
	router := mux.NewRouter().StrictSlash(true)
	router.Use(reqIDMiddleware)

	todosRouter := router.PathPrefix("/todos").Subrouter()
	todoController := controllers.NewController(logger)
	todoRouterHandler(todoController,todosRouter)

	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	handler := corsHandler(router)

	return &Router{Handler: &handler,logger: &routerLogger}
}
