package host

import (
	"fmt"

	"github.com/picatz/hook/pkg/call/vm/state"
	"github.com/picatz/hook/pkg/types/action"
	"github.com/picatz/hook/pkg/types/peer"
)

// Internal function that are called by Envoy into this runtime.

//export proxy_abi_version_0_2_0
func proxyABIVersion() {}

//export malloc
func malloc(size uint) *byte {
	// https://tinygo.org/lang-support/#garbage-collection
	buffer := make([]byte, size)
	start := &buffer[0]
	return start
}

//export proxy_on_vm_start
func proxyOnVMStart(rootContextID uint32, vmConfigurationSize int) bool {
	root := state.GetRootContext(rootContextID)
	state.SetActiveContextID(rootContextID)
	return root.Context.OnVMStart(vmConfigurationSize)
}

//export proxy_on_configure
func proxyOnConfigure(rootContextID uint32, pluginConfigurationSize int) bool {
	root := state.GetRootContext(rootContextID)
	state.SetActiveContextID(rootContextID)
	return root.Context.OnPluginStart(pluginConfigurationSize)
}

//export proxy_on_new_connection
func proxyOnNewConnection(contextID uint32) action.Type {
	ctx := state.GetStreamContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnNewConnection()
}

//export proxy_on_downstream_data
func proxyOnDownstreamData(contextID uint32, dataSize int, endOfStream bool) action.Type {
	ctx := state.GetStreamContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnDownstreamData(dataSize, endOfStream)
}

//export proxy_on_downstream_connection_close
func proxyOnDownstreamConnectionClose(contextID uint32, peerType peer.Type) {
	ctx := state.GetStreamContext(contextID)
	state.SetActiveContextID(contextID)
	ctx.OnDownstreamClose(peerType)
}

//export proxy_on_upstream_data
func proxyOnUpstreamData(contextID uint32, dataSize int, endOfStream bool) action.Type {
	ctx := state.GetStreamContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnUpstreamData(dataSize, endOfStream)
}

//export proxy_on_upstream_connection_close
func proxyOnUpstreamConnectionClose(contextID uint32, peerType peer.Type) {
	ctx := state.GetStreamContext(contextID)
	state.SetActiveContextID(contextID)
	ctx.OnUpstreamClose(peerType)
}

//export proxy_on_request_headers
func proxyOnRequestHeaders(contextID uint32, numHeaders int, endOfStream bool) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnRequestHeaders(numHeaders, endOfStream)
}

//export proxy_on_request_body
func proxyOnRequestBody(contextID uint32, bodySize int, endOfStream bool) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnRequestBody(bodySize, endOfStream)
}

//export proxy_on_request_trailers
func proxyOnRequestTrailers(contextID uint32, numTrailers int) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnRequestTrailers(numTrailers)
}

//export proxy_on_response_headers
func proxyOnResponseHeaders(contextID uint32, numHeaders int, endOfStream bool) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnResponseHeaders(numHeaders, endOfStream)
}

//export proxy_on_response_body
func proxyOnResponseBody(contextID uint32, bodySize int, endOfStream bool) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnResponseBody(bodySize, endOfStream)
}

//export proxy_on_response_trailers
func proxyOnResponseTrailers(contextID uint32, numTrailers int) action.Type {
	ctx := state.GetHTTPContext(contextID)
	state.SetActiveContextID(contextID)
	return ctx.OnResponseTrailers(numTrailers)
}

//export proxy_on_http_call_response
func proxyOnHttpCallResponse(rootContextID, calloutID uint32, numHeaders, bodySize, numTrailers int) {
	root := state.GetRootContext(rootContextID)

	cb := root.HTTPCallbacks[calloutID]
	if cb == nil {
		panic(fmt.Errorf("given an invalid calloutID %d for rootContextID %d", calloutID, rootContextID))
	}

	delete(root.HTTPCallbacks, calloutID)

	state.SetActiveContextID(cb.CallerContextID)

	cb.Callback(numHeaders, bodySize, numTrailers)
}

//export proxy_on_tick
func proxyOnTick(rootContextID uint32) {
	root := state.GetRootContext(rootContextID)
	root.Context.OnTick()
}

//export proxy_on_queue_ready
func proxyOnQueueReady(rootContextID, queueID uint32) {
	root := state.GetRootContext(rootContextID)
	state.SetActiveContextID(rootContextID)
	root.Context.OnQueueReady(queueID)
}

//export proxy_on_context_create
func proxyOnContextCreate(contextID uint32, rootContextID uint32) {
	if rootContextID == 0 {
		state.Current.CreateRootContext(contextID)
	} else if state.Current.NewHttpContext != nil {
		state.Current.CreateHttpContext(contextID, rootContextID)
	} else if state.Current.NewStreamContext != nil {
		state.Current.CreateStreamContext(contextID, rootContextID)
	} else {
		panic(fmt.Errorf("invalid contextID %d for rootContextID %d", contextID, rootContextID))
	}
}

//export proxy_on_done
func proxyOnDone(contextID uint32) bool {
	defer func() {
		delete(state.Current.ContextIDToRootID, contextID)
	}()
	if ctx, ok := state.Current.StreamContexts[contextID]; ok {
		state.SetActiveContextID(contextID)
		delete(state.Current.StreamContexts, contextID)
		ctx.OnStreamDone()
		return true
	} else if ctx, ok := state.Current.HTTPContexts[contextID]; ok {
		state.SetActiveContextID(contextID)
		ctx.OnStreamDone()
		delete(state.Current.HTTPContexts, contextID)
		return true
	} else if ctx, ok := state.Current.RootContexts[contextID]; ok {
		state.SetActiveContextID(contextID)
		response := ctx.Context.OnVMDone()
		delete(state.Current.RootContexts, contextID)
		return response
	} else {
		panic(fmt.Errorf("invalid contextID %d", contextID))
	}
}

// Unimplemented:

//export proxy_on_foreign_function
func proxyOnForeignFunction(rootContextID, funcID, dataSize uint32) {
	// TODO
}

//export proxy_validate_configuration
func proxyValidateConfiguration(rootContextID, configurationSize uint32) uint32 {
	// TODO
	return 1
}
