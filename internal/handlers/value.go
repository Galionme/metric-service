package handlers

import (
	"net/http"
	"strings"

	"github.com/Galionme/metric-service/internal/enums"
	"github.com/Galionme/metric-service/internal/storage"
	"github.com/Galionme/metric-service/internal/util"
	"github.com/go-chi/chi/v5"
)

func ValueMetric(res http.ResponseWriter, req *http.Request) {

	nameMetric := chi.URLParam(req, "name")
	typeMetric := chi.URLParam(req, "type")

	tmp, ok := storage.GlobalMemStorage.Get(nameMetric)
	if !ok {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	switch strings.ToLower(typeMetric) {

	case enums.TypeCounter:
		count, ok := tmp.(int64)
		if !ok {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if strNum, err := util.CounterToString(count); err == nil {
			_, err := res.Write([]byte(strNum))
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
		}

	case enums.TypeGauge:

		count, ok := tmp.(float64)
		if !ok {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		if strNum, err := util.GaugeToString(count); err == nil {
			_, err := res.Write([]byte(strNum))
			if err != nil {
				res.WriteHeader(http.StatusBadRequest)
				return
			}
		}

	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusOK)
}
