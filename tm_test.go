package bbchallenge

import (
	"testing"
)

func TestSimulate(t *testing.T) {
	time, err := TmSimulate(GetBB5Winner())
	if time != BB5 || err != nil {
		t.Error(time, err)
	}
}
