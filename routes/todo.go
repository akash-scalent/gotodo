package routes

import (
	"net/http"

	"github.com/akash-scalent/gotodo/controllers"
	"github.com/gorilla/mux"
)



func  todoRouterHandler(controller *controllers.Controller,router *mux.Router)  {
	router.HandleFunc("/",controller.GetTodo).Methods(http.MethodGet)
	router.HandleFunc("/",controller.AddTodo).Methods(http.MethodPost)
}