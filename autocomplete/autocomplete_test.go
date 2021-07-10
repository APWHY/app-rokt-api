package autocomplete

import (
	"testing"
)

func TestIsInitialSubstring(t *testing.T) {
	type tester struct {
		inSub  string
		inMain string
		ex     bool
	}
	ins := []*tester{
		{
			inSub:  "set",
			inMain: "setup",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "SETUP",
			ex:     true,
		},
		{
			inSub:  "",
			inMain: "",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "set",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "etsset",
			ex:     false,
		},
		{
			inSub:  "set",
			inMain: "se",
			ex:     false,
		},
	}

	for _, in := range ins {
		got := isInitialSubstring(in.inSub, in.inMain)
		if got != in.ex {
			t.Errorf("Expected %v, got %v", in.ex, got)
		}

	}

}

func TestIsSubstringOf(t *testing.T) {
	type tester struct {
		inSub  string
		inMain string
		ex     bool
	}
	ins := []*tester{
		{
			inSub:  "set",
			inMain: "setup",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "SETUP",
			ex:     true,
		},
		{
			inSub:  "",
			inMain: "",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "set",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "etsset",
			ex:     true,
		},
		{
			inSub:  "set",
			inMain: "se",
			ex:     false,
		},
	}

	for _, in := range ins {
		got := isSubstringOf(in.inSub, in.inMain)
		if got != in.ex {
			t.Errorf("Expected %v, got %v", in.ex, got)
		}

	}

}
