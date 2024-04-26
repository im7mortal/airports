package test_cases

import (
	"github.com/im7mortal/airports/pkg/airports/calc"
	"net/http"
)

type TestCase struct {
	Name               string
	Input              [][]string
	Expected           []string
	ExpectedErr        error
	ExpectedStatusCode int
}

var Standard = []TestCase{
	{"Test Case 1", [][]string{{"SFO", "EWR"}}, []string{"SFO", "EWR"}, nil, http.StatusOK},
	{"Test Case 2", [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}}, []string{"SFO", "EWR"}, nil, http.StatusOK},
	{"Test Case 3", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}}, []string{"SFO", "EWR"}, nil, http.StatusOK},

	// Loop. Imagine we have need from Paris to SaoPaolo. And we have cheap flight from Paris to Berlin and from Berlin to SaoPaolo with stop in Paris
	{"Test Case 4.1", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}, nil, http.StatusOK},
	{"Test Case 4.2", [][]string{{"SFO", "ATL"}, {"ATL", "BIR"}, {"BIR", "SFO"}, {"SFO", "GRU"}}, []string{"SFO", "GRU"}, nil, http.StatusOK},
	{"Test Case 4.3", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"IND", "GSO"}, {"GSO", "IND"}, {"ATL", "GSO"}}, []string{"SFO", "EWR"}, nil, http.StatusOK},
	{"Test Case 4.4", [][]string{{"SFO", "ATL"}, {"ATL", "EWR"}, {"EWR", "SFO"}, {"SFO", "EWR"}}, []string{"SFO", "EWR"}, nil, http.StatusOK},

	// NO SOLUTION; sequence is invalid with GTO
	{"Test Case 5", [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"SFO", "GTO"}, {"GSO", "IND"}, {"ATL", "GSO"}}, nil, calc.InvalidSequence, http.StatusBadRequest},

	// NO SOLUTION. It's infinite loop; We can't say where is origin and destination

	{"Test Case 6", [][]string{{"SFO", "ATL"}, {"ATL", "SFO"}}, nil, calc.NotEnoughData, http.StatusBadRequest},
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateMAXSequence() [][]string {
	var sequence [][]string
	// it's just test code so I don't generate it effectively
	sequence = append(sequence, []string{"AAA", "AAB"})
	sequence[0] = []string{"AAA", ""}
	first := true
	var m int
	for _, i := range letters {
		for _, j := range letters {
			for _, k := range letters {
				if first {
					first = false
					continue
				}
				m++
				airport := string([]rune{i, j, k})
				sequence[m-1][1] = airport
				sequence = append(sequence, []string{airport, ""})
			}
		}
	}
	sequence[m][1] = "ZZZ"

	//fmt.Printf("%v\n",sequence)
	return sequence
}

func MaxSequenceRequest() TestCase {
	return TestCase{Input: GenerateMAXSequence(), ExpectedStatusCode: http.StatusOK}
}
