run:
	CGO_ENABLED=1 go run main.go

test:
	CGO_ENABLED=1 go test -v
