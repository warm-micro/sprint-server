package db

import "gorm.io/gorm"

// 스프린트 정의
type Sprint struct {
	gorm.Model
	Name        string `form:"name"`
	WorkspaceId uint   `form:"workspaceId"`
	Status      string `form:"status" gorm:"default:ready"`
	Order       uint
}
