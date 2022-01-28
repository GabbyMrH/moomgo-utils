/*
 * @PackageName: net
 * @Description: 常用net操作
 * @Author: Casso
 * @Date: 2022-01-28 10:54:49
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 11:05:39
 */

package net

import "net"

// LocalIP 获取本地网络IPV4
func LocalIP() string {
	var local = "127.0.0.1"
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return local
	}

	for _, v := range addrs {
		if ipnet, ok := v.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return local
}
