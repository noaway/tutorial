build:
	go build .
gen:
	go generate ./...
all:
	go generate ./... && go build .