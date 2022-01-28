/*
 * @PackageName: time
 * @Description: time testing
 * @Author: Casso
 * @Date: 2022-01-28 11:12:03
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 14:34:28
 */

package time

import "testing"

func TestGetBetweenDates(t *testing.T) {
	res := GetBetweenDates("2021-10-15", "2021-10-20")
	for _, v := range res {
		t.Log(v)
	}
}
