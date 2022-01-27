package bbchallenge

import (
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

	if i <= 0 || i > len(db)/30 {
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
		1, R, 6, 0, L, 1}

}
