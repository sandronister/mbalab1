package di

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/infra/handler"
)

func NewWeatherHandler(env *config.EnviromentVar) *handler.WeatherByCep {
	return handler.NewWeatherByCep(NewWeatherByCep(env))
}
