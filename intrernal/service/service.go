package service

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/pkg/cache"
)

type Fibonacci interface {
	GetSlice(data model.Fibonacci) ([]uint64, error)
}

type Service struct {
	Fibonacci
}

func NewService(rdbCache *cache.Client) *Service {
	return &Service{
		Fibonacci: NewFibonacciService(rdbCache),
	}
}