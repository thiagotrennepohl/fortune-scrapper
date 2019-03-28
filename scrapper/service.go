package scrapper

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	AlreadyExistsMsg = fmt.Sprintf("%s\n", `{"Message":"Message already exists"}`)
)

type scrapperService struct {
	fortuneRepository  FortuneAppRepository
	scrapperRepository ScrapperRepository
}

func NewScrapperService(fortuneRepository FortuneAppRepository, scrapperRepository ScrapperRepository) ScrapperService {
	return &scrapperService{
		fortuneRepository:  fortuneRepository,
		scrapperRepository: scrapperRepository,
	}
}

func (svc *scrapperService) FullSync() error {
	messages, err := svc.scrapperRepository.GetData()
	if err != nil {
		return err
	}
	for _, message := range messages {
		err := svc.fortuneRepository.Save(message)
		if err != nil {
			if err.Error() == AlreadyExistsMsg {
				continue
			}
			return err
		}
	}
	return nil
}

func (svc *scrapperService) SaveMessage() error {
	messages, err := svc.scrapperRepository.GetData()
	if err != nil {
		return err
	}
	index := svc.getRandomIndex(0, len(messages))
	return svc.fortuneRepository.Save(messages[index])
}

func (svc *scrapperService) getRandomIndex(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
