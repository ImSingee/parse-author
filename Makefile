.PHONY: test
test:
	@go test .

.PHONY: format
format:
	@gofmt -s -w *.go