package Controller

import (
	db "Go_app/Config"
	entity "Go_app/Models/entity"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Result string `json:"result"`
}

func createUser(name string, email string, password string) error {
	log.Println("Start_CreateUser_Methods")
	db := db.GormConnect()
	// 検索でレコードが見つからなかった場合
	if err := db.Create(&entity.User{Name: name, Email: email, Password: password, Created_at: time.Now(), Updated_at: time.Now()}).Error; err != nil {
		log.Println("ユーザの登録ができませんでした")
		log.Println(err)
		return err
	}
	return nil
}

func CreateUserAction(c *gin.Context) {
	log.Println("Start_CreateUserAction_Methods")
	var form entity.User
	var u User
	if err := c.Bind(&form); err != nil {
		c.HTML(http.StatusBadRequest, "/goapp/registration.html", gin.H{"err": err})
		c.Abort()
	} else {
		name := c.PostForm("name")
		log.Println("ユーザー名:" + name)
		email := c.PostForm("email")
		log.Println("ユーザーメール:" + email)
		password := c.PostForm("password")
		log.Println("ユーザ-パスワード:" + password)

		u = User{
			Name:   name,
			Email:  email,
			Result: "",
		}

		db := db.GormConnect()
		if err := db.Where("email = ?", email).First(&entity.User{}).Error; err != nil {
			// 検索でレコードが見つからなかった場合
			if err := createUser(name, email, password); err != nil {
				u.Name = name
				u.Email = email
				u.Result = "登録失敗"
				c.JSON(200, u)
			} else {
				u.Name = name
				u.Email = email
				u.Result = "登録成功"
				c.JSON(200, u)
			}
		} else {
			log.Println("ユーザの登録が既にあります")
			u.Name = name
			u.Email = email
			u.Result = "ユーザの登録が既にあります"
			c.JSON(200, u)
		}

		c.Get("/goapp/registration")
	}
	log.Println("リダイレクトするで")
}
