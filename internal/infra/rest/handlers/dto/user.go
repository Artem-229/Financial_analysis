package dto

type RegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Login    string `json:"login"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
