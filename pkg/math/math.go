package math

import "github.com/alex-dev-master/fibonacci.git/pkg/cache"

type Fibonacci interface {
	GetSlice(offset uint64, limit uint64) ([]uint64, error)
}

type MathematicsLibrary struct {
	Fibonacci
}

func NewMathematicsLibrary(rdbCache *cache.Store) *MathematicsLibrary {
	return &MathematicsLibrary{
		Fibonacci: NewFibonacciService(rdbCache),
	}
}