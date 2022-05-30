package api

import (
	"encoding/json"
	"net/http"

	"github.com/YadaYuki/omochi/app/errors"
	"github.com/YadaYuki/omochi/app/errors/code"
	"github.com/YadaYuki/omochi/app/usecase"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TermHandler struct {
	u usecase.ITermUseCase
}

func NewTermHandler(u usecase.ITermUseCase) *TermHandler {
	return &TermHandler{u: u}
}

func (h *TermHandler) FindTermByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uuidStr, ok := vars["uuid"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, errId := uuid.Parse(uuidStr)
	if errId != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	term, err := h.u.FindTermById(r.Context(), id)
	if err != nil {
		covertErrorToResponse(err, w)
		return
	}
	termBody, _ := json.Marshal(term)
	w.Write(termBody)
}

func covertErrorToResponse(err *errors.Error, w http.ResponseWriter) {
	switch err.Code {
	case code.NotExist:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
