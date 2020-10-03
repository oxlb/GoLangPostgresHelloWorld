package model

import (
	"github.com/jinzhu/gorm"
)

type Student struct {
	gorm.Model
	Id     string `json:"id"`
	Name   string `json:"name"`
}
