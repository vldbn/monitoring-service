package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		tokens, err := svc.Auth().Login("testuser", "password")
		assert.NoError(t, err)
		assert.NotNil(t, tokens)
	})
	t.Run("invalid-username", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		tokens, err := svc.Auth().Login("invalid", "password")
		assert.Error(t, err)
		assert.Nil(t, tokens)
	})
	t.Run("invalid-password", func(t *testing.T) {
		svc := ConfigureServicesForTesting(false)
		tokens, err := svc.Auth().Login("testuser", "invalid")
		assert.Error(t, err)
		assert.Nil(t, tokens)
	})
}
