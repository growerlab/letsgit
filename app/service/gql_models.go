// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package service

type AcitvateCodePayload struct {
	Code string `json:"Code"`
}

type NewUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Result struct {
	Ok bool `json:"OK"`
}
