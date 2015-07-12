package etcdtest

import (
	"testing"
)

func TestNewConnection(t *testing.T) {
	machines := []string{"http://127.0.0.1:2379"}
	connection := NewConnection(machines)
	connection.Set("/test", "test")
	connection.Get("/test")
	connection.Delete("/test")
}
