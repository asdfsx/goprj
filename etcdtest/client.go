package etcdtest

import (
	"fmt"
	"github.com/coreos/go-etcd/etcd"
)

type Connection struct {
	client *etcd.Client
}

func NewConnection(machines []string) *Connection {
	client := etcd.NewClient(machines)
	return &Connection{client}
}

func (conn *Connection) Get(Key string) (string, error) {
	response, err := conn.client.Get(Key, false, false)
	return response.Node.Value, err
}

func (conn *Connection) Set(Key string, Value string) (string, error) {
	response, err := conn.client.Set(Key, Value, 0)
	return response.Node.Value, err
}

func (conn *Connection) Delete(Key string) (string, error) {
	response, err := conn.client.Delete(Key, false)
	return response.Node.Value, err
}

func main() {
	fmt.Println("Test etcd")
}
