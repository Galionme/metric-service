package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Options struct {
	Address *string
}

func NewOptions() *Options {
	return &Options{
		Address: flag.String("a", "localhost:8080", "HTTP server start address"),
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
