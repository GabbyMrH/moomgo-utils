/*
 * @PackageName: convertor
 * @Description: 数据类型转换器
 * @Author: gabbymrh
 * @Date: 2021-10-08 12:57:23
 * @Last Modified by: gabbymrh
 * @Last Modified time: 2021-10-08 12:57:23
 */

package convertor

import "strconv"

// Int64ToString 将 int64 转换为 string
func Int64ToString(intNum int64) string {
	return strconv.FormatInt(intNum, 10)
}

// Float64ToString 将 float64 转换为 string
func Float64ToString(floatNum float64) string {
	return strconv.FormatFloat(floatNum, 'f', 10, 64)
}

// BoolToString 将布尔值转换为 string
func BoolToString(boolVal bool) string {
	return strconv.FormatBool(boolVal)
}

// StringToInt 将字符串转换为 int
func StringToInt(str string) (int, error) {
	i, err := strconv.Atoi(str)
	if err != nil {
		return i, err
	}
	return i, nil
}

// StringToBool 将字符串转换为 bool
func StringToBool(str string) (bool, error) {
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return b, nil
}
