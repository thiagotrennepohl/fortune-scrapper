package main

import (
	"log"
	"os"
	"strconv"
	"time"

	fortuneAppRepo "github.com/thiagotrennepohl/fortune-scrapper/internal/repository/fortuneapp_repo"
	scrapperRepo "github.com/thiagotrennepohl/fortune-scrapper/internal/repository/scrapper_repo"
	"github.com/thiagotrennepohl/fortune-scrapper/models"
	"github.com/thiagotrennepohl/fortune-scrapper/scrapper"
)

var (
	crontEnabled                 bool
	cronInterval                 time.Duration
	fortuneAppConfig             models.FortuneAppEndpointConfig
	fortuneMessageSourceEndpoint string
)

func init() {
	if isCronEnabled, ok := os.LookupEnv("ENABLE_CRON"); ok {
		envValue, err := strconv.ParseBool(isCronEnabled)
		if err != nil {
			log.Fatalf("Cannot use %s as ENABLE_CRON environment value", isCronEnabled)
		}
		crontEnabled = envValue
	}

	if EnvCronInterval, ok := os.LookupEnv("CRON_INTERVAL"); ok {
		envValue, err := time.ParseDuration(EnvCronInterval)
		if err != nil {
			log.Fatalf("Cannot use %s as CRON_INTERVAL environment value", EnvCronInterval)
		}
		cronInterval = envValue
	}

	if fortuneAppSaveEndpoint, ok := os.LookupEnv("FORTUNE_APP_SAVE_ENDPOINT"); ok {
		fortuneAppConfig.SaveEndpoint = fortuneAppSaveEndpoint
	} else {
		log.Fatal("Environment FORTUNE_APP_SAVE_ENDPOINT not set ")
	}

	if messageSourceEndpoint, ok := os.LookupEnv("MESSAGE_SOURCE_ENDPOINT"); ok {
		fortuneMessageSourceEndpoint = messageSourceEndpoint
	} else {
		log.Fatal("Environment MESSAGE_SOURCE_ENDPOINT not set ")
	}

}

func main() {

	fortuneAppRepository := fortuneAppRepo.NewFortuneAppRepository(fortuneAppConfig)
	scrapperRepository := scrapperRepo.NewScrapperRepository(fortuneMessageSourceEndpoint)
	fortuneScrapperService := scrapper.NewScrapperService(fortuneAppRepository, scrapperRepository)

	if crontEnabled && cronInterval != 0*time.Second {
		for {
			err := fortuneScrapperService.FullSync()
			if err != nil {
				log.Fatal(err.Error())
			}
			time.Sleep(cronInterval)
		}
	}

	err := fortuneScrapperService.FullSync()
	if err != nil {
		log.Fatal(err.Error())
	}

}
