package state

import (
	"fmt"
	"sync"

	"github.com/picatz/hook/pkg/types/context"
	"github.com/picatz/hook/pkg/types/http/callout"
)

type (
	HTTPCallback = struct {
		CallerContextID uint32
		Callback        callout.Callback
	}

	// TODO: verify and fix this
	GRPCCallback = struct{}

	RootContextState struct {
		Context       context.Root
		HTTPCallbacks map[uint32]*HTTPCallback
		GRPCCallbacks map[uint32]*GRPCCallback
	}

	// internal DSL types

	contextID     = uint32
	rootContextID = contextID

	rootContextFunc   = func(contextID uint32) context.Root
	httpContextFunc   = func(rootContextID, contextID uint32) context.HTTP
	streamContextFunc = func(rootContextID, contextID uint32) context.Stream
	grpcContextFunc   = func(rootContextID, contextID uint32) context.GRPC

	rootContexts   = map[uint32]*RootContextState
	httpContexts   = map[uint32]context.HTTP
	streamContexts = map[uint32]context.Stream
	grpcContexts   = map[uint32]context.GRPC

	contextIDToRootID = map[uint32]uint32
)

type State struct {
	ActiveContextID contextID

	ContextIDToRootID contextIDToRootID

	NewRootContext   rootContextFunc
	NewHttpContext   httpContextFunc
	NewStreamContext streamContextFunc
	NewGRPCContext   grpcContextFunc

	RootContexts   rootContexts
	HTTPContexts   httpContexts
	StreamContexts streamContexts
	GRPCContexts   grpcContexts
}

var currentStateMutex = sync.RWMutex{}

var Current = &State{
	ContextIDToRootID: make(contextIDToRootID),
	RootContexts:      make(rootContexts),
	HTTPContexts:      make(httpContexts),
	GRPCContexts:      make(grpcContexts),
	StreamContexts:    make(streamContexts),
}

func SetActiveContextID(contextID uint32) {
	Current.ActiveContextID = contextID
}

func SetNewRootContext(f rootContextFunc) {
	Current.NewRootContext = f
}

func SetNewHttpContext(f httpContextFunc) {
	Current.NewHttpContext = f
}

func SetNewStreamContext(f streamContextFunc) {
	Current.NewStreamContext = f
}

func RegisterHTTPCallout(calloutID uint32, callback callout.Callback) {
	r := Current.RootContexts[Current.ContextIDToRootID[Current.ActiveContextID]]
	r.HTTPCallbacks[calloutID] = &HTTPCallback{
		Callback:        callback,
		CallerContextID: Current.ActiveContextID,
	}
}

//go:inline
func (s *State) CreateRootContext(contextID uint32) {
	var ctx context.Root
	if s.NewRootContext == nil {
		ctx = &context.RootDefault{}
	} else {
		ctx = s.NewRootContext(contextID)
	}

	s.RootContexts[contextID] = &RootContextState{
		Context:       ctx,
		HTTPCallbacks: map[uint32]*HTTPCallback{},
	}
}

func (s *State) CreateStreamContext(contextID uint32, rootContextID uint32) {
	if _, ok := s.RootContexts[rootContextID]; !ok {
		panic(fmt.Errorf("invalid root context id %d while creating stream context %d", rootContextID, contextID))
	}

	if _, ok := s.StreamContexts[contextID]; ok {
		panic(fmt.Errorf("stream context id %d duplicated", contextID))
	}

	ctx := s.NewStreamContext(rootContextID, contextID)

	s.ContextIDToRootID[contextID] = rootContextID
	s.StreamContexts[contextID] = ctx
}

func (s *State) CreateHttpContext(contextID uint32, rootContextID uint32) {
	if _, ok := s.RootContexts[rootContextID]; !ok {
		panic(fmt.Errorf("invalid root context id %d while creating stream context %d", rootContextID, contextID))
	}

	if _, ok := s.HTTPContexts[contextID]; ok {
		panic(fmt.Errorf("http context id %d duplicated", contextID))
	}

	ctx := s.NewHttpContext(rootContextID, contextID)

	s.ContextIDToRootID[contextID] = rootContextID
	s.HTTPContexts[contextID] = ctx
}

func (s *State) registerHTTPCallout(calloutID uint32, callback callout.Callback) {
	r := s.RootContexts[s.ContextIDToRootID[s.ActiveContextID]]
	r.HTTPCallbacks[calloutID] = &HTTPCallback{
		Callback:        callback,
		CallerContextID: s.ActiveContextID,
	}
}

//go:inline
func (s *State) setActiveContextID(contextID uint32) {
	s.ActiveContextID = contextID
}

func VMStateReset() {
	currentStateMutex.Lock()
	defer currentStateMutex.Unlock()

	Current = &State{
		ContextIDToRootID: make(contextIDToRootID),
		RootContexts:      make(rootContexts),
		HTTPContexts:      make(httpContexts),
		StreamContexts:    make(streamContexts),
		GRPCContexts:      make(grpcContexts),
	}
}

func GetActiveContextID() uint32 {
	return Current.ActiveContextID
}

func GetRootContext(rootContextID uint32) *RootContextState {
	root, ok := Current.RootContexts[rootContextID]
	if !ok {
		panic(fmt.Errorf("invalid root context id %d", rootContextID))
	}
	return root
}

func GetHTTPContext(contextID uint32) context.HTTP {
	http, ok := Current.HTTPContexts[contextID]
	if !ok {
		panic(fmt.Errorf("invalid http context id %d", contextID))
	}
	return http
}

func GetStreamContext(contextID uint32) context.Stream {
	stream, ok := Current.StreamContexts[contextID]
	if !ok {
		panic(fmt.Errorf("invalid stream context id %d", contextID))
	}
	return stream
}
