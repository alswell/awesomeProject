package main

import (
	"fmt"
	"net"
)

func GetIPv4(NIC string) string {
	interfaces, _ := net.Interfaces()
	for _, i := range interfaces {
		if i.Name == NIC {
			addresses, _ := i.Addrs()
			for _, addr := range addresses {
				if v, ok := addr.(*net.IPNet); ok {
					if v.IP.To4() != nil {
						return v.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetIPv4("docker0"))
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//
	//for _, address := range addrs {
	//
	//	// 检查ip地址判断是否回环地址
	//	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//		if ipnet.IP.To4() != nil {
	//			fmt.Println("ip:", ipnet.IP.String())
	//		}
	//
	//	}
	//}
}
