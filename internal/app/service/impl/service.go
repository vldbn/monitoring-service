package impl

import (
	"monitoring-service/internal/app/database"
	"monitoring-service/internal/app/provider"
	"monitoring-service/internal/app/security"
	"monitoring-service/internal/app/service"
)

// ServiceImpl implements service.Service interface
type ServiceImpl struct {
	authService           *authService
	cryptocurrencyService *cryptocurrencyService
	db                    database.Database
	username              string
	password              string
	cryptoAPI             provider.CryptocurrencyAPI
	security              *security.JWTSecurity
}

// Cryptocurrency implements service.Service interface method
func (s *ServiceImpl) Cryptocurrency() service.CryptocurrencyService {
	if s.cryptocurrencyService != nil {
		return s.cryptocurrencyService
	}
	s.cryptocurrencyService = newCryptocurrencyService(s.db, s.cryptoAPI)
	return s.cryptocurrencyService
}

// Auth implements service.Service interface method
func (s *ServiceImpl) Auth() service.AuthService {
	if s.authService != nil {
		return s.authService
	}
	s.authService = newAuthService(s.username, s.password, s.security)
	return s.authService
}

// NewServiceImpl constructor
func NewServiceImpl(
	db database.Database,
	cryptoAPI provider.CryptocurrencyAPI,
	username, password string,
	security *security.JWTSecurity,
) service.Service {
	return &ServiceImpl{
		db:        db,
		cryptoAPI: cryptoAPI,
		username:  username,
		password:  password,
		security:  security,
	}
}
