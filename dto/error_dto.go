package dto

type ErrorDto struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
