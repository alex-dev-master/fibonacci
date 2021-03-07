package math

type FibonacciService struct {
}

func NewFibonacciService() *FibonacciService {
	return &FibonacciService{}
}

func (receiver *FibonacciService) GetSlice(offset uint64, limit uint64) ([]uint64, error) {
	return []uint64{1, 2, 3, 4}, nil
}
