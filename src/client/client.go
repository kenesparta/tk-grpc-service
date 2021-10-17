package main

import (
	"context"
	"github.com/kenesparta/tkGrpcService/gen/multiply"
	"google.golang.org/grpc"
	"log"
)

func main()  {
	conn, err := grpc.Dial(":8085", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%s", err)
	}

	defer conn.Close()

	c := multiply.NewMultiplyServiceClient(conn)

	r, err := c.Multiply(context.Background(), &multiply.Factor{
		First:  2.1,
		Second: 14.2,
	})

	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("%f", r.GetResult())
}
