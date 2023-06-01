PACKAGE=github.com/pedy4000/logstronaut

# clear-log is used to clear previous logs
clear-log:
	@echo "Clearing Previous Logs..."
	@rm *.log

# create-bin-folder is used to create bin folder (if not exists)
create-bin-folder:
	@mkdir -p bin

# Build-All is used to build all services separately
build: create-bin-folder
	@echo "Building Logstronaut Service..."
	go build -o bin/server main.go
	@echo "Build Done!"
	@echo

run: build
	@echo "Running Logstronaut Service..."
	./bin/server
	@echo

.PHONY: clear-log create-bin-folder build run