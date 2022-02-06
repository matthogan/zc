
# Image URL to use all building/pushing image targets
IMG ?= zccn
TAG ?= 0.1.0
SONAR_URL ?= http://localhost:9000
SONAR_KEY ?= github.com.matthogan.zc.cmd.cn
SONAR_TOKEN ?= 36bf40c0d7cd898009d4bf5d2b52483c0743f025
SOURCE ?= pkg/container/digest.go
PACKAGE ?= container
DEST ?= pkg/container/digest.go

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

.PHONY: all
all: build

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: docs
docs: ## Generate docs in doc/
	cmd/help/gendocs.sh

##@ Development

.PHONY: fmt
fmt: ## Run go fmt against code.
	go fmt ./...

.PHONY: vet
vet: ## Run go vet against code.
	go vet ./...

.PHONY: tidy
tidy: ## Run go mod tidy against code.
	go mod tidy

.PHONY: test
test: fmt vet ## Run tests.
	go test ./... -coverprofile cover.out

.PHONY: sonar
sonar: test
## brew install sonar-scanner
	sonar-scanner \
		-Dsonar.projectKey=${SONAR_KEY} \
		-Dsonar.sources=cmd/cn,pkg \
		-Dsonar.tests=cmd/cn,pkg \
		-Dsonar.test.inclusions=**/*_test.go \
		-Dsonar.exclusions=**/*_test.go,pkg/oci/**/* \
		-Dsonar.go.coverage.reportPaths=cover.out \
		-Dsonar.host.url=${SONAR_URL} \
		-Dsonar.login=${SONAR_TOKEN}

.PHONY: mockgen
mockgen:
## go install github.com/golang/mock/mockgen@v1.6.0
## go get github.com/golang/mock/gomock
	~/go/bin/mockgen --source ${SOURCE} \
		-package ${PACKAGE} \
		-copyright_file copyright \
		-destination mock/${DEST}
	ls -l mock/${DEST}

##@ Build

.PHONY: docker-build
docker-build: test ## Build docker image with the manager.
	DOCKER_BUILDKIT=0 docker build --progress tty -t ${IMG}:${TAG} .

.PHONY: docker-push
docker-push: ## Push docker image with the manager.
	docker push ${IMG}:${TAG}

.PHONY: darwin-build
darwin-build: fmt vet ## Run tests.
	goreleaser build --snapshot --rm-dist --id darwin-amd64

.PHONY: linux-build
linux-build: fmt vet ## Run tests.
	goreleaser build --snapshot --rm-dist --id linux

##@ Release
.PHONY: release
release: 
	git tag -a v${TAG} -m "v${TAG}"
	git push origin v${TAG}
	goreleaser release
