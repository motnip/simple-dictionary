.PHONY: all test unit integr mock

mockfolder=mocks
prjroot=github.com/motnip/sermo/

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
	@#https://www.gnu.org/software/bash/manual/html_node/Command-Substitution.html
	$(eval iinterface=$(shell echo $(interface) | tr '[:upper:]' '[:lower:]'))
	@mockgen -destination=$(mockfolder)/$(package)/mock_$(iinterface).go -package=$(mockfolder) $(prjroot)$(package) $(interface)
	@echo "github.com/motnip/sermo/$(mockfolder)/$(package):"
	@ls -1 ./mocks/$(package)
