package service

import (
    "math/rand"
    "time"
)

type StockService interface {
    GetStock(barCode string) (int, error)
}

type MockStockService struct{}

func NewMockStockService() *MockStockService {
    return &MockStockService{}
}

func (s *MockStockService) GetStock(barCode string) (int, error) {
    // Simulate API delay
    time.Sleep(100 * time.Millisecond)
    return rand.Intn(100), nil
}
