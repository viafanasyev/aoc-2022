.PHONY: test
test: main ../sample-1.txt ../sample-2.txt ../input.txt
	./main ../sample-1.txt
	./main ../sample-2.txt
	./main ../input.txt

main: main.go ../../go.mod
	go build -o main
