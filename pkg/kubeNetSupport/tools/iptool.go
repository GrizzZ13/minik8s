package tools

import (
	"fmt"
	"minik8s/pkg/kubeNetSupport/netconfig"
	"net"
	"strings"
)

//获取内网ip
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

//比如, name可以为 ens33
func getIPv4ByInterface(name string) ([]string, error) {
	var ips []string

	iface, err := net.InterfaceByName(name)
	if err != nil {
		return nil, err
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}
func GetEns3IPv4Addr() string {
	val, _ := getIPv4ByInterface(netconfig.ETH_NAME)
	return val[0]
}
func GetDynamicIp() string {
	ips, _ := getIPv4ByInterface(netconfig.ETH_NAME)
	fmt.Println(ips)
	fmt.Println(netconfig.GlobalIpMap)
	return netconfig.GlobalIpMap[ips[0]]
}
func GetBasicIpAndMask(ipAndMask string) string {
	index := strings.Index(ipAndMask, ".")
	a := ipAndMask[:index]
	ipAndMask = ipAndMask[index+1:]
	index = strings.Index(ipAndMask, ".")
	b := ipAndMask[:index]
	return a + "." + b + ".0.0/16"
}
