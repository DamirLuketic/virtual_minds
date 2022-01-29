package security

import (
	"github.com/DamirLuketic/virtual_minds/config"
	"net/http"
)

type Security interface {
	ValidateAuthorizationMiddleware(next http.Handler) http.Handler
}

type SecurityImpl struct {
	ServerConfig *config.ServerConfig
}
