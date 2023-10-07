package shared

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindReplacer(t *testing.T) {
	query := "SELECT * FROM table where column1 = ? and column2 = ?"
	expected := "SELECT * FROM table where column1 = $1 and column2 = $2"

	got := BindReplacer(query)
	assert.Equal(t, expected, got)
}

func TestToBaseUnitAmount(t *testing.T) {
	testCases := []struct {
		input    float64
		expected int64
	}{
		{input: 1.87, expected: 187},
		{input: 1.03, expected: 103},
		{input: 0.01, expected: 1},
		{input: 0.43, expected: 43},
		{input: 1.15, expected: 115},
		{input: 1.65, expected: 165},
		{input: 200.03, expected: 20003},
		{input: 15.33, expected: 1533},
		{input: 2.60, expected: 260},
		{input: 2.13, expected: 213},
		{input: 2.98, expected: 298},
		{input: 1.06, expected: 106},
		{input: 300, expected: 30000},
		{input: 518, expected: 51800},
		{input: 0.14, expected: 14},
		{input: 0.71, expected: 71},
		{input: 0.33, expected: 33},
		{input: 20755, expected: 2075500},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("test-case-%d", idx), func(t *testing.T) {
			resp := ToBaseUnitAmount(tc.input)
			assert.Equal(t, tc.expected, resp)
		})
	}
}
