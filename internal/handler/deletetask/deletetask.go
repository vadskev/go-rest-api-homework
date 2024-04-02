package deletetask

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	ErrMethodRequest = errors.New("DELETE Is not DELETE Request")
)

type TaskStore interface {
	Delete(string) error
}

func New(store TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idTask := chi.URLParam(r, "id")

		err := store.Delete(idTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
