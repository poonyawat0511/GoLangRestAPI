package models

type User struct {
	ID         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
}
