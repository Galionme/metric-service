package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Address        *string
	ReportInterval *int
	PollInterval   *int
}

func NewOptions() *Options {
	return &Options{
		Address:        flag.String("a", "localhost:8080", "HTTP server search address"),
		ReportInterval: flag.Int("r", 10, "frequency of sending metrics to the server"),
		PollInterval:   flag.Int("p", 2, "frequency of polling metrics from the package"),
	}
}

func ParseOptions() error {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Use: %s [-a address]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "Invalid arguments: %v\n", flag.Args())
		flag.Usage()
		return errors.New("error flags")
	}
	return nil
}
