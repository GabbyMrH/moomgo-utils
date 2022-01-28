/*
 * @PackageName: regexp
 * @Description: 常用正则匹配
 * @Author: Casso
 * @Date: 2022-01-28 15:13:29
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 16:41:24
 */

package regexp

import "regexp"

var (
	moblie_ch = `^1[3-9][0-9]{9}$`
	email     = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	name      = `^[a-z0-9_-]{3,16}$`
	pass      = `^[a-z0-9_-]{6,18}$`
	ip        = `(([01]{0,1}\d{0,1}\d|2[0-4]\d|25[0-5])\.){3}([01]{0,1}\d{0,1}\d|2[0-4]\d|25[0-5])`
)

// Check 入参：需要验证字符串，检验类型
func Check(val string, typ int64) bool {
	var rgx *regexp.Regexp
	switch typ {
	case 1: // 手机号
		rgx = regexp.MustCompile(moblie_ch)
	case 2: // 电子邮箱
		rgx = regexp.MustCompile(email)
	case 3: // 用户名
		rgx = regexp.MustCompile(name)
	case 4: // 密码
		rgx = regexp.MustCompile(pass)
	case 5: // ip地址
		rgx = regexp.MustCompile(ip)
	}

	return rgx.MatchString(val)
}
