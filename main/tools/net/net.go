package net

import (
	"io/ioutil"
	"net"
	"net/http"
)

const (
	DEFAULT_PORT = 8333
	PUBLISH_HOPS = 5 //hops（跳数）
)
const (
	NODE_NETWORK = (1 << 0)
)

func GetMyExternalIP() string {
	res, _ := http.Get("http://myexternalip.com/raw")

	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)

	return string(content)
}
func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		panic("")
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}

		}
	}
	return ""
}
func IsRoutable(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
