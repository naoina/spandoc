BIN_NAME="$(notdir $(PWD))"
BUILDFLAGS := -tags netgo -installsuffix netgo -ldflags '-w -s --extldflags "-static"'

.PHONY: all
all:
	CGO_ENABLED=0 go build $(BUILDFLAGS) ./cmd/$(BIN_NAME)

.PHONY: clean
clean:
	$(RM) -f $(BIN_NAME)
