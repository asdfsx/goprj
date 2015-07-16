package geoip

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type geoip_block struct {
	startip    int
	endip      int
	locationid int
}

type blockhouse struct {
	geoip_blockfile  string
	geoip_blocks     []geoip_block
	geoip_blocks_len int
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
	house.geoip_blocks_len = len(house.geoip_blocks)
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

func (house *blockhouse) Sort() {
	sort.Sort(house)
}

func (house *blockhouse) Search(ipaddr int) int {
	cmpfunc := func(i int) bool {
		fmt.Printf("===========ipaddr===%v: %+v, %v\n", ipaddr, house.geoip_blocks[i], house.geoip_blocks[i].startip <= ipaddr && ipaddr <= house.geoip_blocks[i].endip)
		if house.geoip_blocks[i].startip > ipaddr {
			return true
		} else if ipaddr > house.geoip_blocks[i].endip {
			return false
		} else {
			return false
		}
	}
	return sort.Search(house.geoip_blocks_len, cmpfunc)
}
