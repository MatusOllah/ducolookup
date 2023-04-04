GO = go

BINARY = ./bin

FLAGS = -gcflags="-dwarf=false"

GOOS = $(shell $(GO) env GOOS)
GOARCH = $(shell $(GO) env GOARCH)

.PHONY: all
all: build

.PHONY: build
build: clean
	mkdir -p $(BINARY)/$(GOOS)-$(GOARCH)

	$(GO) get

	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(FLAGS) -o $(BINARY)/$(GOOS)-$(GOARCH)

.PHONY: clean
clean:
	rm -rf $(BINARY)/$(GOOS)-$(GOARCH)
