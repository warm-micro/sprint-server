package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"wm/sprint/config"
	"wm/sprint/db"

	"github.com/gin-gonic/gin"
)

func CreateSprint(c *gin.Context) {
	var sprint db.Sprint
	c.Bind(&sprint)
	if sprint.WorkspaceId == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "workspace id is required",
			"body":    nil,
		})
		return
	}
	req, _ := http.NewRequest("GET", config.WORKSPACE_SERVICE+"/workspace/exists?workspaceId="+fmt.Sprint(sprint.WorkspaceId), nil)
	req.Header.Add("Authorization", c.GetHeader("Authorization"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "check workspace failed",
		})
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "workspace check response read error",
		})
		return
	}
	var checkResponse map[string]interface{}
	json.Unmarshal([]byte(data), &checkResponse)
	if checkResponse["message"] == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong workspace id",
			"body":    nil,
		})
		return
	}
	db.DB.Create(&sprint)
	c.JSON(http.StatusOK, gin.H{
		"message": "sprint created",
		"body":    sprint,
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

func CheckSprint(c *gin.Context) {
	sprintId, ok := c.GetQuery("sprintId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "sprint id is required",
			"body":    nil,
		})
		return
	}
	var sprint db.Sprint
	var check bool
	if err := db.DB.Where("ID = ?", sprintId).First(&sprint).Error; err != nil {
		check = false
	} else {
		check = true
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "check sprint id",
		"body":    check,
	})
}

func DeleteSprint(c *gin.Context) {
	sprintId := c.Param("sprintId")
	var sprint db.Sprint
	if err := db.DB.Where("ID = ?", sprintId).First(&sprint).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "wrong sprint Id",
			"body":    sprintId,
		})
		return
	}
	db.DB.Delete(&sprint)
	c.JSON(http.StatusOK, gin.H{
		"message": "sprint deleted",
		"body":    sprint,
	})
}
