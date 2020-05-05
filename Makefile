build:
		@go build -o marriage ./...

format:
		@go fmt ./...

run : build
		./marriage

clean :
		rm -rf marriage generated.yml