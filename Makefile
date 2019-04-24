.PHONY: tests
tests:
	@go test \
		-v \
		-race \
		-count 5 \
		-failfast \
		./...

.PHONY: bench
bench:
	@go test \
		-v \
		-benchtime 10s \
		-count 5 \
		-benchmem \
		-bench=.
