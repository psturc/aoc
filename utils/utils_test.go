package utils

import (
	"reflect"
	"testing"
)

func TestStrSliceToIntSlice(t *testing.T) {
	in := []string{"1", "2", "3"}
	expected := []int{1, 2, 3}

	out := StrSliceToIntSlice(in)

	if !reflect.DeepEqual(out, expected) {
		t.Fatalf("StrSliceToIntSlice() - in: %v, out: %v, expected: %v", in, out, expected)
	}
}
