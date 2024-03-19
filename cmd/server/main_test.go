package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Galionme/metric-service/internal/handlers"
	"github.com/Galionme/metric-service/internal/middleware"
	"github.com/go-chi/chi/v5"
)

func Test_run(t *testing.T) {
	router := chi.NewRouter()
	router.Get("/", handlers.HomeMetrics)
	router.Route("/value", func(r chi.Router) {
		r.With(middleware.Correctness).Get("/{type}/{name}", handlers.ValueMetric)
	})
	router.Route("/update", func(r chi.Router) {
		r.With(middleware.Correctness).Post("/{type}/{name}/{value}", handlers.UpdateMetric)
	})

	ts := httptest.NewServer(router)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL+"/value/counter/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status %d; got %d", http.StatusNotFound, res.StatusCode)
	}

	req, err = http.NewRequest("POST", ts.URL+"/update/gauge/test/20", nil)
	if err != nil {
		t.Fatal(err)
	}

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, res.StatusCode)
	}
}
