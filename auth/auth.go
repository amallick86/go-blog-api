package auth

import (
	"go-blog-api/database"
	"go-blog-api/models"
	"go-blog-api/security"
	"go-blog-api/utils/channels"

	"github.com/jinzhu/gorm"
)

func SignIn(email, password string) (string, error) {
	user := models.User{}
	var err error
	var db *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}
		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return CreateToken(user.ID)
	}
	return "", err
}
