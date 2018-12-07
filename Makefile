APP_NAME = news
BUILD_DIR = output
RELEASE_DIR = release
INSTALL_DIR = /usr/local/bin
DARWIN = darwin
LINUX = linux
VERSION = `git tag |sort -Vr |head -1`
export GO111MODULE=on

build:
	@go mod vendor
	@mkdir -p $(BUILD_DIR)
	@mkdir -p $(RELEASE_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(DARWIN)/$(APP_NAME) -ldflags "-s -w"
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(LINUX)/$(APP_NAME) -ldflags "-s -w"
	@echo "build success!"
	#@tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.gz -C $(BUILD_DIR)/$(DARWIN) .
	#@tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.xz -C $(BUILD_DIR)/$(LINUX) .
	zip -j $(RELEASE_DIR)/$(VERSION).zip $(BUILD_DIR)/$(DARWIN)/*
	tar -cvzf $(RELEASE_DIR)/$(VERSION).tar.gz -C $(BUILD_DIR)/$(LINUX) .
	@echo "release success!"

BUILD_DIR = output
APP_NAME = news
BIN_PATH = /usr/local/bin
GZ = gz
XZ = xz

build-drawin:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(GZ)/$(APP_NAME) -ldflags "-s -w"
	@echo "build success!"
	@#$(BUILD_DIR)/$(APP_NAME) -h

build-linux:
	@mkdir -p $(BUILD_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(XZ)/$(APP_NAME) -ldflags "-s -w"
	@echo "build success!"
	@#$(BUILD_DIR)/$(APP_NAME) -h

install: build
	@cp $(BUILD_DIR)/$(APP_NAME) $(BIN_PATH)

uninstall: install
	@sudo rm -rf $(BIN_PATH)/$(APP_NAME)

clean:
	@rm -vrf $(BUILD_DIR)/*
	@echo "====>clean success!"

.PHONY: help
help:
	@echo "help"
