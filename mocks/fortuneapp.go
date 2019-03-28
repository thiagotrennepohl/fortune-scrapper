package mocks

import (
	"github.com/thiagotrennepohl/fortune-scrapper/models"
)

// type ScrapperRepository interface {
// 	GetData() ([]models.FortuneMessage, error)
// }

// type FortuneAppRepository interface {
// 	Save(models.FortuneMessage) error
// }

// type ScrapperService interface {
// 	SaveMessage() error
// 	FullSync() error
// }

type fortuneAppRepositoryMock struct {
	saveFuncErrBehaviour error
}

func NewFortuneAppRepositoryMock() *fortuneAppRepositoryMock {
	return &fortuneAppRepositoryMock{}
}

func (m *fortuneAppRepositoryMock) SetSaveFuncReturn(err error) {
	m.saveFuncErrBehaviour = err
}

func (m *fortuneAppRepositoryMock) Save(message models.FortuneMessage) error {
	return m.saveFuncErrBehaviour
}
