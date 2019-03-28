package models

import (
	fortuneBackEnd "github.com/thiagotrennepohl/fortune-backend/models"
)

type FortuneMessage fortuneBackEnd.FortuneMessage

type FortuneAppEndpointConfig struct {
	SaveEndpoint string
}
