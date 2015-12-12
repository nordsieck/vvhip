package data

import (
	"testing"
	"time"

	"github.com/nordsieck/defect"
)

func TestParseDancer(t *testing.T) {
	dancer := &Dancer{}
	err := dancer.Parse([]string{})
	defect.Equal(t, err, ErrWrongNumFields)

	err = dancer.Parse([]string{"12132", "Cohen", "Miles"})
	defect.Equal(t, err, nil)
	defect.Equal(t, *dancer, Dancer{
		Number: 12132,
		First:  "Miles",
		Last:   "Cohen",
	})
}

func TestParseCompetition(t *testing.T) {
	comp := &Competition{}
	err := comp.Parse([]string{})
	defect.Equal(t, err, ErrWrongNumFields)

	err = comp.Parse([]string{"56", "Easter Swing", "Bellevue, WA", "Apr 2015"})
	defect.Equal(t, err, nil)
	defect.Equal(t, *comp, Competition{
		Number:   56,
		Name:     "Easter Swing",
		Location: "Bellevue, WA",
		Date:     time.Date(2015, 4, 1, 0, 0, 0, 0, time.UTC),
	})

}

func TestParseResult(t *testing.T) {
	result := &Result{}
	err := result.Parse([]string{})
	defect.Equal(t, err, ErrWrongNumFields)

	err = result.Parse([]string{"12132", "56", "f", "All-Stars", "F", "1"})
	defect.Equal(t, err, nil)
	defect.Equal(t, *result, Result{
		Dancer:      12132,
		Competition: 56,
		Lead:        false,
		Category:    "All-Stars",
		Result:      0,
		Points:      1,
	})

	err = result.Parse([]string{"12132", "56", "l", "Novice", "0", "1"})
	defect.Equal(t, err, nil)
	defect.Equal(t, *result, Result{
		Dancer:      12132,
		Competition: 56,
		Lead:        true,
		Category:    "Novice",
		Result:      0,
		Points:      1,
	})

	err = result.Parse([]string{"12132", "56", "f", "Invitational", "1", "10"})
	defect.Equal(t, err, nil)
	defect.Equal(t, *result, Result{
		Dancer:      12132,
		Competition: 56,
		Lead:        false,
		Category:    "Invitational",
		Result:      1,
		Points:      10,
	})
}
