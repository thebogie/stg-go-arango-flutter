// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type User struct {
	Key       string `json:"key"`
	ID        string `json:"id"`
	Rev       string `json:"rev"`
	Firstname string `json:"firstname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
