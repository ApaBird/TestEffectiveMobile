/*
 * Онлайн библиотека песен
 *
 * тестовое задание
 *
 * API version: 0.0.1
 * Contact: apa.birds@mail.ru
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	sw "SongServer/server"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
)

func main() {
	log.Printf("Server started")

	if err := LoadConfig("db_config.env"); err != nil {
		panic(err)
	}

	if err := LoadConfig("music_info_config.env"); err != nil {
		panic(err)
	}
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func LoadConfig(path string) error {
	if err := godotenv.Load(path); err != nil {
		return err
	}

	return nil
}
