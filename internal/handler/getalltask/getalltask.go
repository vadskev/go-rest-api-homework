package getalltask

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/vadskev/go-rest-api-homework/internal/model"
)

var (
	ErrFailedResponse = errors.New("GET Failed to write response")
	ErrFailedMarshal  = errors.New("GET Failed to marshal json")
)

type TaskStore interface {
	GetAll() (*map[string]model.Task, error)
}

func New(store TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			log.Println(ErrFailedResponse.Error())
			return
		}
	}
}
