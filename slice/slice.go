/*
 * @PackageName: slice
 * @Description: slice操作
 * @Author: Casso
 * @Date: 2022-01-28 15:17:17
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 17:03:55
 */

package slice

// IntInSlice int存在于切片返回true
func IntInSlice(num int, slice []int) (exists bool) {
	if len(slice) == 0 {
		return false
	}

	for i := range slice {
		if slice[i] == num {
			exists = true
			return exists
		}
	}

	return exists
}

// Int64InSlice int64存在于切片返回true
func Int64InSlice(num int64, slice []int64) (exists bool) {
	if len(slice) == 0 {
		return false
	}

	for i := range slice {
		if slice[i] == num {
			exists = true
			return exists
		}
	}

	return exists
}

// StrInSlice 判断字符串是否存在于切片中,存在则返回true
func StrInSlice(str string, slice []string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if str == v {
			return true
		}
	}

	return false
}

// removeRepByLoop 通过两重循环过滤重复元素
func removeRepByLoop(slc []int) []int {
	result := []int{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}

	return result
}

// removeRepByMap 通过map主键唯一的特性过滤重复元素
func removeRepByMap(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}

	return result
}

// RemoveRep 切片元素去重
func RemoveRep(slc []int) []int {
	if len(slc) < 1024 {
		// 切片长度小于1024的时候，循环来过滤
		return removeRepByLoop(slc)
	} else {
		// 大于的时候，通过map来过滤
		return removeRepByMap(slc)
	}
}

// JoinStrSlice 拼接切片字符串: ["cas" "so"] == > "casso"
func JoinStrSlice(s []string, slc string) (res string) {
	for i, v := range s {
		if i != len(s)-1 {
			res += v + slc
		} else {
			res += v
		}
	}
	return
}
