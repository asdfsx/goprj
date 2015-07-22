package geoip

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
)

func TestReadlocation1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `1	O1|O1(not set)|O1(not set)|0.0000|0.0000
2	Asia/Pacific Region|Asia/Pacific Region(not set)|Asia/Pacific Region(not set)|35.0000|105.0000
3	Europe|Europe(not set)|Europe(not set)|47.0000|8.0000
4	Andorra|Andorra(not set)|Andorra(not set)|42.5000|1.5000
5	United Arab Emirates|United Arab Emirates(not set)|United Arab Emirates(not set)|24.0000|54.0000
6	Afghanistan|Afghanistan(not set)|Afghanistan(not set)|33.0000|65.0000
7	Antigua and Barbuda|Antigua and Barbuda(not set)|Antigua and Barbuda(not set)|17.0500|-61.8000
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestReadblock...")

	house, err := NewLocationpthouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", house)

	//delete testfile
	os.Remove(testfile)
}

func TestGetLocation1(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `1	O1|O1(not set)|O1(not set)|0.0000|0.0000
2	Asia/Pacific Region|Asia/Pacific Region(not set)|Asia/Pacific Region(not set)|35.0000|105.0000
3	Europe|Europe(not set)|Europe(not set)|47.0000|8.0000
4	Andorra|Andorra(not set)|Andorra(not set)|42.5000|1.5000
5	United Arab Emirates|United Arab Emirates(not set)|United Arab Emirates(not set)|24.0000|54.0000
6	Afghanistan|Afghanistan(not set)|Afghanistan(not set)|33.0000|65.0000
7	Antigua and Barbuda|Antigua and Barbuda(not set)|Antigua and Barbuda(not set)|17.0500|-61.8000
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestGetLocation...")

	house, err := NewLocationpthouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}

	if val, ok := house.Geoip_locations[5]; ok {
		fmt.Printf("%+v", val)
	}

	if val, ok := house.Geoip_locations[8]; ok {
		fmt.Printf("%+v", val)
	}
	//delete testfile
	os.Remove(testfile)
}
