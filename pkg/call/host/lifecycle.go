package host

func SetEffectiveContext(contextID uint32) {
	ProxySetEffectiveContext(contextID)
}

func Done() {
	ProxyDone()
}
