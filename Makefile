l/build:
	cd ./src/ && go build .

l/up:
	cd ./src/ && protoc --go_out=plugins=grpc:. proto/multiply.proto
	docker-compose down --remove-orphans --rmi all
	docker-compose up --detach --remove-orphans --force-recreate

l/proto:
	@rm -rf ./src/gen
	cd ./src/ && protoc --go_out=plugins=grpc:. proto/multiply.proto

# Run the test coverage report by each file on HTML
# Requires a golang compiler installed
l/tch:
	cd ./src/ ; go test -coverprofile=coverage.out ./... ; go tool cover -html=coverage.out ; rm coverage.out

# Run the overall test coverage report on the console
# Requires a golang compiler installed
l/tco:
	cd ./src/ ; go test -v -coverpkg=./... -coverprofile=profile.cov ./... ; go tool cover -func profile.cov ; rm profile.cov
