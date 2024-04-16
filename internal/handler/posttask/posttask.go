package posttask

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/vadskev/go-rest-api-homework/internal/model"
)

var (
	ErrMethodRequest       = errors.New("POST Is not POST Request")
	ErrFailedDecodeJson    = errors.New("POST failed to decode json")
	ErrFailedDecodeRequest = errors.New("POST failed to decode request body")
)

type TaskStore interface {
	Add(string, model.Task) error
}

func New(store TaskStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task model.Task
		var buf bytes.Buffer

		_, err := buf.ReadFrom(r.Body)
		if err != nil {
			http.Error(w, ErrFailedDecodeRequest.Error(), http.StatusBadRequest)
			return
		}

		if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
			http.Error(w, ErrFailedDecodeJson.Error(), http.StatusBadRequest)
			return
		}

		err = store.Add(task.ID, task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
