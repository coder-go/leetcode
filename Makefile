clean:
	find ./ -name '*.log' -delete

lint: clean
	golangci-lint run --modules-download-mode=vendor --timeout=60m ./...
