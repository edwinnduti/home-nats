package models

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

type NewUser struct {
	Message string
}

type Configs struct {
	NatsUrl string
	Port    string
}

type NewResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
