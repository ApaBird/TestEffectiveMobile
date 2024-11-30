package main

import (
	"SongServer/postgres"
	sw "SongServer/server"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Онлайн библиотека песен
//	@version		0.0.1
//	@description	тестовое задание

//	@host		localhost:8080
//	@BasePath	/

func main() {
	log.Printf("Server started")

	if err := LoadConfig("db_config.env"); err != nil {
		panic(err)
	}

	// По условию не ясно как должна проходить миграция, вроде логично что через какой нибудь Docker Compose, но о docker не слова в задче
	// Поэтому сделал миграцию внутри сервиса
	postgres.Init()

	if err := LoadConfig("music_info_config.env"); err != nil {
		panic(err)
	}
	router := sw.NewRouter()
	router.Handle("/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func LoadConfig(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}
