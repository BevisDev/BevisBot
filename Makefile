# name app build
APP_NAME = bevis

# dir and main
DOC_DIR = ./docs
MAIN_DIR = ./cmd/server
INFRAS_DIR = $(MAIN_DIR)/infrastructure
DI_DIR = ./internal/di
MAIN_FILE = $(MAIN_DIR)/main.go

# default build if not args in make
all: build

# Install and generate swag
swag_install:
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@latest

swag_gen:
	@echo "Generating swag.."
	@swag init --generalInfo $(MAIN_FILE) --output $(DOC_DIR)/swagger

# Install DI
di_install:
	@echo "Installing wire..."
	@go install github.com/google/wire/cmd/wire@latest

di_gen:
	@echo "Generating wire..."
	@cd $(DI_DIR) && wire gen

# Build the application
build:
	@echo "Building the application..."
	@go build -v -o $(APP_NAME) $(MAIN_FILE)

# run the application
run:
	@echo "Running the application..."
	@go run $(MAIN_FILE)