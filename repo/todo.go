package repo

import (
	"context"

	"github.com/akash-scalent/gotodo/models"
	"github.com/akash-scalent/gotodo/utils"
)

func ( rep *TodoRepo) AddTodo(ctx context.Context,todo *models.Todo) (error ){
	requestID := utils.GetRequestID(ctx)
	rep.logger.Info().Str("requestID",requestID,).Msg("Adding todo started on repo layer")
result:=rep.db.Create(todo)
if result.Error != nil {
	rep.logger.Error().Err(result.Error).Str("requestID",requestID).Msg("Error while adding todo")
	return result.Error
}
	return nil
}

func (rep *TodoRepo) GetAllTodos(ctx context.Context) []*models.Todo {
	return []*models.Todo{}
}
func (rep *TodoRepo) GetTodoByID(ctx context.Context,todoID int) *models.Todo {
	return &models.Todo{}
}