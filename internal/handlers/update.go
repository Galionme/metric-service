package handlers

import (
	"net/http"
	"strings"

	"github.com/Galionme/metric-service/internal/enums"
	"github.com/Galionme/metric-service/internal/storage"
	"github.com/Galionme/metric-service/internal/util"
	"github.com/go-chi/chi/v5"
)

func UpdateMetric(res http.ResponseWriter, req *http.Request) {

	nameMetric := chi.URLParam(req, "name")
	typeMetric := chi.URLParam(req, "type")

	switch strings.ToLower(typeMetric) {

	case enums.TypeCounter:

		tmp, _ := storage.GlobalMemStorage.Get(nameMetric)

		if tmp == nil {
			tmp = int64(0)
		}

		count, ok := tmp.(int64)
		if !ok {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		num, err := util.StringToCounter(chi.URLParam(req, "value"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		storage.GlobalMemStorage.Set(nameMetric, count+num)

	case enums.TypeGauge:

		num, err := util.StringToGauge(chi.URLParam(req, "value"))
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		storage.GlobalMemStorage.Set(nameMetric, num)

	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
