package service

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/pkg/cache"
	"github.com/alex-dev-master/fibonacci.git/pkg/math"
)

type FibonacciService struct {
	rdbCache *cache.Client
}

func NewFibonacciService(rdbCache *cache.Client) *FibonacciService {
	return &FibonacciService{rdbCache}
}

func (i *FibonacciService) GetSlice(data model.Fibonacci) ([]uint64, error) {
	res, err := math.NewMathematicsLibrary(i.rdbCache).GetSlice(data.X, data.Y)
	if err != nil {
		return nil, err
	}

	return res, nil
}