package model

// Pagination
// @Description Pagination
type Pagination struct {
	PageSize uint `json:"page_size" query:"page_size"` // Page size
	Offset   uint `json:"offset" query:"offset"`       // Offset
}
