package geoip

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
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

	t.Log("Starting TestReadSort...")

	house, err := NewBlockhouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", house)

	house.Sort()
	fmt.Printf("after sort:\n%+v\n", house)

	//delete testfile
	os.Remove(testfile)
}

func TestBlockSearch(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	testfile := "testfile.txt"

	//create testfile
	content := `startip,endip,location
"1","3","1"
"4","6","2"
"7","9",3
10,12,4
13,15,5
16,18,6
19,21,7
`
	ostream, _ := os.Create(testfile)
	ostream.WriteString(content)
	ostream.Close()

	t.Log("Starting TestReadSearch...")

	house, err := NewBlockhouse(testfile)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}

	fmt.Printf("before sort:\n%+v", house)

	house.Sort()
	fmt.Printf("after sort:\n%+v\n", house)


	if location, ok := house.Search(1); ok {
		fmt.Printf("locationid searched from Blockhouse by ip addr %v: %v\n", 1, house.geoip_blocks[location])
	} else {
		fmt.Printf("locationid searched from Blockhouse by ip addr %v: not found:%v\n", 1, location)
	}



	//delete testfile
	os.Remove(testfile)
}
