package math

import "fmt"

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

var cache = make(map[uint64]uint64, 0)

func calcFibonacciSlice(calcFibonacci calcFibonacciFunc, a, b uint64) []uint64 {
	length := b - a + 1
	slice := make([]uint64, length)
	j := 0
	for i := a; i <= b; i++ {
		if cache[i] != 0 {
			slice[j] = cache[i]
		} else {
			switch i {
			case 0:
				slice[j] = 0
			case 1:
				slice[j] = 1
			default:
				res := calcFibonacci(i)
				slice[j] = res
				cache[i] = res
			}
		}

		j++
	}
	//fmt.Printf("%v", cache)
	fmt.Println(cache)
	return slice
}

func calcFibonacci(n uint64) uint64 {
	var a, b uint64 = 1, 1
	for i := 0; i < int(n); i++ {
		a, b = b, a+b
	}
	return a
}
