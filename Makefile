define make_build
    GOOS=$(1) GOARCH=$(2) go build -o builds/$(1)/$(2)/$(3)
	cp -f *.yaml builds/$(1)/$(2)/
	cd builds/$(1)/$(2) && rm -f webstaticer.zip && zip --recurse-paths --move webstaticer.zip . && cd -
endef

# Batch build
build: deps build-freebsd build-linux build-macosx build-windows

# Dependencies
deps:
	go mod tidy && go mod vendor

# Freebsd
build-freebsd:
	$(call make_build,freebsd,amd64,webstaticer)

# Linux
build-linux:
	$(call make_build,linux,amd64,webstaticer)

# MacOSX
build-macosx:
	$(call make_build,darwin,amd64,webstaticer)
	$(call make_build,darwin,arm64,webstaticer)

# Windows
build-windows:
	$(call make_build,windows,amd64,webstaticer.exe)
