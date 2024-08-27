package util

import (
	"math/rand"
	"strconv"
	"time"
)

// 生成指定长度的随机数字字符串
func RandomNumeric(size int) string {
	// 创建一个随机数生成器，使用当前时间的纳秒数作为种子
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 如果指定长度小于等于0，则抛出异常
	if size <= 0 {
		panic("{ size : " + strconv.Itoa(size) + " } must be more than 0 ")
	}
	// 初始化一个空字符串
	value := ""
	// 循环指定长度次
	for index := 0; index < size; index++ {
		// 每次循环生成一个0-9之间的随机数，并将其转换为字符串，添加到value中
		value += strconv.Itoa(r.Intn(10))
	}

	// 返回生成的随机数字字符串
	return value
}
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, t.Location())
}
