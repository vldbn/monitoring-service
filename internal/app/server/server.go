package server

import (
	"encoding/json"
	"github.com/buaazp/fasthttprouter"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"monitoring-service/internal/app/docs"
	"monitoring-service/internal/app/security"
	"monitoring-service/internal/app/service"
)

// Server HTTP server struct
type Server struct {
	router   *fasthttprouter.Router
	services service.Service
	security *security.JWTSecurity
}

// Router getter
func (s *Server) Router() *fasthttprouter.Router {
	return s.router
}

func (s *Server) response(ctx *fasthttp.RequestCtx, data interface{}, code int) {
	ctx.SetStatusCode(code)
	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			ctx.SetBody([]byte(err.Error()))
		}
		ctx.SetBody(b)
	}
}

// ConfigureRoutes configures HTTP routes
func (s *Server) ConfigureRoutes() {
	docs.SwaggerInfo.Title = "Monitoring service"
	docs.SwaggerInfo.Description = "Monitoring service"
	docs.SwaggerInfo.Version = "1.0.0-0"
	docs.SwaggerInfo.Schemes = []string{"http"}
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/api"

	s.router.GET("/api/cryptocurrencies", s.IsAuthenticated(s.HandlerListCryptocurrencies))
	s.router.POST("/api/cryptocurrencies", s.IsAuthenticated(s.HandlerCreateCryptocurrency))
	s.router.GET("/api/cryptocurrencies/:currencyID", s.IsAuthenticated(s.HandlerGetCryptocurrencyByCurrencyID))
	s.router.DELETE("/api/cryptocurrencies/:currencyID", s.IsAuthenticated(s.HandlerDeleteCryptocurrencyByCurrencyID))

	s.router.POST("/api/auth/login", s.HandlerLogin)

	s.router.GET("/docs/*filepath", fasthttpadaptor.NewFastHTTPHandlerFunc(httpSwagger.WrapHandler))
}

// NewServer constructor for Server
func NewServer(services service.Service, security *security.JWTSecurity) *Server {
	return &Server{
		router:   fasthttprouter.New(),
		services: services,
		security: security,
	}
}
