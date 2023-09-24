package model

type UserID int

type User struct {
	ID       UserID `json:"id"       db:"id"`
	Name     string `json:"name"     db:"name"`
	Password string `json:"password" db:"password"`
	Created  string `json:"created"  db:"created"`
	Modified string `json:"modified" db:"modified"`
}
