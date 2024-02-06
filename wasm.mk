static:
	mkdir -p static

static/multiparty.wasm: static
	GOOS=js GOARCH=wasm \
		go build \
			-ldflags="-s -w" \
			-o static/multiparty.wasm \
			cmd/wasm/main.go

static/multiparty.wasm.gz: static static/multiparty.wasm
	gzip --best -f < static/multiparty.wasm > static/multiparty.wasm.gz

wasm: static/multiparty.wasm static/multiparty.wasm.gz
	@echo "WASM built"

clean:
	rm -rf static
	@echo "\ncleaned"