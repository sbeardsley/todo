//+build wireinject

package main

import (
	"beards.ly/todo/app"
	"beards.ly/todo/database"
	"beards.ly/todo/dispatcher"
	"beards.ly/todo/repository"
	"github.com/google/wire"
)

func initializeApplication() app.Application {
	wire.Build(database.NewPostgresConnectionString, database.NewDB, repository.NewUserRepository, dispatcher.NewSender, app.NewApplication)
	return app.Application{}
}
