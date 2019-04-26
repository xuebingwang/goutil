package goutil

import (
	"math"
	"math/rand"
	"time"
    "strconv"
    "fmt"
)

// 随机数生成
// @Param	min 	int	最小值
// @Param 	max		int	最大值
// @return  int
func RandInt(min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn((max - min)) + min
	if num < min || num > max {
		RandInt(min, max)
	}
	return num
}

//小数四舍五入
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//小数仅保留2位
func Decimal(value float64) float64 {
    value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
    return value
}

//生成随机字符串
func RandString(len int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len; i++ {
		result = append(result, bytes[r.Intn(len)])
	}
	return string(result)
}
