package main

import "fmt"

func contains(src []string, target string) bool {
	for _, v := range src {
		if v == target {
			return true
		}
	}
	return false
}

func serializeMap(m map[string]string) string {
	var res string
	for k, v := range m {
		res += "-" + k + " " + v
	}
	return res
}

func printfErr(err error) {
	fmt.Printf("[ERROR]: %v\n", err)
}
