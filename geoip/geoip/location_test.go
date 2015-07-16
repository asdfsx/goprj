package geoip

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
)

func TestReadlocation(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `Copyright (c) 2012 MaxMind LLC.  All Rights Reserved.
locId,country,region,city,postalCode,latitude,longitude,metroCode,areaCode
1,"O1","","","",0.0000,0.0000,,
2,"AP","","","",35.0000,105.0000,,
3,"EU","","","",47.0000,8.0000,,
4,"AD","","","",42.5000,1.5000,,
5,"AE","","","",24.0000,54.0000,,
6,"AF","","","",33.0000,65.0000,,
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestReadblock...")

	house, err := NewLocationhouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", house)

	//delete testfile
	os.Remove(testfile)
}
