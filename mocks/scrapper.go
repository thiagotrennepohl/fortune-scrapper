package mocks

import (
	"github.com/thiagotrennepohl/fortune-scrapper/models"
)

type scrapperRepositoryMock struct {
	GetDataFuncValueBehaviour []models.FortuneMessage
	GetDataFunErrBehaviour    error
}

func NewScrapperRepositoryMock() *scrapperRepositoryMock {
	return &scrapperRepositoryMock{}
}

func (m *scrapperRepositoryMock) SetGetDataRetun(messages []models.FortuneMessage, err error) {
	m.GetDataFuncValueBehaviour = messages
	m.GetDataFunErrBehaviour = err
}

func (m *scrapperRepositoryMock) GetData() ([]models.FortuneMessage, error) {
	return m.GetDataFuncValueBehaviour, m.GetDataFunErrBehaviour
}
