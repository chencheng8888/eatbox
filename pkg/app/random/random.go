package random

import (
	"math/rand"
	"time"
)

// GenerateRandomNumberString 生成随机数字字符串
func GenerateRandomNumberString(length int) string {
	rand.Seed(time.Now().UnixNano()) // 使用当前时间作为随机数种子
	const digits = "0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}
