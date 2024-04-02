package memstorage

import (
	"errors"

	"github.com/vadskev/go-rest-api-homework/internal/model"
)

var (
	ErrTaskNotFound = errors.New("task not found")
)

type MemStorage struct {
	data map[string]model.Task
}

func InitStorage() *MemStorage {
	return &MemStorage{
		map[string]model.Task{
			"1": {
				ID:          "1",
				Description: "Сделать финальное задание темы REST API",
				Note:        "Если сегодня сделаю, то завтра будет свободный день. Ура!",
				Application: []string{
					"VS Code",
					"Terminal",
					"git",
				},
			},
			"2": {
				ID:          "2",
				Description: "Протестировать финальное задание с помощью Postmen",
				Note:        "Лучше это делать в процессе разработки, каждый раз, когда запускаешь сервер и проверяешь хендлер",
				Application: []string{
					"VS Code",
					"Terminal",
					"git",
					"Postman",
				},
			},
		},
	}
}

func (store *MemStorage) Add(id string, item model.Task) error {
	if _, ok := store.data[id]; ok {
		return model.ErrTaskExists
	}
	store.data[id] = item
	return nil
}

func (store *MemStorage) GetAll() (*map[string]model.Task, error) {
	if store == nil {
		return &map[string]model.Task{}, nil
	}
	return &store.data, nil
}

func (store *MemStorage) Get(id string) (model.Task, error) {
	if _, ok := store.data[id]; !ok {
		return model.Task{}, ErrTaskNotFound
	}
	return store.data[id], nil
}

func (store *MemStorage) Delete(id string) error {
	if _, ok := store.data[id]; !ok {
		return ErrTaskNotFound
	}
	delete(store.data, id)
	return nil
}
