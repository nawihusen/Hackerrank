package main

import (
	"sort"
	"strings"
)

func main() {
	// fmt.Println(60494 + 17284)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
