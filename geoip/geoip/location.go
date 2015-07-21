package geoip

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Geoip_location struct {
	locId      int
	country    string
	region     string
	city       string
	postalCode string
	latitude   string
	longitude  string
	metroCode  string
	areaCode   string
}

type Locationhouse struct {
	Geoip_locationfile string
	Geoip_locations    map[int]*Geoip_location
}

func NewLocationhouse(locationfile string) (*Locationhouse, error) {
	house := &Locationhouse{
		Geoip_locationfile: locationfile,
		Geoip_locations:    make(map[int]*Geoip_location, 0),
	}
	err := house.readlocation()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *Locationhouse) readlocation() error {
	istream, err := os.Open(house.Geoip_locationfile)
	defer istream.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(istream)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tmp := strings.Split(strings.Replace(line, "\"", "", -1), ",")
		if len(tmp) != 9 {
			continue
		}

		locId, err := strconv.Atoi(tmp[0])
		if err != nil {
			continue
		}
		country := tmp[1]
		region := tmp[2]
		city := tmp[3]
		postalCode := tmp[4]
		latitude := tmp[5]
		longitude := tmp[6]
		metroCode := tmp[7]
		areaCode := tmp[8]

		location := Geoip_location{
			locId:      locId,
			country:    country,
			region:     region,
			city:       city,
			postalCode: postalCode,
			latitude:   latitude,
			longitude:  longitude,
			metroCode:  metroCode,
			areaCode:   areaCode,
		}
		house.Geoip_locations[locId] = &location
	}

	return nil
}
