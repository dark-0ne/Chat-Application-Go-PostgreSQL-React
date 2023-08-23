package main

import (
	"log"
	"net/http"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/routers"
)

func main() {

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	routersInit := routers.InitRouter()

	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           ":3000",
		Handler:        routersInit,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening on localhost:3000")

	server.ListenAndServe()
}
