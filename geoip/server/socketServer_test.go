package server

import (
	"fmt"
	"io/ioutil"
	"net"
	"runtime/debug"
	"testing"
	"time"
)

func TestSocketServer(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			t.Errorf("Fatal Error:%s\n", err)
		}
	}()

	t.Log("Starting TestSocketServer...")

	server := NewSocketServer("0.0.0.0:12345")
	err := server.Listen()
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("%+v\n", server)
	go server.Run()

	t.Log("Start connect SocketServer...")

	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		t.Error(err)
	}
	conn.Write([]byte("test\n"))
	fmt.Println("Write Finish")

	result, err := ioutil.ReadAll(conn)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(result))
	time.Sleep(50)
}
