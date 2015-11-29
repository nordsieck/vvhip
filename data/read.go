package data

import (
	"strconv"
	"time"

	"github.com/nordsieck/defect"
)

const (
	ErrWrongNumFields = defect.Error("Wrong number of fields")

	layout = "Jan 2006"
)

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

type Competition struct {
	Number         uint32
	Name, Location string
	Date           time.Time
}

func ParseCompetition(s []string) (*Competition, error) {
	if len(s) != 4 {
		return nil, ErrWrongNumFields
	}
	date, err := time.Parse(layout, s[3])
	if err != nil {
		return nil, err
	}
	num, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return nil, err
	}
	return &Competition{Number: uint32(num), Name: s[1], Location: s[2], Date: date}, nil
}
