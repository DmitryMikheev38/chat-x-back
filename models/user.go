package models

// User ...
type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	EMail        string `json:"eMail"`
	Password     string `json:"password"`
	HashPassword string `json:"hashPassword"`
}
