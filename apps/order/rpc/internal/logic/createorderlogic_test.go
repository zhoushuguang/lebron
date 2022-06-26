package logic

import (
	"testing"
	"time"
)

func TestGenOrderID(t *testing.T) {
	oid := genOrderID(time.Now())
	if len(oid) != 24 {
		t.Failed()
	} else {
		t.Log(oid)
	}
}
