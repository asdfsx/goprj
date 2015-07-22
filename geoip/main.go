package main

import (
	"flag"
	"fmt"
	"geoip/server"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

const (
	usage = `usage:  geoip [flags]
	geoip server query for geoip info via ipaddress
	geoip -h| --help`

	flags = `
	--ipaddr 127.0.0.1:8080
	--block-file geoblock.csv
	--locationi-file geolocation.csv
	--pprof-ipaddr 127.0.0.1:6060`
)

type config struct {
	*flag.FlagSet
	ipaddr          string "ipaddr to listen"
	geoblockfile    string "geoip block filepath"
	geolocationfile string "geoip location filepath"
	pprof_ipaddr    string "pprof ipaddr"
}

func NewConfig() *config {
	cfg := &config{
		ipaddr:          "0.0.0.0:8080",
		geoblockfile:    "geoip.csv",
		geolocationfile: "geolocation.csv",
		pprof_ipaddr:    "0.0.0.0:6060",
	}
	cfg.FlagSet = flag.NewFlagSet("geoip", flag.ContinueOnError)
	fs := cfg.FlagSet

	fs.Usage = func() {
		fmt.Println(usage)
		fmt.Println(flags)
	}

	fs.StringVar(&cfg.geoblockfile, "geoblockfile", "geoblock.csv", "path to the geoip block file")
	fs.StringVar(&cfg.geolocationfile, "geolocationfile", "geolocation.csv", "path to the geoip location file")
	fs.StringVar(&cfg.ipaddr, "ipaddr", "0.0.0.0:8080", "ipaddr to listen")
	fs.StringVar(&cfg.pprof_ipaddr, "pprof_ipaddr", "0.0.0.0:6060", "ipaddr for pprof to listen")
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

func (cfg *config) String() string {
	return fmt.Sprintf("\nipaddr:%v,\npprof_ipaddr:%v,\ngeoblockfile:%v,\ngeolocationfile:%v\n",
		cfg.ipaddr, cfg.pprof_ipaddr, cfg.geoblockfile, cfg.geolocationfile)
}

func main() {
	cfg := NewConfig()
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}

	log.Printf("%+v\n", cfg)

	go func() {
		log.Println(http.ListenAndServe(cfg.pprof_ipaddr, nil))
	}()

	geoipserver, err := server.NewGeoipServer(cfg.geoblockfile, cfg.geolocationfile)

	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}

	socketserver := server.NewSocketServer(cfg.ipaddr)
	socketserver.Handler = geoipserver.HandlerSocket
	err = socketserver.Listen()
	if err != nil {
		log.Println(err)
	}
	socketserver.Run()
}
