package strings

import "github.com/cmaddux/string_manipulation/math"

// CountSpecial counts the number of "special" strings
// given the run length encoding of a string. The 2
// types of special strings are:
//     1. Any combination of repeated values (incl. a single value)
//        for example 'aaa' has 6 special strings: 3 x 'a', 
//        2 x 'aa' and 1 x 'aaa' 
//     2. Any repeated value of equal length with a single value 
//        at the middle point for example 'aba' or 'aabaa'
func CountSpecial(rlencoded [][2]interface{}) int {
	ct := 0
	for idx, item := range rlencoded {
		charct := item[1].(int)

		ct += (charct * (charct + 1)) / 2
		if ((idx > 0) && (idx < (len(rlencoded) - 1)) && (charct == 1)) {
			prev := rlencoded[idx - 1]
			prevchar := prev[0].(rune)
			prevct := prev[1].(int)

			next := rlencoded[idx + 1]
			nextchar := next[0].(rune)
			nextct := next[1].(int)

			if prevchar == nextchar {
				ct += math.MinInt(prevct, nextct)
			}
		}

	}

	return ct
}
