package main

import (
	"os"

	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/di"
	"github.com/sandronister/mbalab1/internal/infra/web"
)

func main() {
	env, err := config.LoadEnviromentVars(".")

	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = env.WebPort
	}

	server := web.NewServer(port)

	handler := di.NewWeatherHandler(env)

	server.RegisterRoutes(handler)

	if err := server.Start(); err != nil {
		panic(err)
	}

}
