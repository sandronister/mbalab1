package di

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/usecase"
)

func NewCepUseCase(env *config.EnviromentVar) *usecase.CepUsacase {
	return usecase.NewWeatherByCep(env.CepServicePath, newHttpService())
}

func NewWeatherUseCase(env *config.EnviromentVar) *usecase.WeatherUseCase {
	return usecase.NewWeatherUseCase(env.WeatherServicePath, newHttpService())
}
