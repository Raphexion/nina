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
