package dto

type SuccessDto struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessListDto struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}
