package main

import (
	"wm/sprint/db"
	"wm/sprint/middlewares"
	"wm/sprint/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect()
}

func main() {
	r := gin.Default()
	r.Use(middlewares.Logger())
	r.Use(middlewares.JwtFilter())
	routers.SetUpRouter(r)
	r.Run(":50052")
}
