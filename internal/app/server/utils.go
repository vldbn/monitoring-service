package server

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"monitoring-service/internal/app/core"
	"monitoring-service/internal/app/database/mockdb"
	"monitoring-service/internal/app/provider/coincapmock"
	"monitoring-service/internal/app/security"
	"monitoring-service/internal/app/service/impl"
	"net"
)

// ConfigureServerForTesting configures services for testing
func ConfigureServerForTesting(raiseError bool) *Server {
	conf := core.NewConfig()
	conf.SetUsername("testuser")
	jwt := security.NewJWTSecurity(conf.SecretKey())
	svc := impl.NewServiceImpl(
		mockdb.NewMockDB(raiseError),
		coincapmock.NewMockCoinCapAPI(),
		conf.Username(),
		conf.Password(),
		security.NewJWTSecurity(conf.SecretKey()),
	)
	return NewServer(svc, jwt)
}

// ListenServerForTesting listens fasthttp server for testing
func ListenServerForTesting(server *Server, port int) net.Listener {
	ln, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("cannot start tcp server: %s", err)
	}
	go fasthttp.Serve(ln, server.Router().Handler)
	return ln
}
