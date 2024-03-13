package main

import (
	"github.com/Galionme/metric-service.git/cmd/agent/options"
	"github.com/Galionme/metric-service.git/internal/agent"
	config "github.com/Galionme/metric-service.git/internal/config/agent"
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
)

func main() {

	options.ParseOptions()

	var cfg config.ConfigAgent
	if err := env.Parse(&cfg); err != nil {
		return
	}

	if cfg.Address != "" && *options.OptionsAgent.Address != "" {
		*options.OptionsAgent.Address = cfg.Address
	}
	if cfg.ReportInterval != 0 && *options.OptionsAgent.ReportInterval != 0 {
		*options.OptionsAgent.ReportInterval = cfg.ReportInterval
	}
	if cfg.PollInterval != 0 && *options.OptionsAgent.PollInterval != 0 {
		*options.OptionsAgent.PollInterval = cfg.PollInterval
	}

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
