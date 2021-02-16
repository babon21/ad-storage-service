package http

import "errors"

var (
	SortParamIsEmpty = errors.New("Sort parameter is empty")
	WrongSortField   = errors.New("Sort field is wrong in sort param")
)
