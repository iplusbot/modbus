package modbus

import (
	"testing"
	"time"
)

func TestRtuOverTcp(t *testing.T) {
	h := NewRtuTcpClientHandler("172.21.5.198:10000", 1)
	h.Logger = LoggerFunc(t.Logf)
	c := NewClient(h)
	for i := 0; i < 10; i++ {
		data, err := c.ReadDiscreteInputs(0, 8)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(data)
			time.Sleep(200 * time.Millisecond)
		}
	}
}
