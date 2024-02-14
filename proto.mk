.PHONY: proto

API=api
PROTO=api/_proto
GET=./scripts/get.sh
DEPS := buf protoc/bin/protoc protoc-gen-go
DEPS_PATHS := $(addprefix $(PROTO)/,$(DEPS))
BUF := cd $(API) && PATH=$(notdir $(PROTO)):$(notdir $(PROTO))/protoc/bin:$$PATH buf

PKGS := \
	multiparty/v1 \
	multiparty/prompt/v1

proto: $(PKGS)

$(PKGS): proto-deps
	$(BUF) generate --path $@ --template $@/buf.gen.yaml
	$(BUF) build . --path $@ --as-file-descriptor-set -o $@/$(notdir $@).descriptor.pb

proto-deps: $(DEPS_PATHS)
$(DEPS_PATHS):
	@$(GET) $(notdir $@)
