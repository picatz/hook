package host

import (
	"github.com/picatz/hook/pkg/types/buffer"
	"github.com/picatz/hook/pkg/types/log"
	"github.com/picatz/hook/pkg/types/metric"
	"github.com/picatz/hook/pkg/types/status"
	"github.com/picatz/hook/pkg/types/stream"
	"github.com/picatz/hook/pkg/types/wmap"
)

//export proxy_log
func ProxyLog(
	level log.Level,
	data *byte,
	size int,
) status.Type

//export proxy_send_local_response
func ProxySendLocalResponse(
	statusCode uint32,
	statusCodeData *byte,
	statusCodeSize int,
	bodyData *byte,
	bodySize int,
	headersData *byte,
	headersSize int,
	grpcStatus int32,
) status.Type

//export proxy_get_shared_data
func ProxyGetSharedData(
	keyData *byte,
	keySize int,
	returnValueData **byte,
	returnValueSize *int,
	returnCas *uint32,
) status.Type

//export proxy_set_shared_data
func ProxySetSharedData(
	keyData *byte,
	keySize int,
	valueData *byte,
	valueSize int,
	cas uint32,
) status.Type

//export proxy_register_shared_queue
func ProxyRegisterSharedQueue(
	nameData *byte,
	nameSize int,
	returnID *uint32,
) status.Type

//export proxy_resolve_shared_queue
func ProxyResolveSharedQueue(
	vmIDData *byte,
	vmIDSize int,
	nameData *byte,
	nameSize int,
	returnID *uint32,
) status.Type

//export proxy_dequeue_shared_queue
func ProxyDequeueSharedQueue(
	queueID uint32,
	returnValueData **byte,
	returnValueSize *int,
) status.Type

//export proxy_enqueue_shared_queue
func ProxyEnqueueSharedQueue(
	queueID uint32,
	valueData *byte,
	valueSize int,
) status.Type

//export proxy_get_header_map_value
func ProxyGetHeaderMapValue(
	wmapType wmap.Type,
	keyData *byte,
	keySize int,
	returnValueData **byte,
	returnValueSize *int,
) status.Type

//export proxy_add_header_map_value
func ProxyAddHeaderMapValue(
	wmapType wmap.Type,
	keyData *byte,
	keySize int,
	valueData *byte,
	valueSize int,
) status.Type

//export proxy_replace_header_map_value
func ProxyReplaceHeaderMapValue(
	wmapType wmap.Type,
	keyData *byte,
	keySize int,
	valueData *byte,
	valueSize int,
) status.Type

//export proxy_remove_header_map_value
func ProxyRemoveHeaderMapValue(
	wmapType wmap.Type,
	keyData *byte,
	keySize int,
) status.Type

//export proxy_get_header_map_pairs
func ProxyGetHeaderMapPairs(
	wmapType wmap.Type,
	returnValueData **byte,
	returnValueSize *int,
) status.Type

//export proxy_set_header_map_pairs
func ProxySetHeaderMapPairs(
	wmapType wmap.Type,
	mapData *byte,
	mapSize int,
) status.Type

//export proxy_get_buffer_bytes
func ProxyGetBufferBytes(
	bt buffer.Type,
	start int,
	maxSize int,
	returnBufferData **byte,
	returnBufferSize *int,
) status.Type

//export proxy_set_buffer_bytes
func ProxySetBufferBytes(
	bt buffer.Type,
	start int,
	maxSize int,
	bufferData *byte,
	bufferSize int,
) status.Type

//export proxy_continue_stream
func ProxyContinueStream(streamType stream.Type) status.Type

//export proxy_close_stream
func ProxyCloseStream(streamType stream.Type) status.Type

//export proxy_http_call
func ProxyHTTPCall(
	upstreamData *byte,
	upstreamSize int,
	headerData *byte,
	headerSize int,
	bodyData *byte,
	bodySize int,
	trailersData *byte,
	trailersSize int,
	timeout uint32,
	calloutIDPtr *uint32,
) status.Type

//export proxy_set_tick_period_milliseconds
func ProxySetTickPeriodMilliseconds(period uint32) status.Type

//export proxy_get_current_time_nanoseconds
func ProxyGetCurrentTimeNanoseconds(returnTime *int64) status.Type

//export proxy_set_effective_context
func ProxySetEffectiveContext(contextID uint32) status.Type

//export proxy_done
func ProxyDone() status.Type

//export proxy_define_metric
func ProxyDefineMetric(
	metricType metric.Type,
	metricNameData *byte,
	metricNameSize int,
	returnMetricIDPtr *uint32,
) status.Type

//export proxy_increment_metric
func ProxyIncrementMetric(metricID uint32, offset int64) status.Type

//export proxy_record_metric
func ProxyRecordMetric(metricID uint32, value uint64) status.Type

//export proxy_get_metric
func ProxyGetMetric(metricID uint32, returnMetricValue *uint64) status.Type

//export proxy_get_property
func ProxyGetProperty(
	pathData *byte,
	pathSize int,
	returnValueData **byte,
	returnValueSize *int,
) status.Type

//export proxy_set_property
func ProxySetProperty(
	pathData *byte,
	pathSize int,
	valueData *byte,
	valueSize int,
) status.Type
