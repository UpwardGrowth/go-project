package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func InitRouter()  {
	engine := gin.Default()
	if engine == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}

	port := viper.GetString("settings.application.port")
	InitAlbumsRouter(engine, port)
}