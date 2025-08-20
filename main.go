package main

import (
	"github.com/kevinoctavian/evodka_backend/cmd/server"
	"github.com/kevinoctavian/evodka_backend/pkg/config"
)

func main() {

	config.LoadAllConfigs(".env")

	server.Serve()

}
