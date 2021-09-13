package routers

import (
	"wm/sprint/controllers"

	"github.com/gin-gonic/gin"
)

func SetUpRouter(router *gin.Engine) {
	router.GET("sprint/ping", controllers.Pong)
	AddSpringRouter(router.Group("sprint"))
}
