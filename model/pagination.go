package model

type Pagination struct {
	PageSize uint `query:"page_size"`
	Offset   uint `query:"offset"`
}
