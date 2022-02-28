package controllers

import (
	"net/http"

	"github.com/akash-scalent/gotodo/apimodel"
	"github.com/akash-scalent/gotodo/models"
	"github.com/akash-scalent/gotodo/utils"
	"github.com/akash-scalent/gotodo/validation"
)

func (ctrl Controller) GetTodo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("working"))
}

func (ctrl Controller) AddTodo(w http.ResponseWriter, r *http.Request) {
	requestID := utils.GetRequestID(r.Context())
	ctrl.logger.Info().Str("requestID", requestID).Msg("Addding todo started on controller layer")
	todo := &models.Todo{}
	err := validation.DecodeAndVaildate(r.Body, todo)
	if err != nil {
		ctrl.logger.Error().Err(err).Str("requestID", requestID).Msg("Error while validating request")
		utils.SendResponseWithError(w, apimodel.NewResponseError(http.StatusBadRequest, "Invalid Data"), nil)
		return
	}
	err = ctrl.todoService.AddTodo(r.Context(), todo)
	if err != nil {
		ctrl.logger.Error().Err(err).Str("requestID", requestID).Msg("Error while addding todo on controller layer")
		utils.SendResponseWithError(w, apimodel.NewResponseError(http.StatusInternalServerError, "Error while adding todo"), nil)
		return
	}
	utils.SendResponseWithData(w, http.StatusOK, "Todo created", todo)
	ctrl.logger.Info().Str("requestID", requestID).Msg("Addding todo completed on controller layer")
}
