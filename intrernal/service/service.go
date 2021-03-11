package service

import (
	"github.com/alex-dev-master/fibonacci.git/intrernal/model"
	"github.com/alex-dev-master/fibonacci.git/pkg/cache"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Fibonacci interface {
	GetSlice(data model.Fibonacci) ([]uint64, error)
}

type Service struct {
	Fibonacci
}

func NewService(rdbCache *cache.Store) *Service {
	return &Service{
		Fibonacci: NewFibonacciService(rdbCache),
	}
}