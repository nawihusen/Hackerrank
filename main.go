package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	tes := "sdfghjmk,lkjhgf"
	t := strings.Split(tes, "")
	sort.Strings(t)
	fmt.Println(t)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
