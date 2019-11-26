package models

type User struct {
	Username string `json:Username`
	Email    string `json:Email`
	Pass     string `json:Pass`
}
