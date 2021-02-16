package postgres

import (
	"fmt"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain"
	"github.com/babon21/ad-storage-service/internal/ad/storage/service"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"time"
)

type postgresAdRepository struct {
	Conn *sqlx.DB
}

func (p *postgresAdRepository) GetAdList(page int, sortField service.SortField, sortOrder service.SortOrder) ([]domain.Ad, error) {
	getAdListQuery := formGetListQuery(page, 10, sortField, sortOrder)

	dbAds, err := p.getAds(getAdListQuery)
	if err != nil {
		return nil, err
	}
	return dbAds, nil
}

func (p *postgresAdRepository) getAds(query string) ([]domain.Ad, error) {
	ads := make([]domain.Ad, 0, 1)
	err := p.Conn.Select(&ads, query)
	return ads, err
}

func formGetListQuery(page int, count int, sortField service.SortField, sortOrder service.SortOrder) string {
	var order string
	switch sortOrder {
	case service.AscOrder:
		order = "ASC"
	case service.DescOrder:
		order = "DESC"
	}

	offset := (page - 1) * count

	return fmt.Sprintf("SELECT title,main_photo,price FROM ad ORDER BY %s %s LIMIT %d OFFSET %d", sortField, order, count, offset)
}

func (p *postgresAdRepository) GetAd(adId string, optionalFields []string) (domain.Ad, error) {
	var ad domain.Ad
	var err error

	if len(optionalFields) == 0 {
		err = p.Conn.QueryRow(getSimpleAdQuery, adId).Scan(&ad.Title, &ad.Price, &ad.Description, &ad.MainPhoto, pq.Array(&ad.Photos))
		return ad, err
	}

	isPhotos := false
	isDesc := false
	for _, field := range optionalFields {
		if field == "photos" {
			isPhotos = true
		} else if field == "description" {
			isDesc = true
		}
	}

	if isPhotos && isDesc {
		err = p.Conn.QueryRow(getAdWithPhotoAndDescQuery, adId).Scan(&ad.Title, &ad.Price, &ad.Description, &ad.MainPhoto, pq.Array(&ad.Photos))
	} else if isPhotos {
		err = p.Conn.QueryRow(getAdWithPhotosQuery, adId).Scan(&ad.Title, &ad.Price, &ad.MainPhoto, pq.Array(&ad.Photos))
	} else if isDesc {
		err = p.Conn.QueryRow(getAdWithDescQuery, adId).Scan(&ad.Title, &ad.Price, &ad.Description, &ad.MainPhoto)
	} else {
		// TODO create error
	}
	return ad, err
}

func (p *postgresAdRepository) CreateAd(ad *domain.Ad) (string, error) {
	var id string
	err := p.Conn.QueryRow(createAdQuery, ad.Title, ad.Description, ad.Price, ad.MainPhoto, pq.Array(ad.Photos), time.Now()).Scan(&id)
	return id, err
}

func NewPostgresAdRepository(conn *sqlx.DB) service.AdRepository {
	return &postgresAdRepository{conn}
}
