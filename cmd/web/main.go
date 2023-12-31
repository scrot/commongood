package main

import (
	"github.com/go-chi/httplog/v2"
	"github.com/scrot/commongood/internal"
)

func main() {
	logger := httplog.NewLogger("commongood")
	srv := internal.NewServer(logger, nil)
	srv.Start()
}
