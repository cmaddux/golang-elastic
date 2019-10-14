package encoding

// RunLength given a string will produce a slice representing
// run length encoding of the string. Each 2-item array in
// the output slice contains a character at index 0 and 
// the run length count of that character at index 1.
func RunLength(str string) [][2]interface{} {
	var memo rune
	var ct int

	var res [][2]interface{}
	for _, char := range str {
		if char == memo {
			ct++
			continue
		}

		if memo != '\x00' {
			res = append(res, [2]interface{}{memo, ct})
		}

		memo = char
		ct = 1
	}

	if memo != '\x00' {
		res = append(res, [2]interface{}{memo, ct})
	}

	return res
}
