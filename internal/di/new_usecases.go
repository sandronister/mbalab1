package di

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/infra/http/types"
	"github.com/sandronister/mbalab1/internal/usecase"
)

func NewCepUseCase(env *config.EnviromentVar, httpService types.IHttpService) *usecase.CepUsacase {
	return usecase.NewWeatherByCep(env.CepServicePath, httpService)
}

func NewWeatherUseCase(env *config.EnviromentVar, httpService types.IHttpService) *usecase.WeatherUseCase {
	return usecase.NewWeatherUseCase(env.WeatherServicePath, httpService)
}
