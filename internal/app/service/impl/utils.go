package impl

import (
	"monitoring-service/internal/app/core"
	"monitoring-service/internal/app/database/mockdb"
	"monitoring-service/internal/app/provider/coincapmock"
	"monitoring-service/internal/app/security"
	"monitoring-service/internal/app/service"
)

// ConfigureServicesForTesting configures services for testing
func ConfigureServicesForTesting(raiseError bool) service.Service {
	conf := core.NewConfig()
	conf.SetUsername("testuser")
	svc := NewServiceImpl(
		mockdb.NewMockDB(raiseError),
		coincapmock.NewMockCoinCapAPI(),
		conf.Username(),
		conf.Password(),
		security.NewJWTSecurity(conf.SecretKey()),
	)
	return svc
}
