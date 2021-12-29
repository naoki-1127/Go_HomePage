package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	serve()
}

func serve() {
	router := gin.Default()

	router.Static("/Views", "./Views")

	router.StaticFS("/goapp/login", http.Dir("./views/static"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
