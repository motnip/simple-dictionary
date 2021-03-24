all: lint fmt

lint:
	@echo "linting...linting"
	golint ./...

fmt:
	@echo "verify formatting..."
	go fmt ./...

test:
	@echo "running all tests..."
	go test ./... -v -cover -failfast -short

unit:
	@echo "running unit tests..."
	go test ./... -v -cover -failfast -short

integr:
	@echo "running integration tests..."
	go test ./... -v -cover -failfast -short

mock:
	@echo "generating mocks"
	# the package path and interface name must be parameters
	# the package root path is github.com/motnip/sermo/ and must be omitted
	# the name of the mock file is the name lowercase name of the interface
	mockgen -destination=mock/mock_test.go -package=mock github.com/motnip/sermo/controller ...
