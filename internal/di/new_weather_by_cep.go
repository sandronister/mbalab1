package di

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/usecase"
)

func NewWeatherByCep(env *config.EnviromentVar) *usecase.WeatherByCep {
	return usecase.NewWeatherByCep(env.CepServicePath, newHttpService())
}
