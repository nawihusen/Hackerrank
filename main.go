package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	test := [][]int{{1, 1}, {2, 2}}
	fmt.Println(len(test))
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//s////
