package models

type Info struct {
	Message string `json:"message"`
}

type NewResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
