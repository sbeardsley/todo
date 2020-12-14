package app

import (
	"beards.ly/todo/config"
	"github.com/sbeardsley/cqrs"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Application used to pass around to httpserver or cli or whatever
type Application struct {
	Dispatcher cqrs.Sender
	Config     *viper.Viper
	DB         *gorm.DB
}

// NewApplication constructs application
func NewApplication(dispatcher cqrs.Sender, db *gorm.DB) Application {
	cfg := config.GetConfig()
	return Application{
		Dispatcher: dispatcher,
		Config:     cfg,
		DB:         db,
	}
}
