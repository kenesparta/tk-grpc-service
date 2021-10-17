FROM golang:1.17-alpine AS go_builder_grpc

WORKDIR /grpcService

COPY ./src .

RUN apk add gcc musl-dev protobuf git && \
    go mod tidy && \
    go get -d -u && \
    go install github.com/golang/protobuf/protoc-gen-go@latest && \
    protoc --go_out=plugins=grpc:. proto/multiply.proto && \
    go test ./... -cover && \
    go build .


FROM alpine:3.14

COPY --from=go_builder_grpc /grpcService/tkGrpcService .

EXPOSE 8085

ENTRYPOINT [ "./tkGrpcService" ]
