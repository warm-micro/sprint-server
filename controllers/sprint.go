package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"wm/sprint/db"

	"github.com/gin-gonic/gin"
)

func CreateSprint(c *gin.Context) {
	workspaceId, err := strconv.ParseUint(c.PostForm("workspaceId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong workspaceId",
		})
		return
	}

	sprint := db.Sprint{
		Name:        c.PostForm("name"),
		WorkspaceId: uint(workspaceId),
	}
	db.DB.Create(&sprint)
	c.JSON(http.StatusOK, gin.H{
		"message": "sprint created",
		"body":    sprint.Name,
	})
}

func ListSprint(c *gin.Context) {
	workspaceId, check := c.GetQuery("workspaceId")
	fmt.Println(workspaceId)
	if !check {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "workspaceId is required",
		})
		return
	}
	var sprints []db.Sprint
	db.DB.Where("workspaceId = ?", workspaceId).Find(&sprints)

	c.JSON(http.StatusOK, gin.H{
		"message": "list sprint",
		"body":    sprints,
	})
}
