package sessionServices

import (
	"errors"
	"time"
	"wejh-go/app/models"


	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetUserSession(c *gin.Context, user *models.User) error {
	webSession := sessions.Default(c)
	webSession.Options(sessions.Options{MaxAge: 3600 * 24 * 7, Path: "/api"})
	webSession.Set("id", user.ID)
	return webSession.Save()
}

func GetUserSession(c *gin.Context) (*models.User, error) {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	if id == nil {
		return nil, errors.New("")
	}
	var user *models.User
	if id.(int) == 1 {
		user = &models.User{
			ID:           1,
			Username:     "wejh",
			JHPassword:     "123456",
			StudentID:    "202103150901",
			PhoneNum:     "12345678901",
			Type: 3,
			LibPassword: "123456",
			ZFPassword: "123456",
			OauthPassword: "123456",
			YxyUid: "123456",
			UnionID: "123456",
			CreateTime: time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),
	
		}
	}else {
		return nil, errors.New("")
	}
	return user, nil
}

func UpdateUserSession(c *gin.Context) (*models.User, error) {
	user, err := GetUserSession(c)
	if err != nil {
		return nil, err
	}
	err = SetUserSession(c, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CheckUserSession(c *gin.Context) bool {
	webSession := sessions.Default(c)
	id := webSession.Get("id")
	return id != nil
}

func ClearUserSession(c *gin.Context) {
	webSession := sessions.Default(c)
	webSession.Delete("id")
	webSession.Save()
}
