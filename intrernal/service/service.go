package service

import "github.com/alex-dev-master/fibonacci.git/intrernal/model"

type Fibonacci interface {
	GetSlice(data model.Fibonacci) ([]uint64, error)
}

type Service struct {
	Fibonacci
}

func NewService() *Service {
	return &Service{
		Fibonacci: NewFibonacciService(),
	}
}