package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	aku := "asvnfkjjvdiljmdklc"
	catch := SortString(aku)
	fmt.Println(catch)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	return strings.Join(s, "")
}
