package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	arr_str := strings.Split(s, " ")
	res := map[string]int{}

	for _, v := range arr_str {
		res[v] += 1
	}

	return res
}

func main() {
	wc.Test(WordCount)
}
