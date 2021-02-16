package api

import "github.com/babon21/ad-storage-service/internal/ad/storage/domain"

type CreateAdRequest struct {
	Title       string   `json:"title" validate:"required,max=200"`
	Description string   `json:"description" validate:"required,max=1000"`
	Price       int      `json:"price" validate:"gt=0"`
	Photos      []string `json:"photos" validate:"required,max=3"`
}

type CreateAdResponse struct {
	Id string `json:"id"`
}

type GetAdResponse struct {
	domain.Ad
}

type GetAdListResponse struct {
	Ads []domain.Ad `json:"ads"`
}
