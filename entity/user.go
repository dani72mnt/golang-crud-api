package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Family    string    `json:"family"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
