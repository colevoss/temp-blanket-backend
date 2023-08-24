package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type API struct {
	App *gin.Engine
}

func NewApi() *API {
	return &API{
		App: gin.Default(),
	}
}

func (a *API) Run() {
	a.App.Run()
}

func (a *API) Init() {
	a.configureCors()
}

func (a *API) configureCors() {
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"*"}
	// corsConfig.AllowAllOrigins = true

	a.App.Use(cors.Default())
}
