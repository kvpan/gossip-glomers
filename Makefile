GO := go
OUTPUT_NANE := gossip-glomers

build:
	$(GO) build -o build/$(OUTPUT_NANE)

echo: build
	@maelstrom test -w echo --bin build/$(OUTPUT_NANE) --node-count 1 --time-limit 10
