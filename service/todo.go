package service

import (
	"context"

	"github.com/akash-scalent/gotodo/models"
	"github.com/akash-scalent/gotodo/utils"
)


func ( todoService *TodoService) AddTodo(ctx context.Context,todo *models.Todo) (error) {
	requestId := utils.GetRequestID(ctx)
	todoService.logger.Info().Str("requestId",requestId).Msg("Adding todo started on repo layer")
	err := todoService.dbRepo.AddTodo(ctx,todo)
	if err != nil {
	todoService.logger.Error().Err(err).Str("requestId",requestId).Msg("Error while adding todo on repo layer")
		return err
	}
	return nil
}