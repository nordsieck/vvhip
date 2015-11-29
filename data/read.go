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

type Result struct {
	Dancer, Competition uint32
	Result, Points      uint8
}

func ParseResult(s []string) (*Result, error) {
	if len(s) != 4 {
		return nil, ErrWrongNumFields
	}
	dancer, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return nil, err
	}
	competition, err := strconv.ParseUint(s[1], 10, 32)
	if err != nil {
		return nil, err
	}
	var result uint64
	if s[2] != "F" {
		result, err = strconv.ParseUint(s[2], 10, 32)
		if err != nil {
			return nil, err
		}
	}
	points, err := strconv.ParseUint(s[3], 10, 32)
	if err != nil {
		return nil, err
	}
	return &Result{Dancer: uint32(dancer), Competition: uint32(competition), Result: uint8(result), Points: uint8(points)}, nil
}
