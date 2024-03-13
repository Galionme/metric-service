package main

import (
	"github.com/Galionme/metric-service.git/internal/agent"
	"github.com/go-resty/resty/v2"
)

func main() {
	service := agent.NewStats(0, 100, sendServer)
	service.InitDoctor(10, 2)
}

func sendServer(typeMetric, nameMetric, valueMetric string) {

	client := resty.New()

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		Post("http://localhost:8080/update/" + typeMetric + "/" + nameMetric + "/" + valueMetric)

	if err != nil {
		panic(err)
	}

}
