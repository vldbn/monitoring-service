package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"monitoring-service/internal/app/model/request"
	"monitoring-service/internal/app/model/response"
	"strconv"
)

// HandlerCreateCryptocurrency HTTP handler
// @Summary Add Cryptocurrency for monitoring
// @Description Adds Cryptocurrency for monitoring
// @Tags cryptocurrencies
// @Accept  json
// @Produce  json
// @Param credentials body request.CreateCryptocurrencyReq true "Add Cryptocurrency"
// @Success 201 {object} response.CryptocurrencyRes
// @Failure 400 {object} response.DefaultRes
// @Failure 401 {object} response.DefaultRes
// @Router /cryptocurrencies [post]
func (s *Server) HandlerCreateCryptocurrency(ctx *fasthttp.RequestCtx) {
	req := request.CreateCryptocurrencyReq{}
	res := response.CryptocurrencyRes{}
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	if req.CurrencyID == "" || req.RefreshInterval <= 0 {
		res.Message = errors.New("invalid currency id or refresh interval").Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	cur, err := s.services.Cryptocurrency().CreateCryptocurrency(req.CurrencyID, req.RefreshInterval)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	res.Cryptocurrency = cur
	s.response(ctx, res, fasthttp.StatusCreated)
	return
}

// HandlerGetCryptocurrencyByCurrencyID HTTP handler
// @Summary Get Cryptocurrency from monitoring
// @Description Gets Cryptocurrency from monitoring by Currency ID
// @Tags cryptocurrencies
// @Produce  json
// @Param  id path string true "Currency ID"
// @Success 200 {object} response.CryptocurrencyRes
// @Failure 404 {object} response.DefaultRes
// @Failure 401 {object} response.DefaultRes
// @Router /cryptocurrencies/{id} [get]
func (s *Server) HandlerGetCryptocurrencyByCurrencyID(ctx *fasthttp.RequestCtx) {
	res := response.CryptocurrencyRes{}
	curId := fmt.Sprintf("%v", ctx.UserValue("currencyID"))
	if curId == "" {
		res.Message = errors.New("invalid currency ID").Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	cur, err := s.services.Cryptocurrency().GetCryptocurrencyByCurrencyID(curId)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusNotFound)
		return
	}
	res.Cryptocurrency = cur
	s.response(ctx, res, fasthttp.StatusOK)
	return
}

// HandlerDeleteCryptocurrencyByCurrencyID HTTP handler
// @Summary Delete Cryptocurrency from monitoring
// @Description Deletes Cryptocurrency from monitoring
// @Tags cryptocurrencies
// @Produce  json
// @Param  id path string true "Currency ID"
// @Success 204 {object} response.CryptocurrencyRes
// @Failure 404 {object} response.DefaultRes
// @Failure 401 {object} response.DefaultRes
// @Router /cryptocurrencies/{id} [delete]
func (s *Server) HandlerDeleteCryptocurrencyByCurrencyID(ctx *fasthttp.RequestCtx) {
	res := response.CryptocurrencyRes{}
	curId := fmt.Sprintf("%v", ctx.UserValue("currencyID"))
	if curId == "" {
		res.Message = errors.New("invalid currency ID").Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	cur, err := s.services.Cryptocurrency().DeleteCryptocurrencyByCurrencyID(curId)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusNotFound)
		return
	}
	res.Cryptocurrency = cur
	fmt.Println(res.Cryptocurrency)
	s.response(ctx, res, fasthttp.StatusNoContent)
	return
}

// HandlerListCryptocurrencies HTTP handler
// @Summary List of Cryptocurrencies in monitoring
// @Description Returns list of cryptocurrencies
// @Tags cryptocurrencies
// @Produce  json
// @Param limit query int false "limit of currencies in response"
// @Param offset query int false "number of currencies to offset"
// @Success 200 {object} response.CryptocurrencyRes
// @Failure 400 {object} response.DefaultRes
// @Failure 401 {object} response.DefaultRes
// @Router /cryptocurrencies [get]
func (s *Server) HandlerListCryptocurrencies(ctx *fasthttp.RequestCtx) {
	res := response.CryptocurrenciesRes{}
	offsetQuery := string(ctx.QueryArgs().Peek("offset"))
	limitQuery := string(ctx.QueryArgs().Peek("limit"))
	if offsetQuery == "" {
		offsetQuery = "0"
	}
	if limitQuery == "" {
		limitQuery = "20"
	}
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	if limit < 0 || offset < 0 {
		res.Message = errors.New("offset or limit can not be negative").Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	cur, err := s.services.Cryptocurrency().ListCryptocurrencies(offset, limit)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, fasthttp.StatusBadRequest)
		return
	}
	res.Cryptocurrencies = cur
	s.response(ctx, res, fasthttp.StatusOK)
	return
}
