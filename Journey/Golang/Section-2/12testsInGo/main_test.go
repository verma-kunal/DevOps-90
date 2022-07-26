package main

import "testing"

// Efficient Way of Testing:

// slice of structs:
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32 // what do we expect as a result from the division
	isErr    bool    //  is an error expected
}{
	// {"test_name", dividend, divisor, expected, isErr}
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0, 0, true},
	{"expect-5", 10.0, 2.0, 5.0, false},
}

func TestDivision(t *testing.T) {
	// ranging over the above created tests:
	for _, tt := range tests {
		gotFromFunction, err := divide(tt.dividend, tt.divisor)

		// running our checks:
		if tt.isErr { // if the test is suppose to throw an error
			if err == nil { // then we fail the test
				t.Error("Expected an error but did not get one!")
			}
		} else {
			if err != nil { 
				t.Error("Did not expect an error but got one", err.Error())
			}
		}
		if gotFromFunction != tt.expected{
			t.Errorf("Expected %f but got %f", tt.expected, gotFromFunction)
		}
	}
}

// ---------------------------------------------
// Manual way of testing out:

// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)
// 	if err != nil {
// 		t.Error("Got an error when we should not have!")
// 	}
// }

// // testing for bad division:

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)
// 	if err == nil {
// 		t.Error("Did not get an error when we should have!")
// 	}
// }
