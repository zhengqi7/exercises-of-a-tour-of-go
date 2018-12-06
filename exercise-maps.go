package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    array := strings.Fields(s)
	m := make(map[string]int)
    for i:=0; i<len(array); i++ {
	if _,ok := m[array[i]]; ok == false {
		m[array[i]] = 1;
	} else {
		m[array[i]] = m[array[i]] + 1;
	}
  }
	return m
}

func main() {
	wc.Test(WordCount)
}