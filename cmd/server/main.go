package main

import (
	"github.com/Galionme/metric-service.git/internal/handlers"
	"github.com/Galionme/metric-service.git/internal/middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		return
	}
}

func run() error {

	router := chi.NewRouter()

	router.Get("/", handlers.HomeMetrics)

	router.Route("/value", func(r chi.Router) {
		r.With(middleware.CorrectnessType, middleware.CorrectnessName).Get("/{type}/{name}", handlers.ValueMetric)
	})
	router.Route("/update", func(r chi.Router) {
		r.With(middleware.CorrectnessType, middleware.CorrectnessName).Post("/{type}/{name}/{value}", handlers.UpdateMetric)
	})

	return http.ListenAndServe(":8080", router)
}
