package interation

import "strings"

func Repeat(s string, count int) string {
	var res strings.Builder
	for i := 0; i < count; i++ {
		res.WriteString(s)
	}
	return res.String()
}
