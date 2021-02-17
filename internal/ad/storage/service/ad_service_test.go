package service_test

import (
	"errors"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain/mocks"
	"github.com/babon21/ad-storage-service/internal/ad/storage/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGetAdList(t *testing.T) {
	mockAdRepo := new(mocks.AdRepository)
	mockAd := domain.Ad{
		Id:          "1",
		Title:       "title",
		Description: "desc",
		Price:       10,
		DateAdded:   "2020-12-30",
		MainPhoto:   "link1",
		Photos:      nil,
	}

	mockAds := make([]domain.Ad, 0)
	mockAds = append(mockAds, mockAd)

	t.Run("success", func(t *testing.T) {
		mockAdRepo.On("GetAdList", mock.AnythingOfType("int"), mock.Anything, mock.Anything).Return(mockAds, nil).Once()
		s := service.NewAdService(mockAdRepo)

		page := 1
		list, err := s.GetAdList(page, service.PriceField, service.AscOrder)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockAds))

		mockAdRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockAdRepo.On("GetAdList", mock.AnythingOfType("int"), mock.Anything, mock.Anything).
			Return(nil, errors.New("Unexpected Error")).Once()
		s := service.NewAdService(mockAdRepo)

		page := 1
		list, err := s.GetAdList(page, service.PriceField, service.AscOrder)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockAdRepo.AssertExpectations(t)
	})
}

func TestCreateAd(t *testing.T) {
	mockAdRepo := new(mocks.AdRepository)
	mockAd := domain.Ad{
		Id:          "1",
		Title:       "title",
		Description: "desc",
		Price:       10,
		DateAdded:   "2020-12-30",
		MainPhoto:   "link1",
		Photos:      nil,
	}

	t.Run("success", func(t *testing.T) {
		mockAdRepo.On("CreateAd", mock.Anything).Return("1", nil)
		s := service.NewAdService(mockAdRepo)
		id, err := s.CreateAd(&mockAd)

		assert.NoError(t, err)
		assert.Equal(t, mockAd.Id, id)
		mockAdRepo.AssertExpectations(t)
	})
}

func TestGetAd(t *testing.T) {
	mockAdRepo := new(mocks.AdRepository)
	mockAd := domain.Ad{
		Id:          "1",
		Title:       "title",
		Description: "desc",
		Price:       10,
		DateAdded:   "2020-12-30",
		MainPhoto:   "link1",
		Photos:      nil,
	}

	mockAds := make([]domain.Ad, 0)
	mockAds = append(mockAds, mockAd)

	t.Run("success", func(t *testing.T) {
		mockAdRepo.On("GetAd", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(mockAd, nil).Once()
		s := service.NewAdService(mockAdRepo)

		fields := []string{"photos"}
		ad, err := s.GetAd(mockAd.Id, fields)
		assert.NoError(t, err)
		assert.Equal(t, mockAd, ad)

		mockAdRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockAdRepo.On("GetAdList", mock.AnythingOfType("int"), mock.Anything, mock.Anything).
			Return(nil, errors.New("Unexpected Error")).Once()
		s := service.NewAdService(mockAdRepo)

		page := 1
		list, err := s.GetAdList(page, service.PriceField, service.AscOrder)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockAdRepo.AssertExpectations(t)
	})
}
