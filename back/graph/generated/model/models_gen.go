// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
