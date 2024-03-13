package main

import (
	"github.com/Galionme/metric-service.git/internal/api/controllers"
	"github.com/Galionme/metric-service.git/internal/middleware"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.Handle(
		"/update/",
		middleware.Checking(http.HandlerFunc(controllers.UpdateMetric)),
	)
	return http.ListenAndServe(":8080", mux)
}
