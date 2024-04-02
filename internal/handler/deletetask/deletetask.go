package deletetask

import (
	"net/http"

	"github.com/vadskev/go-rest-api-homework/internal/storage/memstorage"
)

func New(store *memstorage.MemStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
