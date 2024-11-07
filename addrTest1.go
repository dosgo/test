package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/wlynxg/anet"
)

func main() {
	var subnets []string // 修改为切片以存储多个有效的IP段

	interfaces, err := anet.Interfaces()
	if err != nil {
		return // 修改返回值
	}
	for _, iface := range interfaces {
		if strings.Contains(iface.Name, "wlan") || strings.Contains(iface.Name, "ap") {
			// 获取该接口的所有地址
			addrs, err := anet.InterfaceAddrsByInterface(&iface)

			if err != nil {
				fmt.Printf("Error getting addresses for interface %s: %v\n", iface.Name, err)
				continue
			}

			// 遍历每个地址
			for _, addr := range addrs {
				// 检查地址类型
				switch v := addr.(type) {
				case *net.IPNet:

					subnets = append(subnets, v.String())

				case *net.IPAddr:
					fmt.Printf("  IP Address: %s\n", v.IP)
				default:
					fmt.Printf("  Unknown address type: %v\n", v)
				}
			}
		}
	}
	log.Printf("subnets:%+v\r\n", subnets)

}
