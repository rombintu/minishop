package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rombintu/minishop/config"
	"github.com/rombintu/minishop/internal/store"
	"github.com/sirupsen/logrus"
)

type App struct {
	Logger *logrus.Logger
	Config *config.Config
	Router *gin.Engine
	Store  *store.Store
}

// Return App
func NewApp(config *config.Config) *App {
	return &App{
		Config: config,
		Logger: logrus.New(),
		Router: gin.Default(),
	}
}

// Init
func (s *App) Start() error {
	s.Logger.Debug("Configure Logger")
	if err := s.ConfigureLogger(); err != nil {
		return err
	}
	s.Logger.Debug("Configure Router")
	s.ConfigureRouter()

	s.Logger.Debug("Configure store")
	if err := s.ConfigureStore(); err != nil {
		return err
	}

	s.Logger.Info(fmt.Sprintf(
		"Starting API server on http://%s:%s",
		s.Config.Server.Host,
		s.Config.Server.Port,
	),
	)

	return http.ListenAndServe(
		s.Config.Server.Port,
		s.Router,
	)
}

// Configure Logger, set value from Config file
func (s *App) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.Config.Default.LogLevel)
	if err != nil {
		return err
	}

	s.Logger.SetLevel(level)

	return nil
}

// Add routes
func (s *App) ConfigureRouter() {
	s.Router.GET("/ping", s.Ping())

	s.Router.GET("/user", s.GetUser())
	s.Router.POST("/user", s.CreateUser())

	s.Router.GET("/item", s.GetItem())
	s.Router.GET("/item/all", s.GetItems())
	s.Router.POST("/item", s.CreateItem())

	s.Router.GET("/basket", s.GetBasket())
	s.Router.GET("/basket/all", s.GetBaskets())
	s.Router.POST("/basket", s.UpdateBasket())
}

// Configure db, from Config file
func (s *App) ConfigureStore() error {
	s.Store = &store.Store{
		Config: &s.Config.Postgres,
	}
	return nil
}
