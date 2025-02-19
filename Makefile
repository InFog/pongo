run:
	CGO_ENABLED=1 go run .

test:
	CGO_ENABLED=1 go test -v
