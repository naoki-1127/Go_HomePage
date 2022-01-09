package Controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

func TopPageDisplayAction(c *gin.Context) {
	log.Println("Start_TopPageDisplayAction_Methods")
	c.HTML(200, "index.html", gin.H{})
}

func RegistrationPageDisplayAction(c *gin.Context) {
	log.Println("Start_RegistrationPageDisplayAction_Methods")
	c.HTML(200, "registration.html", gin.H{})
}
