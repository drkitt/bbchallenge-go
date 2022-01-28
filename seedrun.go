package bbchallenge

import (
	"encoding/binary"
	"fmt"
)

// Constants computed during the seeding run
const TOTAL_UNDECIDED = 88664064
const TOTAL_UNDECIDED_TIME = 14322029
const TOTAL_UNDECIDED_SPACE = 74342035

func TestDB(db []byte, withHeader bool) error {

	offset := 0
	if withHeader {
		offset = 1
	}

	if len(db)/30-offset != TOTAL_UNDECIDED {
		return fmt.Errorf("core %d != %d", len(db)/30, TOTAL_UNDECIDED)
	} else if !withHeader {
		return nil
	}

	if binary.BigEndian.Uint32(db[0:4]) != TOTAL_UNDECIDED_TIME {
		return fmt.Errorf("header %d != %d", binary.BigEndian.Uint32(db[0:4]), TOTAL_UNDECIDED_TIME)
	}

	if binary.BigEndian.Uint32(db[4:8]) != TOTAL_UNDECIDED_SPACE {
		return fmt.Errorf("header %d != %d", binary.BigEndian.Uint32(db[4:8]), TOTAL_UNDECIDED_SPACE)
	}

	if binary.BigEndian.Uint32(db[8:12]) != TOTAL_UNDECIDED {
		return fmt.Errorf("header %d != %d", binary.BigEndian.Uint32(db[8:12]), TOTAL_UNDECIDED)
	}

	return nil
}
