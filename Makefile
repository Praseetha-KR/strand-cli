GOCMD = go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean

OS_LINUX = linux
OS_DARWIN = darwin
OS_WINDOWS = windows

BUILD_DIR = bin
BINARY_NAME = strand
BINARY_LINUX := $(BINARY_NAME)_$(OS_LINUX)
BINARY_DARWIN := $(BINARY_NAME)_$(OS_DARWIN)
BINARY_WINDOWS := $(BINARY_NAME)_$(OS_WINDOWS)

VERSION := $(shell git describe --abbrev=0 --tags)
GIT_COMMIT := $(shell git describe --always)
LDFLAGS := -X main.version=${VERSION} -X main.commit=${GIT_COMMIT}

DARWIN_INSTALL_TARGET = /usr/local/bin/strand
LINUX_INSTALL_TARGET := $(DARWIN_INSTALL_TARGET)


all: build

build: build_darwin build_linux build_windows

build_darwin: build_darwin_amd64 build_darwin_arm64
build_linux: build_linux_amd64 build_linux_arm64 build_linux_arm build_linux_386
build_windows: build_windows_amd64 build_windows_386


build_darwin_amd64: mkdir_bin
	GOOS=$(OS_DARWIN) GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_DARWIN)_amd64 -ldflags="$(LDFLAGS)"

build_darwin_arm64: mkdir_bin
	GOOS=$(OS_DARWIN) GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_DARWIN)_arm64 -ldflags="$(LDFLAGS)"


build_linux_amd64: mkdir_bin
	GOOS=$(OS_LINUX) GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_LINUX)_amd64 -ldflags="$(LDFLAGS)"

build_linux_arm64: mkdir_bin
	GOOS=$(OS_LINUX) GOARCH=arm64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_LINUX)_arm64 -ldflags="$(LDFLAGS)"

build_linux_arm: mkdir_bin
	GOOS=$(OS_LINUX) GOARCH=arm $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_LINUX)_arm -ldflags="$(LDFLAGS)"

build_linux_386: mkdir_bin
	GOOS=$(OS_LINUX) GOARCH=386 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_LINUX)_386 -ldflags="$(LDFLAGS)"


build_windows_amd64: mkdir_bin
	GOOS=$(OS_WINDOWS) GOARCH=amd64 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_WINDOWS)_amd64 -ldflags="$(LDFLAGS)"

build_windows_386: mkdir_bin
	GOOS=$(OS_WINDOWS) GOARCH=386 $(GOBUILD) -o $(BUILD_DIR)/$(BINARY_WINDOWS)_386 -ldflags="$(LDFLAGS)"


install_darwin_amd64: build_darwin_amd64
	mv $(BUILD_DIR)/$(BINARY_DARWIN)_amd64 $(DARWIN_INSTALL_TARGET)

install_darwin_arm64: build_darwin_arm64
	mv $(BUILD_DIR)/$(BINARY_DARWIN)_arm64 $(DARWIN_INSTALL_TARGET)

install_darwin_arm: build_darwin_arm
	mv $(BUILD_DIR)/$(BINARY_DARWIN)_arm $(DARWIN_INSTALL_TARGET)


install_linux_amd64: build_linux_amd64
	mv $(BUILD_DIR)/$(BINARY_LINUX)_amd64 $(LINUX_INSTALL_TARGET)

install_linux_arm64: build_linux_arm64
	mv $(BUILD_DIR)/$(BINARY_LINUX)_arm64 $(LINUX_INSTALL_TARGET)

install_linux_arm: build_linux_arm
	mv $(BUILD_DIR)/$(BINARY_LINUX)_arm $(LINUX_INSTALL_TARGET)

install_linux_386: build_linux_386
	mv $(BUILD_DIR)/$(BINARY_LINUX)_386 $(LINUX_INSTALL_TARGET)


mkdir_bin:
	mkdir -p bin


clean_darwin: clean
	rm -f $(DARWIN_INSTALL_TARGET)

clean_linux: clean
	rm -f $(LINUX_INSTALL_TARGET)

clean:
	rm -rf $(BUILD_DIR)
	$(GOCLEAN)
