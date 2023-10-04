.PHONY: proto

API=api
PROTO=api/_proto
GET=./scripts/get.sh
BUF := cd $(API) && PATH=$(notdir $(PROTO)):$(notdir $(PROTO))/protoc/bin:$$PATH buf

PKGS := \
	multiparty/v1

proto: $(PKGS)

$(PKGS): proto-deps
	$(BUF) generate --path $@ --template $@/buf.gen.yaml
	$(BUF) build . --path $@ --as-file-descriptor-set -o $@/$(notdir $@).descriptor.pb

proto-deps: \
 $(PROTO)/buf \
 $(PROTO)/googleapis \
 $(PROTO)/protoc/bin/protoc \
 $(PROTO)/protoc-gen-go

$(PROTO)/buf:
	@$(GET) buf

$(PROTO)/googleapis:
	@git submodule update --quiet --depth=1 --init $@ >/dev/null 2>&1

$(PROTO)/protoc/bin/protoc:
	@$(GET) protoc

$(PROTO)/protoc-gen-go:
	@$(GET) protoc-gen-go


