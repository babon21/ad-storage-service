package postgres_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain"
	"github.com/babon21/ad-storage-service/internal/ad/storage/repository/postgres"
	"github.com/babon21/ad-storage-service/internal/ad/storage/service"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetAdList_Desc(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:       "title1",
			Description: "desc1",
			Price:       10,
			DateAdded:   "2002-03-30",
			MainPhoto:   "link1",
		},
		{
			Title:       "title2",
			Description: "desc2",
			Price:       20,
			DateAdded:   "2002-03-30",
			MainPhoto:   "link2",
		},
	}

	rows := sqlmock.NewRows([]string{"title", "main_photo", "price"}).
		AddRow(mockAds[0].Title, mockAds[0].MainPhoto, mockAds[0].Price).
		AddRow(mockAds[1].Title, mockAds[1].MainPhoto, mockAds[1].Price)

	sortField := service.PriceField
	sortOrder := service.DescOrder
	page := 1
	count := 10
	offset := (page - 1) * count

	query := fmt.Sprintf("SELECT title,main_photo,price FROM ad ORDER BY %s %s LIMIT %d OFFSET %d", sortField, strings.ToUpper(string(sortOrder)), count, offset)

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	list, err := a.GetAdList(page, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

func TestGetAdList_Asc(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:       "title1",
			Description: "desc1",
			Price:       10,
			DateAdded:   "2002-03-30",
			MainPhoto:   "link1",
		},
		{
			Title:       "title2",
			Description: "desc2",
			Price:       20,
			DateAdded:   "2002-03-30",
			MainPhoto:   "link2",
		},
	}

	rows := sqlmock.NewRows([]string{"title", "main_photo", "price"}).
		AddRow(mockAds[0].Title, mockAds[0].MainPhoto, mockAds[0].Price).
		AddRow(mockAds[1].Title, mockAds[1].MainPhoto, mockAds[1].Price)

	sortField := service.PriceField
	sortOrder := service.AscOrder
	page := 1
	count := 10
	offset := (page - 1) * count

	query := fmt.Sprintf("SELECT title,main_photo,price FROM ad ORDER BY %s %s LIMIT %d OFFSET %d", sortField, strings.ToUpper(string(sortOrder)), count, offset)

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	list, err := a.GetAdList(page, sortField, sortOrder)
	assert.NoError(t, err)
	assert.Len(t, list, 2)
}

/*
func TestSaveStatistics(t *testing.T) {
	// TODO
	ad := domain.Ad {
		Title:       "title1",
		Description: "desc1",
		Price:       10,
		MainPhoto:   "link1",
	}
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%ad' was not expected when opening s stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")


	columns := []string{"id"}
	exceptedId := "1"
	mock.ExpectQuery("INSERT INTO ad").WithArgs(ad.Title, ad.Description, ad.Price, ad.MainPhoto, pq.Array(ad.Photos)).WillReturnRows(sqlmock.NewRows(columns).AddRow(exceptedId))

	s := postgres.NewPostgresAdRepository(db)

	id, err := s.CreateAd(&ad)
	assert.NoError(t, err)
	assert.Equal(t, exceptedId, id)
}
*/

func TestGetAdWithDesc(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:       "title1",
			Description: "desc1",
			Price:       10,
			MainPhoto:   "link1",
		},
	}

	rows := sqlmock.NewRows([]string{"title", "price", "description", "main_photo"}).
		AddRow(mockAds[0].Title, mockAds[0].Price, mockAds[0].Description, mockAds[0].MainPhoto)

	id := "1"

	mock.ExpectQuery("SELECT title,price,description,main_photo FROM ad").WithArgs(id).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	ad, err := a.GetAd(id, []string{"description"})
	assert.NoError(t, err)
	assert.Equal(t, mockAds[0].Title, ad.Title)
	assert.Equal(t, mockAds[0].Price, ad.Price)
	assert.Equal(t, mockAds[0].Description, ad.Description)
	assert.Equal(t, mockAds[0].MainPhoto, ad.MainPhoto)
}

func TestGetAdWithPhotos(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:     "title1",
			Photos:    []string{"link1", "link2"},
			Price:     10,
			MainPhoto: "link1",
		},
	}

	rows := sqlmock.NewRows([]string{"title", "price", "main_photo", "photos"}).
		AddRow(mockAds[0].Title, mockAds[0].Price, mockAds[0].MainPhoto, pq.Array(mockAds[0].Photos))

	id := "1"

	mock.ExpectQuery("SELECT title,price,main_photo,photos FROM ad").WithArgs(id).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	ad, err := a.GetAd(id, []string{"photos"})
	assert.NoError(t, err)
	assert.Equal(t, mockAds[0].Title, ad.Title)
	assert.Equal(t, mockAds[0].Price, ad.Price)
	assert.Equal(t, mockAds[0].MainPhoto, ad.MainPhoto)
	assert.Equal(t, mockAds[0].Photos, ad.Photos)
}

func TestGetAdSimple(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:     "title1",
			Price:     10,
			MainPhoto: "link1",
		},
	}

	rows := sqlmock.NewRows([]string{"title", "price", "main_photo"}).
		AddRow(mockAds[0].Title, mockAds[0].Price, mockAds[0].MainPhoto)

	id := "1"

	mock.ExpectQuery("SELECT title,price,main_photo FROM ad").WithArgs(id).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	ad, err := a.GetAd(id, nil)
	assert.NoError(t, err)
	assert.Equal(t, mockAds[0].Title, ad.Title)
	assert.Equal(t, mockAds[0].Price, ad.Price)
	assert.Equal(t, mockAds[0].MainPhoto, ad.MainPhoto)
}

func TestGetAdWithDescAndPhotos(t *testing.T) {
	dbMock, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	db := sqlx.NewDb(dbMock, "")

	mockAds := []domain.Ad{
		{
			Title:       "title1",
			Description: "desc1",
			Price:       10,
			MainPhoto:   "link1",
			Photos:      []string{"link1", "link2"},
		},
	}

	rows := sqlmock.NewRows([]string{"title", "price", "description", "main_photo", "photos"}).
		AddRow(mockAds[0].Title, mockAds[0].Price, mockAds[0].Description, mockAds[0].MainPhoto, pq.Array(mockAds[0].Photos))

	id := "1"

	mock.ExpectQuery("SELECT title,price,description,main_photo,photos FROM ad").WithArgs(id).WillReturnRows(rows)
	a := postgres.NewPostgresAdRepository(db)

	ad, err := a.GetAd(id, []string{"description", "photos"})
	assert.NoError(t, err)
	assert.Equal(t, mockAds[0].Title, ad.Title)
	assert.Equal(t, mockAds[0].Price, ad.Price)
	assert.Equal(t, mockAds[0].Description, ad.Description)
	assert.Equal(t, mockAds[0].Photos, ad.Photos)
	assert.Equal(t, mockAds[0].MainPhoto, ad.MainPhoto)
}
