.DEFAULT_GOAL := run
PROXY_FILE?=proxy_list
SUCCESS_PROXY_FILE?=success_proxy_list

run:
	@go install
	@proxyhealth

.PHONY: proxyfile
proxyfile:
	@cat proxy.json | jq -r '.["Proxies"][] | [.ip, .port|tostring] | join(":")' >> $(PROXY_FILE)
