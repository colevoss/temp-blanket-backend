package api

import (
	"github.com/colevoss/temperature-blanket-backend/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type API struct {
	App *gin.Engine
	Cfg *config.Config
}

func NewApi(cfg *config.Config) *API {
	return &API{
		App: gin.Default(),
		Cfg: cfg,
	}
}

func (a *API) Run() {
	a.App.Run()
}

func (a *API) Init() {
	a.configureCors()
	a.configureRequestId()
}

func (a *API) configureCors() {
	a.App.Use(cors.Default())
}

func (a *API) configureRequestId() {
	a.App.Use(func(c *gin.Context) {
		requestId := uuid.NewString()

		c.Set("requestId", requestId)

		c.Next()
	})
}
