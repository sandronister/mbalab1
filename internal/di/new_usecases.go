package di

import (
	"github.com/sandronister/mbalab1/config"
	"github.com/sandronister/mbalab1/internal/infra/http/types"
	"github.com/sandronister/mbalab1/internal/usecase"
	usecaseTypes "github.com/sandronister/mbalab1/internal/usecase/types"
)

func NewCepUseCase(env *config.EnviromentVar, httpService types.IHttpService) usecaseTypes.ICep {
	return usecase.NewWeatherByCep(env.CepServicePath, httpService)
}

func NewWeatherUseCase(env *config.EnviromentVar, httpService types.IHttpService) usecaseTypes.IWeatherUseCase {
	return usecase.NewWeatherUseCase(env.WeatherServicePath, httpService)
}
