build:
	@go build -o bin/gopanic
run: build;
	@./bin/gopanic
