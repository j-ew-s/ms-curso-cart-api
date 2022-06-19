package sessionServices

import "github.com/j-ew-s/ms-curso-user-api/user"

type Session struct {
	Token           string
	TokenValidation *user.TokenValidation
}
