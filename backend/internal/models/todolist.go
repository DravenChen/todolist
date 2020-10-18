package models

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	TodoNotStartStatus = iota
	TodoInProcessStatus
	TodoCompletedStatus
)

type Todo struct {
	gorm.Model
	status int `gorm:""`
	plant string ``
}

func Test() {
	fmt.Println("hello")
}