package utils

import (
	"log"

	"github.com/DravenChen/todolist/internal/config"
	"github.com/spf13/viper"
)

// InitConfig init config file
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs/")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&config.C); err != nil {
		log.Fatal(err)
	}
	log.Println("configuration initialized")
}