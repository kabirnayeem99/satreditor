GO          ?= go
DLV         ?= dlv
CMD_PKG     ?= ./cmd/satr
BIN_NAME    ?= satr
BIN_DIR     ?= bin
PREFIX      ?= /usr/local
INSTALL_BIN ?= $(PREFIX)/bin/$(BIN_NAME)

.PHONY: run debug build clean install uninstall fmt tidy

run:
	$(GO) run $(CMD_PKG)

# Build the CLI into ./bin/satr
build: $(BIN_DIR)/$(BIN_NAME)

$(BIN_DIR)/$(BIN_NAME):
	@mkdir -p $(BIN_DIR)
	$(GO) build -o $(BIN_DIR)/$(BIN_NAME) $(CMD_PKG)

# Launch the editor under Delve for interactive debugging
# Requires github.com/go-delve/delve/cmd/dlv to be installed.
debug:
	$(DLV) debug $(CMD_PKG)

fmt:
	$(GO) fmt ./...

tidy:
	$(GO) mod tidy

clean:
	rm -rf $(BIN_DIR)

install: build
	install -d $(PREFIX)/bin
	install -m 755 $(BIN_DIR)/$(BIN_NAME) $(INSTALL_BIN)

uninstall:
	rm -f $(INSTALL_BIN)
