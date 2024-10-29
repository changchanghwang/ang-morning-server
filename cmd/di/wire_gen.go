// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"angmorning.com/internal/libs/db"
	"angmorning.com/internal/libs/oauth"
	"angmorning.com/internal/server"
	"angmorning.com/internal/services/users/application"
	"angmorning.com/internal/services/users/infrastructure"
	"angmorning.com/internal/services/users/presentation"
)

// Injectors from wire.go:

func InitializeServer() (*server.Server, error) {
	healthCheckHandler := server.NewHandler()
	sqlDB := db.InitDb()
	userRepository := infrastructure.New(sqlDB)
	oauthClientFactory := oauth.NewFactory()
	userService := application.New(userRepository, oauthClientFactory)
	userHandler := presentation.New(userService)
	serverServer := server.NewServer(healthCheckHandler, userHandler)
	return serverServer, nil
}
