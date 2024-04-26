// Created by Petr Lozhkin
package calc_test

import (
	"github.com/im7mortal/airports/pkg/airports/calc"
	"github.com/im7mortal/airports/pkg/airports/calc/test_cases"
	"reflect"
	"testing"
)

func TestProcessFlights(t *testing.T) {

	for _, tc := range test_cases.Standard {
		t.Run(tc.Name, func(t *testing.T) {
			result, err := calc.ProcessFlights(tc.Input)
			if tc.ExpectedErr != nil {
				if tc.ExpectedErr != err {
					t.Fatalf("Expected error: %s", tc.ExpectedErr)
				}
			}
			if !reflect.DeepEqual(result, tc.Expected) {
				t.Fatalf("Expected: %v, but got: %v", tc.Expected, result)
			}
		})
	}
}
