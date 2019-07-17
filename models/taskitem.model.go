package models

import "github.com/jinzhu/gorm"

type Taskitem struct {
	gorm.Model
	Prefix     string `json:"Prefix"`
	CurrentNum int    `json:"CurrentNum"`
	Notice     string `json:"Notice"`
	Status     int    `json:"Status"`
	TasklistID uint
}
