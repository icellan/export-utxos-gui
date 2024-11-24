APP_NAME = export-utxos
VERSION = 1.0.0

.PHONY: all clean build-mac build-windows

all: clean build-mac build-windows

clean:
	rm -rf build

build-mac:
	mkdir -p build/mac
	GOOS="darwin" GOARCH="amd64" CGO_ENABLED="1" go build -o build/mac/$(APP_NAME)-$(VERSION)-mac main.go
	echo "macOS build completed successfully at build/mac/$(APP_NAME)-$(VERSION)-mac"

build-windows:
	mkdir -p build/windows
	GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build -o build/windows/$(APP_NAME)-$(VERSION)-windows.exe main.go
	echo "Windows build completed successfully at build/windows/$(APP_NAME)-$(VERSION)-windows.exe"
