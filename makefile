build:
	go build -o main cmd/universal_blog/main.go

run: 
	./main

dev: build run