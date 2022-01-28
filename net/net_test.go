/*
 * @PackageName: net
 * @Description: net testing
 * @Author: Casso
 * @Date: 2022-01-28 10:56:41
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 11:01:21
 */

package net

import "testing"

func TestLocalIP(t *testing.T) {
	if LocalIP() == "127.0.0.1" {
		t.Errorf("LocalIP err")
	}
}
