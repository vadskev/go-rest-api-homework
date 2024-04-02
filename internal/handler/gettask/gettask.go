package gettask

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vadskev/go-rest-api-homework/internal/model"
)

var (
	ErrMethodRequest  = errors.New("GET Is not GET Request")
	ErrFailedMarshal  = errors.New("GET Failed to marshal json")
	ErrFailedResponse = errors.New("GET Failed to write response")
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

		data, err := store.Get(idTask)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, ErrFailedMarshal.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, ErrFailedResponse.Error(), http.StatusInternalServerError)
			return
		}
	}
}
