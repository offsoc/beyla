// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package gotracer

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpf_debugConnectionInfoT struct {
	S_addr [16]uint8
	D_addr [16]uint8
	S_port uint16
	D_port uint16
}

type bpf_debugEgressKeyT struct {
	S_port uint16
	D_port uint16
}

type bpf_debugGoAddrKeyT struct {
	Pid  uint64
	Addr uint64
}

type bpf_debugGoroutineMetadata struct {
	Parent    bpf_debugGoAddrKeyT
	Timestamp uint64
}

type bpf_debugGrpcClientFuncInvocationT struct {
	StartMonotimeNs uint64
	Cc              uint64
	Method          uint64
	MethodLen       uint64
	Tp              bpf_debugTpInfoT
	Flags           uint64
}

type bpf_debugGrpcSrvFuncInvocationT struct {
	StartMonotimeNs uint64
	Stream          uint64
	Tp              bpf_debugTpInfoT
}

type bpf_debugGrpcTransportsT struct {
	Conn bpf_debugConnectionInfoT
	_    [4]byte
	Tp   bpf_debugTpInfoT
	Type uint8
	_    [7]byte
}

type bpf_debugHttpClientDataT struct {
	Method        [7]uint8
	Path          [100]uint8
	_             [5]byte
	ContentLength int64
	Pid           struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	_ [4]byte
}

type bpf_debugHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_debugTpInfoT
}

type bpf_debugKafkaClientReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	_               [7]byte
	Conn            bpf_debugConnectionInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
}

type bpf_debugKafkaGoReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Topic           [64]uint8
	_               [7]byte
	Conn            bpf_debugConnectionInfoT
	Tp              bpf_debugTpInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Op uint8
	_  [7]byte
}

type bpf_debugNewFuncInvocationT struct{ Parent uint64 }

type bpf_debugOffTableT struct{ Table [44]uint64 }

type bpf_debugProduceReqT struct {
	MsgPtr          uint64
	ConnPtr         uint64
	StartMonotimeNs uint64
}

type bpf_debugRedisClientReqT struct {
	Type            uint8
	StartMonotimeNs uint64
	EndMonotimeNs   uint64
	Buf             [256]uint8
	_               [7]byte
	Conn            bpf_debugConnectionInfoT
	_               [4]byte
	Tp              bpf_debugTpInfoT
	Pid             struct {
		HostPid uint32
		UserPid uint32
		Ns      uint32
	}
	Err uint8
	_   [3]byte
}

type bpf_debugServerHttpFuncInvocationT struct {
	StartMonotimeNs uint64
	Tp              bpf_debugTpInfoT
	Method          [7]uint8
	Path            [100]uint8
	_               [5]byte
	ContentLength   uint64
	Status          uint64
}

type bpf_debugSqlFuncInvocationT struct {
	StartMonotimeNs uint64
	SqlParam        uint64
	QueryLen        uint64
	Conn            bpf_debugConnectionInfoT
	_               [4]byte
	Tp              bpf_debugTpInfoT
}

type bpf_debugTopicT struct {
	Name [64]int8
	Tp   bpf_debugTpInfoT
}

type bpf_debugTpInfoPidT struct {
	Tp      bpf_debugTpInfoT
	Pid     uint32
	Valid   uint8
	ReqType uint8
	_       [2]byte
}

type bpf_debugTpInfoT struct {
	TraceId  [16]uint8
	SpanId   [8]uint8
	ParentId [8]uint8
	Ts       uint64
	Flags    uint8
	_        [7]byte
}

type bpf_debugTraceMapKeyT struct {
	Conn bpf_debugConnectionInfoT
	Type uint32
}

// loadBpf_debug returns the embedded CollectionSpec for bpf_debug.
func loadBpf_debug() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_Bpf_debugBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf_debug: %w", err)
	}

	return spec, err
}

// loadBpf_debugObjects loads bpf_debug and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpf_debugObjects
//	*bpf_debugPrograms
//	*bpf_debugMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpf_debugObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf_debug()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpf_debugSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugSpecs struct {
	bpf_debugProgramSpecs
	bpf_debugMapSpecs
}

// bpf_debugSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugProgramSpecs struct {
	BeylaUprobeClientConnClose                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ClientConn_Close"`
	BeylaUprobeClientConnInvoke                    *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ClientConn_Invoke"`
	BeylaUprobeClientConnInvokeReturn              *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ClientConn_Invoke_return"`
	BeylaUprobeClientConnNewStream                 *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ClientConn_NewStream"`
	BeylaUprobeClientConnNewStreamReturn           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ClientConn_NewStream_return"`
	BeylaUprobeServeHTTP                           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ServeHTTP"`
	BeylaUprobeServeHTTPReturns                    *ebpf.ProgramSpec `ebpf:"beyla_uprobe_ServeHTTPReturns"`
	BeylaUprobeClientStreamRecvMsgReturn           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_clientStream_RecvMsg_return"`
	BeylaUprobeClientRoundTrip                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_client_roundTrip"`
	BeylaUprobeConnServe                           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_connServe"`
	BeylaUprobeConnServeRet                        *ebpf.ProgramSpec `ebpf:"beyla_uprobe_connServeRet"`
	BeylaUprobeExecDC                              *ebpf.ProgramSpec `ebpf:"beyla_uprobe_execDC"`
	BeylaUprobeGrpcFramerWriteHeaders              *ebpf.ProgramSpec `ebpf:"beyla_uprobe_grpcFramerWriteHeaders"`
	BeylaUprobeGrpcFramerWriteHeadersReturns       *ebpf.ProgramSpec `ebpf:"beyla_uprobe_grpcFramerWriteHeaders_returns"`
	BeylaUprobeHttp2FramerWriteHeaders             *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2FramerWriteHeaders"`
	BeylaUprobeHttp2FramerWriteHeadersReturns      *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2FramerWriteHeaders_returns"`
	BeylaUprobeHttp2ResponseWriterStateWriteHeader *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2ResponseWriterStateWriteHeader"`
	BeylaUprobeHttp2RoundTrip                      *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2RoundTrip"`
	BeylaUprobeHttp2RoundTripConn                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2RoundTripConn"`
	BeylaUprobeHttp2ServerOperateHeaders           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2Server_operateHeaders"`
	BeylaUprobeHttp2ServerProcessHeaders           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2Server_processHeaders"`
	BeylaUprobeHttp2serverConnRunHandler           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_http2serverConn_runHandler"`
	BeylaUprobeNetFdRead                           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_netFdRead"`
	BeylaUprobePersistConnRoundTrip                *ebpf.ProgramSpec `ebpf:"beyla_uprobe_persistConnRoundTrip"`
	BeylaUprobeProcGoexit1                         *ebpf.ProgramSpec `ebpf:"beyla_uprobe_proc_goexit1"`
	BeylaUprobeProcNewproc1                        *ebpf.ProgramSpec `ebpf:"beyla_uprobe_proc_newproc1"`
	BeylaUprobeProcNewproc1Ret                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_proc_newproc1_ret"`
	BeylaUprobeProtocolRoundtrip                   *ebpf.ProgramSpec `ebpf:"beyla_uprobe_protocol_roundtrip"`
	BeylaUprobeProtocolRoundtripRet                *ebpf.ProgramSpec `ebpf:"beyla_uprobe_protocol_roundtrip_ret"`
	BeylaUprobeQueryDC                             *ebpf.ProgramSpec `ebpf:"beyla_uprobe_queryDC"`
	BeylaUprobeQueryReturn                         *ebpf.ProgramSpec `ebpf:"beyla_uprobe_queryReturn"`
	BeylaUprobeReadContinuedLineSliceReturns       *ebpf.ProgramSpec `ebpf:"beyla_uprobe_readContinuedLineSliceReturns"`
	BeylaUprobeReadRequestReturns                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_readRequestReturns"`
	BeylaUprobeReadRequestStart                    *ebpf.ProgramSpec `ebpf:"beyla_uprobe_readRequestStart"`
	BeylaUprobeReaderRead                          *ebpf.ProgramSpec `ebpf:"beyla_uprobe_reader_read"`
	BeylaUprobeReaderReadRet                       *ebpf.ProgramSpec `ebpf:"beyla_uprobe_reader_read_ret"`
	BeylaUprobeReaderSendMessage                   *ebpf.ProgramSpec `ebpf:"beyla_uprobe_reader_send_message"`
	BeylaUprobeRedisProcess                        *ebpf.ProgramSpec `ebpf:"beyla_uprobe_redis_process"`
	BeylaUprobeRedisProcessRet                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_redis_process_ret"`
	BeylaUprobeRedisWithWriter                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_redis_with_writer"`
	BeylaUprobeRedisWithWriterRet                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_redis_with_writer_ret"`
	BeylaUprobeRoundTrip                           *ebpf.ProgramSpec `ebpf:"beyla_uprobe_roundTrip"`
	BeylaUprobeRoundTripReturn                     *ebpf.ProgramSpec `ebpf:"beyla_uprobe_roundTripReturn"`
	BeylaUprobeSaramaBrokerWrite                   *ebpf.ProgramSpec `ebpf:"beyla_uprobe_sarama_broker_write"`
	BeylaUprobeSaramaResponsePromiseHandle         *ebpf.ProgramSpec `ebpf:"beyla_uprobe_sarama_response_promise_handle"`
	BeylaUprobeSaramaSendInternal                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_sarama_sendInternal"`
	BeylaUprobeServerHandleStream                  *ebpf.ProgramSpec `ebpf:"beyla_uprobe_server_handleStream"`
	BeylaUprobeServerHandleStreamReturn            *ebpf.ProgramSpec `ebpf:"beyla_uprobe_server_handleStream_return"`
	BeylaUprobeServerHandlerTransportHandleStreams *ebpf.ProgramSpec `ebpf:"beyla_uprobe_server_handler_transport_handle_streams"`
	BeylaUprobeTransportHttp2ClientNewStream       *ebpf.ProgramSpec `ebpf:"beyla_uprobe_transport_http2Client_NewStream"`
	BeylaUprobeTransportWriteStatus                *ebpf.ProgramSpec `ebpf:"beyla_uprobe_transport_writeStatus"`
	BeylaUprobeWriteSubset                         *ebpf.ProgramSpec `ebpf:"beyla_uprobe_writeSubset"`
	BeylaUprobeWriterProduce                       *ebpf.ProgramSpec `ebpf:"beyla_uprobe_writer_produce"`
	BeylaUprobeWriterWriteMessages                 *ebpf.ProgramSpec `ebpf:"beyla_uprobe_writer_write_messages"`
}

// bpf_debugMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpf_debugMapSpecs struct {
	DebugEvents                   *ebpf.MapSpec `ebpf:"debug_events"`
	Events                        *ebpf.MapSpec `ebpf:"events"`
	FetchRequests                 *ebpf.MapSpec `ebpf:"fetch_requests"`
	GoOffsetsMap                  *ebpf.MapSpec `ebpf:"go_offsets_map"`
	GoTraceMap                    *ebpf.MapSpec `ebpf:"go_trace_map"`
	Http2ServerRequestsTp         *ebpf.MapSpec `ebpf:"http2_server_requests_tp"`
	IncomingTraceMap              *ebpf.MapSpec `ebpf:"incoming_trace_map"`
	KafkaRequests                 *ebpf.MapSpec `ebpf:"kafka_requests"`
	Newproc1                      *ebpf.MapSpec `ebpf:"newproc1"`
	OngoingClientConnections      *ebpf.MapSpec `ebpf:"ongoing_client_connections"`
	OngoingGoHttp                 *ebpf.MapSpec `ebpf:"ongoing_go_http"`
	OngoingGoroutines             *ebpf.MapSpec `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests     *ebpf.MapSpec `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites       *ebpf.MapSpec `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcOperateHeaders     *ebpf.MapSpec `ebpf:"ongoing_grpc_operate_headers"`
	OngoingGrpcRequestStatus      *ebpf.MapSpec `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests     *ebpf.MapSpec `ebpf:"ongoing_grpc_server_requests"`
	OngoingGrpcTransports         *ebpf.MapSpec `ebpf:"ongoing_grpc_transports"`
	OngoingHttpClientRequests     *ebpf.MapSpec `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.MapSpec `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.MapSpec `ebpf:"ongoing_http_server_requests"`
	OngoingKafkaRequests          *ebpf.MapSpec `ebpf:"ongoing_kafka_requests"`
	OngoingProduceMessages        *ebpf.MapSpec `ebpf:"ongoing_produce_messages"`
	OngoingProduceTopics          *ebpf.MapSpec `ebpf:"ongoing_produce_topics"`
	OngoingRedisRequests          *ebpf.MapSpec `ebpf:"ongoing_redis_requests"`
	OngoingServerConnections      *ebpf.MapSpec `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.MapSpec `ebpf:"ongoing_sql_queries"`
	OngoingStreams                *ebpf.MapSpec `ebpf:"ongoing_streams"`
	OutgoingTraceMap              *ebpf.MapSpec `ebpf:"outgoing_trace_map"`
	ProduceRequests               *ebpf.MapSpec `ebpf:"produce_requests"`
	ProduceTraceparents           *ebpf.MapSpec `ebpf:"produce_traceparents"`
	RedisWrites                   *ebpf.MapSpec `ebpf:"redis_writes"`
	TraceMap                      *ebpf.MapSpec `ebpf:"trace_map"`
}

// bpf_debugObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugObjects struct {
	bpf_debugPrograms
	bpf_debugMaps
}

func (o *bpf_debugObjects) Close() error {
	return _Bpf_debugClose(
		&o.bpf_debugPrograms,
		&o.bpf_debugMaps,
	)
}

// bpf_debugMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugMaps struct {
	DebugEvents                   *ebpf.Map `ebpf:"debug_events"`
	Events                        *ebpf.Map `ebpf:"events"`
	FetchRequests                 *ebpf.Map `ebpf:"fetch_requests"`
	GoOffsetsMap                  *ebpf.Map `ebpf:"go_offsets_map"`
	GoTraceMap                    *ebpf.Map `ebpf:"go_trace_map"`
	Http2ServerRequestsTp         *ebpf.Map `ebpf:"http2_server_requests_tp"`
	IncomingTraceMap              *ebpf.Map `ebpf:"incoming_trace_map"`
	KafkaRequests                 *ebpf.Map `ebpf:"kafka_requests"`
	Newproc1                      *ebpf.Map `ebpf:"newproc1"`
	OngoingClientConnections      *ebpf.Map `ebpf:"ongoing_client_connections"`
	OngoingGoHttp                 *ebpf.Map `ebpf:"ongoing_go_http"`
	OngoingGoroutines             *ebpf.Map `ebpf:"ongoing_goroutines"`
	OngoingGrpcClientRequests     *ebpf.Map `ebpf:"ongoing_grpc_client_requests"`
	OngoingGrpcHeaderWrites       *ebpf.Map `ebpf:"ongoing_grpc_header_writes"`
	OngoingGrpcOperateHeaders     *ebpf.Map `ebpf:"ongoing_grpc_operate_headers"`
	OngoingGrpcRequestStatus      *ebpf.Map `ebpf:"ongoing_grpc_request_status"`
	OngoingGrpcServerRequests     *ebpf.Map `ebpf:"ongoing_grpc_server_requests"`
	OngoingGrpcTransports         *ebpf.Map `ebpf:"ongoing_grpc_transports"`
	OngoingHttpClientRequests     *ebpf.Map `ebpf:"ongoing_http_client_requests"`
	OngoingHttpClientRequestsData *ebpf.Map `ebpf:"ongoing_http_client_requests_data"`
	OngoingHttpServerRequests     *ebpf.Map `ebpf:"ongoing_http_server_requests"`
	OngoingKafkaRequests          *ebpf.Map `ebpf:"ongoing_kafka_requests"`
	OngoingProduceMessages        *ebpf.Map `ebpf:"ongoing_produce_messages"`
	OngoingProduceTopics          *ebpf.Map `ebpf:"ongoing_produce_topics"`
	OngoingRedisRequests          *ebpf.Map `ebpf:"ongoing_redis_requests"`
	OngoingServerConnections      *ebpf.Map `ebpf:"ongoing_server_connections"`
	OngoingSqlQueries             *ebpf.Map `ebpf:"ongoing_sql_queries"`
	OngoingStreams                *ebpf.Map `ebpf:"ongoing_streams"`
	OutgoingTraceMap              *ebpf.Map `ebpf:"outgoing_trace_map"`
	ProduceRequests               *ebpf.Map `ebpf:"produce_requests"`
	ProduceTraceparents           *ebpf.Map `ebpf:"produce_traceparents"`
	RedisWrites                   *ebpf.Map `ebpf:"redis_writes"`
	TraceMap                      *ebpf.Map `ebpf:"trace_map"`
}

func (m *bpf_debugMaps) Close() error {
	return _Bpf_debugClose(
		m.DebugEvents,
		m.Events,
		m.FetchRequests,
		m.GoOffsetsMap,
		m.GoTraceMap,
		m.Http2ServerRequestsTp,
		m.IncomingTraceMap,
		m.KafkaRequests,
		m.Newproc1,
		m.OngoingClientConnections,
		m.OngoingGoHttp,
		m.OngoingGoroutines,
		m.OngoingGrpcClientRequests,
		m.OngoingGrpcHeaderWrites,
		m.OngoingGrpcOperateHeaders,
		m.OngoingGrpcRequestStatus,
		m.OngoingGrpcServerRequests,
		m.OngoingGrpcTransports,
		m.OngoingHttpClientRequests,
		m.OngoingHttpClientRequestsData,
		m.OngoingHttpServerRequests,
		m.OngoingKafkaRequests,
		m.OngoingProduceMessages,
		m.OngoingProduceTopics,
		m.OngoingRedisRequests,
		m.OngoingServerConnections,
		m.OngoingSqlQueries,
		m.OngoingStreams,
		m.OutgoingTraceMap,
		m.ProduceRequests,
		m.ProduceTraceparents,
		m.RedisWrites,
		m.TraceMap,
	)
}

// bpf_debugPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpf_debugObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpf_debugPrograms struct {
	BeylaUprobeClientConnClose                     *ebpf.Program `ebpf:"beyla_uprobe_ClientConn_Close"`
	BeylaUprobeClientConnInvoke                    *ebpf.Program `ebpf:"beyla_uprobe_ClientConn_Invoke"`
	BeylaUprobeClientConnInvokeReturn              *ebpf.Program `ebpf:"beyla_uprobe_ClientConn_Invoke_return"`
	BeylaUprobeClientConnNewStream                 *ebpf.Program `ebpf:"beyla_uprobe_ClientConn_NewStream"`
	BeylaUprobeClientConnNewStreamReturn           *ebpf.Program `ebpf:"beyla_uprobe_ClientConn_NewStream_return"`
	BeylaUprobeServeHTTP                           *ebpf.Program `ebpf:"beyla_uprobe_ServeHTTP"`
	BeylaUprobeServeHTTPReturns                    *ebpf.Program `ebpf:"beyla_uprobe_ServeHTTPReturns"`
	BeylaUprobeClientStreamRecvMsgReturn           *ebpf.Program `ebpf:"beyla_uprobe_clientStream_RecvMsg_return"`
	BeylaUprobeClientRoundTrip                     *ebpf.Program `ebpf:"beyla_uprobe_client_roundTrip"`
	BeylaUprobeConnServe                           *ebpf.Program `ebpf:"beyla_uprobe_connServe"`
	BeylaUprobeConnServeRet                        *ebpf.Program `ebpf:"beyla_uprobe_connServeRet"`
	BeylaUprobeExecDC                              *ebpf.Program `ebpf:"beyla_uprobe_execDC"`
	BeylaUprobeGrpcFramerWriteHeaders              *ebpf.Program `ebpf:"beyla_uprobe_grpcFramerWriteHeaders"`
	BeylaUprobeGrpcFramerWriteHeadersReturns       *ebpf.Program `ebpf:"beyla_uprobe_grpcFramerWriteHeaders_returns"`
	BeylaUprobeHttp2FramerWriteHeaders             *ebpf.Program `ebpf:"beyla_uprobe_http2FramerWriteHeaders"`
	BeylaUprobeHttp2FramerWriteHeadersReturns      *ebpf.Program `ebpf:"beyla_uprobe_http2FramerWriteHeaders_returns"`
	BeylaUprobeHttp2ResponseWriterStateWriteHeader *ebpf.Program `ebpf:"beyla_uprobe_http2ResponseWriterStateWriteHeader"`
	BeylaUprobeHttp2RoundTrip                      *ebpf.Program `ebpf:"beyla_uprobe_http2RoundTrip"`
	BeylaUprobeHttp2RoundTripConn                  *ebpf.Program `ebpf:"beyla_uprobe_http2RoundTripConn"`
	BeylaUprobeHttp2ServerOperateHeaders           *ebpf.Program `ebpf:"beyla_uprobe_http2Server_operateHeaders"`
	BeylaUprobeHttp2ServerProcessHeaders           *ebpf.Program `ebpf:"beyla_uprobe_http2Server_processHeaders"`
	BeylaUprobeHttp2serverConnRunHandler           *ebpf.Program `ebpf:"beyla_uprobe_http2serverConn_runHandler"`
	BeylaUprobeNetFdRead                           *ebpf.Program `ebpf:"beyla_uprobe_netFdRead"`
	BeylaUprobePersistConnRoundTrip                *ebpf.Program `ebpf:"beyla_uprobe_persistConnRoundTrip"`
	BeylaUprobeProcGoexit1                         *ebpf.Program `ebpf:"beyla_uprobe_proc_goexit1"`
	BeylaUprobeProcNewproc1                        *ebpf.Program `ebpf:"beyla_uprobe_proc_newproc1"`
	BeylaUprobeProcNewproc1Ret                     *ebpf.Program `ebpf:"beyla_uprobe_proc_newproc1_ret"`
	BeylaUprobeProtocolRoundtrip                   *ebpf.Program `ebpf:"beyla_uprobe_protocol_roundtrip"`
	BeylaUprobeProtocolRoundtripRet                *ebpf.Program `ebpf:"beyla_uprobe_protocol_roundtrip_ret"`
	BeylaUprobeQueryDC                             *ebpf.Program `ebpf:"beyla_uprobe_queryDC"`
	BeylaUprobeQueryReturn                         *ebpf.Program `ebpf:"beyla_uprobe_queryReturn"`
	BeylaUprobeReadContinuedLineSliceReturns       *ebpf.Program `ebpf:"beyla_uprobe_readContinuedLineSliceReturns"`
	BeylaUprobeReadRequestReturns                  *ebpf.Program `ebpf:"beyla_uprobe_readRequestReturns"`
	BeylaUprobeReadRequestStart                    *ebpf.Program `ebpf:"beyla_uprobe_readRequestStart"`
	BeylaUprobeReaderRead                          *ebpf.Program `ebpf:"beyla_uprobe_reader_read"`
	BeylaUprobeReaderReadRet                       *ebpf.Program `ebpf:"beyla_uprobe_reader_read_ret"`
	BeylaUprobeReaderSendMessage                   *ebpf.Program `ebpf:"beyla_uprobe_reader_send_message"`
	BeylaUprobeRedisProcess                        *ebpf.Program `ebpf:"beyla_uprobe_redis_process"`
	BeylaUprobeRedisProcessRet                     *ebpf.Program `ebpf:"beyla_uprobe_redis_process_ret"`
	BeylaUprobeRedisWithWriter                     *ebpf.Program `ebpf:"beyla_uprobe_redis_with_writer"`
	BeylaUprobeRedisWithWriterRet                  *ebpf.Program `ebpf:"beyla_uprobe_redis_with_writer_ret"`
	BeylaUprobeRoundTrip                           *ebpf.Program `ebpf:"beyla_uprobe_roundTrip"`
	BeylaUprobeRoundTripReturn                     *ebpf.Program `ebpf:"beyla_uprobe_roundTripReturn"`
	BeylaUprobeSaramaBrokerWrite                   *ebpf.Program `ebpf:"beyla_uprobe_sarama_broker_write"`
	BeylaUprobeSaramaResponsePromiseHandle         *ebpf.Program `ebpf:"beyla_uprobe_sarama_response_promise_handle"`
	BeylaUprobeSaramaSendInternal                  *ebpf.Program `ebpf:"beyla_uprobe_sarama_sendInternal"`
	BeylaUprobeServerHandleStream                  *ebpf.Program `ebpf:"beyla_uprobe_server_handleStream"`
	BeylaUprobeServerHandleStreamReturn            *ebpf.Program `ebpf:"beyla_uprobe_server_handleStream_return"`
	BeylaUprobeServerHandlerTransportHandleStreams *ebpf.Program `ebpf:"beyla_uprobe_server_handler_transport_handle_streams"`
	BeylaUprobeTransportHttp2ClientNewStream       *ebpf.Program `ebpf:"beyla_uprobe_transport_http2Client_NewStream"`
	BeylaUprobeTransportWriteStatus                *ebpf.Program `ebpf:"beyla_uprobe_transport_writeStatus"`
	BeylaUprobeWriteSubset                         *ebpf.Program `ebpf:"beyla_uprobe_writeSubset"`
	BeylaUprobeWriterProduce                       *ebpf.Program `ebpf:"beyla_uprobe_writer_produce"`
	BeylaUprobeWriterWriteMessages                 *ebpf.Program `ebpf:"beyla_uprobe_writer_write_messages"`
}

func (p *bpf_debugPrograms) Close() error {
	return _Bpf_debugClose(
		p.BeylaUprobeClientConnClose,
		p.BeylaUprobeClientConnInvoke,
		p.BeylaUprobeClientConnInvokeReturn,
		p.BeylaUprobeClientConnNewStream,
		p.BeylaUprobeClientConnNewStreamReturn,
		p.BeylaUprobeServeHTTP,
		p.BeylaUprobeServeHTTPReturns,
		p.BeylaUprobeClientStreamRecvMsgReturn,
		p.BeylaUprobeClientRoundTrip,
		p.BeylaUprobeConnServe,
		p.BeylaUprobeConnServeRet,
		p.BeylaUprobeExecDC,
		p.BeylaUprobeGrpcFramerWriteHeaders,
		p.BeylaUprobeGrpcFramerWriteHeadersReturns,
		p.BeylaUprobeHttp2FramerWriteHeaders,
		p.BeylaUprobeHttp2FramerWriteHeadersReturns,
		p.BeylaUprobeHttp2ResponseWriterStateWriteHeader,
		p.BeylaUprobeHttp2RoundTrip,
		p.BeylaUprobeHttp2RoundTripConn,
		p.BeylaUprobeHttp2ServerOperateHeaders,
		p.BeylaUprobeHttp2ServerProcessHeaders,
		p.BeylaUprobeHttp2serverConnRunHandler,
		p.BeylaUprobeNetFdRead,
		p.BeylaUprobePersistConnRoundTrip,
		p.BeylaUprobeProcGoexit1,
		p.BeylaUprobeProcNewproc1,
		p.BeylaUprobeProcNewproc1Ret,
		p.BeylaUprobeProtocolRoundtrip,
		p.BeylaUprobeProtocolRoundtripRet,
		p.BeylaUprobeQueryDC,
		p.BeylaUprobeQueryReturn,
		p.BeylaUprobeReadContinuedLineSliceReturns,
		p.BeylaUprobeReadRequestReturns,
		p.BeylaUprobeReadRequestStart,
		p.BeylaUprobeReaderRead,
		p.BeylaUprobeReaderReadRet,
		p.BeylaUprobeReaderSendMessage,
		p.BeylaUprobeRedisProcess,
		p.BeylaUprobeRedisProcessRet,
		p.BeylaUprobeRedisWithWriter,
		p.BeylaUprobeRedisWithWriterRet,
		p.BeylaUprobeRoundTrip,
		p.BeylaUprobeRoundTripReturn,
		p.BeylaUprobeSaramaBrokerWrite,
		p.BeylaUprobeSaramaResponsePromiseHandle,
		p.BeylaUprobeSaramaSendInternal,
		p.BeylaUprobeServerHandleStream,
		p.BeylaUprobeServerHandleStreamReturn,
		p.BeylaUprobeServerHandlerTransportHandleStreams,
		p.BeylaUprobeTransportHttp2ClientNewStream,
		p.BeylaUprobeTransportWriteStatus,
		p.BeylaUprobeWriteSubset,
		p.BeylaUprobeWriterProduce,
		p.BeylaUprobeWriterWriteMessages,
	)
}

func _Bpf_debugClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_debug_arm64_bpfel.o
var _Bpf_debugBytes []byte
