package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"log"
	"monitoring-service/internal/app/model/request"
	"net/http"
	"testing"
)

func TestServer_HandlerLogin(t *testing.T) {
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.POST(url, srv.HandlerLogin)
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		userCreds := request.AuthLoginReq{
			Username: "testuser",
			Password: "password",
		}
		j, _ := json.Marshal(userCreds)
		b := bytes.NewBuffer(j)
		req, _ := http.NewRequest(http.MethodPost, fullUrl, b)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusOK, res.StatusCode)
	})
	t.Run("invalid-request-json", func(t *testing.T) {
		userCreds := request.AuthLoginReq{}
		j, _ := json.Marshal(userCreds)
		b := bytes.NewBuffer(j)
		req, _ := http.NewRequest(http.MethodPost, fullUrl, b)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusBadRequest, res.StatusCode)
	})
	t.Run("nil-body-request", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodPost, fullUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusBadRequest, res.StatusCode)
	})
}
