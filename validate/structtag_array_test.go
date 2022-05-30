package validate

import (
	"testing"
	"time"
)

type testStructArray struct {
	A []int
	B []string
	C []bool
	D []float64
	E []time.Duration
	F []int64
	G []testSubArray `cfgRequired:"true"`
}

type testSubArray struct {
	A int        `cfg:"A" cfgDefault:"300"`
	B []string   `cfg:"C" cfgRequired:"true"`
	S []testSubArraySub
}

type testSubArraySub struct {
	A int      `cfg:"A" cfgDefault:"500"`
	B []string `cfg:"SOMENAME"`
}

// Test functions

func TestParseArray(t *testing.T) {

	Setup("cfg", "cfgDefault")

	s := &testStructArray{
		A: []int{ 1,2,3 },
		B: []string{ "a", "b" },
		G: []testSubArray{{
			A: 1,
			B: []string{ "", "a" },
			S: []testSubArraySub{{
				B: []string{ "test" },
			},},
		},},
	}

	err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}

	// Check value in the array
	if s.A[1] != 2 {
		t.Fatal("Expected 2 in the array, but got:", s.A)
	}

	// Test to check cfgRequired
	s = &testStructArray{}
	err = Parse(s)
	if err == nil { // TODO: support cfgRequired
		t.Fatal("expected error about required field but got nil")
	}

	// Test to check cfgRequired in SubArray
	s = &testStructArray{G: []testSubArray{{A: 1,}}}
	err = Parse(s)
	if err == nil { // TODO: support cfgRequired
		t.Fatal("expected error about required field but got nil")
	}

}
