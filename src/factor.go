package main

import (
	"context"
	"errors"
	"github.com/kenesparta/multiplyLogic"
	"github.com/kenesparta/tkGrpcService/gen/multiply"
	"log"
)

type Server struct{}

func (s *Server) Multiply(ctx context.Context, product *multiply.Factor) (*multiply.Product, error) {
	var (
		fp = product.GetFirst()
		sp = product.GetSecond()
		m = multiplyLogic.Factor{
			First:  &fp,
			Second: &sp,
		}
	)
	r := m.Product()
	if m.AreValidNumbers() {
		return nil, errors.New("")
	}
	if m.IsProductInfinite() {
		return nil, errors.New("")
	}
	log.Printf("Product: %f", r)
	return &multiply.Product{Result: r}, nil
}
