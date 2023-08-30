package model

import (
	"gorm.io/gorm"
)

type Students struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
	Status  string `json:"status"`
}

type NewStudent struct {
	Nama    string `json:"nama"`
	Jurusan string `json:"jurusan"`
	Status  string `json:"status"`
}