package scrapper

import "github.com/thiagotrennepohl/fortune-scrapper/models"

type ScrapperRepository interface {
	GetData() (models.FortuneMessage, error)
}

type FortuneAppRepository interface {
	Save(models.FortuneMessage) error
}

type ScrapperService interface {
	SaveMessage() error
	FullSync() error
}
