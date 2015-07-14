package main

import (
	"fmt"
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

type geoip_blockhouse struct {
	geoip_blockfile string
	geoip_blocks    []geoip_block
}

type geoip_locationhouse struct{
	geoip_locaitionfile
	geoip_locations [] geoip_location
}


