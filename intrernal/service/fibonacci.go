package service

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/pkg/math"
)

type FibonacciService struct {

}

func NewFibonacciService() *FibonacciService {
	return &FibonacciService{}
}

func (i *FibonacciService) GetSlice(data model.Fibonacci) ([]uint64, error) {
	res, err := math.NewFibonacciService().GetSlice(data.X, data.Y)
	if err != nil {
		return nil, err
	}

	return res, nil
}