package main

import (
	"context"
	"github.com/kenesparta/multiplyLogic"
	"github.com/kenesparta/tkGrpcService/gen/multiply"
	"log"
)

type Server struct{}

func (s *Server) Multiply(ctx context.Context, product *multiply.Factor) (*multiply.Product, error) {
	r := multiplyLogic.Multiply(product.GetFirst(), product.GetSecond())
	log.Printf("Product: %f", r)
	return &multiply.Product{Result: r}, nil
}
