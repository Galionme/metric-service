package middleware

import (
	"net/http"

	"github.com/Galionme/metric-service/internal/enums"

	"github.com/go-chi/chi/v5"
)

func Correctness(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if typeMetric := chi.URLParam(req, "type"); typeMetric != enums.TypeCounter && typeMetric != enums.TypeGauge {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		if nameMetric := chi.URLParam(req, "name"); nameMetric == "" {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		handler.ServeHTTP(res, req)
	})
}

// right?

//type Correctness struct {
//	handler http.Handler
//}
//
//func (c *Correctness) ServeHTTP(res http.ResponseWriter, req *http.Request) {
//	if nameMetric := chi.URLParam(req, "name"); nameMetric == "" {
//		res.WriteHeader(http.StatusNotFound)
//		return
//	}
//	if typeMetric := chi.URLParam(req, "type"); typeMetric != enums.TypeCounter && typeMetric != enums.TypeGauge {
//		res.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	c.handler.ServeHTTP(res, req)
//}
//
//func NewCorrectness(handler http.Handler) *Correctness {
//	return &Correctness{
//		handler: handler,
//	}
//}
