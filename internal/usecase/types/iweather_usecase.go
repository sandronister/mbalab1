package types

import (
	"github.com/sandronister/mbalab1/internal/dto"
	"github.com/sandronister/mbalab1/internal/infra/http/types"
)

type IWeatherUseCase interface {
	Get(localidade string) (*dto.ResultDTO, *types.HttpError)
}
