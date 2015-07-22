package geoip

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Geoip_block struct {
	Startip    int
	Endip      int
	Locationid int
}

type Blockhouse struct {
	Geoip_blockfile  string
	Geoip_blocks     []Geoip_block
	Geoip_blocks_len int
}

//const seprate = ","
const seprate = "\t"

func NewBlockhouse(blockfile string) (*Blockhouse, error) {
	house := &Blockhouse{
		Geoip_blockfile: blockfile,
		Geoip_blocks:    make([]Geoip_block, 0),
	}

	err := house.readblock()
	if err != nil {
		return house, err
	}
	return house, nil
}

func (house *Blockhouse) readblock() error {
	istream, err := os.Open(house.Geoip_blockfile)
	defer istream.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(istream)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tmp := strings.Split(strings.Replace(line, "\"", "", -1), seprate)
		if len(tmp) != 3 {
			continue
		}
		Startip, err := strconv.Atoi(tmp[0])
		if err != nil {
			continue
		}
		Endip, err := strconv.Atoi(tmp[1])
		if err != nil {
			continue
		}
		Locationid, err := strconv.Atoi(tmp[2])
		if err != nil {
			continue
		}

		block := Geoip_block{
			Startip:    Startip,
			Endip:      Endip,
			Locationid: Locationid,
		}

		house.Geoip_blocks = append(house.Geoip_blocks, block)
	}
	house.Geoip_blocks_len = len(house.Geoip_blocks)
	return nil
}

func (house *Blockhouse) Len() int {
	return len(house.Geoip_blocks)
}

func (house *Blockhouse) Swap(i, j int) {
	house.Geoip_blocks[i], house.Geoip_blocks[j] = house.Geoip_blocks[j], house.Geoip_blocks[i]
}

func (house *Blockhouse) Less(i, j int) bool {
	if house.Geoip_blocks[i].Startip < house.Geoip_blocks[j].Startip {
		return true
	} else if house.Geoip_blocks[i].Startip > house.Geoip_blocks[j].Startip {
		return false
	} else if house.Geoip_blocks[i].Endip < house.Geoip_blocks[j].Endip {
		return true
	} else if house.Geoip_blocks[i].Endip > house.Geoip_blocks[j].Endip {
		return false
	} else {
		return true
	}
}

func (house *Blockhouse) Sort() {
	sort.Sort(house)
}

func (house *Blockhouse) Search(ipaddr int) (int, bool) {
	if house.Geoip_blocks_len > 0 {

		i, j := 0, house.Geoip_blocks_len
		for i <= j {
			h := i + (j-i)/2 // avoid overflow when computing h
			// i ≤ h < j
			if house.Geoip_blocks[h].Startip <= ipaddr && ipaddr <= house.Geoip_blocks[h].Endip {
				return house.Geoip_blocks[h].Locationid, true
			} else if ipaddr < house.Geoip_blocks[h].Startip {
				j = h - 1
			} else if ipaddr > house.Geoip_blocks[h].Endip {
				i = h + 1
			}
		}
	}
	return 0, false
}
