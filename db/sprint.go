package db

import "gorm.io/gorm"

type Sprint struct {
	gorm.Model
	Name        string `form:"name"`
	WorkspaceId uint   `form:"workspaceId"`
}
