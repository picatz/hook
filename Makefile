install-tinygo:
	@echo "installing tinygo"
	@brew tap tinygo-org/tools
	@brew install tinygo
install-getenvoy:
	@echo "installing getenvoy"
	@brew tap tetratelabs/getenvoy
	@brew install getenvoy
install-envoy:
	@echo "installing envoy"
	@getenvoy fetch wasm:1.15/darwin
install-deps: install-tinygo install-getenvoy install-envoy