package security

import (
	"github.com/DamirLuketic/virtual_minds/config"
	"net/http"
)

const (
	ValidationUnauthorizedError = "unauthorized"
)

func NewSecurity(conf *config.ServerConfig) Security {
	return &SecurityImpl{
		ServerConfig: conf,
	}
}

func (m *SecurityImpl) ValidateAuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, isOk := r.BasicAuth()
		if isOk {
			if m.isAuthValid(username, password) {
				next.ServeHTTP(w, r)
				return
			}
		}
		http.Error(w, ValidationUnauthorizedError, http.StatusUnauthorized)
	})
}
