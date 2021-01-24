.PHONY: build
build:
	@go build ./...


.PHONY: test
test:
	@go test ./...


.PHONY: coverage
coverage:
	@go test -cover ./...


.PHONY: coverage-report
coverage-report:
	@go test -coverprofile=/tmp/coverage.txt ./...
	@go tool cover -html=/tmp/coverage.txt


.PHONY: lint
lint:
	@gofmt -l .
	@go vet ./...


.PHONY: docs-server
docs-server:
	@echo "Documentation @ http://127.0.0.1:6060"
	@godoc -http=:6060


.PHONY: find_todo
find_todo:
	@grep --color=always --include=\*.go -PnRe '(//|/*).*TODO' ./ || true


.PHONY: count
count:
	@echo "Lines of code:"
	@find . -type f -name "*.go" | xargs wc -l
