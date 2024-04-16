package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vadskev/go-rest-api-homework/internal/handler/deletetask"
	"github.com/vadskev/go-rest-api-homework/internal/handler/getalltask"
	"github.com/vadskev/go-rest-api-homework/internal/handler/gettask"
	"github.com/vadskev/go-rest-api-homework/internal/handler/posttask"
	"github.com/vadskev/go-rest-api-homework/internal/storage/memstorage"
)

const (
	getAllTaskPostfix = "/tasks"
	postTaskPostfix   = "/tasks"
	getTaskPostfix    = "/tasks/{id}"
	deleteTaskPostfix = "/tasks/{id}"
)

func New(store *memstorage.MemStorage) *chi.Mux {
	router := chi.NewRouter()

	router.Get(getAllTaskPostfix, getalltask.New(store))
	router.Post(postTaskPostfix, posttask.New(store))
	router.Get(getTaskPostfix, gettask.New(store))
	router.Delete(deleteTaskPostfix, deletetask.New(store))

	return router
}
