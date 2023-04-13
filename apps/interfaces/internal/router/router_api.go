package router

import (
	"lark/apps/interfaces/dig"
	"lark/apps/interfaces/internal/ctrl/ctrl_auth"

	"github.com/gin-gonic/gin"
)

func Register(e *gin.Engine) {
	openGroup := e.Group("open")
	registerOpenRouter(openGroup)

	api := e.Group("api")
	registerApiRouter(api)
}

func registerOpenRouter(rg *gin.RouterGroup) {
	authGroup := rg.Group("auth")

	var (
		ctrl *ctrl_auth.AuthCtrl
	)

	dig.Invoke(func(c *ctrl_auth.AuthCtrl) {
		ctrl = c
	})

	authGroup.POST("sign_in", ctrl.SignIn)
	authGroup.POST("sign_up", ctrl.SignUp)
}

func registerApiRouter(rg *gin.RouterGroup) {
	authGroup := rg.Group("api")

	var (
		ctrl *ctrl_auth.AuthCtrl
	)

	dig.Invoke(func(c *ctrl_auth.AuthCtrl) {
		ctrl = c
	})

	authGroup.POST("sign_in", ctrl.SignOut)
}
