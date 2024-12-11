dev:
	go run cmd/main.go 

build:
	mkdir build
	go build -o build/chess cmd/main.go 
