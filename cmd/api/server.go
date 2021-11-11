package api

import (
	"fmt"
	"github.com/spf13/viper"
	"go-project/app/web/router"
)

func Init()  {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	viper.SetConfigName("settings")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("File not found: %s \n", err))
		} else {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}

	setup()
}

func setup()  {
	router.InitRouter()
}
