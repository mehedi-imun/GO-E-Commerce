package util

import "net/http"

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type Pagination struct {
	Limit      int64 `json:"limit"`
	Page       int64 `json:"page"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int64 `json:"totalPages"`
}

func SendPage(w http.ResponseWriter, data any, limit int64, page int64, totalItems int64) {
	totalPages := totalItems / limit
	if totalItems%limit != 0 {
		totalPages++
	}
	resp := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Limit:      limit,
			Page:       page,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}

	SendData(w, http.StatusOK, resp)

}
