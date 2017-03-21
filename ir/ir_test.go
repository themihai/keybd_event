package ir

import (
	"reflect"
	"testing"
)

func TestToKeys(t *testing.T) {
	var cases = []struct {
		in  string
		out []int
	}{

		{
			in:  "xYz",
			out: []int{0, 1},
		},
		{
			in:  "Hello World",
			out: []int{0, 1},
		},
	}

	for k, cs := range cases {
		keys, err := ToKeys(cs.in)
		if err != nil {
			t.Errorf("case %v: err %v", k, err)
		}
		if !reflect.DeepEqual(keys, cs.out) {
			t.Errorf("case %v: e %#v, r #%v", k, cs.out, keys)
		}
	}

}
