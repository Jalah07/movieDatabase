package main

import (
	"log"
	"movieProj/handlers"
	"movieProj/repo"
	"movieProj/service"
	"net/http"
	"path/filepath"
)

func main() {
	fn := "C:/Users/Jalah/Desktop/movieProj/moviedb.json"

	ext := filepath.Ext(fn)
	if ext != ".json" {
		log.Fatalln("File extension invaild")
	}

	r := repo.NewRepo(fn)

	svc := service.NewService(r)

	hdlr := handlers.NewMovieHandler(svc)

	router := handlers.ConfigureRouter(hdlr)

	svr := &http.Server {
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}

	log.Fatalln(svr.ListenAndServe())

	
}