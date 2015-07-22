package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

const (
	usage = `usage:  geoipclient [flags]
	geoipserver's client query for geoip from geoipserver
	geoipclient -h| --help`

	flags = `
	--serveraddr 127.0.0.1:8080
	--ipaddr 123`
)

type config struct {
	*flag.FlagSet
	ipaddr     string "ipaddr for query"
	serveraddr string "ipaddr of geoipserver"
}

func NewConfig() *config {
	cfg := &config{
		ipaddr:     "1",
		serveraddr: "0.0.0.0:8080",
	}
	cfg.FlagSet = flag.NewFlagSet("geoipclient", flag.ContinueOnError)
	fs := cfg.FlagSet

	fs.Usage = func() {
		fmt.Println(usage)
		fmt.Println(flags)
	}

	fs.StringVar(&cfg.serveraddr, "serveraddr", "0.0.0.0:8080", "geoipserver ipaddr")
	fs.StringVar(&cfg.ipaddr, "ipaddr", "1", "ipaddr to query")
	return cfg
}

func (cfg *config) parse(arguments []string) error {
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
		log.Printf("%v\n", err)
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", cfg.serveraddr)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte(cfg.ipaddr + "\n"))

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("result:", string(result))
}
