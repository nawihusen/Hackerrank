package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	tes := 10

	fmt.Println(tes / 3)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
