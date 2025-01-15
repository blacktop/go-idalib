.PHONY: bump
bump:
	@echo "🚀 Bumping Version"
	git tag $(shell svu patch)
	git push --tags

.PHONY: build
build:
	@echo "🚀 Building Version $(shell svu current)"
	@cd cmd/ida; go build -o ../../ida main.go

.PHONY: clean
clean:
	@echo "🚀 Cleaning"
	@go clean --cache && go build -x --work