package models

type Book struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type Author struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
