BINARY := commongood
.PHONY: build
build: 
	go build -o=/tmp/bin/${BINARY} ./cmd/web

run: build
	/tmp/bin/${BINARY}

run/live:
		go run github.com/cosmtrek/air@v1.49.0 \
			--build.cmd "make build" --build.bin "/tmp/bin/${BINARY}" --build.delay "100" \
			--build.exclude_dir "" \
			--build.include_ext "go, ico" \
			--misc.clean_on_exit "true"
