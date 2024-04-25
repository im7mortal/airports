package calc

import "testing"

var flights = [][]string{{"SFO", "EWR"}, {"ATL", "EWR"}, {"SFO", "ATL"}, {"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}

func BenchmarkProcessFlights(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProcessFlights(flights)
	}
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateMAXSequence() [][]string {
	var sequence [][]string

	for _, i := range letters {
		for _, j := range letters {
			for _, k := range letters {
				airport := string([]rune{i, j, k})
				sequence = append(sequence, []string{airport, airport})
			}
		}
	}

	return sequence
}

var maxSequence = generateMAXSequence()

func BenchmarkProcessFlightsMax(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ProcessFlights(maxSequence)
	}
}
