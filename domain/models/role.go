package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"

	"gorm.io/gorm"
)

// Role model
type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"size:255"`
	IsActive    bool         `json:"is_active" gorm:"default:false"`
	Permissions []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}

func (*Role) TableName() string {
	return "roles"
}
