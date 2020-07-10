BINARY		= fund
PackAge		= ca-services
GitBaseUrl 	= github.com
PackAge		= fund
GithubUser  = travelliu
PACKAGE  	= $(GitBaseUrl)/travelliu/${PackAge}
BASE     	= $(GOPATH)/src/$(PACKAGE)
GithubRep   = ${PackAge}
Version  	= $(shell grep 'var Version' version.go | sed -E 's/.*"(.+)"$$/v\1/')
BuildDate   ?= $(shell date +%FT%T%z)
GitTag  	?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)
PKGS     	= $(or $(PKG),$(shell env GO111MODULE=auto $(GO) list ./...|grep -v obs))
TESTPKGS 	= $(shell env GO111MODULE=auto $(GO) list -f '{{ if or .TestGoFiles .XTestGoFiles }}{{ .ImportPath }}{{ end }}' $(PKGS))
BIN      	= $(CURDIR)/bin
GO      	= go
TIMEOUT 	= 15
V 			= 0
Q 			= $(if $(filter 1,$V),,@)
M 			= $(shell printf "\033[34;1m▶\033[0m")
# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags '-X ${GitBaseUrl}/${PackAge}/Version=$(Version) \
			-X ${GitBaseUrl}/${PackAge}/BuildDate=$(BuildDate) \
			-X ${GitBaseUrl}/${PackAge}/GitTag=$(GitTag)'
.PHONY: all
all: version build

export GOPROXY=https://goproxy.cn


#$(BASE): ; $(info creating local GOPATH ...)
#	@mkdir -p $(dir $@)
#	@ln -sf $(CURDIR) $@
#	env


.PHONY: version
version: ; $(info) @ ## Print Version info
	$(info $(M) Will Build ${PackAge} Version   : $(Version))
	$(info $(M) Will Build ${PackAge} GitTag    : $(GitTag))
	$(info $(M) Will Build ${PackAge} BuildDate : $(BuildDate))


### ################################################
### tools
### ################################################

$(BIN):
	@mkdir -p $@

$(BIN)/%: | $(BIN) ; $(info $(M) building $(REPOSITORY)…)
	$(info $(M) building $(REPOSITORY)…)
	$Q tmp=$$(mktemp -d); \
	   GOPATH=$$tmp $(GO) get $(REPOSITORY) \
		|| ret=$$?; \
	   rm -rf $$tmp ; exit $$ret

GOLINT = golint
#$(BIN)/golint: REPOSITORY=golang.org/x/lint/golint
#
GOCOVMERGE = gocovmerge
#$(BIN)/gocovmerge: REPOSITORY=github.com/wadey/gocovmerge
#
GOCOV = gocov
#$(BIN)/gocov: REPOSITORY=github.com/axw/gocov/...
#
GOCOVXML = gocov-xml
#$(BIN)/gocov-xml: REPOSITORY=github.com/AlekSi/gocov-xml
#
GO2XUNIT = go2xunit
#$(BIN)/go2xunit: REPOSITORY=github.com/tebeka/go2xunit

CHGLOG = git-chglog

.PHONY: fmt lint generate

fmt: ; $(info $(M) running gofmt) @ ## Run go fmt on all source files
	$Q go fmt ./...

lint: fmt ; $(info $(M) running golint) @ ## Run golint
	$Q golint -set_exit_status $(PKGS)

generate: fmt ; $(info $(M) running go generate) @ ## Run go generate
	$Q go generate -v $(PKGS)


# ################################################
# Building
# ###############################################$
PREFIX?=
SUFFIX=
ifeq ($(GOOS),windows)
    SUFFIX=.exe
endif

# export GO111MODULE = auto
build: | $(BASE) # vendor swag
	$(info $(M) build executable $(PREFIX)bin/${PackAge}$(SUFFIX) begin)
	$Q $(GO) build \
		-tags debug \
		${LDFLAGS} \
		-o ${BINARY} main.go
	$(info $(M) build executable $(PREFIX)bin/${PackAge}$(SUFFIX) finished)

.PHONY: build

.PHONY: docker
docker:
	docker build -t ${GithubUser}/${BINARY} .

.PHONY: dockerBuildx
dockerBuildx:
	docker buildx build -t ${GithubUser}/${BINARY} --platform=linux/arm64,linux/amd64 --push .

# ################################################
# Cleaning
# ################################################

.PHONY: clean cleandocker
clean: ; $(info $(M) cleaning)	@ ## Run cleanup everything
	@rm -rf ${BINARY}
	@rm -rf $(GOPATH)
	@rm -rf test/coverage.*


#################################################
# help
#################################################

.PHONY: help
help: ; $(info) @ ## Print help info
	@grep -E '^[ a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'




# ################################################
# Tests
# ################################################
.PHONY: sonar
sonar:
	sonar-scanner

TEST_TARGETS := test-default test-bench test-short test-verbose test-race
.PHONY: $(TEST_TARGETS) test-xml check test tests
test-bench:   ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-short:   ARGS=-short        ## Run only short tests
test-verbose: ARGS=-v            ## Run tests in verbose mode with coverage reporting
test-race:    ARGS=-race         ## Run tests with race detector
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: fmt ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests
	$Q $(GO) test -v -count=1 -p 1 $(ARGS) $(TESTPKGS)

test-xml: fmt ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	@echo $(TESTPKGS)
	$Q mkdir -p test
	$Q 2>&1 $(GO) test -v -count=1 -p 1 $(TESTPKGS) | tee test/tests.output
#	$Q 2>&1 $(GO) test -timeout 20s | tee test/tests.output
	$(GO2XUNIT) -fail -input test/tests.output -output test/tests.xml

test-json: fmt ; $(info $(M) running $(NAME:%=% )tests…) @ ## Run tests with xUnit output
	$Q mkdir -p test
	$Q 2>&1 $(GO) test -json -count=1 -p 1 $(TESTPKGS) | tee test/test-report.json

COVERAGE_MODE = atomic
COVERAGE_PROFILE = test/profile.out
COVERAGE_XML = $(COVERAGE_DIR)/coverage.xml
COVERAGE_HTML = $(COVERAGE_DIR)/index.html
.PHONY: test-coverage
test-coverage: COVERAGE_DIR := $(CURDIR)/test/coverage.$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
test-coverage: clean fmt  ; $(info $(M) running coverage tests…) @ ## Run coverage tests
	$Q mkdir -p $(COVERAGE_DIR)/coverage
	$Q for pkg in $(TESTPKGS); do \
		$(GO) test -v -count=1 -p 1\
			-coverpkg=$$($(GO) list -f '{{ join .Deps "\n" }}' $$pkg | \
					grep '^$(PACKAGE)/' | \
					tr '\n' ',')$$pkg \
			-covermode=$(COVERAGE_MODE) \
			-coverprofile="$(COVERAGE_DIR)/coverage/`echo $$pkg | tr "/" "-"`.cover" $$pkg ;\
	 done
	$Q $(GOCOVMERGE) $(COVERAGE_DIR)/coverage/*.cover > $(COVERAGE_PROFILE)
	$Q $(GO) tool cover -html=$(COVERAGE_PROFILE) -o $(COVERAGE_HTML)
	$Q $(GOCOV) convert $(COVERAGE_PROFILE) | $(GOCOVXML) > $(COVERAGE_XML)

# ################################################
# Change Log
# ################################################
.PHONY: changelog
changelog: ; $(info $(M) updating changelog...)	@ ## Run updating changelog
	$(CHGLOG) --output CHANGELOG.md v2.2..
