package api

type Todo struct {
	Title     string `json:"title,omitempty"`
	Completed bool `json:"completed,omitempty"`
	Date      string `json:"date,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}
