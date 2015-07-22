package geoip

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Geoip_locationpt struct {
	LocId   int
	LocInfo string
}

type Locationpthouse struct {
	Geoip_locationfile string
	Geoip_locations    map[int]*Geoip_locationpt
}

func NewLocationpthouse(locationfile string) (*Locationpthouse, error) {
	house := &Locationpthouse{
		Geoip_locationfile: locationfile,
		Geoip_locations:    make(map[int]*Geoip_locationpt, 0),
	}
	err := house.readlocation()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *Locationpthouse) readlocation() error {
	istream, err := os.Open(house.Geoip_locationfile)
	defer istream.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(istream)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tmp := strings.Split(strings.Replace(line, "\"", "", -1), "\t")
		if len(tmp) != 2 {
			continue
		}

		LocId, err := strconv.Atoi(tmp[0])
		if err != nil {
			continue
		}
		LocInfo := tmp[1]

		location := Geoip_locationpt{
			LocId:   LocId,
			LocInfo: LocInfo,
		}
		house.Geoip_locations[LocId] = &location
	}

	return nil
}
