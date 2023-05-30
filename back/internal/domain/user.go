// internal/domain/user.go
package domain

type User struct {
	Key       string `json:"_key"`
	Id        string `json:"_id"`
	Rev       string `json:"_rev"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
