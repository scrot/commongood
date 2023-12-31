package internal

import (
	"embed"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/caarlos0/env/v10"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"8080"`
}

type Server struct {
	Port   int
	Logger *httplog.Logger
}

func NewServer(logger *httplog.Logger, config *ServerConfig) *Server {
	if config == nil {
		config = &ServerConfig{}
		if err := env.Parse(config); err != nil {
			logger.Error("server", slog.Any("error", err))
		}
	}

	return &Server{Port: config.Port, Logger: logger}
}

//go:embed favicon.ico
var favicon embed.FS

func (svr Server) Start() {
	mux := chi.NewRouter()
	mux.Use(httplog.RequestLogger(svr.Logger))

	mux.Get("/", RootHandler)
	mux.Handle("/favicon.ico", http.FileServer(http.FS(favicon)))

	svr.Logger.Info("starting server", slog.Int("addr", svr.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", svr.Port), mux)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	log.Debug("Hello from RootHandler")

	fmt.Fprintf(w, "Under construction...")
}
