package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcFibonacci(t *testing.T) {
	input := map[uint64]uint64{
		0: 0,
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
		6: 8,
		7: 13,
	}

	for key, value := range input {
		valueFunc := CalcFibonacci(key)
		assert.Equal(t, valueFunc, value)
	}

}
