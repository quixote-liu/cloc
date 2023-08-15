package main

import "fmt"

func stringsContains(src []string, target string) bool {
	for _, s := range src {
		if target == s {
			return true
		}
	}
	return false
}

func printfErr(err error) {
	fmt.Printf("[ERROR]: %v\n", err)
}

