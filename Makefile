build:
	go build -o bin/gobank

run: build
	./bin/gobank

test:
	g test -v ./...

clean:
	rm -rf bin