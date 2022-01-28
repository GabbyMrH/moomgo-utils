/*
 * @PackageName: regexp
 * @Description: regexp testing
 * @Author: Casso
 * @Date: 2022-01-28 16:20:07
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 16:41:48
 */

package regexp

import "testing"

func TestCheck(t *testing.T) {
	mobiles := []string{
		"13456897894",
		"15623568956",
		"14523564578",
	}
	for _, v := range mobiles {
		if !Check(v, 1) {
			t.Error("mobile check fail")
		}
	}

	email := []string{
		"casso.@163.com",
		"12345622@qq.com",
	}
	for _, v := range email {
		if !Check(v, 2) {
			t.Error("mobile check fail")
		}
	}

	names := []string{
		"casscom",
		"12345622com",
	}
	for _, v := range names {
		if !Check(v, 3) {
			t.Error("mobile check fail")
		}
	}

	pass := []string{
		"cassosss",
		"12345622",
	}
	for _, v := range pass {
		if !Check(v, 4) {
			t.Error("mobile check fail")
		}
	}

	ips := []string{
		"192.168.1.12",
		"19.16.1.12",
	}
	for _, v := range ips {
		if !Check(v, 5) {
			t.Error("mobile check fail")
		}
	}
}
