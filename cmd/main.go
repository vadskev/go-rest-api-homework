package main

import (
	"fmt"
	"net/http"

	"github.com/vadskev/go-rest-api-homework/internal/router"
	"github.com/vadskev/go-rest-api-homework/internal/storage/memstorage"
)

func main() {
	store := memstorage.InitStorage()

	if err := http.ListenAndServe(":8080", router.New(store)); err != nil {
		fmt.Printf("Ошибка при запуске сервера: %s", err.Error())
		return
	}

}
