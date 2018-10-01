GO ?= GO111MODULE=on go

depsdev: ghch
	$(GO) get github.com/laher/goxc

run:
	$(GO) run cmd/hoi/main.go

goxc:
	GO111MODULE=on goxc -tasks='xc archive' -bc 'linux windows darwin' -d $(WERCKER_OUTPUT_DIR)/ -resources-include='README*'

build:
	$(GO) build -o hoi cmd/hoi/main.go

test:
	$(GO) test -cover -v $(shell go list ./... | grep -v vendor)
