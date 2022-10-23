package main

import (
	"math"
	"sort"
)

func appendAndDelete(s string, t string, k int32) string {
	// Write your code here
	result := ""
	temp := ""
	if len(s) < len(t) {
		for k, v := range s {
			if string(v) == string(t[k]) {
				temp += string(v)
			} else if string(v) != string(t[k]) {
				break
			}
		}

		min := int32(math.Abs(float64(len(s) - len(temp))))
		plus := int32(math.Abs(float64(len(temp) - len(t))))
		total := min + plus
		if (k-total)%2 == 0 {
			return "Yes"
		} else if (k-total)%2 != 0 {
			return "No"
		}

	}

	for k, v := range t {
		if string(v) == string(s[k]) {
			temp += string(v)
		} else if string(v) != string(s[k]) {
			break
		}
	}
	min := int32(math.Abs(float64(len(s) - len(temp))))
	plus := int32(math.Abs(float64(len(temp) - len(t))))
	total := min + plus
	if k-total >= 0 {
		return "Yes"
	} else if k-total < 0 {
		return "No"
	}
	return result
}

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	// Write your code here
	var cost int32 = 0
	var dc int32 = (d1 - d2) * 15
	var mc int32 = (m1 - m2) * 500
	var yc int32 = (y1 - y2) * 10000

	if y1 > y2 {
		// if m1 > m2{
		//     if d1 > d2 {
		//         cost = yc + mc + dc
		//     }else if d1 <= d2{
		//         cost = yc + mc
		//     }
		// }else if m1 == m2 {
		//     if d1 > d2 {
		//         cost = yc + dc
		//     }else if d1 <= d2{
		//         cost = yc
		//     }
		// }else if m1 < m2{
		cost = yc
		// }
	} else if y1 == y2 {
		if m1 > m2 {
			// if d1 > d2 {
			// cost = mc + dc
			// }else if d1 <= d2{
			cost = mc
			// }
		} else if m1 == m2 {
			if d1 > d2 {
				cost = dc
			} else if d1 <= d2 {
				cost = 0
			}
		} else if m1 < m2 {
			cost = 0
		}
	} else if y1 < y2 {
		cost = 0
	}

	return cost
}

func repeatedString(s string, n int64) int64 {
	// Write your code here
	p := int64(len(s))
	count := int64(n / p)
	sisa := int64(n % p)
	result := int64(0)
	do := false
	for _, v := range s {
		if string(v) == "a" {
			do = true
			break
		}
	}

	if do {
		if len(s) == 1 {
			return n
		}
		sum := 0 // jumlah a pada s
		for _, v := range s {
			if string(v) == "a" {
				sum += 1
			}
		}

		result = count * int64(sum)
		if sisa > 0 {
			slisisa := s[:sisa]
			for _, v := range slisisa {
				if string(v) == "a" {
					result += 1
				}
			}
		}
		return result
	} else {
		return 0
	}
}

func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	steps := int32(0)
	for i := 0; i != len(c)-1; {

		i += 2

		if i >= len(c) {
			steps += 1
			break
		} else if c[i] == 1 {
			i -= 1
		}

		steps += 1
		// fmt.Println("i",i)
		// fmt.Println("steps",steps)
	}

	return steps
}

func equalizeArray(arr []int32) int32 {
	// Write your code here
	sumall := map[int32]int32{}
	keys := []int32{}

	for _, v := range arr {
		_, isExist := sumall[v]
		if !isExist {
			keys = append(keys, v)
		}
		sumall[v] += 1
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return sumall[keys[i]] > sumall[keys[j]]
	})

	total := int32(0)
	for k, v1 := range keys {
		if k == 0 {
			continue
		}

		total += sumall[v1]
	}

	return total
}

func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {
	// Write your code here
	costarr := []int64{}
	tob := (int64(b+w) * int64(bc)) + (int64(w) * int64(z))
	tow := (int64(b+w) * int64(wc)) + (int64(b) * int64(z))
	o := (int64(b) * int64(bc)) + (int64(w) * int64(wc))

	costarr = append(costarr, tob, tow, o)

	sort.SliceStable(costarr, func(i, j int) bool {
		return costarr[i] < costarr[j]
	})

	return costarr[0]
}
