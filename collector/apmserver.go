package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

//Apmserver json structure
type Apmserver struct {
	Acm       Acm       `json:"acm"`
	Decoder   Decoder   `json:"decoder"`
	Jaeger    Jaeger    `json:"jaeger"`
	Processor Processor `json:"processor"`
	Profile   Profile   `json:"profile"`
	Root      Root      `json:"root"`
	Sampling  Sampling  `json:"sampling"`
	Server    Server    `json:"server"`
	Sourcemap Sourcemap `json:"sourcemap"`
}

type Request struct {
	Count float64 `json:"count"`
}

type Errors struct {
	Closed       float64 `json:"closed"`
	Count        float64 `json:"count"`
	Decode       float64 `json:"decode"`
	Forbidden    float64 `json:"forbidden"`
	Internal     float64 `json:"internal"`
	Invalidquery float64 `json:"invalidquery"`
	Method       float64 `json:"method"`
	Notfound     float64 `json:"notfound"`
	Queue        float64 `json:"queue"`
	Ratelimit    float64 `json:"ratelimit"`
	Toolarge     float64 `json:"toolarge"`
	Unauthorized float64 `json:"unauthorized"`
	Unavailable  float64 `json:"unavailable"`
	Validate     float64 `json:"validate"`
	Invalid      float64 `json:"invalid"`
	Server       float64 `json:"server"`
}

type Valid struct {
	Accepted    float64 `json:"accepted"`
	Count       float64 `json:"count"`
	Notmodified float64 `json:"notmodified"`
	Ok          float64 `json:"ok"`
}

type Response struct {
	Count  float64 `json:"count"`
	Errors Errors  `json:"errors"`
	Valid  Valid   `json:"valid"`
}

type Acm struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Unset    float64  `json:"unset"`
}

type Deflate struct {
	ContentLength float64 `json:"content-length"`
	Count         float64 `json:"count"`
}

type Gzip struct {
	ContentLength float64 `json:"content-length"`
	Count         float64 `json:"count"`
}

type MissingContentLength struct {
	Count float64 `json:"count"`
}

type Reader struct {
	Count float64 `json:"count"`
}

type Uncompressed struct {
	ContentLength float64 `json:"content-length"`
	Count         float64 `json:"count"`
}

type Decoder struct {
	Deflate              Deflate              `json:"deflate"`
	Gzip                 Gzip                 `json:"gzip"`
	MissingContentLength MissingContentLength `json:"missing-content-length"`
	Reader               Reader               `json:"reader"`
	Uncompressed         Uncompressed         `json:"uncompressed"`
}

type Dropped struct {
	Count float64 `json:"count"`
}

type Received struct {
	Count float64 `json:"count"`
}

type Event struct {
	Dropped  Dropped  `json:"dropped"`
	Received Received `json:"received"`
}

type Collect struct {
	Event    Event    `json:"event"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Sampling struct {
	Event               Event    `json:"event"`
	Request             Request  `json:"request"`
	Response            Response `json:"response"`
	TransactionsDropped float64  `json:"transactions_dropped"`
}
type Grpc struct {
	Collect  Collect  `json:"collect"`
	Sampling Sampling `json:"sampling"`
}

type HTTP struct {
	Event    Event    `json:"event"`
	Request  Request  `json:"request"`
	Response Response `json:"response"`
}

type Jaeger struct {
	Grpc Grpc `json:"grpc"`
	HTTP HTTP `json:"http"`
}
type Error struct {
	Frames          float64 `json:"frames"`
	Stacktraces     float64 `json:"stacktraces"`
	Transformations float64 `json:"transformations"`
}

type Metric struct {
	Transformations float64 `json:"transformations"`
}

type Decoding struct {
	Count  float64 `json:"count"`
	Errors float64 `json:"errors"`
}

type Validation struct {
	Count  float64 `json:"count"`
	Errors float64 `json:"errors"`
}

type Sourcemap struct {
	Counter    float64    `json:"counter"`
	Decoding   Decoding   `json:"decoding"`
	Validation Validation `json:"validation"`
	Request    Request    `json:"request"`
	Response   Response   `json:"response"`
	Unset      float64    `json:"unset"`
}

type Span struct {
	Frames          float64 `json:"frames"`
	Stacktraces     float64 `json:"stacktraces"`
	Transformations float64 `json:"transformations"`
}

type Stream struct {
	Accepted float64 `json:"accepted"`
	Errors   Errors  `json:"errors"`
}

type Transaction struct {
	Transformations float64 `json:"transformations"`
}

type Processor struct {
	Error       Error       `json:"error"`
	Metric      Metric      `json:"metric"`
	Sourcemap   Sourcemap   `json:"sourcemap"`
	Span        Span        `json:"span"`
	Stream      Stream      `json:"stream"`
	Transaction Transaction `json:"transaction"`
}

type Profile struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Unset    float64  `json:"unset"`
}

type Root struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Unset    float64  `json:"unset"`
}

type Server struct {
	Request  Request  `json:"request"`
	Response Response `json:"response"`
	Unset    float64  `json:"unset"`
}

type apmserverCollector struct {
	beatInfo *BeatInfo
	stats    *Stats
	metrics  exportedMetrics
}

// NewApmserverCollector constructor
func NewApmserverCollector(beatInfo *BeatInfo, stats *Stats) prometheus.Collector {
	return &apmserverCollector{
		beatInfo: beatInfo,
		stats:    stats,
		metrics: exportedMetrics{
			// ACM
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "request_count"),
					"apm-server.acm.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_count"),
					"apm-server.acm.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors_count"),
					"apm-server.acm.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "decode"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Decode },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "forbidden"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Forbidden },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "internal"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Internal },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "invalidquery"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Invalidquery },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "method"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Method },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "notfound"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Notfound },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "ratelimit"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Ratelimit },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "unauthorized"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Unauthorized },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "unavailable"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Unavailable },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_errors"),
					"apm-server.acm.response.errors",
					nil, prometheus.Labels{"error": "validate"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Errors.Validate },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_valid"),
					"apm-server.acm.response.valid",
					nil, prometheus.Labels{"status": "accepted"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Valid.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_valid_count"),
					"apm-server.acm.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_valid"),
					"apm-server.acm.response.valid",
					nil, prometheus.Labels{"status": "notmodified"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Valid.Notmodified },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "response_valid"),
					"apm-server.acm.response.valid",
					nil, prometheus.Labels{"status": "ok"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Response.Valid.Ok },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "acm", "unset"),
					"apm-server.acm.unset",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Acm.Unset },
				valType: prometheus.CounterValue,
			},
			// DECODER
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "deflate"),
					"apm-server.decoder.deflate",
					nil, prometheus.Labels{"content_length": "bytes"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Deflate.ContentLength },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "deflate_count"),
					"apm-server.decoder.deflate.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Deflate.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "gzip"),
					"apm-server.decoder.gzip",
					nil, prometheus.Labels{"content_length": "bytes"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Gzip.ContentLength },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "gzip_count"),
					"apm-server.decoder.gzip.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Gzip.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "missing_content_length_count"),
					"apm-server.decoder.missing-content-length.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.MissingContentLength.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "reader_count"),
					"apm-server.decoder.reader.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Reader.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "uncompressed"),
					"apm-server.decoder.uncompressed",
					nil, prometheus.Labels{"content_length": "bytes"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Uncompressed.ContentLength },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "decoder", "uncompressed_count"),
					"apm-server.decoder.uncompressed.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Decoder.Uncompressed.Count },
				valType: prometheus.CounterValue,
			},
			// JAEGER
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_event_dropped_count"),
					"apm-server.jaeger.grpc.collect.event.dropped.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Event.Dropped.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_event_received_count"),
					"apm-server.jaeger.grpc.collect.event.received.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Event.Received.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_request_count"),
					"apm-server.jaeger.grpc.collect.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_response_count"),
					"apm-server.jaeger.grpc.collect.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_response_errors_count"),
					"apm-server.jaeger.grpc.collect.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_collect_response_valid_count"),
					"apm-server.jaeger.grpc.collect.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Collect.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_event_dropped_count"),
					"apm-server.jaeger.grpc.sampling.event.dropped.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Event.Dropped.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_event_received_count"),
					"apm-server.jaeger.grpc.sampling.event.received.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Event.Received.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_request_count"),
					"apm-server.jaeger.grpc.sampling.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_response_count"),
					"apm-server.jaeger.grpc.sampling.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_response_errors_count"),
					"apm-server.jaeger.grpc.sampling.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "grpc_sampling_response_valid_count"),
					"apm-server.jaeger.grpc.sampling.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.Grpc.Sampling.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_event_dropped_count"),
					"apm-server.jaeger.http.event.dropped.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Event.Dropped.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_event_received_count"),
					"apm-server.jaeger.http.received.dropped.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Event.Received.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_request_count"),
					"apm-server.jaeger.http.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_response_count"),
					"apm-server.jaeger.http.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_response_errors_count"),
					"apm-server.jaeger.http.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "jaeger", "http_response_valid_count"),
					"apm-server.jaeger.http.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Jaeger.HTTP.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			// PROCESSOR
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "errors"),
					"apm-server.processor.error",
					nil, prometheus.Labels{"error": "frames"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Error.Frames },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "errors"),
					"apm-server.processor.error",
					nil, prometheus.Labels{"error": "stacktraces"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Error.Stacktraces },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "errors"),
					"apm-server.processor.error",
					nil, prometheus.Labels{"error": "transformations"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Error.Transformations },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "metric_transformations"),
					"apm-server.processor.metric.transformations",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Metric.Transformations },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "sourcemap_counter"),
					"apm-server.processor.sourcemap.counter",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Sourcemap.Counter },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "sourcemap_decoding"),
					"apm-server.processor.sourcemap.decoding",
					nil, prometheus.Labels{"decoding": "count"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Sourcemap.Decoding.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "sourcemap_decoding"),
					"apm-server.processor.sourcemap.decoding",
					nil, prometheus.Labels{"decoding": "errors"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Sourcemap.Decoding.Errors },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "sourcemap_validation"),
					"apm-server.processor.sourcemap.validation",
					nil, prometheus.Labels{"validation": "count"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Sourcemap.Validation.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "sourcemap_validation"),
					"apm-server.processor.sourcemap.validation",
					nil, prometheus.Labels{"validation": "errors"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Sourcemap.Validation.Errors },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "spans"),
					"apm-server.processor.span",
					nil, prometheus.Labels{"span": "frames"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Span.Frames },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "spans"),
					"apm-server.processor.span",
					nil, prometheus.Labels{"span": "stacktraces"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Span.Stacktraces },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "spans"),
					"apm-server.processor.span",
					nil, prometheus.Labels{"span": "transformations"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Span.Transformations },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_accepted"),
					"apm-server.processor.stream.accepted",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_errors"),
					"apm-server.processor.stream.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_errors"),
					"apm-server.processor.stream.errors",
					nil, prometheus.Labels{"error": "invalid"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Errors.Invalid },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_errors"),
					"apm-server.processor.stream.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_errors"),
					"apm-server.processor.stream.errors",
					nil, prometheus.Labels{"error": "server"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Errors.Server },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "stream_errors"),
					"apm-server.processor.stream.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Stream.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "processor", "transaction_transformations"),
					"apm-server.processor.stream.transaction.transformations",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Processor.Transaction.Transformations },
				valType: prometheus.CounterValue,
			},
			// PROFILE
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "request_count"),
					"apm-server.profile.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_count"),
					"apm-server.profile.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors_count"),
					"apm-server.profile.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "decode"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Decode },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "forbidden"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Forbidden },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "internal"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Internal },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "invalidquery"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Invalidquery },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "method"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Method },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "notfound"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Notfound },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "ratelimit"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Ratelimit },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "unauthorized"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Unauthorized },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "unavailable"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Unavailable },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_errors"),
					"apm-server.profile.response.errors",
					nil, prometheus.Labels{"error": "validate"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Errors.Validate },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_valid"),
					"apm-server.profile.response.valid",
					nil, prometheus.Labels{"status": "accepted"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Valid.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_valid_count"),
					"apm-server.profile.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_valid"),
					"apm-server.profile.response.valid",
					nil, prometheus.Labels{"status": "notmodified"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Valid.Notmodified },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "response_valid"),
					"apm-server.profile.response.valid",
					nil, prometheus.Labels{"status": "ok"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Response.Valid.Ok },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "profile", "unset"),
					"apm-server.profile.unset",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Profile.Unset },
				valType: prometheus.CounterValue,
			},
			// ROOT
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "request_count"),
					"apm-server.root.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_count"),
					"apm-server.root.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors_count"),
					"apm-server.root.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "decode"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Decode },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "forbidden"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Forbidden },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "internal"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Internal },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "invalidquery"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Invalidquery },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "method"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Method },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "notfound"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Notfound },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "ratelimit"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Ratelimit },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "unauthorized"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Unauthorized },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "unavailable"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Unavailable },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_errors"),
					"apm-server.root.response.errors",
					nil, prometheus.Labels{"error": "validate"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Errors.Validate },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_valid"),
					"apm-server.root.response.valid",
					nil, prometheus.Labels{"status": "accepted"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Valid.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_valid_count"),
					"apm-server.root.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_valid"),
					"apm-server.root.response.valid",
					nil, prometheus.Labels{"status": "notmodified"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Valid.Notmodified },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "response_valid"),
					"apm-server.root.response.valid",
					nil, prometheus.Labels{"status": "ok"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Response.Valid.Ok },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "root", "unset"),
					"apm-server.root.unset",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Root.Unset },
				valType: prometheus.CounterValue,
			},
			// SAMPLING
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sampling", "transactions_dropped"),
					"apm-server.sampling.transactions_dropped",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sampling.TransactionsDropped },
				valType: prometheus.CounterValue,
			},
			// SERVER
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "request_count"),
					"apm-server.server.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_count"),
					"apm-server.server.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors_count"),
					"apm-server.server.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "decode"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Decode },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "forbidden"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Forbidden },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "internal"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Internal },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "invalidquery"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Invalidquery },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "method"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Method },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "notfound"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Notfound },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "ratelimit"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Ratelimit },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "unauthorized"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Unauthorized },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "unavailable"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Unavailable },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_errors"),
					"apm-server.server.response.errors",
					nil, prometheus.Labels{"error": "validate"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Errors.Validate },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_valid"),
					"apm-server.server.response.valid",
					nil, prometheus.Labels{"status": "accepted"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Valid.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_valid_count"),
					"apm-server.server.response.valid.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_valid"),
					"apm-server.server.response.valid",
					nil, prometheus.Labels{"status": "notmodified"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Valid.Notmodified },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "response_valid"),
					"apm-server.server.response.valid",
					nil, prometheus.Labels{"status": "ok"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Response.Valid.Ok },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "server", "unset"),
					"apm-server.server.unset",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Server.Unset },
				valType: prometheus.CounterValue,
			},
			// SOURCEMAP
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "request_count"),
					"apm-server.sourcemap.request.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Request.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_count"),
					"apm-server.sourcemap.response.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "closed"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Closed },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors_count"),
					"apm-server.sourcemap.response.errors.count",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "decode"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Decode },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "forbidden"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Forbidden },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "internal"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Internal },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "invalidquery"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Invalidquery },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "method"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Method },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "notfound"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Notfound },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "queue"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Queue },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "ratelimit"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Ratelimit },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "toolarge"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Toolarge },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "unauthorized"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Unauthorized },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "unavailable"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Unavailable },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_errors"),
					"apm-server.sourcemap.response.errors",
					nil, prometheus.Labels{"error": "validate"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Errors.Validate },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_valid"),
					"apm-server.sourcemap.response.valid",
					nil, prometheus.Labels{"valid": "accepted"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Valid.Accepted },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_valid"),
					"apm-server.sourcemap.response.valid",
					nil, prometheus.Labels{"valid": "count"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Valid.Count },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_valid"),
					"apm-server.sourcemap.response.valid",
					nil, prometheus.Labels{"valid": "notmodified"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Valid.Notmodified },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "response_valid"),
					"apm-server.sourcemap.response.valid",
					nil, prometheus.Labels{"valid": "ok"},
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Response.Valid.Ok },
				valType: prometheus.CounterValue,
			},
			{
				desc: prometheus.NewDesc(
					prometheus.BuildFQName(beatInfo.Beat, "sourcemap", "unset"),
					"apm-server.sourcemap.unset",
					nil, nil,
				),
				eval:    func(stats *Stats) float64 { return stats.Apmserver.Sourcemap.Unset },
				valType: prometheus.CounterValue,
			},
		},
	}
}

// Describe returns all descriptions of the collector.
func (c *apmserverCollector) Describe(ch chan<- *prometheus.Desc) {

	for _, metric := range c.metrics {
		ch <- metric.desc
	}

}

// Collect returns the current state of all metrics of the collector.
func (c *apmserverCollector) Collect(ch chan<- prometheus.Metric) {

	for _, i := range c.metrics {
		ch <- prometheus.MustNewConstMetric(i.desc, i.valType, i.eval(c.stats))
	}

}
