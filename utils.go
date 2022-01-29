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
