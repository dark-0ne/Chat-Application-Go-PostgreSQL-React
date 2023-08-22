package main

import (
	"log"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/services"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}
	db.DB()

	router := gin.Default()

	router.GET("/users", services.GetUsers)
	router.GET("/user/:id", services.GetUser)
	router.POST("/user", services.PostUser)
	router.PUT("/user/:id", services.UpdateUser)
	router.DELETE("/user/:id", services.DeleteUser)

	log.Fatal(router.Run(":3000"))
}
