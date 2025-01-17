# pinned versions
FABRIC_VERSION=2.2

TOP = .

all: install-tools checks unit-tests #integration-tests

.PHONY: install-tools
install-tools:
# Thanks for great inspiration https://marcofranssen.nl/manage-go-tools-via-go-modules
	@echo Installing tools from tools/tools.go
	@cd tools; cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

# include the checks target
include $(TOP)/checks.mk

.PHONY: unit-tests
unit-tests:
	@go test -cover $(shell go list ./... | grep -v '/integration/')
	cd integration/nwo/; go test -cover ./...

.PHONY: unit-tests-race
unit-tests-race:
	@export GORACE=history_size=7; go test -race -cover $(shell go list ./... | grep -v '/integration/')
	cd integration/nwo/; go test -cover ./...

.PHONY: docker-images
docker-images: fabric-docker-images orion-server-images monitoring-docker-images

.PHONY: fabric-docker-images
fabric-docker-images:
	docker pull hyperledger/fabric-baseos:$(FABRIC_VERSION)
	docker image tag hyperledger/fabric-baseos:$(FABRIC_VERSION) hyperledger/fabric-baseos:latest
	docker pull hyperledger/fabric-ccenv:$(FABRIC_VERSION)
	docker image tag hyperledger/fabric-ccenv:$(FABRIC_VERSION) hyperledger/fabric-ccenv:latest

.PHONY: monitoring-docker-images
monitoring-docker-images:
	docker pull hyperledger/explorer-db:latest
	docker pull hyperledger/explorer:latest
	docker pull prom/prometheus:latest
	docker pull grafana/grafana:latest

.PHONY: orion-server-images
orion-server-images:
	docker pull orionbcdb/orion-server:latest

.PHONY: integration-tests-dlog-fabric
integration-tests-dlog-fabric:
	cd ./integration/token/fungible/dlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-fabtoken-fabric
integration-tests-fabtoken-fabric:
	cd ./integration/token/fungible/fabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-dlog-orion
integration-tests-dlog-orion:
	cd ./integration/token/fungible/odlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-fabtoken-orion
integration-tests-fabtoken-orion:
	cd ./integration/token/fungible/ofabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-nft-dlog
integration-tests-nft-dlog:
	cd ./integration/token/nft/dlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-nft-fabtoken
integration-tests-nft-fabtoken:
	cd ./integration/token/nft/fabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-nft-dlog-orion
integration-tests-nft-dlog-orion:
	cd ./integration/token/nft/odlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-nft-fabtoken-orion
integration-tests-nft-fabtoken-orion:
	cd ./integration/token/nft/ofabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-dvp-fabtoken
integration-tests-dvp-fabtoken:
	cd ./integration/token/dvp/fabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-dvp-dlog
integration-tests-dvp-dlog:
	cd ./integration/token/dvp/dlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-interop-fabtoken
integration-tests-interop-fabtoken:
	cd ./integration/token/interop/fabtoken; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: integration-tests-interop-dlog
integration-tests-interop-dlog:
	cd ./integration/token/interop/dlog; ginkgo -keepGoing --slowSpecThreshold 60 .

.PHONY: tidy
tidy:
	@go mod tidy -compat=1.17

.PHONY: clean
clean:
	docker network prune -f
	docker container prune -f
	rm -rf ./integration/token/fungible/dlog/cmd/
	rm -rf ./integration/token/fungible/fabtoken/cmd/
	rm -rf ./integration/token/fungible/odlog/cmd/
	rm -rf ./integration/token/fungible/ofabtoken/cmd/
	rm -rf ./integration/token/nft/dlog/cmd/
	rm -rf ./integration/token/nft/fabtoken/cmd/
	rm -rf ./integration/token/nft/odlog/cmd/
	rm -rf ./integration/token/nft/ofabtoken/cmd/
	rm -rf ./integration/token/dvp/dlog/cmd/
	rm -rf ./integration/token/dvp/fabtoken/cmd/
	rm -rf ./integration/token/interop/fabtoken/cmd/
	rm -rf ./integration/token/interop/dlog/cmd/
	rm -rf ./samples/fungible/cmd
	rm -rf ./samples/nft/cmd

.PHONY: tokengen
tokengen:
	@go install ./cmd/tokengen
