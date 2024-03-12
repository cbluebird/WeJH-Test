package router

import (
	"wejh-go/app/controllers/userController"

	"github.com/gin-gonic/gin"
)

func userRouterInit(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/login/wechat", userController.WeChatLogin)
		user.POST("/login", userController.AuthByPassword)
		user.POST("/login/session", userController.AuthBySession)
	}
}
