package net

import (
	"fmt"
	"net"
	"testing"
)

func Test_GetMyExternalIP(t *testing.T) {
	if GetMyExternalIP() == "" {
		t.Errorf("Get IP Failed")
	}
}

func Test_GetIntranetIp(t *testing.T) {

	if GetIntranetIp() == "" {
		t.Errorf("Get IP Failed")
	}
}
func Test_IsRoutable(t *testing.T) {
	fmt.Printf("%t\n", IsRoutable(net.ParseIP(GetMyExternalIP())))
}
