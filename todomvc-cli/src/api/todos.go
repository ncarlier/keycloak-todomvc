package api

type EmbeddedResponse struct {
	Todos []Todo `json:"todos,omitempty"`
}

type TodosResponse struct {
	Embedded EmbeddedResponse `json:"_embedded,omitempty"`
	Page     Page `json:"page,omitempty"`
}
