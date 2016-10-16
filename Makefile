VERBOSE_FLAG = $(if $(VERBOSE),-verbose)

BUILD_FLAGS = -ldflags "\
	      -s -w \
	      "
TARGET_OSARCH="linux/amd64"

all: build sha256

build:
	mkdir -p build
	cd build && \
	  gox $(VERBOSE_FLAG) $(BUILD_FLAGS) \
	    -osarch=$(TARGET_OSARCH) ..

sha256:
	cd build && \
	  shasum -a 256 * | awk '{print "* " $$2 "\n    * " $$1}'

gox:
	go get github.com/mitchellh/gox

clean:
	if [ -d build ]; then \
	  rm -f build/mackerel-plugin-*; \
	  rmdir build; \
	fi

.PHONY: all build sha256 gox clean
