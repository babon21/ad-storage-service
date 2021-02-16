package service

type SortField string

const (
	PriceField     SortField = "price"
	DateAddedField SortField = "date_added"
)

type SortOrder string

const (
	AscOrder  SortOrder = "asc"
	DescOrder SortOrder = "desc"
)
