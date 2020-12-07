CONSUL_CONFIG_PATH:=$(shell realpath ../consul)
ENVOY_PATH:=$(shell realpath ~/.getenvoy/builds/wasm/1.15/darwin/bin/envoy)

require-proper-consul-version:
	@consul version | grep -q "Revision 2b5365b4c" && echo "✅ proper consul version" || echo "❌ invalid consul version"
consul-version:
	@consul version
envoy-path:
	@echo "$(ENVOY_PATH)"
build:
	@echo "🛠 building filter example as replace_response.wasm"
	@tinygo build -o replace_response.wasm -scheduler=none -target=wasi -wasm-abi=generic main.go
	@mkdir -p /tmp/filters
	@cp replace_response.wasm /tmp/filters/replace_response.wasm
consul-dev-server: require-proper-consul-version
# Temporarily using/assuming a custom Consul version from the wasm-filters
# branch which includes the small, but neccessary changes to enable the
# wasm_filter proxy config options found in the service.hcl file.
#
# https://github.com/hashicorp/consul/tree/wasm-filters
#
# $ consul version
# Consul v1.9.0-dev
# Revision 2b5365b4c
	@echo "🌸 running consul dev agent"
	@consul agent -dev -config-dir=$(CONSUL_CONFIG_PATH)
consul-fake-service: require-proper-consul-version
	@echo "🐍 running simple python server on port 8080"
	@echo "📝 register consul service with replace_response.wasm filter"
	@consul services register service.hcl
	@mkdir -p /tmp/fake-service && cd /tmp/fake-service && python3 -m http.server 8080
consul-service-sidecar: require-proper-consul-version build
	@echo "💘 running consul connect envoy sidecar proxy for service 'web'"
	@echo "🕸 service 'web' is now available in the mesh using envoy"
	@consul connect envoy -sidecar-for="web" -envoy-binary="$(ENVOY_PATH)" -- -l debug
consul-connect: require-proper-consul-version
	@echo "💘 running consul connect development proxy"
	@echo "⬆️ upstream 'web' is available on on localhost:9090 through the consul service mesh"
	@consul connect proxy -service example -upstream web:9090
consul-test:
	@echo "🧪 running test with curl to verify custom filter is applied"
	@curl localhost:9090 -v 2>&1 | grep -q "replaced with wasm" && echo "✅ service is enabled with custom wasm filter" || echo "❌ service test failed, replaced body not found or not available (try again in a few seconds if just reloaded)"
# If it failed, remember this assumes these commands are all running separate terminals in parallel:
#
# 1. make consul-dev-server
# 2. make consul-fake-service
# 3. make consul-connect
#
# Use the Consul UI (http://localhost:8500) to help with debugging to verify everything is running (green checkmarks!)