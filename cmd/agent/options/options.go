package options

import (
	"flag"
	"fmt"
	"os"
)

var (
	OptionsAgent *Options
)

type Options struct {
	Address        *string
	ReportInterval *int
	PollInterval   *int
}

func NewOptions() *Options {
	return &Options{}
}

func init() {
	OptionsAgent = NewOptions()
	OptionsAgent.Address = flag.String("a", "localhost:8080", "HTTP server search address")
	OptionsAgent.ReportInterval = flag.Int("r", 10, "frequency of sending metrics to the server")
	OptionsAgent.PollInterval = flag.Int("p", 2, "frequency of polling metrics from the package")
}

func ParseOptions() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Use: %s [-a address]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Invalid arguments: %v\n", flag.Args())
		flag.Usage()
		os.Exit(1)
	}
}
