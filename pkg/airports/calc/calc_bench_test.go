package calc

import "testing"

var flights = [][]string{{"SFO", "EWR"}, {"ATL", "EWR"}, {"SFO", "ATL"}, {"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}

func BenchmarkProcessFlights(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProcessFlights(flights)
	}
}
