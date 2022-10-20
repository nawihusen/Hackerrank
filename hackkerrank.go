package main

import "math"

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
