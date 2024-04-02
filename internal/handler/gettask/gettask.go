package gettask

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vadskev/go-rest-api-homework/internal/model"
)

var (
	ErrMethodRequest = errors.New("GET Is not GET Request")
	ErrTaskNotFound  = errors.New("GET task not found")
)

type TaskStore interface {
	Get(string) (model.Task, error)
}

func New(store TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, ErrMethodRequest.Error(), http.StatusBadRequest)
			return
		}

		idTask := chi.URLParam(r, "id")

		_, err := store.Get(idTask)
		if err != nil {
			http.Error(w, ErrTaskNotFound.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
	}
}
