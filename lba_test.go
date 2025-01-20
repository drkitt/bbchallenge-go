package bbchallenge

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetMachine(t *testing.T) {
	const DB_PATH = "all_5_states_undecided_machines_with_global_header"

	DB, err := ioutil.ReadFile(DB_PATH)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	n := 0
	m, err := GetMachineI(DB[:], n, true)

	if err != nil {
		t.Fail()
	}
	fmt.Println(m.ToAsciiTable(5))

	n = 10
	m, err = GetMachineI(DB[:], n, true)

	if err != nil {
		t.Fail()
	}
	fmt.Println(m.ToAsciiTable(5))

	n = 4888230
	m, err = GetMachineI(DB[:], n, true)

	if err != nil {
		t.Fail()
	}
	fmt.Println(m.ToAsciiTable(5))
}

func TestSimulate(t *testing.T) {
	time, err := LbaSimulate(GetBB5Winner())
	if time != BB5 || err != nil {
		t.Error(time, err)
	}
}
