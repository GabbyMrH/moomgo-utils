/*
 * @PackageName: regexp
 * @Description: 常用正则匹配
 * @Author: Casso
 * @Date: 2022-01-28 15:13:29
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 16:10:38
 */

package regexp

import "regexp"

var (
	moblie_ch    = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	mobile_short = `^((0\d{2,3})-)(\d{7,8})(-(\d{3,}))?$`
	email        = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	name         = `^[a-z0-9_-]{3,16}$`
	pass         = `^[a-z0-9_-]{6,18}$`
)

// Check 入参：需要验证字符串，检验类型
func Check(val string, typ int64) bool {
	var rgx *regexp.Regexp
	switch typ {
	case 1: // 手机号
		rgx = regexp.MustCompile(moblie_ch)
	case 2: // 座机号
		rgx = regexp.MustCompile(mobile_short)
	case 3: // 电子邮箱
		rgx = regexp.MustCompile(email)
	case 4: // 用户名
		rgx = regexp.MustCompile(name)
	case 5: // 密码
		rgx = regexp.MustCompile(pass)
	}

	return rgx.MatchString(val)
}
