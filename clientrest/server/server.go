package server

import (
	"github.com/ervitis/logme"
	"net/http"
	"os"
	"os/signal"
)

type (
	Server struct {
		address, port string
		router        *http.ServeMux
		log           logme.Loggerme
	}

	Options func(*Server)
)

func WithPort(port string) Options {
	return func(server *Server) {
		server.port = port
	}
}

func WithAddress(address string) Options {
	return func(server *Server) {
		server.address = address
	}
}

func WithLogger(logger logme.Loggerme) Options {
	return func(server *Server) {
		server.log = logger
	}
}

func WithRouter(r *http.ServeMux) Options {
	return func(server *Server) {
		server.router = r
	}
}

func defaultOptions() *Server {
	return &Server{address: "127.0.0.1", port: "8080"}
}

func (s *Server) Listen() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	errChannel := make(chan error)

	go func() {
		s.log.L().Infof("server running on %s", s.getStringConnection())
		if err := http.ListenAndServe(s.getStringConnection(), s.router); err != nil {
			errChannel <- err
		}
	}()

	go func() {
		for {
			select {
			case err := <-errChannel:
				s.log.L().Fatalf("error received from handler: %s", err.Error())
				signals <- os.Interrupt
			}
		}
	}()

	<-signals
}

func (s *Server) getStringConnection() string {
	return s.address + ":" + s.port
}

func CreateServer(serverOptions ...Options) *Server {
	opts := defaultOptions()

	for _, opt := range serverOptions {
		opt(opts)
	}

	if opts.router == nil {
		panic("handler router not set with WithRouter")
	}

	return opts
}
