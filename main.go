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
	fmt.Println(aku)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
