package storage

import "github.com/vadskev/go-rest-api-homework/internal/model"

// Storage
type Storage interface {
	GetAll() (map[string]model.Task, error)
	Add(string, model.Task) error
	Get(key string) (model.Task, error)
	Delete(key string) error
}
