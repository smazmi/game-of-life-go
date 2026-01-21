APP_NAME = game-of-life-go
BUILD_DIR = build

CMD_DIR_TERM = ./cmd/terminal
CMD_DIR_EBITEN = ./cmd/ebiten

TERM_BIN = $(BUILD_DIR)/$(APP_NAME)-term
EBITEN_BIN = $(BUILD_DIR)/$(APP_NAME)-ebiten

.PHONY: build-term build-ebiten run-term run-ebiten clean

build-term:
	mkdir -p $(BUILD_DIR)
	go build -o $(TERM_BIN) $(CMD_DIR_TERM)

build-ebiten:
	mkdir -p $(BUILD_DIR)
	go build -o $(EBITEN_BIN) $(CMD_DIR_EBITEN)

run-term: build-term
	./$(TERM_BIN)

run-ebiten: build-ebiten
	./$(EBITEN_BIN)

clean:
	rm -rf $(BUILD_DIR)
