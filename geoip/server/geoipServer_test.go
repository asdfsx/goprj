package server

import (
	"fmt"
	"os"
	"runtime/debug"
	"testing"
	"net"
	"io/ioutil"
	"time"
)

const (
	blockfilename = "testblock.txt"
	blockcontent  = `startip,endip,location
"123","123","123"
"a","b","c"
1,2,3
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
	createfile(locationfilename, locationcontent)

	t.Log("Starting TestNewGeoipServer...")

	server, err := NewGeoipServer(blockfilename, locationfilename)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Println("=============blocks:", server.bhouse)
	fmt.Println("=====================get location 123: ",server.GetLocation(123))
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
	createfile(locationfilename, locationcontent)

	t.Log("Starting TesthandlerFunc...")

	server, err := NewGeoipServer(blockfilename, locationfilename)
	if err != nil {
		t.Errorf("Fatal Error:%s\n", err)
	}
	fmt.Println("=============blocks:", server.bhouse)
	fmt.Println("=====================get location 123: ",server.GetLocation(123))

	socketserver := NewSocketServer("0.0.0.0:12345")
	socketserver.handler = server.handlerFunc
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
	conn.Write([]byte("123\n"))
	fmt.Println("=======Write Finish")

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("----------------------result",string(result))
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
