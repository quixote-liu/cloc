package util

import "fmt"

func StringsContains(src []string, target string) bool {
	for _, s := range src {
		if target == s {
			return true
		}
	}
	return false
}

func PrintfErr(err error) {
	fmt.Printf("[ERROR]: %v\n", err)
}
