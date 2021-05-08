package ipaddress

import (
	"testing"
)

func TestOutbound(t *testing.T) {
	ip, err := Outbound()
	if err != nil {
		t.Fatal(err)
	}
	if ip.String() == "" {
		t.Fatal("address is empty")
	}
}
