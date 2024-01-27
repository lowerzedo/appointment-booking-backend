package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	FullName string
	Role     string
	Status   bool
}
