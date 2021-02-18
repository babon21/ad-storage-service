package http_test

import (
	"encoding/json"
	"errors"
	adHttp "github.com/babon21/ad-storage-service/internal/ad/storage/delivery/http"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain/mocks"
	"github.com/bxcodec/faker/v3"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAdList(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)
	mockAds := make([]domain.Ad, 0)
	mockAds = append(mockAds, mockAd)

	mockService.On("GetAdList", mock.AnythingOfType("int"), mock.Anything, mock.Anything).Return(mockAds, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads?page=1&sort=-price", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAdList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetAdList_SortParamInvalid(t *testing.T) {
	mockService := new(mocks.AdService)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/rooms?/ads?page=1&sort=abracadabra", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAdList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetAdList_PageParamInvalid(t *testing.T) {
	mockService := new(mocks.AdService)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/rooms?/ads?page=some_str&sort=date_added", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAdList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestGetAdList_ServiceError(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)
	mockAds := make([]domain.Ad, 0)
	mockAds = append(mockAds, mockAd)

	mockService.On("GetAdList", mock.AnythingOfType("int"), mock.Anything, mock.Anything).Return(nil, errors.New("Unexpected Error"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads?page=1&sort=-price", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAdList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetAdList_SortParamIsEmpty(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)
	mockAds := make([]domain.Ad, 0)
	mockAds = append(mockAds, mockAd)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads?page=1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAdList(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateAd(t *testing.T) {
	mockAd := domain.Ad{
		Title:       "title",
		Description: "desc",
		Price:       10,
		Photos:      []string{"link1"},
	}

	tempMockAd := mockAd
	mockService := new(mocks.AdService)

	j, err := json.Marshal(tempMockAd)
	assert.NoError(t, err)

	mockService.On("CreateAd", mock.AnythingOfType("*domain.Ad")).Return("1", nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.CreateAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestCreateAd_RequestBodyInvalid(t *testing.T) {
	mockService := new(mocks.AdService)

	mockAd := domain.Ad{
		Title:       "title",
		Description: "desc",
		Price:       -10,
		Photos:      []string{"link1"},
	}

	j, err := json.Marshal(mockAd)
	assert.NoError(t, err)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.CreateAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateAd_ServiceError(t *testing.T) {
	mockAd := domain.Ad{
		Title:       "title",
		Description: "desc",
		Price:       10,
		Photos:      []string{"link1"},
	}

	tempMockAd := mockAd
	mockService := new(mocks.AdService)

	j, err := json.Marshal(tempMockAd)
	assert.NoError(t, err)

	mockService.On("CreateAd", mock.AnythingOfType("*domain.Ad")).Return("", errors.New("Unexpected Error"))

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/ads", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.CreateAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetAd(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)

	mockService.On("GetAd", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(mockAd, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads/1?fields=description,photos", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetAd_WithoutFields(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)

	mockService.On("GetAd", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(mockAd, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads/1", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockService.AssertExpectations(t)
}

func TestGetAd_ServiceError(t *testing.T) {
	var mockAd domain.Ad
	err := faker.FakeData(&mockAd)
	assert.NoError(t, err)
	mockService := new(mocks.AdService)

	mockService.On("GetAd", mock.AnythingOfType("string"), mock.AnythingOfType("[]string")).Return(domain.Ad{}, errors.New("Unexpected Error"))

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/ads/1?fields=description,photos", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := adHttp.AdHandler{
		AdService: mockService,
	}
	err = handler.GetAd(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
	mockService.AssertExpectations(t)
}
