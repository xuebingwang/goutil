package goutil

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
