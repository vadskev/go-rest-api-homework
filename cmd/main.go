package main

import (
	"fmt"
	"net/http"

	"github.com/vadskev/go-rest-api-homework/internal/router"
	"github.com/vadskev/go-rest-api-homework/internal/storage/memstorage"
)

func main() {

	store := memstorage.InitStorage()

	/*
		item := model.Task{
			ID:          "1",
			Description: "sdfsdf",
			Note:        "Еyuiyui",
			Application: []string{
				"fgh",
				"uium",
				"456rth",
			},
		}
	*/

	if err := http.ListenAndServe(":8080", router.New(store)); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
