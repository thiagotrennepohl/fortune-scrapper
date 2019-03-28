package scrapper_repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thiagotrennepohl/fortune-scrapper/models"
	"github.com/thiagotrennepohl/fortune-scrapper/scrapper"
)

// type ScrapperRepository interface {
// 	GetData() (models.FortuneMessage, error)
// }

// type FortuneAppRepository interface {
// 	Save(models.FortuneMessage) error
// }

// type ScrapperService interface {
// 	SaveMessage() error
// 	FullSync() error
// }

var httpClient = http.Client{
	Timeout: 30 * time.Second,
}

type scrapperRepository struct {
	url string
}

func NewScrapperRepository(url string) scrapper.ScrapperRepository {
	return &scrapperRepository{
		url: url,
	}
}

func (repository *scrapperRepository) GetData() ([]models.FortuneMessage, error) {
	messages := []models.FortuneMessage{}
	req, err := http.NewRequest("GET", repository.url, nil)
	if err != nil {
		return messages, &models.ErrCouldNotRetrieveMessages{Message: "Failed to create http request " + err.Error()}
	}

	resp, err := httpClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return messages, &models.ErrCouldNotRetrieveMessages{Message: "Failed to perform request " + err.Error()}
	}

	if resp.StatusCode != 200 {
		return messages, &models.ErrCouldNotRetrieveMessages{Message: fmt.Sprintf("The fortune service smees not to be working, Status Code: %d", resp.StatusCode)}
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return messages, &models.ErrCouldNotRetrieveMessages{Message: "Failed to read response body " + err.Error()}
	}

	err = json.Unmarshal(responseBody, &messages)
	if err != nil {
		return messages, &models.ErrCouldNotRetrieveMessages{Message: "Failed to bind json " + err.Error()}
	}

	responseBody = nil

	return messages, nil

}
