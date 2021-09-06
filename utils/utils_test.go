package utils

import (
	"testing"
)

func TestClosestMatch(t *testing.T) {
	errText := "unable to pick out best match"

	tests := []struct {
		name    string
		names   []string
		best    string
		errText string
	}{
		{"a", []string{"a", "b"}, "a", ""},
		{"b", []string{"a", "b"}, "b", ""},
		{"a", []string{"a", "a"}, "", errText},
		{"foo", []string{"foobar", "barfoo"}, "", errText},
		{"foo", []string{"foobar1111", "barfoo1112"}, "", errText},
	}

	for i, tt := range tests {
		best, err := ClosestMatch(tt.name, tt.names)

		if err != nil && err.Error() != tt.errText {
			t.Fatalf("tests[%d]: wrong error. expected: %s. got: %s", i, tt.errText, err.Error())
		}

		if err == nil && best != tt.best {
			t.Fatalf("tests[%d]: wrong best. expected: %s. got: %s", i, tt.best, best)
		}
	}
}

func TestMinutesFromHMFormat(t *testing.T) {
	tests := []struct {
		input   string
		minutes int
		errText string
	}{
		{"12", 0, "incorrect format: 12. please use 1h23m"},
		{"m", 0, "incorrect format: zero time. please use 1h23m"},
		{"h", 0, "incorrect format: h. please use 1h23m"},
		{"2X3m", 0, "incorrect character: X. please use 1h23m"},
		{"12m3h", 0, "incorrect format: 12m3h. please use 1h23m"},
		{"1m2m", 0, "incorrect format: 1m2m. please use 1h23m"},
		{"1hh2m", 0, "incorrect format: 1hh2m. please use 1h23m"},
		{"21m", 21, ""},
		{"1h2m", 62, ""},
		{"0h3m", 3, ""},
		{"0h120m", 120, ""},
	}

	for i, tt := range tests {
		minutes, err := MinutesFromHMFormat(tt.input)

		if err == nil && len(tt.errText) > 0 {
			t.Fatalf("tests[%d]: missing error. expected: %s", i, tt.errText)
		}

		if err != nil && len(tt.errText) == 0 {
			t.Fatalf("test[%d]: got unexpected error message: %s", i, err.Error())
		}

		if err != nil {
			if err.Error() != tt.errText {
				t.Fatalf("tests[%d]: wrong error. expected: %s. got %s", i, tt.errText, err.Error())
			}
		} else if tt.minutes != minutes {
			t.Fatalf("tests[%d]: incorrect result. expected: %d. got %d", i, tt.minutes, minutes)
		}
	}
}
