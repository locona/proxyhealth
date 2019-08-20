.DEFAULT_GOAL := run

run:
	@go install
	@proxyhealth
