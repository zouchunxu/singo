package jwt

import "errors"

var (
	TokenExpired     = errors.New("token已经过期")
	TokenNotValidYet = errors.New("token还没生效")
	TokenMalformed   = errors.New("无效的token")
	TokenInvalid     = errors.New("无效的token")
)
