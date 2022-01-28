/*
 * @PackageName: copy
 * @Description: 数据复制相关
 * @Author: Casso
 * @Date: 2022-01-28 11:11:23
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 13:58:53
 */

package copy

import "reflect"

// CopyStruct 复制两个结构体内相同字段
func CopyStruct(src, dst *reflect.Value) *reflect.Value {
	for i := 0; i < src.NumField(); i++ {
		filed := src.Type().Field(i).Name
		cfiled := dst.FieldByName(filed)
		// 可用的属性且可 赋值
		if cfiled.IsValid() && cfiled.CanSet() {
			dst.Field(i).Set(src.Field(i))
		}
	}

	return dst
}
