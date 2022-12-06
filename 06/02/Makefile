.PHONY: test
test: main ../sample-1.txt ../sample-2.txt ../sample-3.txt ../sample-4.txt ../sample-5.txt ../input.txt
	./main ../sample-1.txt
	./main ../sample-2.txt
	./main ../sample-3.txt
	./main ../sample-4.txt
	./main ../sample-5.txt
	./main ../input.txt

main: main.go ../../go.mod
	go build -o main
