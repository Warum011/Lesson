package unpacking

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	error1 = errors.New("invalid string: backslash at the end of the string")
	error2 = errors.New("invalid string: invalid escape sequence")
	error3 = errors.New("invalid string: string cannot start with a digit or a digit must not be consecutive")
	error4 = errors.New("invalid string: a number consisting of more than one digit is encountered")
)

func Unpack(input string) (string, error) {
	var builder strings.Builder
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length; {
		var fl rune // flag
		if runes[i] == '\\' {
			if i+1 >= length {
				return "", error1
			}
			next := runes[i+1]
			if (next >= '0' && next <= '9') || next == '\\' {
				fl = next
				i += 2
			} else {
				return "", error2
			}
		} else {
			if runes[i] >= '0' && runes[i] <= '9' {
				return "", error3
			}
			fl = runes[i]
			i++
		}
		repeatCount := 1
		if i < length && (runes[i] >= '0' && runes[i] <= '9') {
			repeatCount = int(runes[i] - '0')
			i++
		}
		if i < length && (runes[i] >= '0' && runes[i] <= '9') {
			return "", error4
		}
		if repeatCount > 0 {
			builder.WriteString(strings.Repeat(string(fl), repeatCount))
		}
	}
	return builder.String(), nil

}

func RunDaemon() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Daemon Mode:")
	for {
		fmt.Print("Введите строку:")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		input = strings.TrimRight(input, "\n")
		result, err := Unpack(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		} else {
			fmt.Println(result)
		}
	}
}
