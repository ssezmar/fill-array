
build:
	go build -o ./.bin/array .

run: build
	./.bin/array

test:
	go test ./... -coverprofile=NUL /dev/null