
all: cover

test:
	go test -v ./...

cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | grep total
	go tool cover -html=coverage.out

clean:
	rm -rf cmd/server/server
	rm -rf cmd/agent/agent
	rm -rf coverage.out
