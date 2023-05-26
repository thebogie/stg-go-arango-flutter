// internal/domain/auth.go
package domain

type AuthPayload struct {
	Token string
	User  User
}
