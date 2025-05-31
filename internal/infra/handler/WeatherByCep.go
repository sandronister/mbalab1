package handler

import (
	"github.com/labstack/echo"
	"github.com/sandronister/mbalab1/internal/usecase"
)

type WeatherByCep struct {
	WeatherByCepUseCase *usecase.WeatherByCep
}

func NewWeatherByCep(weatherByCepUseCase *usecase.WeatherByCep) *WeatherByCep {
	return &WeatherByCep{
		WeatherByCepUseCase: weatherByCepUseCase,
	}
}

func (w *WeatherByCep) Get(c echo.Context) error {
	cep := c.Param("cep")

	resp, err := w.WeatherByCepUseCase.Get(cep)
	if err != nil {
		return c.JSON(err.StatusCode, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(200, map[string]interface{}{
		"data": resp,
	})
}
