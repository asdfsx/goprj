package server

import (
	"strings"
	"fmt"
	"geoip/geoip"
	"net"
	"strconv"
	"io"
)

type geoipserver struct {
	blockfile    string
	locationfile string
	bhouse       *geoip.Blockhouse
	lhouse       *geoip.Locationhouse
}

const noLimit int64 = (1 << 63) - 1
const limit1k int64 = 1024

func NewGeoipServer(blockfile, locationfile string) (server *geoipserver, err error) {
	server = &geoipserver{blockfile: blockfile, locationfile: locationfile}
	bhouse, err := geoip.NewBlockhouse(blockfile)
	if err != nil {
		return
	}
	lhouse, err := geoip.NewLocationhouse(locationfile)
	if err != nil {
		return
	}
	bhouse.Sort()
	server.bhouse = bhouse
	server.lhouse = lhouse

	return
}

func (server *geoipserver) GetLocation(ipaddr int) string {
	if locationid, ok := server.bhouse.Search(ipaddr); ok {
		if location, ok := server.lhouse.Geoip_locations[locationid]; ok {
			return fmt.Sprintf("%v", location)
		}
	}
	return "not found"
}

func (server *geoipserver) HandlerSocket(conn net.Conn) {
	defer conn.Close()
	reader := newBufioReader(io.LimitReader(conn, limit1k))
	ipstr, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	ipaddr, err := strconv.Atoi(strings.TrimSpace(ipstr))
	if err != nil{
		conn.Write([]byte(fmt.Sprintf("ipaddr format error: %v\n",ipstr)))
		return
	}

	result := server.GetLocation(ipaddr)
	_, err = conn.Write([]byte(result))
	if err != nil {
		return
	}
}
