package core

import (
	"fmt"
)

func FizzBuzz(int1, int2 int, str1, str2 string, limit int) []string {
	fb := make([]string, limit)
	for i := range fb {
		n := i + 1
		switch {
		case n%int1 == 0 && n%int2 == 0:
			fb[i] = str1 + str2
		case n%int1 == 0:
			fb[i] = str1
		case n%int2 == 0:
			fb[i] = str2
		default:
			fb[i] = fmt.Sprint(n)
		}
	}
	return fb
}
