.PHONY: tool
tool:
	@aqua install

.PHONY: ymlfmt
ymlfmt:
	@yamlfmt

.PHONY: ymlint
ymlint:
	@yamlfmt -lint

.PHONY: module
module: ## go modules update
	@go get -u -t ./...
	@go mod tidy
	@go mod vendor
