/*
 * @PackageName: net
 * @Description: 常用net操作
 * @Author: Casso
 * @Date: 2022-01-28 10:54:49
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 15:00:16
 */

package net

import (
	"net"
	"net/http"
	"strings"
)

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

// GetClientIP 获取http请求客户端IP
func GetClientIP(req *http.Request) string {
	var ips []string
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		ips = strings.Split(ip, ",")
	}

	if len(ips) > 0 && ips[0] != "" {
		rip, _, err := net.SplitHostPort(ips[0])
		if err != nil {
			rip = ips[0]
		}
		return rip
	}

	if ip, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		return ip
	}

	return req.RemoteAddr
}
