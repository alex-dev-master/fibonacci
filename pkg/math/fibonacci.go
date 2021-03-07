package math

type FibonacciService struct {
}

func NewFibonacciService() *FibonacciService {
	return &FibonacciService{}
}

func (receiver *FibonacciService) GetSlice(offset uint64, limit uint64) ([]uint64, error) {
	res := calcFibonacciSlice(calcFibonacci, offset, limit)

	return res, nil
}

type calcFibonacciFunc func(uint64) uint64

func calcFibonacciSlice(calcFibonacci calcFibonacciFunc, a, b uint64) []uint64 {
	length := b - a + 1
	slice := make([]uint64, length)
	j := 0
	for i := a; i <= b; i++ {
		slice[j] = calcFibonacci(i)
		j++
	}
	return slice
}

func calcFibonacci(n uint64) uint64 {
	var a, b uint64 = 1, 1
	for i := 0; i < int(n); i++ {
		a, b = b, a+b
	}
	return a
}
