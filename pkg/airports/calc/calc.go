// Created by Petr Lozhkin
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

	for i := range flights {

		// let it panic if input is invalid
		// TODO I didn't decide yet if I want to have duplicate here or only validation on JSON
		origin := flights[i][0]
		destin := flights[i][1]

		if _, exist := m[origin]; !exist {
			m[origin] = &Metadata{Origin: 1}
		} else {
			m[origin].Origin += 1
		}
		if _, exist := m[destin]; !exist {
			m[destin] = &Metadata{Destination: 1}
		} else {
			m[destin].Destination += 1
		}
	}
	//fmt.Printf("%v\n", m)
	for k, v := range m {
		// if sum is odd then we have origin or destination
		if (v.Origin+v.Destination)%2 == 1 {
			// if sum is bigger than 1 then we have a loop, in this case we choose even number
			if v.Origin+v.Destination > 1 {
				if v.Origin%2 == 0 {
					//fmt.Println(k)
					if findOrig != "" {
						return nil, InvalidSequence
					} else {
						findOrig = k
					}
				}
				if v.Destination%2 == 0 {
					if findDest != "" {
						return nil, InvalidSequence
					} else {
						findDest = k
					}
				}
			} else {
				if v.Origin%2 == 1 {
					//fmt.Println(k)
					if findOrig != "" {
						return nil, InvalidSequence
					} else {
						findOrig = k
					}
				}
				if v.Destination%2 == 1 {
					if findDest != "" {
						return nil, InvalidSequence
					} else {
						findDest = k
					}
				}
			}
		}
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
