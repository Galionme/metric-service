package handlers

import (
	"errors"
	"github.com/Galionme/metric-service.git/internal/storage"
	"net/http"
	"strconv"
	"strings"
)

func UpdateMetric(res http.ResponseWriter, req *http.Request) {

	urlParts := strings.Split(req.URL.Path, "/")

	if len(urlParts) < 3 {

		//log.Println("400 - Bad Request", req.URL, "no metric type")
		res.WriteHeader(http.StatusBadRequest)
		return
	} else if len(urlParts) < 5 {

		//log.Println("404 - Not Found", req.URL, "no metric name")
		res.WriteHeader(http.StatusNotFound)
		return
	}

	if err := pushStorage(urlParts[2], urlParts[3], urlParts[4]); err != nil {

		//log.Println("400 - Bad Request", req.URL, "no metric type")
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	//log.Println("200 - Good!")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

func pushStorage(typeMetric, nameMetric, valMetric string) (err error) {
	var val interface{}

	//log.Println(typeMetric, nameMetric, valMetric)

	switch typeMetric {
	case "gauge":
		val, err = strconv.ParseFloat(valMetric, 64)
	case "counter":
		val, err = strconv.ParseInt(valMetric, 0, 64)
	default:
		err = errors.New("unsupported metric type")
	}

	if err != nil {
		return errors.New("not valid argument")
	}

	storage.GlobalMemStorage.Set(nameMetric, val)

	//log.Println(nameMetric, val)

	return nil
}
