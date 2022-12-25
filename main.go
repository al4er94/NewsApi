package main

import (
	"NewsApi/pkg/api"
	"NewsApi/pkg/db"
	"log"
	"net/http"
)

func main() {
	// инициализация зависимостей приложения
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Run mysql")

	api := api.New(db)

	log.Println("Run API")

	// запуск веб-сервера с API и приложением
	err = http.ListenAndServe(":81", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
