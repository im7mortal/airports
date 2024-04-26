// Created by Petr Lozhkin
package calc_test

import (
	"github.com/im7mortal/airports/pkg/airports/calc"
	"github.com/im7mortal/airports/pkg/airports/calc/test_cases"
	"testing"
)

var flights = [][]string{{"SFO", "EWR"}, {"ATL", "EWR"}, {"SFO", "ATL"}, {"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}

func BenchmarkProcessFlights(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		calc.ProcessFlights(flights)
	}
}

var maxSequence = test_cases.GenerateMAXSequence()

func BenchmarkProcessFlightsMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		calc.ProcessFlights(maxSequence)
	}
}
