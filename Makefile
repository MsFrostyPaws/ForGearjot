MAIN_PROJECT_NAME:=gearjot
ifneq ($(OS),Windows_NT)
PROJECT_DIR:=$(shell pwd)
else
PROJECT_DIR:=$(shell cd)
endif

run: 
	go mod download && go build -o ./.bin/app ./cmd/server/main.go
	 ./.bin/app

# run tests and create test coverage
test:
	go test -coverprofile=coverage.out ./...

# `go mod vendor` create directory that contains copies of all packages needed to support builds
create-vendor:
	go mod vendor

clean-vendor:
	rm -rf $(PROJECT_DIR)/vendor

# `go mod tidy` remove unused modules 
tidy:
	go mod tidy
