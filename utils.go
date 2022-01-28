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
