package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"sync"
)

type socketServer struct {
	ipaddr   string
	listener net.Listener
	handler  func(net.Conn)
}

var (
	bufioReaderPool   sync.Pool
	bufioWriter2kPool sync.Pool
	bufioWriter4kPool sync.Pool
)

const noLimit int64 = (1 << 63) - 1

func bufioWriterPool(size int) *sync.Pool {
	switch size {
	case 2 << 10:
		return &bufioWriter2kPool
	case 4 << 10:
		return &bufioWriter4kPool
	}
	return nil
}

func newBufioReader(r io.Reader) *bufio.Reader {
	if v := bufioReaderPool.Get(); v != nil {
		br := v.(*bufio.Reader)
		br.Reset(r)
		return br
	}
	return bufio.NewReader(r)
}

func putBufioReader(br *bufio.Reader) {
	br.Reset(nil)
	bufioReaderPool.Put(br)
}

func NewSocketServer(ipaddr string) *socketServer {
	return &socketServer{ipaddr: ipaddr, handler: handlerFunc}
}

func (server *socketServer) Listen() error {
	listener, err := net.Listen("tcp", server.ipaddr)
	if err != nil {
		return err
	}
	server.listener = listener
	return nil
}

func (server *socketServer) Accept() (c net.Conn, err error) {
	c, err = server.listener.Accept()
	return
}

func (server *socketServer) Run() {
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go server.handler(conn)
	}
}

func handlerFunc(conn net.Conn) {
	defer conn.Close()
	reader := newBufioReader(io.LimitReader(conn, noLimit))
	result, err := reader.ReadString('\n')

	if err != nil {
		return
	}
	_, err = conn.Write([]byte(result))
	if err != nil {
		return
	}
}
