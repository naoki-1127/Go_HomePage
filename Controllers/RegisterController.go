package Controller

import (
	db "Go_app/Models/db"
	entity "Go_app/Models/entity"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createUser(username string, password string) []error {
	db := db.GormConnect()
	if err := db.Create(&entity.User{Email: username, Password: password, Created_at: time.Now(), Updated_at: time.Now()}).Error; err != nil {
		log.Println("error")
	}
	return nil

}

func CreateUserAction(c *gin.Context) {
	log.Println("ContactController_test")
	var form entity.User
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "/goapp/registration.html", gin.H{"err": err})
		c.Abort()
	} else {
		username := c.PostForm("username")
		log.Println("ユーザーメール:" + username)
		password := c.PostForm("password")
		log.Println("ユーザ-パスワード:" + password)
		// 登録ユーザーが重複していた場合にはじく処理
		if err := createUser(username, password); err != nil {
			c.HTML(http.StatusBadRequest, "/goapp/registration.html", gin.H{"err": err})
		}
		c.Redirect(302, "/goapp/registration.html")
	}
	log.Println("リダイレクトするで")

}
