package main

import (
	"github.com/kevinoctavian/evodka_backend/cmd/server"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
)

func main() {
	// load all configs from .env
	config.LoadAllConfigs(".env")

	// serve the server
	server.Serve()
}
