package main

import (
	Config "Go_app/Config"
	Utils "Go_app/Utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	Utils.LoggingSetting(Config.Config.LogFile)
	log.Println("Sample")
	serve()
}

func serve() {
	log.Println("サーバーを起動しました")
	router := gin.Default()
	router.Static("/Views", "./Views")

	router.StaticFS("/goapp/login", http.Dir("./views/static"))

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
