/*
 * @PackageName: big
 * @Description: 大数操作
 * @Author: Casso
 * @Date: 2022-01-28 14:43:32
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 14:57:26
 */

package big

import (
	"fmt"
	"math/big"
)

var (
	seedLocal = "1000000000000000000" // 19位
)

// ParserPrice 用于小数点左移19位，返回 price * seed 的完整字符串
func ParserPrice(price float64, seed string) string {
	var a, _ = new(big.Rat).SetString(fmt.Sprintf("%f", price))

	// 使用包内默认值 19 位
	if seed == "" {
		seed = seedLocal
	}
	var b, _ = new(big.Rat).SetString(seed)

	// 相乘
	a.Mul(a, b)
	var c, _ = new(big.Rat).SetString(a.String())

	return c.RatString()
}
