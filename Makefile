APP_NAME = game-of-life-go
BUILD_DIR = build
CMD_DIR = ./cmd

.PHONY: build run clean

build:
	mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

run: build
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)
