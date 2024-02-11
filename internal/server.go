package internal

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bthkn/basilisk/internal/appconfig"
)

type server struct {
	server *http.Server
	config *appconfig.AppConfig
}

type Server interface {
	Run() error
}

var _ Server = (*server)(nil)

func (s server) Run() error {
	fmt.Printf("Listen on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func NewServer(config *appconfig.AppConfig) Server {
	return &server{
		server: &http.Server{
			Addr:    config.Host + ":" + strconv.Itoa(config.Port),
			Handler: MakeHandler(CardHandler),
		},
		config: config,
	}
}
