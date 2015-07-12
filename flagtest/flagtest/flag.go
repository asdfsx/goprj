package flagtest

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	*flag.FlagSet

	config_file string "config file"
	ipaddr      string "ip address"
	port        uint   "port number"
}

func NewConfig() *config {

	cfg := &config{
		config_file: "config.toml",
		ipaddr:      "127.0.0.1",
		port:        8080,
	}

	cfg.FlagSet = flag.NewFlagSet("etcd", flag.ContinueOnError)
	fs := cfg.FlagSet
	fs.Usage = func() {
		fmt.Println(usage)
		fmt.Println(flags)
	}

	fs.StringVar(&cfg.config_file, "config-file", "config.toml", "path to the config file")
	fs.StringVar(&cfg.ipaddr, "ipaddr", "127.0.0.1", "ipaddr to listen")
	fs.UintVar(&cfg.port, "port", 8080, "port to listen")

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
