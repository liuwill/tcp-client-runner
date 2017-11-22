install:
	go get

run:
	cd bin/ && go run main.go

.PHONY: run install
