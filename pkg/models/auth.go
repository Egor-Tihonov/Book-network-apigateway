package models

type RegistrationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	AuthString string `json:"authstring"`
	Password   string `json:"password"`
}
