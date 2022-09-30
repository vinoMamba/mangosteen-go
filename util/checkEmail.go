package util

import "regexp"

func CheckEmail(email string) bool {
	//匹配电子邮箱
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
