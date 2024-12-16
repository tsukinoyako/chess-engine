dev:
	go run .

build: clean
	mkdir build
	go build -o bin/chess 

clean: 
	rm -rf bin 
