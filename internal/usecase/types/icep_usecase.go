package types

import (
	"github.com/sandronister/mbalab1/internal/dto"
	"github.com/sandronister/mbalab1/internal/infra/http/types"
)

type ICep interface {
	Get(cep string) (*dto.CepResultDTO, *types.HttpError)
}
