package impl

import (
	"errors"
	"monitoring-service/internal/app/model/entity"
	"monitoring-service/internal/app/security"
	"time"
)

type authService struct {
	username string
	password string
	security *security.JWTSecurity
}

// Login implements service.AuthService interface method
func (a *authService) Login(username, password string) (*entity.Tokens, error) {
	if a.username != username || a.password != password {
		return nil, errors.New("invalid username or password")
	}
	access, err := a.security.GenerateToken(username, time.Minute*10)
	if err != nil {
		return nil, err
	}
	refresh, err := a.security.GenerateToken(username, time.Hour*48)
	if err != nil {
		return nil, err
	}
	tokens := entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}
	return &tokens, nil
}

func newAuthService(username string, password string, security *security.JWTSecurity) *authService {
	return &authService{username: username, password: password, security: security}
}
