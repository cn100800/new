BUILD_DIR = output

build :
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/new
	@echo "build success!"
	@$(BUILD_DIR)/new -h
