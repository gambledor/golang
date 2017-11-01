// Package binarysearch provides ...
package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct {
		key   string
		slice []string
		want  int
	}{
		{
			key:   "baz",
			slice: []string{"bar", "baz", "foo", "foobar", "foobaz"},
			want:  1,
		},
		{
			key:   "foobar",
			slice: []string{"bar", "baz", "foo", "foobar", "foobaz"},
			want:  3,
		},
		{
			key:   "rab",
			slice: []string{"bar", "baz", "foo", "foobar", "foobaz"},
			want:  -1,
		},
	}

	for _, c := range cases {
		t.Logf("Testing binary search expecting %d", c.want)
		got := FindStringInSortedSlice(c.key, c.slice)
		if got != c.want {
			t.Errorf("BinarySearch(%s) got %d, want %d", c.key, got, c.want)
		}
	}
}
