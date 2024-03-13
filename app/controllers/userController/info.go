package userController

import (
	"github.com/gin-gonic/gin"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetUserInfo(c *gin.Context) {
	_, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	user = models.User{
		ID:            1,
		Username:      "wejh",
		JHPassword:    "123456",
		StudentID:     "202103150901",
		PhoneNum:      "12345678901",
		Type:          3,
		LibPassword:   "123456",
		ZFPassword:    "123456",
		OauthPassword: "123456",
		YxyUid:        "123456",
		UnionID:       "123456",
		CreateTime:    time.Date(2024, 3, 8, 0, 0, 0, 0, time.UTC),
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"studentID": user.StudentID,
			"bind": gin.H{
				"zf": user.ZFPassword != "",
				//"lib":   user.LibPassword != "",
				"lib":   false,
				"yxy":   user.YxyUid != "",
				"oauth": user.OauthPassword != "",
			},
			"userType":   user.Type,
			"phoneNum":   user.PhoneNum,
			"createTime": user.CreateTime,
		},
	})

}
