package security

import "crypto/subtle"

func (m *SecurityImpl) isAuthValid(username, password string) bool {
	if subtle.ConstantTimeCompare([]byte(m.ServerConfig.APIUser), []byte(username)) == 1 &&
		subtle.ConstantTimeCompare([]byte(m.ServerConfig.APIPassword), []byte(password)) == 1 {
		return true
	}
	return false
}
