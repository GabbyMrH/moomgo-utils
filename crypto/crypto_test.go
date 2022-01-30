/*
 * @PackageName: crypto
 * @Description: crypto testing
 * @Author: Casso-Wong
 * @Date: 2022-01-30 14:03:48
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2022-01-30 14:03:48
 */

package crypto

import "testing"

func TestRandStringRunes(t *testing.T) {
	t.Log(RandStringRunes(5), RandStringRunes(8))
}

func TestRandInt64(t *testing.T) {
	t.Log(RandInt64(2, 5), RandInt64(100, 90))
}

func TestRandStr(t *testing.T) {
	t.Log(RandStr(2), RandStr(8))
}
