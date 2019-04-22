package state

import "testing"

func TestState(t *testing.T) {
	o := NewOrderInfo(0)
	o.Refund()
}
