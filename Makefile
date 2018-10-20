BUILD_DIR = output
APP_NAME = new

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME)
	@echo "build success!"
	@$(BUILD_DIR)/$(APP_NAME) -h

clean:
	@rm -rf $(BUILD_DIR)/*
	@echo "clean success!"

.PHONY: help
help:
	@echo "help"
