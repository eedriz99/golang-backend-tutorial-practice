package models

type Post struct {
	ID      int    `json:"id"`
	Date    string `json:"date"`
	Creator User   `json:"creator"`
	Content string `json:"content"`
}
