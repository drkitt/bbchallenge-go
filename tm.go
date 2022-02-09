package bbchallenge

import (
	"encoding/binary"
	"errors"
	"strconv"

	tabulate "github.com/rgeoghegan/tabulate"
)

const R = 0
const L = 1

type TM [2 * 5 * 3]byte

func tmTransitionToStr(b1 byte, b2 byte, b3 byte) (toRet string) {

	if b3 == 0 {
		return "???"
	}

	toRet = strconv.Itoa(int(b1))

	if b2 == 0 {
		toRet += "R"
	} else {
		toRet += "L"
	}

	toRet += string(rune(int('A') + int(b3) - 1))

	return toRet
}

func (tm TM) ToAsciiTable(nbStates byte) (toRet string) {

	var table [][]string

	for i := byte(0); i < nbStates; i += 1 {

		table = append(table, []string{string(rune(int('A') + int(i))),
			tmTransitionToStr(tm[6*i], tm[6*i+1], tm[6*i+2]),
			tmTransitionToStr(tm[6*i+3], tm[6*i+4], tm[6*i+5])})
	}

	layout := &tabulate.Layout{Headers: []string{"-", "0", "1"}, Format: tabulate.SimpleFormat}
	asText, _ := tabulate.Tabulate(
		table, layout,
	)

	return asText
}

func GetMachineI(db []byte, i int, hasHeader bool) (tm TM, err error) {

	if i < 0 || i > len(db)/30 {
		err := errors.New("invalid db index")
		return tm, err
	}

	offset := 0
	if hasHeader {
		offset = 1
	}

	copy(tm[:], db[30*(i+offset):30*(i+offset+1)])
	return tm, nil
}

func GetMachineIFromIndex(db []byte, i int, hasHeader bool, undecidedMachinesIndex []byte) (tm TM, err error) {

	if i < 0 || i > len(undecidedMachinesIndex)/4 {
		err := errors.New("invalid index of undecided machines index")
		return tm, err
	}

	indexInDb := binary.BigEndian.Uint32(undecidedMachinesIndex[i*4 : (i+1)*4])

	if indexInDb < 0 || indexInDb > uint32(len(db)/30) {
		err := errors.New("invalid db index")
		return tm, err
	}

	offset := uint32(0)
	if hasHeader {
		offset = 1
	}

	copy(tm[:], db[30*(indexInDb+offset):30*(indexInDb+offset+1)])
	return tm, nil
}

func GetBB5Winner() TM {
	// +---+-----+-----+
	// | - |  0  |  1  |
	// +---+-----+-----+
	// | A | 1RB | 1LC |
	// | B | 1RC | 1RB |
	// | C | 1RD | 0LE |
	// | D | 1LA | 1LD |
	// | E | 1RH | 0LA |
	// +---+-----+-----+

	return TM{
		1, R, 2, 1, L, 3,
		1, R, 3, 1, R, 2,
		1, R, 4, 0, L, 5,
		1, L, 1, 1, L, 4,
		1, R, 0, 0, L, 1}

}

const MAX_MEMORY = 40000

type SimpleTape struct {
	tape [MAX_MEMORY]byte
}

func TmStep(tm TM, read byte, currState byte, currPos int, currTime int) (write byte, nextState byte, nextPos int) {

	tmTransition := 6*(currState-1) + 3*read
	write = tm[tmTransition]

	move := tm[tmTransition+1]
	nextState = tm[tmTransition+2]

	if move == R {
		nextPos = currPos + 1

	} else {
		nextPos = currPos - 1

	}

	return write, nextState, nextPos
}

func TmSimulate(tm TM) (int, error) {
	currPos := MAX_MEMORY / 2
	nextPos := currPos
	currState := byte(1)
	currTime := 0
	write := byte(0)
	var tape SimpleTape

	var err error

	for err == nil && currState != 0 {

		if currPos < 0 || currPos >= len(tape.tape) {
			err = errors.New("memory exceeded")
			continue
		}

		read := tape.tape[currPos]

		write, currState, nextPos = TmStep(tm, read, currState, currPos, currTime)

		tape.tape[currPos] = write
		currPos = nextPos
		currTime += 1
	}

	return currTime, err
}
