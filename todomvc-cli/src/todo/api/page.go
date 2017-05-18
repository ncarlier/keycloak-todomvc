package api

type Page struct {
	Size          int `json:"size,omitempty"`
	TotalElements int `json:"totalElements,omitempty"`
	TotalPages    int `json:"totalPages,omitempty"`
	number        int `json:"number,omitempty"`
}
