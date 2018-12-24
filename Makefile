.PHONY: fmt lint vet test check build

fmt:
	@set -e; \
	trap 'rm -f .fmt.diff' EXIT; \
	goimports -d . > .fmt.diff; \
	if [ -s .fmt.diff ]; then \
		cat .fmt.diff; \
		exit 1; \
	fi

lint:
	@set -e; \
	golint -set_exit_status ./...

vet:
	@set -e; \
	go vet ./...

test:
	@set -e; \
	go test -v ./...

check: fmt lint vet test

build:
	@set -e; \
	cd cmd; \
	for CMD in $$(ls); \
	do \
		cd $$CMD; \
		CGO_ENABLED=0 go build -ldflags "-X main.buildID=$$(git rev-parse HEAD)"; \
		cd ..; \
	done