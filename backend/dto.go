package main

import "gorm.io/gorm"

type DTO struct {
	UID            string `json:"UID"`
	FName          string `json:"FName"`
	RollNo         string `json:"RollNo"`
	CountOfAddress string `json:"countadd"`
	gorm.Model
}
