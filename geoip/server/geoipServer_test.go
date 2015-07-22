package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"runtime/debug"
	"testing"
	"time"
)

const (
	blockfilename = "testblock.txt"
	blockcontent  = `startip,endip,location
"123"	"123"	"123"
"a","b","c"
1	2	3
1,2,3,4
1,2,3,
`
	locationfilename = "testlocation.txt"
	locationcontent  = `Copyright (c) 2012 MaxMind LLC.  All Rights Reserved.
locId,country,region,city,postalCode,latitude,longitude,metroCode,areaCode
1,"O1","","","",0.0000,0.0000,,
2,"AP","","","",35.0000,105.0000,,
3,"EU","","","",47.0000,8.0000,,
4,"AD","","","",42.5000,1.5000,,
5,"AE","","","",24.0000,54.0000,,
6,"AF","","","",33.0000,65.0000,,
`

	location1content = `1	O1|O1(not set)|O1(not set)|0.0000|0.0000
2	Asia/Pacific Region|Asia/Pacific Region(not set)|Asia/Pacific Region(not set)|35.0000|105.0000
3	Europe|Europe(not set)|Europe(not set)|47.0000|8.0000
4	Andorra|Andorra(not set)|Andorra(not set)|42.5000|1.5000
5	United Arab Emirates|United Arab Emirates(not set)|United Arab Emirates(not set)|24.0000|54.0000
6	Afghanistan|Afghanistan(not set)|Afghanistan(not set)|33.0000|65.0000
7	Antigua and Barbuda|Antigua and Barbuda(not set)|Antigua and Barbuda(not set)|17.0500|-61.8000
`
)

func TestNewGeoipServer(t *testing.T) {
	defer func() {
		deletefile(blockfilename)
		deletefile(locationfilename)
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	createfile(blockfilename, blockcontent)
	createfile(locationfilename, locationcontent)

	t.Log("Starting TestNewGeoipServer...")

	server, err := NewGeoipServer(blockfilename, locationfilename)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Printf("%+v", server)
}

func TestGetLocation(t *testing.T) {
	defer func() {
		deletefile(blockfilename)
		deletefile(locationfilename)
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	createfile(blockfilename, blockcontent)
	createfile(locationfilename, location1content)

	t.Log("Starting TestNewGeoipServer...")

	server, err := NewGeoipServer(blockfilename, locationfilename)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Println("=============blocks:", server.bhouse)
	fmt.Println("=====================get location 1: ", server.GetLocation(1))
}

func TestHandlerFunc(t *testing.T) {
	defer func() {
		deletefile(blockfilename)
		deletefile(locationfilename)
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	createfile(blockfilename, blockcontent)
	createfile(locationfilename, location1content)

	t.Log("Starting TesthandlerFunc...")

	server, err := NewGeoipServer(blockfilename, locationfilename)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Println("=============blocks:", server.bhouse)
	fmt.Println("=====================get location 1: ", server.GetLocation(1))

	socketserver := NewSocketServer("0.0.0.0:12345")
	socketserver.Handler = server.HandlerSocket
	err = socketserver.Listen()
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%+v\n", server)

	go socketserver.Run()

	t.Log("Start connect SocketServer...")
	t.Log("Start connect SocketServer...")

	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		t.Error(err)
	}
	conn.Write([]byte("1\n"))
	fmt.Println("=======Write Finish")

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("----------------------result", string(result))
	time.Sleep(50)
	conn.Close()
}

func createfile(filename, content string) error {
	ostream, err := os.Create(filename)
	defer ostream.Close()
	if err != nil {
		return err
	}
	ostream.WriteString(content)
	return nil
}

func deletefile(filename string) error {
	return os.Remove(filename)
}
