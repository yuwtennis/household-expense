
test:
	go test -v ./tests/...

build:
	docker build -t household-expense:test .