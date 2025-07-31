.PHONY: dev build run clean

# Development target - runs air for hot reloading
dev:
	air main.go

# Build the application
build:
	go build -o bin/server main.go

# Run the built application
run: build
	./bin/server

# Clean build artifacts
clean:
	rm -rf bin/
