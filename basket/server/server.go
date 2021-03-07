package server

import (
	"github.com/ervitis/backend-challenge/basket/endpoint"
	protopb "github.com/ervitis/backend-challenge/proto"
	"github.com/ervitis/logme"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"time"
)

type (
	Server struct {
		address, port string
		router        endpoint.IRouter
		log           logme.Loggerme
		srv           *grpc.Server
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

func WithRouter(r endpoint.IRouter) Options {
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

	listen, err := net.Listen("tcp", ":" + s.port)
	if err != nil {
		panic(err)
	}

	protopb.RegisterBasketServer(s.srv, s.router.GetRouter())

	go func() {
		s.log.L().Infof("server running on %s", s.getStringConnection())
		if err := s.srv.Serve(listen); err != nil {
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

	s.srv.GracefulStop()
	os.Exit(0)
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

	timeout := 15 * time.Second
	if os.Getenv("DEBUG") != "" {
		timeout = 240 * time.Second
	}

	opts.srv = grpc.NewServer(grpc.ConnectionTimeout(timeout))

	return opts
}
