package controllers

import (
	"github.com/akash-scalent/gotodo/service"
	"github.com/rs/zerolog"
)

type Controller struct {
	logger *zerolog.Logger
	todoService *service.TodoService
}

func NewController(logger *zerolog.Logger) *Controller {
	controllerLogger := logger.With().Str("component","controller").Logger()
	todoService := service.NewTodoService(logger)

	return &Controller{logger: &controllerLogger,todoService: todoService}
}
