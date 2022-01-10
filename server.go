package main

import (
	Config "Go_app/Config"
	Controllers "Go_app/Controllers"
	db "Go_app/Models/db"
	Utils "Go_app/Utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	Utils.LoggingSetting(Config.Config.LogFile)
	log.Println("Sample")
	db.Open()
	serve()
}

func serve() {
	log.Println("サーバーを起動しました")
	router := gin.Default()

	router.LoadHTMLGlob("Views/static/*.html")
	router.Static("/Views", "./Views")

	router.GET("/goapp", Controllers.TopPageDisplayAction)
	router.GET("/goapp/registration", Controllers.RegistrationPageDisplayAction)
	router.POST("/goapp/registration", Controllers.CreateUserAction)

	/* v1 := router.Group("/api/v1")
	{
		v1.POST("/register", Controllers.CreateUserAction)
	} */

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
