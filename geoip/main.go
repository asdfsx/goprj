package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	usage = `usage: flagtest [flags]
	do some flag test
	flagtest -h| --help`

	flags = `
	--ipaddr 127.0.0.1
	--port 8080
	--confg-file config.ini`
)

type config struct {
	*flag.FlagSet
	ipaddr    string "ipaddr to listen"
	port      uint   "port to listen"
	geoipfile string "geoip filepath"
}

func NewConfig() *config {
	cfg := &config{
		ipaddr:    "0.0.0.0",
		port:      12345,
		geoipfile: "geoip.csv",
	}
	cfg.FlagSet = flag.NewFlagSet("geoip", flag.ContinueOnError)
	fs := cfg.FlagSet

	fs.Usage = func() {
		fmt.Println(usage)
		fmt.Println(flags)
	}

	fs.StringVar(&cfg.geoipfile, "geoipfile", "geoip.csv", "path to the geoip file")
	fs.StringVar(&cfg.ipaddr, "ipaddr", "0.0.0.0", "ipaddr to listen")
	fs.UintVar(&cfg.port, "port", 12345, "port to listen")
	return cfg
}

func (cfg *config) Parse(arguments []string) error {
	perr := cfg.FlagSet.Parse(arguments)
	switch perr {
	case nil:
	case flag.ErrHelp:
		os.Exit(0)
	default:
		os.Exit(2)
	}
	if len(cfg.FlagSet.Args()) != 0 {
		return fmt.Errorf("'%s' is not a valid flag", cfg.FlagSet.Arg(0))
	}
	return nil
}

func main() {
	cfg := NewConfig()
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", cfg)
}
