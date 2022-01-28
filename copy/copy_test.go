/*
 * @PackageName: copy
 * @Description: copy testing
 * @Author: Casso
 * @Date: 2022-01-28 11:13:22
 * @LastModifiedBy: Casso
 * @LastEditTime: 2022-01-28 11:57:23
 */

package copy

import (
	"reflect"
	"testing"
)

func TestCopyStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int64
	}

	type Stu struct {
		Name string
		Age  int64
	}

	var A = Person{
		Name: "casso",
		Age:  20,
	}
	var B Stu
	dst := reflect.ValueOf(&A).Elem()
	src := reflect.ValueOf(&B).Elem()
	B2 := CopyStruct(&dst, &src).Interface().(Stu)

	if A.Name != B2.Name || A.Age != B2.Age {
		t.Errorf("want: %#v; got: %#v", A, B)
	}
}
