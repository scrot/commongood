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
	Addr string `env:"ADDRESS" envDefault:":3000"`
}

type Server struct {
	Addr   string
	Logger *httplog.Logger
}

func NewServer(logger *httplog.Logger, config *ServerConfig) *Server {
	if config == nil {
		config = &ServerConfig{}
		if err := env.Parse(config); err != nil {
			logger.Error("server", slog.Any("error", err))
		}
	}

	return &Server{Addr: config.Addr, Logger: logger}
}

//go:embed favicon.ico
var favicon embed.FS

func (svr Server) Start() {
	mux := chi.NewRouter()
	mux.Use(httplog.RequestLogger(svr.Logger))

	mux.Get("/", RootHandler)
	mux.Handle("/favicon.ico", http.FileServer(http.FS(favicon)))

	svr.Logger.Info("starting server", slog.String("addr", svr.Addr))
	http.ListenAndServe(svr.Addr, mux)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	log.Debug("Hello from RootHandler")

	fmt.Fprintf(w, "Under construction...")
}
