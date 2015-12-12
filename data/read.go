package data

import (
	"strconv"
	"time"

	"github.com/nordsieck/defect"
)

const (
	ErrWrongNumFields    = defect.Error("Wrong number of fields")
	ErrInvalidLeadFollow = defect.Error("Invalid lead of follow value")

	layout = "Jan 2006"
)

type Parser interface {
	Parse([]string) error
}

type Dancer struct {
	Number      uint32
	First, Last string
}

var _ Parser = &Dancer{}

func (d *Dancer) Parse(s []string) error {
	if len(s) != 3 {
		return ErrWrongNumFields
	}
	num, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return err
	}
	d.Number = uint32(num)
	d.First = s[2]
	d.Last = s[1]
	return nil
}

type Competition struct {
	Number         uint32
	Name, Location string
	Date           time.Time
}

var _ Parser = &Competition{}

func (c *Competition) Parse(s []string) error {
	if len(s) != 4 {
		return ErrWrongNumFields
	}
	date, err := time.Parse(layout, s[3])
	if err != nil {
		return err
	}
	num, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return err
	}
	c.Number = uint32(num)
	c.Name = s[1]
	c.Location = s[2]
	c.Date = date
	return nil
}

type Result struct {
	Dancer, Competition uint32
	Lead                bool
	Category            string
	Result, Points      uint8
}

var _ Parser = &Result{}

func (r *Result) Parse(s []string) error {
	if len(s) != 6 {
		return ErrWrongNumFields
	}
	dancer, err := strconv.ParseUint(s[0], 10, 32)
	if err != nil {
		return err
	}
	competition, err := strconv.ParseUint(s[1], 10, 32)
	if err != nil {
		return err
	}
	lead := true
	if s[2] == "f" {
		lead = false
	} else if s[2] != "l" {
		return ErrInvalidLeadFollow
	}
	var result uint64
	if s[4] != "F" {
		result, err = strconv.ParseUint(s[4], 10, 32)
		if err != nil {
			return err
		}
	}
	points, err := strconv.ParseUint(s[5], 10, 32)
	if err != nil {
		return err
	}
	r.Dancer = uint32(dancer)
	r.Competition = uint32(competition)
	r.Lead = lead
	r.Category = s[3]
	r.Result = uint8(result)
	r.Points = uint8(points)
	return nil
}
