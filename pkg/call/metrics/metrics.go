package metrics

import (
	"github.com/picatz/hook/pkg/call/host"
	"github.com/picatz/hook/pkg/call/log"
	"github.com/picatz/hook/pkg/call/utils"
	"github.com/picatz/hook/pkg/types/metric"
	"github.com/picatz/hook/pkg/types/status"
)

type (
	Counter   uint32
	Gauge     uint32
	Histogram uint32
)

func (m Counter) ID() uint32 {
	return uint32(m)
}

func (m Counter) Get() uint64 {
	var val uint64
	st := host.ProxyGetMetric(m.ID(), &val)
	if err := status.AsError(st); err != nil {
		log.Criticalf("error while getting metric of %d: %v", m.ID(), status.AsError(st))
		return 0
	}
	return val
}

func (m Counter) Increment(offset uint64) {
	if err := status.AsError(host.ProxyIncrementMetric(m.ID(), int64(offset))); err != nil {
		log.Criticalf("increment %d by %d: %v", m.ID(), offset, err)
	}
}

func (m Gauge) ID() uint32 {
	return uint32(m)
}

func (m Gauge) Get() uint64 {
	var val uint64
	st := host.ProxyGetMetric(m.ID(), &val)
	if err := status.AsError(st); err != nil {
		log.Criticalf("error while getting metric of %d: %v", m.ID(), status.AsError(st))
		return 0
	}
	return val
}

func (m Gauge) Increment(offset uint64) {
	if err := status.AsError(host.ProxyIncrementMetric(m.ID(), int64(offset))); err != nil {
		log.Criticalf("error while incrementing %d by %d: %v", m.ID(), offset, err)
	}
}

func (m Histogram) ID() uint32 {
	return uint32(m)
}

func (m Histogram) Get() uint64 {
	var val uint64
	st := host.ProxyGetMetric(m.ID(), &val)
	if err := status.AsError(st); err != nil {
		log.Criticalf("error while getting metric of %d: %v", m.ID(), status.AsError(st))
		return 0
	}
	return val
}

func (m Histogram) Increment(offset uint64) {
	if err := status.AsError(host.ProxyIncrementMetric(m.ID(), int64(offset))); err != nil {
		log.Criticalf("error while incrementing %d by %d: %v", m.ID(), offset, err)
	}
}

func DefineCounter(name string) Counter {
	var id uint32
	ptr := utils.StringToBytePtr(name)
	st := host.ProxyDefineMetric(metric.Counter, ptr, len(name), &id)
	if err := status.AsError(st); err != nil {
		log.Criticalf("could not define metric of name %s: %v", name, status.AsError(st))
		return Counter(0)
	}
	return Counter(id)
}

func DefineGauge(name string) Gauge {
	var id uint32
	ptr := utils.StringToBytePtr(name)
	st := host.ProxyDefineMetric(metric.Gauge, ptr, len(name), &id)
	if err := status.AsError(st); err != nil {
		log.Criticalf("could not define metric of name %s: %v", name, status.AsError(st))
		return Gauge(0)
	}
	return Gauge(id)
}

func DefineHistogramMetric(name string) Histogram {
	var id uint32
	ptr := utils.StringToBytePtr(name)
	st := host.ProxyDefineMetric(metric.Histogram, ptr, len(name), &id)
	if err := status.AsError(st); err != nil {
		log.Criticalf("could not define metric of name %s: %v", name, status.AsError(st))
		return Histogram(0)
	}
	return Histogram(id)
}
