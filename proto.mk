GET="./scripts/get.sh"

PROTO=_proto
BUF := PATH=$(PROTO):$(PROTO)/protoc/bin:$$PATH $(PROTO)/buf

proto: api/multiparty/v1/buf.gen.yaml \
 $(PROTO)/buf \
 $(PROTO)/googleapis \
 $(PROTO)/protoc/bin/protoc \
 $(PROTO)/protoc-gen-go
	$(BUF) generate --path api/multiparty/v1 --template $<
	$(BUF) build . --path api/multiparty/v1 --as-file-descriptor-set -o api/multiparty/v1/v1.descriptor.pb

$(PROTO)/buf:
	@$(GET) buf

$(PROTO)/googleapis:
	@git submodule update --quiet --depth=1 --init $@ >/dev/null 2>&1

$(PROTO)/protoc/bin/protoc:
	@$(GET) protoc

$(PROTO)/protoc-gen-go:
	@$(GET) protoc-gen-go


