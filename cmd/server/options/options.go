package options

import (
	"flag"
	"fmt"
	"os"
)

var (
	OptionsServer *Options
)

type Options struct {
	Address *string
}

func NewOptions() *Options {
	return &Options{}
}

func init() {
	OptionsServer = NewOptions()
	OptionsServer.Address = flag.String("a", "localhost:8080", "HTTP server start address")
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
