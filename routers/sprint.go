package routers

import (
	"wm/sprint/controllers"

	"github.com/gin-gonic/gin"
)

func AddSpringRouter(router *gin.RouterGroup) {
	router.POST("/", controllers.CreateSprint)
	router.GET("", controllers.ListSprint)
	router.GET("/check", controllers.CheckSprint)
}
