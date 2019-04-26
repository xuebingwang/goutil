package goutil

import (
	"sort"
)

//去掉数组数组中的重复值
func RemoveDuplicatesInt(a []int) (ret []int) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func MapSortByKey(list map[string]interface{}) []interface{} {
	var ssLice []string
	for key := range list {
		ssLice = append(ssLice, key)
	}
	sort.Strings(ssLice)

	var newList []interface{}
	//在将key输出
	for _, v := range ssLice {

		newList = append(newList,list[v])
	}
	return newList
}