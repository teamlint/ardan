package pkg

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}

// UpperFirst 字符串首字母转换为大写
func UpperFirst(s string) string {
	if len(s) < 2 {
		return strings.ToUpper(s)
	}
	for _, c := range s {
		sc := string(c)
		return strings.ToUpper(sc) + s[len(sc):]
	}
	return ""
}

// LowerFirst 字符串首字母转换为小写
func LowerFirst(s string) string {
	if len(s) < 2 {
		return strings.ToLower(s)
	}
	for _, c := range s {
		sc := string(c)
		return strings.ToLower(sc) + s[len(sc):]
	}
	return ""
}

func Lower(s string) string {
	return strings.ToLower(s)
}
