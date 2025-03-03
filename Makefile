system-deps:
	aptitude install libc6-dev libgl1-mesa-dev libxcursor-dev libxi-dev libxinerama-dev libxrandr-dev libxxf86vm-dev libasound2-dev pkg-config

run:
	CGO_ENABLED=1 go run .

test:
	CGO_ENABLED=1 go test -v
