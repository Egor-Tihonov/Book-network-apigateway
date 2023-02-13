package models

type Post struct {
	Title   string `json:"title"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Content string `json:"content"`
}
