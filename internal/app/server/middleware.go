package server

import (
	"errors"
	"github.com/valyala/fasthttp"
	"log"
	"monitoring-service/internal/app/model/response"
)

// IsAuthenticated authenticated middleware
func (s *Server) IsAuthenticated(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		accessToken := s.security.ExtractToken(ctx)
		claims, err := s.security.ValidateToken(accessToken)
		if err != nil {
			res := response.DefaultRes{}
			res.Message = errors.New("you must login").Error()
			s.response(ctx, res, fasthttp.StatusUnauthorized)
			return
		}
		log.Println("Logged in with username: ", claims["username"])
		next(ctx)
	}
}
