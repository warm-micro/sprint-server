package main

import (
	"wm/sprint/db"
	"wm/sprint/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Connect()
}

func main() {
	r := gin.Default()
	routers.SetUpRouter(r)
	r.Run(":50052")
}
