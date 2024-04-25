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
	}{
		{"Test Case 1", [][]string{{"SFO", "EWR"}}, []string{"SFO", "EWR"}},
		{"Test Case 2", [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}, []string{"SFO", "EWR"}},
		{"Test Case 3", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}, []string{"SFO", "EWR"}},

		// NO SOLUTION; sequence is invalid with GTO
		//{"Test Case 4", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"SFO", "GTO"}, {"GSO", "IND"}, {"ATL", "GSO"}}, []string{"SFO", "EWR"}},

		// Loop. Imagine we have need from Paris to SaoPaolo. And we have cheap flight from Paris to Berlin and from Berlin to SaoPaolo with stop in Paris
		{"Test Case 4.1", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}},
		{"Test Case 4.2", [][]string{{"SFO", "ATL"}, {"ATL", "BIR"}, {"BIR", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}},

		// NO SOLUTION. It's infinite loop; We can't say where is origin and destination
		//{"Test Case 4", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}}, []string{}},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ProcessFlights(tc.input)
			if err != nil {
				t.Fatalf("Error processing flights: %v", err)
			}
			if !reflect.DeepEqual(result, tc.expected) {
				t.Fatalf("Expected: %v, but got: %v", tc.expected, result)
			}
		})
	}
}
