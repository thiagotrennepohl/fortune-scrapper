package scrapper_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/thiagotrennepohl/fortune-scrapper/mocks"
	"github.com/thiagotrennepohl/fortune-scrapper/models"
	"github.com/thiagotrennepohl/fortune-scrapper/scrapper"
)

var (
	scrapperRepository   = mocks.NewScrapperRepositoryMock()
	fortuneAppRepository = mocks.NewFortuneAppRepositoryMock()
	scrapperService      = scrapper.NewScrapperService(fortuneAppRepository, scrapperRepository)
	validFortuneMessage  = []models.FortuneMessage{
		models.FortuneMessage{
			ID:      "12krikr02491289eurfcmf",
			Message: "CI/CD Without tests is just bug delivery",
		},
	}
)

func TestSaveMessage(t *testing.T) {
	//Happy path
	fortuneAppRepository.SetSaveFuncReturn(nil)
	scrapperRepository.SetGetDataRetun(validFortuneMessage, nil)
	err := scrapperService.SaveMessage()
	assert.NoError(t, err)

	//Err retrieving messages
	scrapperRepository.SetGetDataRetun([]models.FortuneMessage{}, &models.ErrCouldNotRetrieveMessages{Message: "Service is offline"})
	err = scrapperService.SaveMessage()
	assert.Error(t, err)
}

func TestFullSync(t *testing.T) {
	//HappyPath
	scrapperRepository.SetGetDataRetun(validFortuneMessage, nil)
	err := scrapperService.FullSync()
	assert.NoError(t, err)

	//Err retieving messages
	scrapperRepository.SetGetDataRetun([]models.FortuneMessage{}, &models.ErrCouldNotRetrieveMessages{Message: "Service is offline"})
	err = scrapperService.FullSync()
	assert.Error(t, err)

	//Error Message Already Exists
	scrapperRepository.SetGetDataRetun(validFortuneMessage, nil)
	fortuneAppRepository.SetSaveFuncReturn(fmt.Errorf("%s", scrapper.AlreadyExistsMsg))
	err = scrapperService.FullSync()
	assert.NoError(t, err)

	//Generic Error
	fortuneAppRepository.SetSaveFuncReturn(fmt.Errorf("%s", "Lib error"))
	err = scrapperService.FullSync()
	assert.Error(t, err)

}
