package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CorrectnessType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if typeMetric := chi.URLParam(req, "type"); typeMetric != "counter" && typeMetric != "gauge" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		handler.ServeHTTP(res, req)
	})
}

func CorrectnessName(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if nameMetric := chi.URLParam(req, "name"); nameMetric == "" {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		handler.ServeHTTP(res, req)
	})
}
