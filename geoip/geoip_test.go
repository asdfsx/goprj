package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"sort"
)

func TestReadblock(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `startip,endip,location
"123","123","123"
"a","b","c"
1,2,3
1,2,3,4
1,2,3,
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestReadblock...")

	house, err := NewBlockhouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", house)

	//delete testfile
	os.Remove(testfile)
}

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

func TestBlockSort(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `startip,endip,location
"123","123","123"
"a","b","c"
"4","5",3
4,6,3
4,8,3
1,2,3,4
1,2,3,
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestReadblock...")

	house, err := NewBlockhouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", house)

	sort.Sort(house)
	fmt.Printf("after sort:\n%+v", house)

	//delete testfile
	os.Remove(testfile)
}
