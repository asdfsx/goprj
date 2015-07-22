package geoip

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Geoip_block struct {
	startip    int
	endip      int
	locationid int
}

type Blockhouse struct {
	geoip_blockfile  string
	geoip_blocks     []Geoip_block
	geoip_blocks_len int
}

func NewBlockhouse(blockfile string) (*Blockhouse, error) {
	house := &Blockhouse{
		geoip_blockfile: blockfile,
		geoip_blocks:    make([]Geoip_block, 0),
	}

	err := house.readblock()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *Blockhouse) readblock() error {
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

		block := Geoip_block{
			startip:    startip,
			endip:      endip,
			locationid: locationid,
		}

		house.geoip_blocks = append(house.geoip_blocks, block)
	}
	house.geoip_blocks_len = len(house.geoip_blocks)
	return nil
}

func (house *Blockhouse) Len() int {
	return len(house.geoip_blocks)
}

func (house *Blockhouse) Swap(i, j int) {
	house.geoip_blocks[i], house.geoip_blocks[j] = house.geoip_blocks[j], house.geoip_blocks[i]
}

func (house *Blockhouse) Less(i, j int) bool {
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

func (house *Blockhouse) Sort() {
	sort.Sort(house)
}

func (house *Blockhouse) Search(ipaddr int) (int, bool) {
	i, j := 0, house.geoip_blocks_len
	for i <= j {
		h := i + (j-i)/2 // avoid overflow when computing h
		// i ≤ h < j
		if house.geoip_blocks[h].startip <= ipaddr && ipaddr <= house.geoip_blocks[h].endip {
			return house.geoip_blocks[h].locationid, true
		} else if ipaddr < house.geoip_blocks[h].startip {
			j = h - 1
		} else if ipaddr > house.geoip_blocks[h].endip {
			i = h + 1
		}
	}
	return 0, false
}
