package main

import (
	"github.com/Galionme/metric-service.git/cmd/agent/options"
	"github.com/Galionme/metric-service.git/internal/agent"
	"github.com/go-resty/resty/v2"
)

func main() {

	options.ParseOptions()

	service := agent.NewStats(0, 100, sendServer, *options.OptionsAgent.Address)
	service.InitDoctor(*options.OptionsAgent.ReportInterval, *options.OptionsAgent.PollInterval)
}

func sendServer(typeMetric, nameMetric, valueMetric, address string) {

	client := resty.New()

	_, err := client.R().
		SetHeader("Content-Type", "application/json").
		Post("http://" + address + "/update/" + typeMetric + "/" + nameMetric + "/" + valueMetric)

	if err != nil {
		panic(err)
	}

}
