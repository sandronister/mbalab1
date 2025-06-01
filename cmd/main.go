package main

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/di"
	"github.com/sandronister/mbalab1/internal/infra/web"
)

func main() {
	env, err := config.LoadEnviromentVars(".")

	if err != nil {
		panic(err)
	}

	server := web.NewServer(env.WebPort)

	handler := di.NewWeatherHandler(env)

	server.RegisterRoutes(handler)

	if err := server.Start(); err != nil {
		panic(err)
	}

}
