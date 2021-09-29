GOPATH=$(shell go env GOPATH)
GOLANGCI_LINT=$(GOPATH)/bin/golangci-lint
GOFUZZBUILD = $(GOPATH)/bin/go-fuzz-build
GOFUZZ = $(GOPATH)/bin/go-fuzz

.PHONY: test
test:
	@go test .

.PHONY: format
format:
	@gofmt -s -w *.go

.PHONY: fuzz
fuzz: $(GOFUZZBUILD) $(GOFUZZ)
	@go get -d github.com/dvyukov/go-fuzz/go-fuzz-dep
	@GO111MODULE=on $(GOFUZZBUILD)
	@go mod tidy
	@GO111MODULE=on $(GOFUZZ) -workdir=_fuzz

$(GOFUZZBUILD):
	cd / && go get -u github.com/dvyukov/go-fuzz/go-fuzz-build

$(GOFUZZ):
	cd / && go get -u github.com/dvyukov/go-fuzz/go-fuzz github.com/dvyukov/go-fuzz/go-fuzz-dep