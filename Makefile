.PHONY: all test unit integr mock

mockfolder=mocks
prjroot=github.com/motnip/sermo/

all: test build run

lint:
	@echo "  >  Linting..."
	@golint ./...

fmt:
	@echo "  >  Verify formatting..."
	@go fmt ./...

test:
	@echo "  >  Running all tests..."
	@go test ./... -v -cover -failfast

unit:
	@echo "  >  Running unit tests..."
	@go test ./... -v -cover -failfast -short

integr:
	@echo "  >  Running integration tests..."
	@go test -v ./... -run TestIntegration --failfast

mock:
	@echo "  >  Generating mocks..."
	@#https://www.gnu.org/software/bash/manual/html_node/Command-Substitution.html
	$(eval iinterface=$(shell echo $(interface) | tr '[:upper:]' '[:lower:]'))
	@mockgen -destination=$(mockfolder)/$(package)/mock_$(iinterface).go -package=$(mockfolder) $(prjroot)$(package) $(interface)
	@echo "github.com/motnip/sermo/$(mockfolder)/$(package):"
	@ls -1 ./mocks/$(package)

build:
	@echo "  >  Building binary..."
	@go build

run: build
	@echo "  >  Running..."
	@go run .

docker-build:
	@echo " > Building Docker image"
	@docker docker build . -t montip/sermo

docker-run:
	@echo " > Runnin Docker image"
	@docker run --name sermo -d -p 8080:3000 montip/sermo
