CONTRACT_DIR ?= ../medici-contracts/artifacts/contracts

updateabi:
	cat $(CONTRACT_DIR)/Modules/Strategies/AbstractStrategy.sol/AbstractStrategy.json \
		| jq '.abi' > abis/AbstractStrategy.json
	cat $(CONTRACT_DIR)/Interfaces/ERC20.sol/ERC20.json \
		| jq '.abi' > abis/ERC20.json

.PHONY: updateabi

abigen:
	abigen --abi abis/AbstractStrategy.json --pkg strategy --type Strategy --out pkg/contracts/strategy/strategy.go
	abigen --abi abis/ERC20.json --pkg token --type Token --out pkg/contracts/token/token.go
	abigen --abi abis/GnosisSafe.json --pkg safe --type GnosisSafe --out pkg/contracts/safe/safe.go

.PHONY: abigen

mock:
	mockgen -package=testdata -destination=pkg/harvests/testdata/writer_mock.go -source=pkg/harvests/writers/writer.go Writer
	mockgen -package=testdata -destination=pkg/balances/testdata/writer_mock.go -source=pkg/balances/writers/writer.go Writer
	mockgen -package=testdata -destination=pkg/automator/testdata/sender_mock.go -source=pkg/automator/sender/sender.go Sender
	mockgen -package=testdata -destination=pkg/automator/testdata/oracle_mock.go -source=pkg/oracles/priceoracle.go PriceOracle
	mockgen -package=testdata -destination=pkg/automator/testdata/logger_mock.go -source=pkg/automator/logger/logger.go Logger
	mockgen -package=testdata -destination=pkg/automator/testdata/aggregator_mock.go -source=pkg/aggregators/aggregator.go Aggregator

.PHONY: mock

contract:
	abigen --sol pkg/testutils/strategy_mock.sol --pkg testutils > pkg/testutils/strategy_mock.go

build:
	 env GOOS=linux GOARCH=amd64 go build -o medici
.PHONY: build

cover:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out
	rm -f cover.out
.PHONY: cover
