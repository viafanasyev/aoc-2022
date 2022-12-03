.PHONY: test
test: main sample.txt input.txt
	./main sample.txt
	./main input.txt

main: main.go go.mod
	go build -o main
