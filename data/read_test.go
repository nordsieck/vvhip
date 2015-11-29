package data

import (
	"testing"

	"github.com/nordsieck/defect"
)

func TestParseDancer(t *testing.T) {
	_, err := ParseDancer([]string{})
	defect.Equal(t, err, ErrWrongNumFields)

	dancer, err := ParseDancer([]string{"12132", "Cohen", "Miles"})
	defect.DeepEqual(t, dancer, &Dancer{
		Number: 12132,
		First:  "Miles",
		Last:   "Cohen",
	})
}
