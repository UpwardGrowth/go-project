package router

import (
	"github.com/gin-gonic/gin"
	"go-project/app/web/apis"
)

func InitAlbumsRouter(engine *gin.Engine, port string)  {
	engine.GET("/albums", apis.GetAlbums)
	engine.GET("/albums/:id", apis.GetAlbumByID)
	engine.POST("/albums", apis.PostAlbums)
	engine.Run(":"+port)
}