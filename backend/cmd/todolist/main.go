package main

import (
	"github.com/DravenChen/todolist/internal/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
}
