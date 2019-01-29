package goutil

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)
// Empty 判断一个值是否为空(0, "", false, 空数组等)。
// []string{""}空数组里套一个空字符串，不会被判断为空。
func Empty(expr interface{}) bool {
	if expr == nil {
		return true
	}

	switch v := expr.(type) {
	case bool:
		return !v
	case int:
		return 0 == v
	case int8:
		return 0 == v
	case int16:
		return 0 == v
	case int32:
		return 0 == v
	case int64:
		return 0 == v
	case uint:
		return 0 == v
	case uint8:
		return 0 == v
	case uint16:
		return 0 == v
	case uint32:
		return 0 == v
	case uint64:
		return 0 == v
	case string:
		return len(v) == 0
	case float32:
		return 0 == v
	case float64:
		return 0 == v
	case time.Time:
		return v.IsZero()
	case *time.Time:
		return v.IsZero()
	}

	// 符合 IsNil 条件的，都为 Empty
	if IsNil(expr) {
		return true
	}

	// 长度为 0 的数组也是 empty
	v := reflect.ValueOf(expr)
	switch v.Kind() {
	case reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return 0 == v.Len()
	}

	return false
}

// IsNil 判断一个值是否为 nil。
// 当特定类型的变量，已经声明，但还未赋值时，也将返回 true
func IsNil(expr interface{}) bool {
	if nil == expr {
		return true
	}

	v := reflect.ValueOf(expr)
	k := v.Kind()

	return k >= reflect.Chan && k <= reflect.Slice && v.IsNil()
}

// IsEqual 判断两个值是否相等。
//
// 除了通过 reflect.DeepEqual() 判断值是否相等之外，一些类似
// 可转换的数值也能正确判断，比如以下值也将会被判断为相等：
//  int8(5)                     == int(5)
//  []int{1,2}                  == []int8{1,2}
//  []int{1,2}                  == [2]int8{1,2}
//  []int{1,2}                  == []float32{1,2}
//  map[string]int{"1":"2":2}   == map[string]int8{"1":1,"2":2}
//
//  // map 的键值不同，即使可相互转换也判断不相等。
//  map[int]int{1:1,2:2}        != map[int8]int{1:1,2:2}
func IsEqual(v1, v2 interface{}) bool {
	if reflect.DeepEqual(v1, v2) {
		return true
	}

	vv1 := reflect.ValueOf(v1)
	vv2 := reflect.ValueOf(v2)

	// NOTE: 这里返回 false，而不是 true
	if !vv1.IsValid() || !vv2.IsValid() {
		return false
	}

	if vv1 == vv2 {
		return true
	}

	vv1Type := vv1.Type()
	vv2Type := vv2.Type()

	// 过滤掉已经在 reflect.DeepEqual() 进行处理的类型
	switch vv1Type.Kind() {
	case reflect.Struct, reflect.Ptr, reflect.Func, reflect.Interface:
		return false
	case reflect.Slice, reflect.Array:
		// vv2.Kind() 与 vv1 的不相同
		if vv2.Kind() != reflect.Slice && vv2.Kind() != reflect.Array {
			// 虽然类型不同，但可以相互转换成 vv1 的，如：vv2 是 string，vv2 是 []byte，
			if vv2Type.ConvertibleTo(vv1Type) {
				return IsEqual(vv1.Interface(), vv2.Convert(vv1Type).Interface())
			}
			return false
		}

		// reflect.DeepEqual() 未考虑类型不同但是类型可转换的情况，比如：
		// []int{8,9} == []int8{8,9}，此处重新对 slice 和 array 做比较处理。
		if vv1.Len() != vv2.Len() {
			return false
		}

		for i := 0; i < vv1.Len(); i++ {
			if !IsEqual(vv1.Index(i).Interface(), vv2.Index(i).Interface()) {
				return false
			}
		}
		return true // for 中所有的值比较都相等，返回 true
	case reflect.Map:
		if vv2.Kind() != reflect.Map {
			return false
		}

		if vv1.IsNil() != vv2.IsNil() {
			return false
		}
		if vv1.Len() != vv2.Len() {
			return false
		}
		if vv1.Pointer() == vv2.Pointer() {
			return true
		}

		// 两个 map 的键名类型不同
		if vv2Type.Key().Kind() != vv1Type.Key().Kind() {
			return false
		}

		for _, index := range vv1.MapKeys() {
			vv2Index := vv2.MapIndex(index)
			if !vv2Index.IsValid() {
				return false
			}

			if !IsEqual(vv1.MapIndex(index).Interface(), vv2Index.Interface()) {
				return false
			}
		}
		return true // for 中所有的值比较都相等，返回 true
	case reflect.String:
		if vv2.Kind() == reflect.String {
			return vv1.String() == vv2.String()
		}
		if vv2Type.ConvertibleTo(vv1Type) { // 考虑 v1 是 string，v2 是 []byte 的情况
			return IsEqual(vv1.Interface(), vv2.Convert(vv1Type).Interface())
		}

		return false
	}

	if vv1Type.ConvertibleTo(vv2Type) {
		return vv2.Interface() == vv1.Convert(vv2Type).Interface()
	} else if vv2Type.ConvertibleTo(vv1Type) {
		return vv1.Interface() == vv2.Convert(vv1Type).Interface()
	}

	return false
}

//通过interface{}获取字符串
func GetString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

//通过interface{}获取数值型数据
//此获取比较灵活，转换规则如下
//1、如果接收数据为浮点string，则返回浮点数的整数部分，如果是整型string，则返回整数，如果是纯字符串，则返回0
//2、如果接收数据是float型，则返回float的整数部分
//3、如果接收数据是int, int32, int64型，则返回int
func GetInt(val interface{}) int {
	switch v := val.(type) {
	case int:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			fval, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return 0
			}
			return int(fval)
		}
		return int(n)
	case float32:
		return int(v)
	case float64:
		return int(v)
	default:
		return 0
	}
}

//通过interface{}获取小数型数据
//此获取比较灵活，转换规则如下
//1、如果接收数据为浮点string，则将字符串转换为浮点数
//2、如果接收数据是float型，则返回float数据
//3、如果接收数据是int, int32, int64型，则转义成float类型
//4、返回的数据结果统一为float64
func GetFloat(val interface{}) float64 {
	switch v := val.(type) {
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case int32:
		return float64(v)
	case float64:
		return v
	case float32:
		return float64(v)
	case string:
		result, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0
		}
		return result
	}
	return 0
}

//字符串驼峰转蛇形
func HumpToSnake(word string) (outWord string) {
	chars := make([]byte, 0)
	for i, v := range word {
		asciiNum := v
		if asciiNum >= 65 && asciiNum <= 90 {
			if i != 0 {
				chars = append(chars, '_')
			}
			chars = append(chars, byte(rune(asciiNum+32)))
		} else {
			chars = append(chars, byte(asciiNum))
		}
	}
	outWord = string(chars)
	return
}

//sturct to map
func StructToMap(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[HumpToSnake(obj1.Field(i).Name)] = obj2.Field(i).Interface()
	}
	return data
}
