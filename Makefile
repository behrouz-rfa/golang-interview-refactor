

.PHONY:
export
test-integration:
	go test -v -count=1 -tags=integration ./tests/it/...

test-e2e:
	go test -v -count=1 -tags=e2e ./tests/e2e/...

run:
	go run ./cmd/web-api/...

lint:
	gci write .
	gofumpt -l -w .
	golangci-lint run  -v