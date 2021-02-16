package service

import "github.com/babon21/ad-storage-service/internal/ad/storage/domain"

type AdRepository interface {
	GetAdList(page int, sortField SortField, sortOrder SortOrder) ([]domain.Ad, error)
	GetAd(adId string, optionalFields []string) (domain.Ad, error)
	CreateAd(ad *domain.Ad) (string, error)
}
