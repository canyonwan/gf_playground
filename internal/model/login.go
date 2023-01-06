package model

type LoginInput struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type LoginOutput struct {
	Info string `json:"info"`
}
