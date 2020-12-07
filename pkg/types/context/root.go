package context

type Root interface {
	OnQueueReady(queueID uint32)
	OnTick()
	OnVMStart(vmConfigurationSize int) bool
	OnPluginStart(pluginConfigurationSize int) bool
	OnVMDone() bool
}

type RootDefault struct{}

func (*RootDefault) OnQueueReady(uint32)    {}
func (*RootDefault) OnTick()                {}
func (*RootDefault) OnVMStart(int) bool     { return true }
func (*RootDefault) OnPluginStart(int) bool { return true }
func (*RootDefault) OnVMDone() bool         { return true }
