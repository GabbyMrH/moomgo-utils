/*
 * @PackageName: crypto
 * @Description: 常用随机/加密
 * @Author: Casso-Wong
 * @Date: 2022-01-30 13:59:03
 * @Last Modified by: Casso-Wong
 * @Last Modified time: 2022-01-30 13:59:03
 */
package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// RandStringRunes 获取n位随机字符
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// RandInt64 获取入参范围内随机数字（范围：1-100）
func RandInt64(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	if min >= max || min == 0 || max == 0 {
		min = 1
		max = 100
	}
	return rand.Int63n(max-min) + min
}

// RandStr 获取n位随机大写字母字符
func RandStr(n int) string {
	return strings.ToUpper(RandStringRunes(n))
}

// GetMD5Encode 返回一个32位md5加密后的字符串
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
