ifneq (,$(wildcard ./.env))
	include .env
	export
endif

GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)
.DEFAULT_GOAL := help

help: ## Show this help.
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  ${YELLOW}%-24s${GREEN}%s${RESET}\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## run tests
	go test -v ./...

bench: ## run benchmarks
	go test -bench=. ./...

tag: ## tag the current commit
	git tag -a $(VERSION) -m "Release $(VERSION)"

tag-verify: ## verify the tag
	git tag -v $(VERSION)

tag-push: ## push the tag
	git push origin $(VERSION)

release-local: ## dry run a release
	goreleaser release --clean --skip=publish

release: ## release a new version
	goreleaser release --clean
