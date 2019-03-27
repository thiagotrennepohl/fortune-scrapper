package scrapper

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
	return nil
}

func (svc *scrapperService) SaveMessage() error {
	message, err := svc.scrapperRepository.GetData()
	if err != nil {
		return err
	}

	return svc.fortuneRepository.Save(message)
}
