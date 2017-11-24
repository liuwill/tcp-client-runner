install:
	go get

run:
	cd bin/custom && go run main.go

run-demo:
	cd bin/demo && go run main.go

.PHONY: run install
