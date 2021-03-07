package math

type Fibonacci interface {
	GetSlice(offset uint64, limit uint64) ([]uint64, error)
}

type MathematicsLibrary struct {
	Fibonacci
}

func NewMathematicsLibrary() *MathematicsLibrary {
	return &MathematicsLibrary{
		Fibonacci: NewFibonacciService(),
	}
}