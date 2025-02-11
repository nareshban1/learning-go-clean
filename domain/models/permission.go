package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"

	"gorm.io/gorm"
)

// Permission model
type Permission struct {
	gorm.Model
	Name     string `json:"name" gorm:"size:255"`
	IsActive bool   `json:"is_active" gorm:"default:false"`
	Roles    []Role `json:"-" gorm:"many2many:role_permissions"`
}

func (*Permission) TableName() string {
	return "permissions"
}
