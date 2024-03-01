all: deps tidy fmt build test lint

op=op run --env-file="./.env" --

opact=op run --env-file="./.env.act" --

#-------------------------------------------------------------------------
# Variables
# ------------------------------------------------------------------------
env=CGO_ENABLED=0

pyenv=.venv/bin

#-------------------------------------------------------------------------
# Targets
#-------------------------------------------------------------------------
deps:
	python3 -m venv .venv --clear

	$(pyenv)/pip install --upgrade pre-commit
	$(pyenv)/pip install --upgrade detect-secrets
	$(pyenv)/detect-secrets scan > .secrets.baseline
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/goreleaser/goreleaser@latest
	go install github.com/google/go-licenses@latest
	go install github.com/google/addlicense@latest

ci-test-deps:
	# install act
	if [ ! -d .act ]; then make install-act; git clone git@github.com:nektos/act.git .act; fi
	cd .act && git pull && sudo make install

test: lint
	CGO_ENABLED=1 go test -cover -failfast -race ./...

fuzz:
	$(pyenv)/python ./scripts/fuzz_go.py

bench:
	go test -bench=. -benchmem ./...

lint: tidy
	$(pyenv)/python ./scripts/license.py
	golangci-lint run
	$(pyenv)/pre-commit run --all-files

build: update upgrade tidy lint test
	$(env) go build ./...

release: build-ci
	goreleaser release --snapshot --clean

upgrade:
	$(pyenv)/pre-commit autoupdate
	go get -u ./...

update:
	git submodule update --recursive

fmt:
	gofmt -s -w .

tidy: fmt
	go mod tidy

clean:
	rm -rf dist coverage .act .venv

#-------------------------------------------------------------------------
# Git targets
#-------------------------------------------------------------------------

tag:
	$(pyenv)/python ./scripts/tag.py

#-------------------------------------------------------------------------
# CI targets
#-------------------------------------------------------------------------
build-ci: deps
	$(env) go build ./...
	CGO_ENABLED=1 go test \
				-cover \
				-covermode=atomic \
				-coverprofile=coverage.txt \
				-failfast \
				-race ./...
	make fuzz FUZZ_TIME=10

bench-ci: build-ci
	go test -bench=. ./... | tee output.txt

test-ci:
	DOCKER_HOST=$(shell docker context inspect --format='{{json .Endpoints.docker.Host}}' $(shell docker context show)) \
				$(opact) act \
					-s GIT_CREDENTIALS \
					-s GITHUB_TOKEN="$(shell gh auth token)" \
					--var GO_VERSION \
					--var ALERT_CC_USERS

#-------------------------------------------------------------------------
# Force targets
#-------------------------------------------------------------------------

FORCE:

#-------------------------------------------------------------------------
# Phony targets
#-------------------------------------------------------------------------

.PHONY: build test lint fuzz bench fmt tidy clean release update upgrade deps translate test-act ci-test-deps
