install:
	go get

run:
	export DEFAULT_PORT=50001 && cd bin/custom && go run main.go

run-demo:
	export DEFAULT_PORT=50001 && cd bin/demo && go run main.go

mock:
	export DEFAULT_PORT=50001 && cd bin/mock && go run main.go

.PHONY: run mock install
