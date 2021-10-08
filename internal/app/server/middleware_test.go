package server

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"log"
	"monitoring-service/internal/app/core"
	"monitoring-service/internal/app/security"
	"net/http"
	"testing"
	"time"
)

func TestServer_IsAuthenticated(t *testing.T) {
	conf := core.NewConfig()
	jwt := security.NewJWTSecurity(conf.SecretKey())
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.GET(url, srv.IsAuthenticated(srv.HandlerListCryptocurrencies))
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		token, _ := jwt.GenerateToken("testuser", time.Hour)
		req, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
		header := map[string][]string{
			"Authorization": {fmt.Sprintf("Bearer %s", token)},
		}
		req.Header = header
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusOK, res.StatusCode)
	})
	t.Run("invalid-token", func(t *testing.T) {
		token, _ := jwt.GenerateToken("testuser", time.Nanosecond)
		time.Sleep(time.Second)
		req, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
		header := map[string][]string{
			"Authorization": {fmt.Sprintf("Bearer %s", token)},
		}
		req.Header = header
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusUnauthorized, res.StatusCode)
	})
}
