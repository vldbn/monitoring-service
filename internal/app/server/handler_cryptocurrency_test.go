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

func TestServer_HandlerCreateCryptocurrency(t *testing.T) {
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.POST(url, srv.HandlerCreateCryptocurrency)
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		createReq := request.CreateCryptocurrencyReq{
			CurrencyID:      "currency",
			RefreshInterval: 10,
		}
		j, _ := json.Marshal(createReq)
		b := bytes.NewBuffer(j)
		req, _ := http.NewRequest(http.MethodPost, fullUrl, b)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusCreated, res.StatusCode)
	})
	t.Run("invalid-request-params", func(t *testing.T) {
		createReq := request.CreateCryptocurrencyReq{
			CurrencyID:      "",
			RefreshInterval: -10,
		}
		j, _ := json.Marshal(createReq)
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

func TestServer_HandlerGetCryptocurrencyByCurrencyID(t *testing.T) {
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.GET(fmt.Sprintf("%s/:currencyID", url), srv.HandlerGetCryptocurrencyByCurrencyID)
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		bitcoinUrl := fmt.Sprintf("%s/bitcoin", fullUrl)
		req, _ := http.NewRequest(http.MethodGet, bitcoinUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusOK, res.StatusCode)
	})
	t.Run("failure-404", func(t *testing.T) {
		bitcoinUrl := fmt.Sprintf("%s/bitcoin22", fullUrl)
		req, _ := http.NewRequest(http.MethodGet, bitcoinUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusNotFound, res.StatusCode)
	})
}

func TestServer_HandlerDeleteCryptocurrencyByCurrencyID(t *testing.T) {
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.DELETE(fmt.Sprintf("%s/:currencyID", url), srv.HandlerDeleteCryptocurrencyByCurrencyID)
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		bitcoinUrl := fmt.Sprintf("%s/bitcoin", fullUrl)
		req, _ := http.NewRequest(http.MethodDelete, bitcoinUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusNoContent, res.StatusCode)
	})
	t.Run("failure-404", func(t *testing.T) {
		bitcoinUrl := fmt.Sprintf("%s/bitcoin22", fullUrl)
		req, _ := http.NewRequest(http.MethodDelete, bitcoinUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusNotFound, res.StatusCode)
	})
}

func TestServer_HandlerListCryptocurrencies(t *testing.T) {
	port := 9000
	host := "http://localhost"
	url := "/test"
	fullUrl := fmt.Sprintf("%s:%d%s", host, port, url)
	srv := ConfigureServerForTesting(false)
	srv.router.GET(url, srv.HandlerListCryptocurrencies)
	ln := ListenServerForTesting(srv, port)
	defer func() {
		if err := ln.Close(); err != nil {
			log.Println(err)
			return
		}
	}()
	t.Run("success", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fullUrl, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusOK, res.StatusCode)
	})
	t.Run("success-query", func(t *testing.T) {
		query := fmt.Sprintf("%s?offset=1&limit=20", fullUrl)
		req, _ := http.NewRequest(http.MethodGet, query, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusOK, res.StatusCode)
	})
	t.Run("invalid-query", func(t *testing.T) {
		query := fmt.Sprintf("%s?offset=-1&limit=-20", fullUrl)
		req, _ := http.NewRequest(http.MethodGet, query, nil)
		client := http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, fasthttp.StatusBadRequest, res.StatusCode)
	})
}
