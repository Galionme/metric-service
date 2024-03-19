package main

import (
	"fmt"
	"net/http"

	config "github.com/Galionme/metric-service/internal/config/server"

	"github.com/Galionme/metric-service/internal/handlers"
	"github.com/Galionme/metric-service/internal/middleware"

	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
)

func main() {

	options := NewOptions()

	err := ParseOptions()
	if err != nil {
		fmt.Println(err)
		return
	}

	var cfg config.ConfigServer
	if err := env.Parse(&cfg); err != nil {
		fmt.Println(err)
		return
	}

	if cfg.Address != "" && *options.Address != "" {
		*options.Address = cfg.Address
	}

	err = run(*options.Address)
	if err != nil {
		return
	}
}

func run(address string) error {
	return http.ListenAndServe(address, getRouter())
}

func getRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/", handlers.HomeMetrics)

	router.Route("/value", func(r chi.Router) {
		r.With(middleware.Correctness).Get("/{type}/{name}", handlers.ValueMetric)
	})
	router.Route("/update", func(r chi.Router) {
		r.With(middleware.Correctness).Post("/{type}/{name}/{value}", handlers.UpdateMetric)
	})

	return router
}
