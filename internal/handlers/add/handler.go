package add

import (
	"encoding/json"
	"errors"
	"github.com/forinterns/srv/internal"
	"github.com/forinterns/srv/internal/model"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		internal.ErrorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e model.Employee
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			internal.ErrorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			internal.ErrorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	// написать код добавления записи в БД
	internal.ErrorResponse(w, "result", http.StatusOK)
	return
}
