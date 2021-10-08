package server

import (
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"monitoring-service/internal/app/model/request"
	"monitoring-service/internal/app/model/response"
	"net/http"
)

// HandlerLogin HTTP handler
// @Summary Login
// @Description Login
// @Tags auth
// @Param credentials body request.AuthLoginReq true "Username and Password"
// @Accept json
// @Produce json
// @Success 200 {object} response.AuthLoginRes
// @Success 400 {object} response.DefaultRes
// @Router /auth/login [post]
func (s *Server) HandlerLogin(ctx *fasthttp.RequestCtx) {
	req := request.AuthLoginReq{}
	res := response.AuthLoginRes{}
	if err := json.Unmarshal(ctx.PostBody(), &req); err != nil {
		res.Message = err.Error()
		s.response(ctx, res, http.StatusBadRequest)
		return
	}
	if req.Username == "" || req.Password == "" {
		res.Message = errors.New("empty username or password").Error()
		s.response(ctx, res, http.StatusBadRequest)
		return
	}
	tokens, err := s.services.Auth().Login(req.Username, req.Password)
	if err != nil {
		res.Message = err.Error()
		s.response(ctx, res, http.StatusBadRequest)
		return
	}
	res.Username = req.Username
	res.Tokens = tokens
	s.response(ctx, res, http.StatusOK)
	return
}
