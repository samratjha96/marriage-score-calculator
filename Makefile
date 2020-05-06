PHONY: clean

build:
	@go build -o marriage main.go

format:
	@go fmt ./...

run : build
	./marriage

clean :
	rm -rf marriage generated.yml

score : build
	./marriage -score
