package models

type User struct {
	ID         string   `json:"user_id"`
	LastName   string   `json:"last_name"`
	FirsttName string   `json:"first_name"`
	Hobbies    []string `json:"hobbies"`
}
