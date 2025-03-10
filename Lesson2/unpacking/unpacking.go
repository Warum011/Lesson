package unpacking

import (
	"errors"
	"unicode"
)

func UnpackString(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	runes := []rune(str)
	var res []rune
	var counts []int
	flag := false

	for index, value := range runes {
		if unicode.IsDigit(value) {
			if index == 0 {
				return "", errors.New("line starts with a number")
			}

			if flag {
				return "", errors.New("consecutive digits are not allowed")
			}

			digit := int(value - '0')

			if len(res) == 0 {
				return "", errors.New("there is no symbol to repeat before the number")
			}

			if digit > 0 {
				counts[len(counts)-1] = digit
			} else {
				res = res[:len(res)-1]
				counts = counts[:len(counts)-1]
			}

			flag = true
		} else {
			res = append(res, value)
			counts = append(counts, 1)
			flag = false
		}
	}

	if flag {
		return "", errors.New("string cannot end with a digit")
	}

	var output []rune
	for i, r := range res {
		count := counts[i]
		for j := 0; j < count; j++ {
			output = append(output, r)
		}
	}

	return string(output), nil
}
