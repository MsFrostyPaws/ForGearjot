
run: 
	go mod download && go build -o ./.bin/app ./cmd/server/main.go
	 ./.bin/app

test:
	go test -coverprofile=coverage.out ./...

