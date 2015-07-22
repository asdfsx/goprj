package main

import (
	"time"
	"bufio"
	"flag"
	"fmt"
	"io"
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
	ipaddr      string "ipaddr for query"
	serveraddr  string "ipaddr of geoipserver"
	connections int    "how many connections to create"
	second      int    "how long to do the test"
}

func NewConfig() *config {
	cfg := &config{
		ipaddr:      "16779264",
		serveraddr:  "0.0.0.0:8080",
		connections: 100,
		second:      30,
	}
	cfg.FlagSet = flag.NewFlagSet("geoipclient", flag.ContinueOnError)
	fs := cfg.FlagSet

	fs.Usage = func() {
		fmt.Println(usage)
		fmt.Println(flags)
	}

	fs.StringVar(&cfg.serveraddr, "serveraddr", "0.0.0.0:8080", "geoipserver ipaddr")
	fs.StringVar(&cfg.ipaddr, "ipaddr", "16779264", "ipaddr to query")
	fs.IntVar(&cfg.connections, "connections", 100, "how many connections to create")
	fs.IntVar(&cfg.second, "second", 30, "how many connections to create")
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

func doquery(serveraddr, ipaddr string) {
	conn, err := net.Dial("tcp", serveraddr)
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer conn.Close()
	for {
		conn.Write([]byte(ipaddr + "\n"))

		reader := bufio.NewReader(io.LimitReader(conn, 1024))
		_, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println("result:", string(result))
	}
}

func main() {
	cfg := NewConfig()
	err := cfg.Parse(os.Args[1:])
	if err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	fmt.Println("connect to:", cfg.serveraddr)

    i := 0
    for i < cfg.connections{
	    go doquery(cfg.serveraddr, cfg.ipaddr)
		i += 1
	}

    i = 0
    for i < cfg.second{
		fmt.Println("========time:",i)
		time.Sleep(1 * time.Second)
		i += 1
	}
}
