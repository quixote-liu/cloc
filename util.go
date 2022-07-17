package main

func contains(src []string, target string) bool {
	for _, v := range src {
		if v == target {
			return true
		}
	}
	return false
}
