package main

import (
	"github.com/ehabterra/go-site/internal/database"
	"github.com/ehabterra/go-site/internal/environment"
	"github.com/ehabterra/go-site/pkg/web_api"
)

func main() {
	environment.LoadEnv()

	db := database.Connect()
	server := web_api.NewServer(db)

	server.Serve()
}
