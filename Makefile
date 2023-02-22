.PHONY: tool
tool:
	@aqua install

.PHONY: ymlfmt
ymlfmt:
	@yamlfmt

.PHONY: ymlint
ymlint:
	@yamlfmt -lint
