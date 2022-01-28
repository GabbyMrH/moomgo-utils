/*
 * @PackageName: big
 * @Description: big testing
 * @Author: Casso
 * @Date: 2022-01-28 14:48:21
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 14:55:13
 */

package big

import "testing"

func TestParserPrice(t *testing.T) {
	t.Log(ParserPrice(1.23648, "1000000000000000000")) // 1236480000000000000
	want := ParserPrice(1.23648, "1000000000000000000")
	got := "1236480000000000000"
	if want != got {
		t.Errorf("want: %s;got: %s", want, got)
	}
	t.Log(1000000000000000000 * 1.23648) // 1.23648e+18
}
