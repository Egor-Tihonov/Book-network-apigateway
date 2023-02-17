package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdate struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	Status      string `json:"status"`
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpassword"`
}

type SecurityUserUpdate struct {
	OldPassword string `json:"oldpassword"`
	NewPassword string `json:"newpasswrod"`
}
