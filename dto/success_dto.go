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

type SuccessUploadDto struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
}

type MultiResponseDto struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    []StatusDto `json:"data"`
}

type StatusDto struct {
	Filename string `json:"filename"`
	Status   int    `json:"status"`
	Message  string `json:"message"`
}
