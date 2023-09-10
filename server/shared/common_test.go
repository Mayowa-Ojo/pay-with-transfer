package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBindReplacer(t *testing.T) {
	query := "SELECT * FROM table where column1 = ? and column2 = ?"
	expected := "SELECT * FROM table where column1 = $1 and column2 = $2"

	got := BindReplacer(query)
	assert.Equal(t, expected, got)
}
