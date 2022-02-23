package bbchallenge

import (
	"strings"

	uuid "github.com/nu7hatch/gouuid"
)

func GetRunName() string {
	id, _ := uuid.NewV4()

	split := strings.Split(id.String(), "-")
	return "run-" + split[len(split)-1]
}

func MaxI(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinI(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Following 3 functions are used by heuristics deciders

// Equivalent of l[beginIndex::sampleStep] in python
func SampleList(l []interface{}, beginIndex int, sampleStep int) (toReturn []interface{}) {
	for i := beginIndex; i < len(l); i += sampleStep {
		toReturn = append(toReturn, l[i])
	}
	return toReturn
}

func AllZero(l []int) bool {
	for i := range l {
		if l[i] != 0 {
			return false
		}
	}
	return true
}

// Discrete difference is computing sequence ofl[i+1]-l[i]
func discreteDifferenceOperator(l []int) (toReturn []int) {
	if len(l) < 2 {
		return toReturn
	}

	for i := 1; i < len(l); i += 1 {
		toReturn = append(toReturn, l[i]-l[i-1])
	}
	return toReturn
}

// Return the nth discrete difference of l
func DiscreteDifference(l []int, n int) (toReturn []int) {
	toReturn = make([]int, len(l))
	copy(toReturn, l)

	for i := 0; i < n; i += 1 {
		toReturn = discreteDifferenceOperator(toReturn)
	}

	return toReturn
}
