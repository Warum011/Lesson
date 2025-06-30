package unpacking

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	ErrInvalidStringTrailingBackslash = errors.New(
		"invalid string: backslash at the end of the string",
	)

	ErrInvalidStringEscapeSequence = errors.New(
		"invalid string: invalid escape sequence",
	)

	ErrInvalidStringLeadingOrConsecutiveDigit = errors.New(
		"invalid string: string cannot start with a digit or a digit must not be consecutive",
	)

	ErrInvalidStringMultiDigitNumber = errors.New(
		"invalid string: a number consisting of more than one digit is encountered",
	)
)

func Unpack(input string) (string, error) {
	var builder strings.Builder
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length; {
		var fl rune // flag
		if runes[i] == '\\' {
			if i+1 >= length {
				return "", ErrInvalidStringTrailingBackslash
			}
			next := runes[i+1]
			if (next >= '0' && next <= '9') || next == '\\' {
				fl = next
				i += 2
			} else {
				return "", ErrInvalidStringEscapeSequence
			}
		} else {
			if runes[i] >= '0' && runes[i] <= '9' {
				return "", ErrInvalidStringLeadingOrConsecutiveDigit
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
			return "", ErrInvalidStringMultiDigitNumber
		}
		if repeatCount > 0 {
			builder.WriteString(strings.Repeat(string(fl), repeatCount))
		}
	}
	return builder.String(), nil

}

func Pack(input string) string {
	if input == "" {
		return ""
	}
	var builder strings.Builder
	runes := []rune(input)
	n := len(runes)
	count := 1

	escapeRune := func(r rune) string {
		if (r >= '0' && r <= '9') || r == '\\' {
			return "\\" + string(r)
		}
		return string(r)
	}

	current := runes[0]
	for i := 1; i < n; i++ {
		if runes[i] == current {
			count++
		} else {
			if count == 1 {
				builder.WriteString(escapeRune(current))
			} else if count > 1 && count < 10 {
				builder.WriteString(escapeRune(current))
				builder.WriteString(fmt.Sprintf("%d", count))
			} else {
				for j := 0; j < count; j++ {
					builder.WriteString(escapeRune(current))
				}
			}
			current = runes[i]
			count = 1
		}
	}
	if count == 1 {
		builder.WriteString(escapeRune(current))
	} else if count > 1 && count < 10 {
		builder.WriteString(escapeRune(current))
		builder.WriteString(fmt.Sprintf("%d", count))
	} else {
		for j := 0; j < count; j++ {
			builder.WriteString(escapeRune(current))
		}
	}
	return builder.String()
}

func RunUnpackDaemon() {
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

func RunPackDaemon() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Daemon Mode (Pack):")
	for {
		fmt.Print("Введите строку: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		input = strings.TrimRight(input, "\n")
		result := Pack(input)
		fmt.Println(result)
	}
}
