package calc

import (
	"reflect"
	"testing"
)

func TestProcessFlights(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]string
		expected []string
		err      error
	}{
		{"Test Case 1", [][]string{{"SFO", "EWR"}}, []string{"SFO", "EWR"}, nil},
		{"Test Case 2", [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}, []string{"SFO", "EWR"}, nil},
		{"Test Case 3", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}, []string{"SFO", "EWR"}, nil},

		// NO SOLUTION; sequence is invalid with GTO
		{"Test Case 4", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"SFO", "GTO"}, {"GSO", "IND"}, {"ATL", "GSO"}}, nil, InvalidSequence},

		// Loop. Imagine we have need from Paris to SaoPaolo. And we have cheap flight from Paris to Berlin and from Berlin to SaoPaolo with stop in Paris
		{"Test Case 5.1", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}, nil},
		{"Test Case 5.2", [][]string{{"SFO", "ATL"}, {"ATL", "BIR"}, {"BIR", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}, nil},

		// NO SOLUTION. It's infinite loop; We can't say where is origin and destination
		{"Test Case 6", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}}, nil, NotEnoughData},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ProcessFlights(tc.input)
			if tc.err != nil {
				if tc.err != err {
					t.Fatalf("Expected error: %s", tc.err)
				}
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
