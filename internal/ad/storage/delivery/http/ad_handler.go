package http

import (
	"errors"
	"fmt"
	"github.com/babon21/ad-storage-service/internal/ad/storage/domain"
	"github.com/babon21/ad-storage-service/internal/ad/storage/service"
	"github.com/babon21/ad-storage-service/pkg/delivery/http/api"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
	"strings"
)

// ResponseError represent the response error struct
type ResponseError struct {
	Message string `json:"message"`
}

// AdHandler  represent the httphandler for ad
type AdHandler struct {
	adService service.AdService
}

func (h *AdHandler) GetAdList(c echo.Context) error {
	sortFieldStr := c.QueryParam("sort")
	pageParam := c.QueryParam("page")

	sortField, sortOrder, err := parseSortParam(sortFieldStr)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: "page param is invalid, must be natural number"}, "  ")
	}

	// TODO validate page > 0

	ads, err := h.adService.GetAdList(page, sortField, sortOrder)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	response := api.GetAdListResponse{Ads: ads}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func parseSortParam(sortParam string) (service.SortField, service.SortOrder, error) {
	if sortParam == "" {
		return "", "", SortParamIsEmpty
	}

	sortOrder := service.AscOrder
	if sortParam[0] == '-' {
		sortOrder = service.DescOrder
		sortParam = sortParam[1:]
	}

	switch sortParam {
	case string(service.PriceField):
		return service.PriceField, sortOrder, nil
	case string(service.DateAddedField):
		return service.DateAddedField, sortOrder, nil
	default:
		return "", "", WrongSortField
	}
}

func (h *AdHandler) CreateAd(c echo.Context) error {
	var request api.CreateAdRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSONPretty(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()}, "  ")
	}

	err = validateCreateAdRequest(request)
	if err != nil {
		return c.JSONPretty(http.StatusBadRequest, ResponseError{Message: err.Error()}, "  ")
	}

	ad := domain.Ad{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Photos:      request.Photos,
		MainPhoto:   request.Photos[0],
	}

	id, err := h.adService.CreateAd(&ad)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	response := api.CreateAdResponse{Id: id}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

func validateCreateAdRequest(request api.CreateAdRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		var errMessage string
		for _, err := range err.(validator.ValidationErrors) {
			errMessage = fmt.Sprintf("%s must %s %s, but actual value %v", err.Field(), err.Tag(), err.Param(), err.Value())
		}
		return errors.New(errMessage)
	}

	return nil
}

func (h *AdHandler) GetAd(c echo.Context) error {
	adId := c.Param("id")
	fieldsParam := c.QueryParam("fields")
	fields := strings.Split(fieldsParam, ",")

	if fieldsParam == "" {
		fields = nil
	}

	ad, err := h.adService.GetAd(adId, fields)
	if err != nil {
		return c.JSONPretty(getStatusCode(err), ResponseError{Message: err.Error()}, "  ")
	}

	response := api.GetAdResponse{
		Ad: ad,
	}

	return c.JSONPretty(http.StatusOK, response, "  ")
}

// NewAdHandler will initialize the ads/ resources endpoint
func NewAdHandler(e *echo.Echo, s service.AdService) {
	handler := &AdHandler{
		adService: s,
	}

	e.GET("/ads", handler.GetAdList)
	e.GET("/ads/:id", handler.GetAd)
	e.POST("/ads", handler.CreateAd)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Error().Msg(err.Error())
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
