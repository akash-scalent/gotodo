package service

import (
	"github.com/akash-scalent/gotodo/repo"
	"github.com/rs/zerolog"
)


type TodoService struct {
	logger *zerolog.Logger
	dbRepo repo.TodoRepository
}

func NewTodoService(logger *zerolog.Logger) *TodoService {
	serviceLogger := logger.With().Str("component","service").Logger()
	dbRepo := repo.NewTodoRepo(logger)
	return &TodoService{logger: &serviceLogger,dbRepo: dbRepo}
	
}