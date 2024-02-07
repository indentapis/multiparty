static:
	mkdir -p static

static/multiparty.js.wasm: static
	GOOS=js GOARCH=wasm \
		go build \
			-ldflags="-s -w" \
			-o static/multiparty.js.wasm \
			cmd/wasm-js/main.go

static/multiparty.js.wasm.gz: static static/multiparty.js.wasm
	gzip --best -f < static/multiparty.js.wasm > static/multiparty.js.wasm.gz

wasm: static/multiparty.js.wasm static/multiparty.js.wasm.gz
	@echo "WASM built"

static/multiparty.wasm: static
	GOOS=wasip1 GOARCH=wasm \
		go build \
			-ldflags="-s -w" \
			-o static/multiparty.wasm \
			cmd/wasm/main.go

static/multiparty.wasm.gz: static static/multiparty.wasm
	gzip --best -f < static/multiparty.wasm > static/multiparty.wasm.gz

wasi: static/multiparty.wasm static/multiparty.wasm.gz
	@echo "WASI built"

clean:
	rm -rf static
	@echo "\ncleaned"