package main

import (
	"github.com/Galionme/metric-service.git/cmd/server/options"
	config "github.com/Galionme/metric-service.git/internal/config/server"
	"github.com/Galionme/metric-service.git/internal/handlers"
	"github.com/Galionme/metric-service.git/internal/middleware"
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	options.ParseOptions()

	var cfg config.ConfigServer
	if err := env.Parse(&cfg); err != nil {
		return
	}

	if cfg.Address != "" && *options.OptionsServer.Address != "" {
		*options.OptionsServer.Address = cfg.Address
	}

	err := run(*options.OptionsServer.Address)
	if err != nil {
		return
	}
}

func run(address string) error {

	router := chi.NewRouter()

	router.Get("/", handlers.HomeMetrics)

	router.Route("/value", func(r chi.Router) {
		r.With(middleware.CorrectnessType, middleware.CorrectnessName).Get("/{type}/{name}", handlers.ValueMetric)
	})
	router.Route("/update", func(r chi.Router) {
		r.With(middleware.CorrectnessType, middleware.CorrectnessName).Post("/{type}/{name}/{value}", handlers.UpdateMetric)
	})

	return http.ListenAndServe(address, router)
}
