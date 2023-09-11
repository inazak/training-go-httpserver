package model

type UserID int

type User struct {
	ID       UserID `json:"id"       db:"id"`
	Name     string `json:"name"     db:"name"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role"     db:"role"`
	Created  string `json:"created"  db:"created"`
	Modified string `json:"modified" db:"modified"`
}
