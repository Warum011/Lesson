package unpacking

import (
	"errors"
	"testing"
)

func TestUnpack(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{
			name:     "simple case with repeats",
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
			err:      nil,
		},
		{
			name:     "no repeats",
			input:    "abcd",
			expected: "abcd",
			err:      nil,
		},
		{
			name:  "string starting with digit",
			input: "3abc",
			err:   error3,
		},
		{
			name:  "digit-only at beginning",
			input: "45",
			err:   error3,
		},
		{
			name:  "more than one digit in multiplier",
			input: "aaa10b",
			err:   error4,
		},
		{
			name:     "zero multiplier",
			input:    "aaa0b",
			expected: "aab",
			err:      nil,
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
			err:      nil,
		},
		{
			name:     "newline repeat",
			input:    "d\n5abc",
			expected: "d\n\n\n\n\nabc",
			err:      nil,
		},
		{
			name:     "escape digit: qwe\\4\\5 => qwe45",
			input:    "qwe\\4\\5",
			expected: "qwe45",
			err:      nil,
		},
		{
			name:     "escape with multiplier: qwe\\45 => qwe44444",
			input:    "qwe\\45",
			expected: "qwe44444",
			err:      nil,
		},
		{
			name:  "invalid escape sequence: qw\\ne",
			input: "qw\\ne",
			err:   error2,
		},
		{
			name:  "backslash at end",
			input: "abc\\",
			err:   error1,
		},
		{
			name:  "digit following digit without escape",
			input: "a12",
			err:   error4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := Unpack(tc.input)
			if tc.err != nil {
				if err == nil {
					t.Errorf("error %q was expected for %q, but no errors were received", tc.input, tc.err)
				} else if !errors.Is(err, tc.err) && err.Error() != tc.err.Error() {
					t.Errorf("error %q was expected for %q, we got%q", tc.input, tc.err, err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for %q: %v", tc.input, err)
				}
				if result != tc.expected {
					t.Errorf("for %q we expected %q, we got %q", tc.input, tc.expected, result)
				}
			}
		})
	}
}
