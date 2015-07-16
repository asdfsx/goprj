package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type geoip_block struct {
	startip    int
	endip      int
	locationid int
}

type geoip_location struct {
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

type blockhouse struct {
	geoip_blockfile string
	geoip_blocks    []geoip_block
}

type locationhouse struct {
	geoip_locationfile string
	geoip_locations    []geoip_location
}

func NewLocationhouse(locationfile string) (*locationhouse, error) {
	house := &locationhouse{
		geoip_locationfile: locationfile,
		geoip_locations:    make([]geoip_location, 0),
	}
	err := house.readlocation()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *locationhouse) readlocation() error {
	istream, err := os.Open(house.geoip_locationfile)
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

		location := geoip_location{
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
		house.geoip_locations = append(house.geoip_locations, location)
	}

	return nil
}

func NewBlockhouse(blockfile string) (*blockhouse, error) {
	house := &blockhouse{
		geoip_blockfile: blockfile,
		geoip_blocks:    make([]geoip_block, 0),
	}

	err := house.readblock()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *blockhouse) readblock() error {
	istream, err := os.Open(house.geoip_blockfile)
	defer istream.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(istream)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tmp := strings.Split(strings.Replace(line, "\"", "", -1), ",")
		if len(tmp) != 3 {
			continue
		}
		startip, err := strconv.Atoi(tmp[0])
		if err != nil {
			continue
		}
		endip, err := strconv.Atoi(tmp[1])
		if err != nil {
			continue
		}
		locationid, err := strconv.Atoi(tmp[2])
		if err != nil {
			continue
		}

		block := geoip_block{
			startip:    startip,
			endip:      endip,
			locationid: locationid,
		}

		house.geoip_blocks = append(house.geoip_blocks, block)
	}

	return nil
}

func (house *blockhouse) Len() int {
	return len(house.geoip_blocks)
}

func (house *blockhouse) Swap(i, j int) {
	house.geoip_blocks[i], house.geoip_blocks[j] = house.geoip_blocks[j], house.geoip_blocks[i]
}

func (house *blockhouse) Less(i, j int) bool {
	if house.geoip_blocks[i].startip < house.geoip_blocks[j].startip {
		return true
	} else if house.geoip_blocks[i].startip > house.geoip_blocks[j].startip {
		return false
	} else if house.geoip_blocks[i].endip < house.geoip_blocks[j].endip {
		return true
	} else if house.geoip_blocks[i].endip > house.geoip_blocks[j].endip {
		return false
	} else {
		return true
	}
}
