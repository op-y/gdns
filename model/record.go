package model

import (
	"gorm.io/gorm"
)

type A struct {
	gorm.Model
	Domain     string `gorm:"column:domain;index"`
	RecordType string `gorm:"column:record_type"`
	Record     string `gorm:"column:record"`
}

func (a *A) TableName() string {
	return "record_a"
}
