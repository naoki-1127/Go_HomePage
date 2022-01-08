package main

import (
	Config "Go_app/Config"
	Controllers "Go_app/Controllers"
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

	/* 	router.LoadHTMLGlob("Views/static/*.html")

	router.GET("/goapp", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	}) */
	router.Static("/Views", "./Views")

	router.StaticFS("/goapp", http.Dir("./views/static"))
	router.POST("/goapp/registration.html", Controllers.CreateUserAction)

	/* v1 := router.Group("/api/v1")
	{
		v1.POST("/register", Controllers.CreateUserAction)
	} */

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
