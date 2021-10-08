package security

import (
	"github.com/stretchr/testify/assert"
	"monitoring-service/internal/app/core"
	"testing"
	"time"
)

func TestJWTSecurity_GenerateToken(t *testing.T) {
	conf := core.NewConfig()
	jwt := NewJWTSecurity(conf.SecretKey())
	t.Run("generate-token", func(t *testing.T) {
		token, err := jwt.GenerateToken("username", time.Hour)
		assert.NoError(t, err)
		assert.NotNil(t, token)
	})
}

func TestJWTSecurity_ValidateToken(t *testing.T) {
	conf := core.NewConfig()
	jwt := NewJWTSecurity(conf.SecretKey())
	t.Run("success", func(t *testing.T) {
		token, _ := jwt.GenerateToken("username", time.Hour)
		claims, err := jwt.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, "username", claims["username"])
	})
	t.Run("expired-token", func(t *testing.T) {
		token, _ := jwt.GenerateToken("username", time.Nanosecond*1)
		time.Sleep(time.Second * 1)
		_, err := jwt.ValidateToken(token)
		assert.Error(t, err)
	})
}
