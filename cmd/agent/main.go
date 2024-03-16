package main

import (
	"github.com/Galionme/metric-service/internal/agent"
	config "github.com/Galionme/metric-service/internal/config/agent"
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
)

func main() {

	options := NewOptions()

	err := ParseOptions()
	if err != nil {
		panic(err)
	}

	var cfg config.ConfigAgent
	if err := env.Parse(&cfg); err != nil {
		return
	}

	if cfg.Address != "" && *options.Address != "" {
		*options.Address = cfg.Address
	}
	if cfg.ReportInterval != 0 && *options.ReportInterval != 0 {
		*options.ReportInterval = cfg.ReportInterval
	}
	if cfg.PollInterval != 0 && *options.PollInterval != 0 {
		*options.PollInterval = cfg.PollInterval
	}

	service := agent.NewStats(0, 100, sendServer, *options.Address)
	service.InitDoctor(*options.ReportInterval, *options.PollInterval)
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
