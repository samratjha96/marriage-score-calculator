PHONY: clean

install:
	@go install

format:
	@go fmt ./...

run : install
	marriage

clean :
	rm -rf marriage game.yml

