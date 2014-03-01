all:
	@go build .

fmt:
	@go fmt ./...

release:
	@go build .
	@zip newrelic-rpm-check_linux64.zip newrelic-rpm-check
