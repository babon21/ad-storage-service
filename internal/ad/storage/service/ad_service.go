package service

import "github.com/babon21/ad-storage-service/internal/ad/storage/domain"

type AdService interface {
	GetAdList(page int, sortField SortField, sortOrder SortOrder) ([]domain.Ad, error)
	GetAd(adId string, optionalFields []string) (domain.Ad, error)
	CreateAd(ad *domain.Ad) (string, error)
}

type adService struct {
	adRepo AdRepository
}

func (a *adService) GetAdList(page int, sortField SortField, sortOrder SortOrder) ([]domain.Ad, error) {
	list, err := a.adRepo.GetAdList(page, sortField, sortOrder)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (a *adService) GetAd(adId string, optionalFields []string) (domain.Ad, error) {
	return a.adRepo.GetAd(adId, optionalFields)
}

func (a *adService) CreateAd(ad *domain.Ad) (string, error) {
	return a.adRepo.CreateAd(ad)
}

func NewAdService(repository AdRepository) AdService {
	return &adService{adRepo: repository}
}
