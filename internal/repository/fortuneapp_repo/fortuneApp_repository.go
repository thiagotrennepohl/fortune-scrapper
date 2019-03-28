package fortuneapp_repo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/thiagotrennepohl/fortune-scrapper/models"
	"github.com/thiagotrennepohl/fortune-scrapper/scrapper"
)

type fortuneAppRepository struct {
	config models.FortuneAppEndpointConfig
}

var httpClient = http.Client{
	Timeout: 30 * time.Second,
}

func NewFortuneAppRepository(config models.FortuneAppEndpointConfig) scrapper.FortuneAppRepository {
	return &fortuneAppRepository{
		config: config,
	}
}

func (repository *fortuneAppRepository) Save(message models.FortuneMessage) error {
	requestBody, err := json.Marshal(message)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", repository.config.SaveEndpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("%s", string(responseBody))
	}

	return nil

}
