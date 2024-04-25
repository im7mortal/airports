package calc

import (
	"errors"
	"fmt"
)

type Metadata struct {
	Origin, Destination int
}

func (m *Metadata) String() string {
	return fmt.Sprintf("%d:%d", m.Origin, m.Destination)
}

var InvalidSequence = errors.New("invalid sequence")
var NotEnoughData = errors.New("not enough data")

func ProcessFlights(flights [][]string) ([]string, error) {
	var findOrig, findDest string
	m := map[string]*Metadata{}

	destCandidates := map[string]*Metadata{}
	origCandidates := map[string]*Metadata{}

	var exist bool
	var dest *Metadata
	var orig *Metadata

	for i := range flights {

		// let it panic if input is invalid
		// TODO I didn't decide yet if I want to have duplicate here or only validation on JSON
		origin := flights[i][0]
		destin := flights[i][1]

		if orig, exist = m[origin]; !exist {
			orig = &Metadata{}
			m[origin] = orig
		}

		(*orig).Origin++

		if ((*orig).Origin+(*orig).Destination)%2 == 1 && (*orig).Origin%2 == 0 {
			origCandidates[origin] = orig
		} else {
			delete(origCandidates, origin)
		}

		if dest, exist = m[origin]; !exist {
			dest = &Metadata{}
			m[destin] = dest
		}

		(*dest).Destination++

		if ((*dest).Origin+(*dest).Destination)%2 == 1 && (*dest).Destination%2 == 0 {
			destCandidates[destin] = dest
		} else {
			delete(destCandidates, destin)
		}

	}

	if len(destCandidates) > 1 || len(origCandidates) > 1 {
		return nil, InvalidSequence
	}

	for findDest = range destCandidates {
	}
	for findOrig = range origCandidates {
	}

	// It's a ring
	if findOrig == "" && findDest == "" {
		return nil, NotEnoughData
	}

	return []string{findOrig, findDest}, nil

}

//type Node struct {
//	Origin, Destination *Node
//
//	Name string
//}
//func ProcessFlights(flights [][]string) ([]string, error) {
//	m := map[string]*Node{}
//
//	head := nil
//	tail := nil
//
//	for i := range flights {
//		origin := flights[i][0]
//		destin := flights[i][1]
//
//		if val, ok := m[flights[i][0]]; !ok {
//			m[flights[i][0]]
//		}
//	}
//
//	return nil, nil
//}
