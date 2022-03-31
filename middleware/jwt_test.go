package middleware

import "testing"

func TestValidateToken(t *testing.T) {
	token, _ := GenToken("张三", "123456")

	ValidateToken(token)
}
