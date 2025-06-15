package server

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/blacksails/rocket-messages/pkg/message"
	"github.com/blacksails/rocket-messages/pkg/rocket"
)

type Server struct {
	mux           *http.ServeMux
	messageStore  message.Store
	rocketService rocket.Service
	log           *slog.Logger
}

type Option func(*Server)

func WithMessageStore(messageStore message.Store) Option {
	return func(s *Server) {
		s.messageStore = messageStore
	}
}

func WithRocketService(rocketService rocket.Service) Option {
	return func(s *Server) {
		s.rocketService = rocketService
	}
}

func WithLogger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.log = logger
	}
}

func New(opts ...Option) *Server {
	s := &Server{}

	messageStore := message.NewInMemoryStore()
	rocketService := rocket.NewService(messageStore)
	logger := slog.Default()

	defaultOpts := []Option{
		WithMessageStore(messageStore),
		WithRocketService(rocketService),
		WithLogger(logger),
	}

	opts = append(defaultOpts, opts...)
	for _, opt := range opts {
		opt(s)
	}

	s.routes()

	return s
}

func (s *Server) Run() error {
	fmt.Println("Listening on :8080")
	return http.ListenAndServe(":8080", s.mux)
}

func (s *Server) routes() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /messages", s.postMessageHandler())
	mux.HandleFunc("GET /rockets/{id}", s.getRocketHandler())
	mux.HandleFunc("GET /rockets", s.listRocketsHandler())

	s.mux = mux
}
