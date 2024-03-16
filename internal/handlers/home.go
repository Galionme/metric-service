package handlers

import (
	"html/template"
	"net/http"

	"github.com/Galionme/metric-service.git/internal/storage"
)

func HomeMetrics(res http.ResponseWriter, req *http.Request) {

	data := storage.GlobalMemStorage.GetAll()
	tmpl, err := template.ParseFiles("internal/templates/HomeList.html")
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(res, data)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)
}
