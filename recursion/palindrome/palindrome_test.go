package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"anna", true},
		{"level", true},
		{"", true},
		{"a", true},
		{"hello", false},
	}

	for _, c := range cases {
		got := IsPalindrome(c.in)
		if got != c.want {
			t.Errorf("Palindrome(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
