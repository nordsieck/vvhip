package data

import (
	"strconv"

	"github.com/nordsieck/defect"
)

const ErrWrongNumFields = defect.Error("Wrong number of fields")

type Dancer struct {
	Number      uint32
	First, Last string
}

func ParseDancer(s []string) (*Dancer, error) {
	if len(s) != 3 {
		return nil, ErrWrongNumFields
	}
	num, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return nil, err
	}
	return &Dancer{Number: uint32(num), First: s[2], Last: s[1]}, nil
}
