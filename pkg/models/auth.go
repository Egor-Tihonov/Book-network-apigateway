package models

type RegistrationRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	AuthString string `json:"authstring"`
	Password   string `json:"password"`
}
