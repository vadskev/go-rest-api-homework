package model

import "errors"

var (
	ErrNotFound   = errors.New("not found")
	ErrTaskExists = errors.New("task already exists")
)

type Task struct {
	ID          string   `json:"id"`          // ID задачи
	Description string   `json:"description"` // Заголовок
	Note        string   `json:"note"`        // Описание задачи
	Application []string `json:"application"` // Приложения, которыми будете пользоваться
}
