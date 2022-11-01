package main

import (
	"fmt"
	"sort"
	"strings"
)

func Bitw(a, b string) string {
	result := ""
	for k, _ := range a {
		if string(a[k]) == "1" || string(b[k]) == "1" {
			result += "1"
		} else if string(a[k]) == "0" && string(b[k]) == "0" {
			result += "0"
		}
	}

	return result
}

func Count(a string) int {
	result := 0
	for _, v := range a {
		if string(v) == "1" {
			result += 1
		}
	}
	return result
}

func main() {
	satu := "00101"
	dua := "00011"
	tiga := Bitw(satu, dua)
	con := Count(tiga)
	fmt.Println(satu)
	fmt.Println(dua)
	fmt.Println(tiga)
	fmt.Println(con)
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//s///////
