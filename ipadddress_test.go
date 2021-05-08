package ipaddress

import (
	"fmt"
	"testing"
)

func TestOutbound(t *testing.T) {
	ip, err := Outbound()
	if err != nil {
		t.Fatal(err)
	}
	if ip == "" {
		t.Fatal("address is empty")
	}
	fmt.Println(ip)
}

func TestPublic(t *testing.T) {
	ip, err := Public()
	if err != nil {
		t.Fatal(err)
	}
	if ip == "" {
		t.Fatal("address is empty")
	}
	fmt.Println(ip)
}
