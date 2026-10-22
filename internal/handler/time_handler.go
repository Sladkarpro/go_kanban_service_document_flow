package handler

import (
	"encoding/json"
	"net/http"

	"go_kanban_service/internal/apperr"
	"go_kanban_service/internal/dto"
	"go_kanban_service/internal/helper"
	"go_kanban_service/internal/service"
)

type TimeHandler struct {
	service service.TimeServiceInterface
}

func NewTimeHandler(s service.TimeServiceInterface) *TimeHandler {
	return &TimeHandler{service: s}
}

func (h *TimeHandler) SetEstimate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cardID, err := helper.IDParam(r, "id")
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		var req dto.SetEstimateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helper.WriteError(w, apperr.New(apperr.CodeInvalidJSON, "invalid json"))
			return
		}
		res, err := h.service.SetEstimate(r.Context(), cardID, req)
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		helper.WriteJSON(w, http.StatusOK, res)
	}
}

func (h *TimeHandler) CreateEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cardID, err := helper.IDParam(r, "cardId")
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		var req dto.CreateTimeEntryRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			helper.WriteError(w, apperr.New(apperr.CodeInvalidJSON, "invalid json"))
			return
		}
		res, err := h.service.CreateEntry(r.Context(), cardID, req)
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		helper.WriteJSON(w, http.StatusCreated, res)
	}
}

func (h *TimeHandler) DeleteEntry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cardID, err := helper.IDParam(r, "cardId")
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		entryID, err := helper.IDParam(r, "entryId")
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		res, err := h.service.DeleteEntry(r.Context(), cardID, entryID)
		if err != nil {
			helper.WriteError(w, err)
			return
		}
		helper.WriteJSON(w, http.StatusOK, res)
	}
}
