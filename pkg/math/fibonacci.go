package math

import (
	"fmt"
	cache2 "github.com/alex-dev-master/fibonacci.git/pkg/cache"
	"math"
)

type FibonacciService struct {
	rdbCache *cache2.Client
}

func NewFibonacciService(rdbCache *cache2.Client) *FibonacciService {
	return &FibonacciService{
		rdbCache: rdbCache,
	}
}

func (receiver *FibonacciService) GetSlice(offset uint64, limit uint64) ([]uint64, error) {
	res := calcFibonacciSlice(calcFibonacci, offset, limit, receiver.rdbCache)

	return res, nil
}

type calcFibonacciFunc func(uint64) uint64

func calcFibonacciSlice(calcFibonacci calcFibonacciFunc, a, b uint64, rdbCache *cache2.Client) []uint64 {
	length := b - a + 1
	slice := make([]uint64, length)
	j := 0
	for i := a; i <= b; i++ {
		val, err := rdbCache.Get(fmt.Sprint(i))
		if err == nil {
			slice[j] = val
		} else {
			res := calcFibonacci(i)
			slice[j] = res
			rdbCache.Set(fmt.Sprint(i), res, 0)
		}

		j++
	}
	return slice
}

func calcFibonacci(n uint64) uint64 {
	g := (1 + math.Sqrt(5)) / 2
	ret := (math.Pow(g, float64(n)) - math.Pow(1-g, float64(n))) / math.Sqrt(5)
	return uint64(ret)
}
