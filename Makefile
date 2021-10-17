l/build:
	cd ./src/ && go build .

l/up:
	docker-compose down --remove-orphans --rmi all
	docker-compose up --detach --remove-orphans --force-recreate

l/proto:
	@rm -rf ./src/gen
	cd ./src/ && protoc --go_out=plugins=grpc:. proto/multiply.proto