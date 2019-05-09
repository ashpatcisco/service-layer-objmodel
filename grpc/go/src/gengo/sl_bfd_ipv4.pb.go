// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sl_bfd_ipv4.proto

package service_layer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// BFD unique key identifier.
type SLBfdv4Key struct {
	// BFD Session Type
	Type SLBfdType `protobuf:"varint,1,opt,name=Type,enum=service_layer.SLBfdType" json:"Type,omitempty"`
	// VRF name.
	VrfName string `protobuf:"bytes,2,opt,name=VrfName" json:"VrfName,omitempty"`
	// Neighbor Ipv4 address.
	NbrAddr uint32 `protobuf:"varint,3,opt,name=NbrAddr" json:"NbrAddr,omitempty"`
	// Interface name and handle, needed for single-hop BFD
	Interface *SLInterface `protobuf:"bytes,4,opt,name=Interface" json:"Interface,omitempty"`
	// Source Ipv4 address, needed for multi-hop BFD
	SourceAddr uint32 `protobuf:"varint,5,opt,name=SourceAddr" json:"SourceAddr,omitempty"`
}

func (m *SLBfdv4Key) Reset()                    { *m = SLBfdv4Key{} }
func (m *SLBfdv4Key) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4Key) ProtoMessage()               {}
func (*SLBfdv4Key) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SLBfdv4Key) GetType() SLBfdType {
	if m != nil {
		return m.Type
	}
	return SLBfdType_SL_BFD_RESERVED
}

func (m *SLBfdv4Key) GetVrfName() string {
	if m != nil {
		return m.VrfName
	}
	return ""
}

func (m *SLBfdv4Key) GetNbrAddr() uint32 {
	if m != nil {
		return m.NbrAddr
	}
	return 0
}

func (m *SLBfdv4Key) GetInterface() *SLInterface {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (m *SLBfdv4Key) GetSourceAddr() uint32 {
	if m != nil {
		return m.SourceAddr
	}
	return 0
}

// BFD session information.
type SLBfdv4SessionCfg struct {
	// BFD unique key identifier.
	Key *SLBfdv4Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	// BFD common features.
	Config *SLBfdConfigCommon `protobuf:"bytes,2,opt,name=Config" json:"Config,omitempty"`
}

func (m *SLBfdv4SessionCfg) Reset()                    { *m = SLBfdv4SessionCfg{} }
func (m *SLBfdv4SessionCfg) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4SessionCfg) ProtoMessage()               {}
func (*SLBfdv4SessionCfg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *SLBfdv4SessionCfg) GetKey() *SLBfdv4Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SLBfdv4SessionCfg) GetConfig() *SLBfdConfigCommon {
	if m != nil {
		return m.Config
	}
	return nil
}

// Contains a List of BFD session objects.
type SLBfdv4Msg struct {
	// BFD Object Operations
	Oper SLObjectOp `protobuf:"varint,1,opt,name=Oper,enum=service_layer.SLObjectOp" json:"Oper,omitempty"`
	// List of BFD session objects
	Sessions []*SLBfdv4SessionCfg `protobuf:"bytes,2,rep,name=Sessions" json:"Sessions,omitempty"`
}

func (m *SLBfdv4Msg) Reset()                    { *m = SLBfdv4Msg{} }
func (m *SLBfdv4Msg) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4Msg) ProtoMessage()               {}
func (*SLBfdv4Msg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SLBfdv4Msg) GetOper() SLObjectOp {
	if m != nil {
		return m.Oper
	}
	return SLObjectOp_SL_OBJOP_RESERVED
}

func (m *SLBfdv4Msg) GetSessions() []*SLBfdv4SessionCfg {
	if m != nil {
		return m.Sessions
	}
	return nil
}

// BFD result
type SLBfdv4Res struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus" json:"ErrStatus,omitempty"`
	// BFD unique key.
	Key *SLBfdv4Key `protobuf:"bytes,2,opt,name=Key" json:"Key,omitempty"`
}

func (m *SLBfdv4Res) Reset()                    { *m = SLBfdv4Res{} }
func (m *SLBfdv4Res) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4Res) ProtoMessage()               {}
func (*SLBfdv4Res) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *SLBfdv4Res) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLBfdv4Res) GetKey() *SLBfdv4Key {
	if m != nil {
		return m.Key
	}
	return nil
}

// BFD bulk result
type SLBfdv4MsgRsp struct {
	// Summary result of the bulk operation (refer to enum SLErrorStatus)
	// In general, the StatusSummary implies one of 3 things:
	// 1. SL_SUCCESS: signifies that the entire bulk operation was successful.
	//         In this case, the Results list is empty.
	// 2. SL_SOME_ERR: signifies that the operation failed for one or more
	//         entries. In this case, Results holds the result for
	//         each individual entry in the bulk.
	// 3. SL_RPC_XXX: signifies that the entire bulk operation failed.
	//         In this case, the Results list is empty.
	StatusSummary *SLErrorStatus `protobuf:"bytes,1,opt,name=StatusSummary" json:"StatusSummary,omitempty"`
	// In case of errors, this field indicates which entry in the bulk was
	// erroneous.
	Results []*SLBfdv4Res `protobuf:"bytes,2,rep,name=Results" json:"Results,omitempty"`
}

func (m *SLBfdv4MsgRsp) Reset()                    { *m = SLBfdv4MsgRsp{} }
func (m *SLBfdv4MsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4MsgRsp) ProtoMessage()               {}
func (*SLBfdv4MsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *SLBfdv4MsgRsp) GetStatusSummary() *SLErrorStatus {
	if m != nil {
		return m.StatusSummary
	}
	return nil
}

func (m *SLBfdv4MsgRsp) GetResults() []*SLBfdv4Res {
	if m != nil {
		return m.Results
	}
	return nil
}

// BFD Get Message
type SLBfdv4GetMsg struct {
	// BFD key.
	// If the Key is not specified, then request up to the first
	// 'EntriesCount' entries.
	Key *SLBfdv4Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	// Global BFD event sequence number.
	// Return all BFD sessions with sequence number >= SeqNum.
	SeqNum uint64 `protobuf:"varint,2,opt,name=SeqNum" json:"SeqNum,omitempty"`
	// Number of entries requested
	EntriesCount uint32 `protobuf:"varint,3,opt,name=EntriesCount" json:"EntriesCount,omitempty"`
	// if GetNext is FALSE:
	//     request up to 'EntriesCount' entries starting from the key
	// If GetNext is TRUE, or if the key exact match is not found:
	//     request up to 'EntriesCount' entries starting from the key's next
	GetNext bool `protobuf:"varint,4,opt,name=GetNext" json:"GetNext,omitempty"`
}

func (m *SLBfdv4GetMsg) Reset()                    { *m = SLBfdv4GetMsg{} }
func (m *SLBfdv4GetMsg) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4GetMsg) ProtoMessage()               {}
func (*SLBfdv4GetMsg) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *SLBfdv4GetMsg) GetKey() *SLBfdv4Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SLBfdv4GetMsg) GetSeqNum() uint64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

func (m *SLBfdv4GetMsg) GetEntriesCount() uint32 {
	if m != nil {
		return m.EntriesCount
	}
	return 0
}

func (m *SLBfdv4GetMsg) GetGetNext() bool {
	if m != nil {
		return m.GetNext
	}
	return false
}

// BFD Session Config and State info
type SLBfdv4SessionCfgState struct {
	// BFD unique key identifier.
	Key *SLBfdv4Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	// BFD common features.
	Config *SLBfdConfigCommon `protobuf:"bytes,2,opt,name=Config" json:"Config,omitempty"`
	// BFD session State
	State *SLBfdCommonState `protobuf:"bytes,3,opt,name=State" json:"State,omitempty"`
}

func (m *SLBfdv4SessionCfgState) Reset()                    { *m = SLBfdv4SessionCfgState{} }
func (m *SLBfdv4SessionCfgState) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4SessionCfgState) ProtoMessage()               {}
func (*SLBfdv4SessionCfgState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *SLBfdv4SessionCfgState) GetKey() *SLBfdv4Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SLBfdv4SessionCfgState) GetConfig() *SLBfdConfigCommon {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *SLBfdv4SessionCfgState) GetState() *SLBfdCommonState {
	if m != nil {
		return m.State
	}
	return nil
}

// BFD Get Message Response
type SLBfdv4GetMsgRsp struct {
	// End Of File.
	// When set to True, it indicates that the server has returned M, where
	// M < N, of the original N requested Entries.
	Eof bool `protobuf:"varint,1,opt,name=Eof" json:"Eof,omitempty"`
	// Status of the Get operation
	ErrStatus *SLErrorStatus `protobuf:"bytes,2,opt,name=ErrStatus" json:"ErrStatus,omitempty"`
	// Returned entries as requested in the Get operation.
	// if Error is SL_SUCCESS, Entries contains the info requested
	Entries []*SLBfdv4SessionCfgState `protobuf:"bytes,3,rep,name=Entries" json:"Entries,omitempty"`
}

func (m *SLBfdv4GetMsgRsp) Reset()                    { *m = SLBfdv4GetMsgRsp{} }
func (m *SLBfdv4GetMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4GetMsgRsp) ProtoMessage()               {}
func (*SLBfdv4GetMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *SLBfdv4GetMsgRsp) GetEof() bool {
	if m != nil {
		return m.Eof
	}
	return false
}

func (m *SLBfdv4GetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLBfdv4GetMsgRsp) GetEntries() []*SLBfdv4SessionCfgState {
	if m != nil {
		return m.Entries
	}
	return nil
}

// BFD Session and State info
type SLBfdv4SessionState struct {
	// BFD unique key identifier.
	Key *SLBfdv4Key `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	// BFD session State
	State *SLBfdCommonState `protobuf:"bytes,2,opt,name=State" json:"State,omitempty"`
}

func (m *SLBfdv4SessionState) Reset()                    { *m = SLBfdv4SessionState{} }
func (m *SLBfdv4SessionState) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4SessionState) ProtoMessage()               {}
func (*SLBfdv4SessionState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *SLBfdv4SessionState) GetKey() *SLBfdv4Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SLBfdv4SessionState) GetState() *SLBfdCommonState {
	if m != nil {
		return m.State
	}
	return nil
}

// BFD Session and State info Message
type SLBfdv4Notif struct {
	// Event Type
	EventType SLBfdNotifType `protobuf:"varint,1,opt,name=EventType,enum=service_layer.SLBfdNotifType" json:"EventType,omitempty"`
	// Further info based on EventType
	//
	// Types that are valid to be assigned to Event:
	//	*SLBfdv4Notif_ErrStatus
	//	*SLBfdv4Notif_Session
	Event isSLBfdv4Notif_Event `protobuf_oneof:"Event"`
}

func (m *SLBfdv4Notif) Reset()                    { *m = SLBfdv4Notif{} }
func (m *SLBfdv4Notif) String() string            { return proto.CompactTextString(m) }
func (*SLBfdv4Notif) ProtoMessage()               {}
func (*SLBfdv4Notif) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type isSLBfdv4Notif_Event interface {
	isSLBfdv4Notif_Event()
}

type SLBfdv4Notif_ErrStatus struct {
	ErrStatus *SLErrorStatus `protobuf:"bytes,2,opt,name=ErrStatus,oneof"`
}
type SLBfdv4Notif_Session struct {
	Session *SLBfdv4SessionState `protobuf:"bytes,3,opt,name=Session,oneof"`
}

func (*SLBfdv4Notif_ErrStatus) isSLBfdv4Notif_Event() {}
func (*SLBfdv4Notif_Session) isSLBfdv4Notif_Event()   {}

func (m *SLBfdv4Notif) GetEvent() isSLBfdv4Notif_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *SLBfdv4Notif) GetEventType() SLBfdNotifType {
	if m != nil {
		return m.EventType
	}
	return SLBfdNotifType_SL_BFD_EVENT_TYPE_RESERVED
}

func (m *SLBfdv4Notif) GetErrStatus() *SLErrorStatus {
	if x, ok := m.GetEvent().(*SLBfdv4Notif_ErrStatus); ok {
		return x.ErrStatus
	}
	return nil
}

func (m *SLBfdv4Notif) GetSession() *SLBfdv4SessionState {
	if x, ok := m.GetEvent().(*SLBfdv4Notif_Session); ok {
		return x.Session
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SLBfdv4Notif) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SLBfdv4Notif_OneofMarshaler, _SLBfdv4Notif_OneofUnmarshaler, _SLBfdv4Notif_OneofSizer, []interface{}{
		(*SLBfdv4Notif_ErrStatus)(nil),
		(*SLBfdv4Notif_Session)(nil),
	}
}

func _SLBfdv4Notif_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SLBfdv4Notif)
	// Event
	switch x := m.Event.(type) {
	case *SLBfdv4Notif_ErrStatus:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ErrStatus); err != nil {
			return err
		}
	case *SLBfdv4Notif_Session:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Session); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SLBfdv4Notif.Event has unexpected type %T", x)
	}
	return nil
}

func _SLBfdv4Notif_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SLBfdv4Notif)
	switch tag {
	case 2: // Event.ErrStatus
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SLErrorStatus)
		err := b.DecodeMessage(msg)
		m.Event = &SLBfdv4Notif_ErrStatus{msg}
		return true, err
	case 3: // Event.Session
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SLBfdv4SessionState)
		err := b.DecodeMessage(msg)
		m.Event = &SLBfdv4Notif_Session{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SLBfdv4Notif_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SLBfdv4Notif)
	// Event
	switch x := m.Event.(type) {
	case *SLBfdv4Notif_ErrStatus:
		s := proto.Size(x.ErrStatus)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *SLBfdv4Notif_Session:
		s := proto.Size(x.Session)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*SLBfdv4Key)(nil), "service_layer.SLBfdv4Key")
	proto.RegisterType((*SLBfdv4SessionCfg)(nil), "service_layer.SLBfdv4SessionCfg")
	proto.RegisterType((*SLBfdv4Msg)(nil), "service_layer.SLBfdv4Msg")
	proto.RegisterType((*SLBfdv4Res)(nil), "service_layer.SLBfdv4Res")
	proto.RegisterType((*SLBfdv4MsgRsp)(nil), "service_layer.SLBfdv4MsgRsp")
	proto.RegisterType((*SLBfdv4GetMsg)(nil), "service_layer.SLBfdv4GetMsg")
	proto.RegisterType((*SLBfdv4SessionCfgState)(nil), "service_layer.SLBfdv4SessionCfgState")
	proto.RegisterType((*SLBfdv4GetMsgRsp)(nil), "service_layer.SLBfdv4GetMsgRsp")
	proto.RegisterType((*SLBfdv4SessionState)(nil), "service_layer.SLBfdv4SessionState")
	proto.RegisterType((*SLBfdv4Notif)(nil), "service_layer.SLBfdv4Notif")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SLBfdv4Oper service

type SLBfdv4OperClient interface {
	// SLBfdRegMsg.Oper = SL_REGOP_REGISTER:
	//     Global BFD registration.
	//     A client Must Register BEFORE BFD sessions can be added/modified.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//     Global BFD un-registration.
	//     This call is used to end all BFD notifications and unregister any
	//     interest in BFD session configuration.
	//     This call cleans up all BFD sessions previously requested.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_EOF:
	//     BFD End Of File.
	//     After Registration, the client is expected to send an EOF
	//     message to convey the end of replay of the client's known objects.
	//     This is especially useful under certain restart scenarios when the
	//     client and the server are trying to synchronize their BFD sessions.
	SLBfdv4RegOp(ctx context.Context, in *SLBfdRegMsg, opts ...grpc.CallOption) (*SLBfdRegMsgRsp, error)
	// Used to retrieve global BFD info from the server.
	SLBfdv4Get(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetMsgRsp, error)
	// Used to retrieve global BFD stats from the server.
	SLBfdv4GetStats(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetStatsMsgRsp, error)
	// This call is used to get a stream of session state notifications.
	// The caller must maintain the GRPC channel as long as
	// there is interest in BFD session notifications. Only sessions that were
	// created through this API will be notified to caller.
	// This call can be used to get "push" notifications for session states.
	// It is advised that the caller register for notifications before any
	// sessions are created to avoid any loss of notifications.
	SLBfdv4GetNotifStream(ctx context.Context, in *SLBfdGetNotifMsg, opts ...grpc.CallOption) (SLBfdv4Oper_SLBfdv4GetNotifStreamClient, error)
	// SLBfdv4Msg.Oper = SL_OBJOP_ADD:
	//     Add one or multiple BFD sessions.
	//
	// SLBfdv4Msg.Oper = SL_OBJOP_UPDATE:
	//     Update one or multiple BFD sessions.
	//
	// SLBfdv4Msg.Oper = SL_OBJOP_DELETE:
	//     Delete one or multiple BFD sessions.
	SLBfdv4SessionOp(ctx context.Context, in *SLBfdv4Msg, opts ...grpc.CallOption) (*SLBfdv4MsgRsp, error)
	// Retrieve BFD session attributes and state.
	// This call can be used to "poll" the current state of a session.
	SLBfdv4SessionGet(ctx context.Context, in *SLBfdv4GetMsg, opts ...grpc.CallOption) (*SLBfdv4GetMsgRsp, error)
}

type sLBfdv4OperClient struct {
	cc *grpc.ClientConn
}

func NewSLBfdv4OperClient(cc *grpc.ClientConn) SLBfdv4OperClient {
	return &sLBfdv4OperClient{cc}
}

func (c *sLBfdv4OperClient) SLBfdv4RegOp(ctx context.Context, in *SLBfdRegMsg, opts ...grpc.CallOption) (*SLBfdRegMsgRsp, error) {
	out := new(SLBfdRegMsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLBfdv4Oper/SLBfdv4RegOp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv4OperClient) SLBfdv4Get(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetMsgRsp, error) {
	out := new(SLBfdGetMsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLBfdv4Oper/SLBfdv4Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv4OperClient) SLBfdv4GetStats(ctx context.Context, in *SLBfdGetMsg, opts ...grpc.CallOption) (*SLBfdGetStatsMsgRsp, error) {
	out := new(SLBfdGetStatsMsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLBfdv4Oper/SLBfdv4GetStats", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv4OperClient) SLBfdv4GetNotifStream(ctx context.Context, in *SLBfdGetNotifMsg, opts ...grpc.CallOption) (SLBfdv4Oper_SLBfdv4GetNotifStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SLBfdv4Oper_serviceDesc.Streams[0], c.cc, "/service_layer.SLBfdv4Oper/SLBfdv4GetNotifStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &sLBfdv4OperSLBfdv4GetNotifStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SLBfdv4Oper_SLBfdv4GetNotifStreamClient interface {
	Recv() (*SLBfdv4Notif, error)
	grpc.ClientStream
}

type sLBfdv4OperSLBfdv4GetNotifStreamClient struct {
	grpc.ClientStream
}

func (x *sLBfdv4OperSLBfdv4GetNotifStreamClient) Recv() (*SLBfdv4Notif, error) {
	m := new(SLBfdv4Notif)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *sLBfdv4OperClient) SLBfdv4SessionOp(ctx context.Context, in *SLBfdv4Msg, opts ...grpc.CallOption) (*SLBfdv4MsgRsp, error) {
	out := new(SLBfdv4MsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLBfdv4Oper/SLBfdv4SessionOp", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sLBfdv4OperClient) SLBfdv4SessionGet(ctx context.Context, in *SLBfdv4GetMsg, opts ...grpc.CallOption) (*SLBfdv4GetMsgRsp, error) {
	out := new(SLBfdv4GetMsgRsp)
	err := grpc.Invoke(ctx, "/service_layer.SLBfdv4Oper/SLBfdv4SessionGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SLBfdv4Oper service

type SLBfdv4OperServer interface {
	// SLBfdRegMsg.Oper = SL_REGOP_REGISTER:
	//     Global BFD registration.
	//     A client Must Register BEFORE BFD sessions can be added/modified.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_UNREGISTER:
	//     Global BFD un-registration.
	//     This call is used to end all BFD notifications and unregister any
	//     interest in BFD session configuration.
	//     This call cleans up all BFD sessions previously requested.
	//
	// SLBfdRegMsg.Oper = SL_REGOP_EOF:
	//     BFD End Of File.
	//     After Registration, the client is expected to send an EOF
	//     message to convey the end of replay of the client's known objects.
	//     This is especially useful under certain restart scenarios when the
	//     client and the server are trying to synchronize their BFD sessions.
	SLBfdv4RegOp(context.Context, *SLBfdRegMsg) (*SLBfdRegMsgRsp, error)
	// Used to retrieve global BFD info from the server.
	SLBfdv4Get(context.Context, *SLBfdGetMsg) (*SLBfdGetMsgRsp, error)
	// Used to retrieve global BFD stats from the server.
	SLBfdv4GetStats(context.Context, *SLBfdGetMsg) (*SLBfdGetStatsMsgRsp, error)
	// This call is used to get a stream of session state notifications.
	// The caller must maintain the GRPC channel as long as
	// there is interest in BFD session notifications. Only sessions that were
	// created through this API will be notified to caller.
	// This call can be used to get "push" notifications for session states.
	// It is advised that the caller register for notifications before any
	// sessions are created to avoid any loss of notifications.
	SLBfdv4GetNotifStream(*SLBfdGetNotifMsg, SLBfdv4Oper_SLBfdv4GetNotifStreamServer) error
	// SLBfdv4Msg.Oper = SL_OBJOP_ADD:
	//     Add one or multiple BFD sessions.
	//
	// SLBfdv4Msg.Oper = SL_OBJOP_UPDATE:
	//     Update one or multiple BFD sessions.
	//
	// SLBfdv4Msg.Oper = SL_OBJOP_DELETE:
	//     Delete one or multiple BFD sessions.
	SLBfdv4SessionOp(context.Context, *SLBfdv4Msg) (*SLBfdv4MsgRsp, error)
	// Retrieve BFD session attributes and state.
	// This call can be used to "poll" the current state of a session.
	SLBfdv4SessionGet(context.Context, *SLBfdv4GetMsg) (*SLBfdv4GetMsgRsp, error)
}

func RegisterSLBfdv4OperServer(s *grpc.Server, srv SLBfdv4OperServer) {
	s.RegisterService(&_SLBfdv4Oper_serviceDesc, srv)
}

func _SLBfdv4Oper_SLBfdv4RegOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdRegMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv4OperServer).SLBfdv4RegOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv4Oper/SLBfdv4RegOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv4OperServer).SLBfdv4RegOp(ctx, req.(*SLBfdRegMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv4Oper_SLBfdv4Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv4OperServer).SLBfdv4Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv4Oper/SLBfdv4Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv4OperServer).SLBfdv4Get(ctx, req.(*SLBfdGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv4Oper_SLBfdv4GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdGetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv4OperServer).SLBfdv4GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv4Oper/SLBfdv4GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv4OperServer).SLBfdv4GetStats(ctx, req.(*SLBfdGetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv4Oper_SLBfdv4GetNotifStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SLBfdGetNotifMsg)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SLBfdv4OperServer).SLBfdv4GetNotifStream(m, &sLBfdv4OperSLBfdv4GetNotifStreamServer{stream})
}

type SLBfdv4Oper_SLBfdv4GetNotifStreamServer interface {
	Send(*SLBfdv4Notif) error
	grpc.ServerStream
}

type sLBfdv4OperSLBfdv4GetNotifStreamServer struct {
	grpc.ServerStream
}

func (x *sLBfdv4OperSLBfdv4GetNotifStreamServer) Send(m *SLBfdv4Notif) error {
	return x.ServerStream.SendMsg(m)
}

func _SLBfdv4Oper_SLBfdv4SessionOp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdv4Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv4OperServer).SLBfdv4SessionOp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv4Oper/SLBfdv4SessionOp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv4OperServer).SLBfdv4SessionOp(ctx, req.(*SLBfdv4Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _SLBfdv4Oper_SLBfdv4SessionGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SLBfdv4GetMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SLBfdv4OperServer).SLBfdv4SessionGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_layer.SLBfdv4Oper/SLBfdv4SessionGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SLBfdv4OperServer).SLBfdv4SessionGet(ctx, req.(*SLBfdv4GetMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _SLBfdv4Oper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service_layer.SLBfdv4Oper",
	HandlerType: (*SLBfdv4OperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SLBfdv4RegOp",
			Handler:    _SLBfdv4Oper_SLBfdv4RegOp_Handler,
		},
		{
			MethodName: "SLBfdv4Get",
			Handler:    _SLBfdv4Oper_SLBfdv4Get_Handler,
		},
		{
			MethodName: "SLBfdv4GetStats",
			Handler:    _SLBfdv4Oper_SLBfdv4GetStats_Handler,
		},
		{
			MethodName: "SLBfdv4SessionOp",
			Handler:    _SLBfdv4Oper_SLBfdv4SessionOp_Handler,
		},
		{
			MethodName: "SLBfdv4SessionGet",
			Handler:    _SLBfdv4Oper_SLBfdv4SessionGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SLBfdv4GetNotifStream",
			Handler:       _SLBfdv4Oper_SLBfdv4GetNotifStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "sl_bfd_ipv4.proto",
}

func init() { proto.RegisterFile("sl_bfd_ipv4.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 715 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0x9c, 0xa4, 0x4d, 0x3b, 0x69, 0x3f, 0xda, 0xad, 0x5a, 0x85, 0xd0, 0xd2, 0xca, 0x12,
	0x12, 0x12, 0x50, 0xa1, 0xb4, 0x48, 0x08, 0x10, 0x88, 0x46, 0x55, 0xa9, 0x0a, 0x89, 0xd8, 0x20,
	0xb8, 0x8c, 0xf2, 0xb3, 0x8e, 0x82, 0xea, 0x1f, 0xd6, 0xeb, 0x88, 0xf0, 0x04, 0x5c, 0xf3, 0x06,
	0xbc, 0x04, 0xcf, 0xc0, 0x35, 0x0f, 0xc2, 0x33, 0xb0, 0x3b, 0x5e, 0xdb, 0x31, 0xb1, 0x43, 0xc3,
	0x05, 0x77, 0xd9, 0x9d, 0x33, 0x67, 0xe6, 0x9c, 0x19, 0x6f, 0x60, 0xd3, 0xbf, 0xec, 0xf4, 0xac,
	0x41, 0x67, 0xe4, 0x8d, 0x8f, 0x0f, 0x3d, 0xee, 0x0a, 0x97, 0xac, 0xfb, 0x8c, 0x8f, 0x47, 0x7d,
	0xd6, 0xb9, 0xec, 0x4e, 0x18, 0xaf, 0x6d, 0x4b, 0x44, 0xdf, 0xb5, 0x6d, 0xd7, 0xe9, 0x88, 0x89,
	0xc7, 0xfc, 0x10, 0x55, 0xdb, 0xd2, 0x89, 0x61, 0x28, 0xbc, 0x34, 0xbf, 0x1b, 0x00, 0xed, 0x97,
	0x27, 0xd6, 0x60, 0x7c, 0x7c, 0xc1, 0x26, 0xe4, 0x2e, 0x94, 0xde, 0xc8, 0x94, 0xaa, 0x71, 0x60,
	0xdc, 0xfe, 0xbf, 0x5e, 0x3d, 0x4c, 0x11, 0x1f, 0x22, 0x50, 0xc5, 0x29, 0xa2, 0x48, 0x15, 0xca,
	0x6f, 0xb9, 0xd5, 0xec, 0xda, 0xac, 0x5a, 0x90, 0x09, 0xab, 0x34, 0x3a, 0xaa, 0x48, 0xb3, 0xc7,
	0x9f, 0x0f, 0x06, 0xbc, 0x5a, 0x94, 0x91, 0x75, 0x1a, 0x1d, 0xc9, 0x43, 0x58, 0x3d, 0x77, 0x04,
	0xe3, 0x56, 0xb7, 0xcf, 0xaa, 0x25, 0x19, 0xab, 0xd4, 0x6b, 0x33, 0x65, 0x62, 0x04, 0x4d, 0xc0,
	0xe4, 0xa6, 0xec, 0xd4, 0x0d, 0x78, 0x9f, 0x21, 0xed, 0x12, 0xd2, 0x4e, 0xdd, 0x98, 0x9f, 0x60,
	0x53, 0x2b, 0x69, 0x33, 0xdf, 0x1f, 0xb9, 0x4e, 0xc3, 0x1a, 0x92, 0x3b, 0x50, 0x94, 0xba, 0x50,
	0x4f, 0xa5, 0x7e, 0x3d, 0x4b, 0x0f, 0x0a, 0xa7, 0x0a, 0x25, 0x7b, 0x5b, 0x6e, 0xb8, 0x8e, 0x35,
	0x1a, 0xa2, 0x9c, 0x4a, 0xfd, 0x20, 0x0b, 0x1f, 0x22, 0x1a, 0x68, 0x22, 0xd5, 0x78, 0x73, 0x12,
	0xbb, 0xf8, 0xca, 0x1f, 0x92, 0x7b, 0x50, 0x6a, 0x79, 0x8c, 0x6b, 0x17, 0x67, 0xab, 0xb6, 0x7a,
	0xef, 0x59, 0x5f, 0xb4, 0x3c, 0x8a, 0x30, 0xf2, 0x04, 0x56, 0x74, 0xc7, 0xbe, 0x2c, 0x5c, 0xcc,
	0x2b, 0x3c, 0xad, 0x8b, 0xc6, 0x19, 0x66, 0x10, 0x97, 0xa6, 0xcc, 0x27, 0x8f, 0x60, 0xf5, 0x94,
	0xf3, 0xb6, 0xe8, 0x8a, 0xc0, 0xd7, 0xaa, 0x77, 0x67, 0xc8, 0x24, 0xc2, 0xd5, 0x18, 0x9a, 0xc0,
	0x23, 0xaf, 0x0a, 0x57, 0xf1, 0xca, 0xfc, 0x6c, 0xc0, 0x7a, 0x22, 0x99, 0xfa, 0x1e, 0x39, 0x91,
	0x17, 0x48, 0xd4, 0x0e, 0x6c, 0xbb, 0xcb, 0x27, 0x57, 0x2a, 0x9f, 0x4e, 0x21, 0x47, 0x50, 0x96,
	0x2a, 0x82, 0x4b, 0x11, 0x39, 0x91, 0xd3, 0x86, 0x04, 0xd1, 0x08, 0x69, 0x7e, 0x49, 0x5a, 0x39,
	0x63, 0x42, 0x0d, 0x60, 0xa1, 0xa9, 0xef, 0xc0, 0x72, 0x9b, 0x7d, 0x68, 0x06, 0x36, 0x2a, 0x2f,
	0x51, 0x7d, 0x22, 0x26, 0xac, 0x9d, 0x3a, 0x82, 0x8f, 0x98, 0xdf, 0x70, 0x03, 0x47, 0xe8, 0x45,
	0x4e, 0xdd, 0xa9, 0x3d, 0x97, 0x25, 0x9b, 0xec, 0xa3, 0xc0, 0x5d, 0x5e, 0xa1, 0xd1, 0xd1, 0xfc,
	0x66, 0xc0, 0xce, 0xcc, 0xd8, 0x94, 0x58, 0xf6, 0x8f, 0x76, 0x92, 0x3c, 0x80, 0x25, 0xac, 0x87,
	0x8d, 0x57, 0xea, 0xfb, 0xd9, 0x89, 0x2a, 0x05, 0x61, 0x34, 0x44, 0x9b, 0x5f, 0x0d, 0xd8, 0x48,
	0xb9, 0xa9, 0x66, 0xbb, 0x01, 0xc5, 0x53, 0xd7, 0xc2, 0x96, 0x57, 0xa8, 0xfa, 0x99, 0x5e, 0xb4,
	0xc2, 0x62, 0x8b, 0xf6, 0x0c, 0xca, 0xda, 0x45, 0xd9, 0x9b, 0x9a, 0xf2, 0xad, 0x3f, 0xed, 0x7b,
	0xd8, 0x61, 0x94, 0x25, 0x3f, 0xb7, 0xad, 0x34, 0xe4, 0x2f, 0x8c, 0x8d, 0xed, 0x29, 0x2c, 0x64,
	0xcf, 0x0f, 0x03, 0xd6, 0x34, 0x55, 0xd3, 0x15, 0x23, 0x8b, 0x3c, 0x96, 0x46, 0x8c, 0x99, 0x23,
	0xa6, 0xde, 0xcd, 0xbd, 0x2c, 0x2e, 0x44, 0xe3, 0xe3, 0x99, 0xe0, 0xe5, 0xa7, 0xbf, 0x98, 0x8b,
	0x2f, 0xfe, 0x9b, 0xf6, 0xf1, 0x29, 0x94, 0xb5, 0x7e, 0x3d, 0x63, 0x73, 0xae, 0x8f, 0x28, 0x40,
	0x32, 0x44, 0x49, 0x27, 0x65, 0x58, 0xc2, 0x56, 0xea, 0x3f, 0x8b, 0x50, 0xd1, 0x58, 0x7c, 0x91,
	0xce, 0x63, 0x8d, 0x94, 0x0d, 0x5b, 0x1e, 0xa9, 0x65, 0xf1, 0xca, 0x90, 0xdc, 0x8e, 0xda, 0x5e,
	0x7e, 0x4c, 0x6d, 0xce, 0x59, 0xfc, 0x3c, 0xc9, 0x6d, 0xca, 0x26, 0x0a, 0xd7, 0x2c, 0x9b, 0x28,
	0x59, 0xc1, 0xd7, 0x70, 0x2d, 0x21, 0x52, 0x52, 0xfc, 0xb9, 0x6c, 0x66, 0x4e, 0x0c, 0x33, 0x35,
	0xe5, 0x3b, 0xd8, 0x4e, 0x28, 0x71, 0x3e, 0x6d, 0xc1, 0x59, 0xd7, 0x26, 0xfb, 0x39, 0xc9, 0x88,
	0x51, 0xec, 0x37, 0xb2, 0x8d, 0xc6, 0xf8, 0x7d, 0x83, 0x5c, 0xc4, 0x9f, 0x90, 0xb6, 0x5a, 0x7a,
	0x98, 0xb3, 0x8f, 0x8a, 0x6d, 0x37, 0x37, 0xa4, 0xba, 0xa4, 0xbf, 0xff, 0xaf, 0x29, 0x23, 0x73,
	0x52, 0xb4, 0xf8, 0xfd, 0x79, 0x51, 0xc9, 0xd9, 0x5b, 0xc6, 0x7f, 0xff, 0xa3, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xd3, 0xb6, 0x59, 0x09, 0x4d, 0x08, 0x00, 0x00,
}
