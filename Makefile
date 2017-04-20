NOVENDOR = $(shell glide novendor | grep -v node_modules | grep -v static)
GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/chaki

all: codegen fix vet lint test

fix:
	go fix $(NOVENDOR)
.PHONY: fix

vet:
	go vet $(NOVENDOR)
.PHONY: vet

test:
	go test -v -cover $(NOVENDOR)
.PHONY: test

lint:
	printf "%s\n" "$(NOVENDOR)" | xargs -I {} sh -c 'golint -set_exit_status {}'
.PHONY: lint

clean:
	rm -rf $(DOCKER_CMD)
.PHONY: clean

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) main.go

dockerbin: $(DOCKER_CMD)
.PHONY: dockerbin

clean-codegen:
	rm static/build.go || true
.PHONY: clean-codegen

# The go-install is to cache the binary data build
static/build.go:
	yarn build
	go-bindata -o static/build.go -pkg static build/...
	go install

codegen: clean-codegen static/build.go
.PHONY: codegen

