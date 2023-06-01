build:
	go build -o target/engine src/*.go

run-dev:
	go run ./src/main.go

run-prod:
	./target/engine

##### TESTING
run-test:
	go test -v ./test/...

run-test-coverage:
	go test -v ./test/... -coverprofile=./coverage.out & go tool cover -html=./coverage.out

mock-user-repository:
	mockery --dir=src/businesses/users --name=Repository --filename=repository.go --output=src/businesses/users/mocks --outpkg=mocks

mock-user-usecase:
	mockery --dir=src/businesses/users --name=UseCase --filename=usecase.go --output=src/businesses/users/mocks --outpkg=mocks

##### CONTAINER
run-docker:
	docker compose up

stop-docker:
	docker compose down

#### CURL
curl-post:
	curl -X POST -H "Content-Type: application/json" -d @./data/api/order/create_order.json http://localhost:8080/api/v1/orders