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
