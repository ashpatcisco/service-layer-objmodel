// Code generated by protoc-gen-go.
// source: sl_route_common.proto
// DO NOT EDIT!

package service_layer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Route Globals Get Message
type SLRouteGlobalsGetMsg struct {
}

func (m *SLRouteGlobalsGetMsg) Reset()                    { *m = SLRouteGlobalsGetMsg{} }
func (m *SLRouteGlobalsGetMsg) String() string            { return proto.CompactTextString(m) }
func (*SLRouteGlobalsGetMsg) ProtoMessage()               {}
func (*SLRouteGlobalsGetMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

// Route Globals Get Message Response
type SLRouteGlobalsGetMsgRsp struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Max VrfReg objects within a single VrfRegMsg message.
	MaxVrfregPerVrfregmsg uint32 `protobuf:"varint,2,opt,name=MaxVrfregPerVrfregmsg,json=maxVrfregPerVrfregmsg" json:"MaxVrfregPerVrfregmsg,omitempty"`
	// Max Route objects within a single RouteMsg message.
	MaxRoutePerRoutemsg uint32 `protobuf:"varint,3,opt,name=MaxRoutePerRoutemsg,json=maxRoutePerRoutemsg" json:"MaxRoutePerRoutemsg,omitempty"`
}

func (m *SLRouteGlobalsGetMsgRsp) Reset()                    { *m = SLRouteGlobalsGetMsgRsp{} }
func (m *SLRouteGlobalsGetMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLRouteGlobalsGetMsgRsp) ProtoMessage()               {}
func (*SLRouteGlobalsGetMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *SLRouteGlobalsGetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

// Route Global Stats Get Message
type SLRouteGlobalStatsGetMsg struct {
}

func (m *SLRouteGlobalStatsGetMsg) Reset()                    { *m = SLRouteGlobalStatsGetMsg{} }
func (m *SLRouteGlobalStatsGetMsg) String() string            { return proto.CompactTextString(m) }
func (*SLRouteGlobalStatsGetMsg) ProtoMessage()               {}
func (*SLRouteGlobalStatsGetMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

// Route Global Stats Get Message Response
type SLRouteGlobalStatsGetMsgRsp struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Num VRFs registered through the service layer.
	VrfCount uint32 `protobuf:"varint,2,opt,name=VrfCount,json=vrfCount" json:"VrfCount,omitempty"`
	// Num Routes added through the service layer.
	RouteCount uint32 `protobuf:"varint,3,opt,name=RouteCount,json=routeCount" json:"RouteCount,omitempty"`
}

func (m *SLRouteGlobalStatsGetMsgRsp) Reset()                    { *m = SLRouteGlobalStatsGetMsgRsp{} }
func (m *SLRouteGlobalStatsGetMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLRouteGlobalStatsGetMsgRsp) ProtoMessage()               {}
func (*SLRouteGlobalStatsGetMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *SLRouteGlobalStatsGetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

// VRF Registration message
type SLVrfReg struct {
	// VRF Name.
	VrfName string `protobuf:"bytes,1,opt,name=VrfName,json=vrfName" json:"VrfName,omitempty"`
	// Default Admin distance for routes programmed by this application
	// Range [0-255]
	// This default value is used if route objects' AdminDistance is 0.
	// Refer to SLRouteCommon
	AdminDistance uint32 `protobuf:"varint,2,opt,name=AdminDistance,json=adminDistance" json:"AdminDistance,omitempty"`
	// In case the Service Layer -> RIB connection is lost, this specifies the
	// timeout period after which RIB will automatically purge the installed
	// routes, unless the service layer:
	//    1. Re-registers (VRF)
	//    2. Replay all routes
	//    3. And send EOF, before the purge timeout
	VrfPurgeIntervalSeconds uint32 `protobuf:"varint,3,opt,name=VrfPurgeIntervalSeconds,json=vrfPurgeIntervalSeconds" json:"VrfPurgeIntervalSeconds,omitempty"`
}

func (m *SLVrfReg) Reset()                    { *m = SLVrfReg{} }
func (m *SLVrfReg) String() string            { return proto.CompactTextString(m) }
func (*SLVrfReg) ProtoMessage()               {}
func (*SLVrfReg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

// VRF Registration messages.
type SLVrfRegMsg struct {
	// Registration Operation
	Oper SLRegOp `protobuf:"varint,1,opt,name=Oper,json=oper,enum=service_layer.SLRegOp" json:"Oper,omitempty"`
	// List of VRF registrations
	VrfRegMsgs []*SLVrfReg `protobuf:"bytes,2,rep,name=VrfRegMsgs,json=vrfRegMsgs" json:"VrfRegMsgs,omitempty"`
}

func (m *SLVrfRegMsg) Reset()                    { *m = SLVrfRegMsg{} }
func (m *SLVrfRegMsg) String() string            { return proto.CompactTextString(m) }
func (*SLVrfRegMsg) ProtoMessage()               {}
func (*SLVrfRegMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

func (m *SLVrfRegMsg) GetVrfRegMsgs() []*SLVrfReg {
	if m != nil {
		return m.VrfRegMsgs
	}
	return nil
}

// VRF Registration message Result
type SLVrfRegMsgRes struct {
	// Corresponding error code
	ErrStatus *SLErrorStatus `protobuf:"bytes,1,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// VRF Name
	VrfName string `protobuf:"bytes,2,opt,name=VrfName,json=vrfName" json:"VrfName,omitempty"`
}

func (m *SLVrfRegMsgRes) Reset()                    { *m = SLVrfRegMsgRes{} }
func (m *SLVrfRegMsgRes) String() string            { return proto.CompactTextString(m) }
func (*SLVrfRegMsgRes) ProtoMessage()               {}
func (*SLVrfRegMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *SLVrfRegMsgRes) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

// VRF Registration message Response.
type SLVrfRegMsgRsp struct {
	// Summary result of the bulk operation (refer to enum SLErrorStatus)
	// In general, the StatusSummary implies one of 3 things:
	// 1. SL_SUCCESS: signifies that the entire bulk operation was successful.
	//         In this case, the Results list is empty.
	// 2. SL_SOME_ERR: signifies that the operation failed for one or more
	//         entries. In this case, Results holds the result for
	//         each individual entry in the bulk.
	// 3. SL_RPC_XXX: signifies that the entire bulk operation failed.
	//         In this case, the Results list is empty.
	StatusSummary *SLErrorStatus `protobuf:"bytes,1,opt,name=StatusSummary,json=statusSummary" json:"StatusSummary,omitempty"`
	// In case of errors, this field indicates which entry in the bulk was
	// erroneous.
	Results []*SLVrfRegMsgRes `protobuf:"bytes,2,rep,name=Results,json=results" json:"Results,omitempty"`
}

func (m *SLVrfRegMsgRsp) Reset()                    { *m = SLVrfRegMsgRsp{} }
func (m *SLVrfRegMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLVrfRegMsgRsp) ProtoMessage()               {}
func (*SLVrfRegMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *SLVrfRegMsgRsp) GetStatusSummary() *SLErrorStatus {
	if m != nil {
		return m.StatusSummary
	}
	return nil
}

func (m *SLVrfRegMsgRsp) GetResults() []*SLVrfRegMsgRes {
	if m != nil {
		return m.Results
	}
	return nil
}

// VRF Get Message
type SLVrfRegGetMsg struct {
	// VRF name (key).
	// If the Key is not specified, then request up to the first
	// 'EntriesCount' entries.
	VrfName string `protobuf:"bytes,1,opt,name=VrfName,json=vrfName" json:"VrfName,omitempty"`
	// Number of entries requested
	EntriesCount uint32 `protobuf:"varint,2,opt,name=EntriesCount,json=entriesCount" json:"EntriesCount,omitempty"`
	// if GetNext is FALSE:
	//     request up to 'EntriesCount' entries starting from the key
	// If GetNext is TRUE, or if the key exact match is not found:
	//     request up to 'EntriesCount' entries starting from the key's next
	GetNext bool `protobuf:"varint,3,opt,name=GetNext,json=getNext" json:"GetNext,omitempty"`
}

func (m *SLVrfRegGetMsg) Reset()                    { *m = SLVrfRegGetMsg{} }
func (m *SLVrfRegGetMsg) String() string            { return proto.CompactTextString(m) }
func (*SLVrfRegGetMsg) ProtoMessage()               {}
func (*SLVrfRegGetMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

// VRF Get Message Response
type SLVrfRegGetMsgRsp struct {
	// End Of File.
	// When set to True, it indicates that the server has returned M, where
	// M < N, of the original N requested Entries.
	Eof bool `protobuf:"varint,1,opt,name=Eof,json=eof" json:"Eof,omitempty"`
	// Status of the Get operation
	ErrStatus *SLErrorStatus `protobuf:"bytes,2,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Returned entries as requested in the Get operation.
	// if ErrStatus is SL_SUCCESS, Entries contains the info requested
	Entries []*SLVrfReg `protobuf:"bytes,3,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *SLVrfRegGetMsgRsp) Reset()                    { *m = SLVrfRegGetMsgRsp{} }
func (m *SLVrfRegGetMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLVrfRegGetMsgRsp) ProtoMessage()               {}
func (*SLVrfRegGetMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *SLVrfRegGetMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLVrfRegGetMsgRsp) GetEntries() []*SLVrfReg {
	if m != nil {
		return m.Entries
	}
	return nil
}

// VRF Get Stats Message Response
type SLVRFGetStatsMsgRes struct {
	// VRF name as key
	VrfName string `protobuf:"bytes,1,opt,name=VrfName,json=vrfName" json:"VrfName,omitempty"`
	// Num VRF Routes added through the service layer.
	RouteCount uint32 `protobuf:"varint,2,opt,name=RouteCount,json=routeCount" json:"RouteCount,omitempty"`
}

func (m *SLVRFGetStatsMsgRes) Reset()                    { *m = SLVRFGetStatsMsgRes{} }
func (m *SLVRFGetStatsMsgRes) String() string            { return proto.CompactTextString(m) }
func (*SLVRFGetStatsMsgRes) ProtoMessage()               {}
func (*SLVRFGetStatsMsgRes) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{10} }

// VRF Get Stats Message Response
type SLVRFGetStatsMsgRsp struct {
	// End Of File.
	// When set to True, it indicates that the server has returned M, where
	// M < N, of the original N requested Entries.
	Eof bool `protobuf:"varint,1,opt,name=Eof,json=eof" json:"Eof,omitempty"`
	// Status of the Get Stats operation
	ErrStatus *SLErrorStatus `protobuf:"bytes,2,opt,name=ErrStatus,json=errStatus" json:"ErrStatus,omitempty"`
	// Returned entries as requested in the Get Stats operation.
	// if ErrStatus is SL_SUCCESS, Entries contains the info requested
	Entries []*SLVRFGetStatsMsgRes `protobuf:"bytes,3,rep,name=Entries,json=entries" json:"Entries,omitempty"`
}

func (m *SLVRFGetStatsMsgRsp) Reset()                    { *m = SLVRFGetStatsMsgRsp{} }
func (m *SLVRFGetStatsMsgRsp) String() string            { return proto.CompactTextString(m) }
func (*SLVRFGetStatsMsgRsp) ProtoMessage()               {}
func (*SLVRFGetStatsMsgRsp) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{11} }

func (m *SLVRFGetStatsMsgRsp) GetErrStatus() *SLErrorStatus {
	if m != nil {
		return m.ErrStatus
	}
	return nil
}

func (m *SLVRFGetStatsMsgRsp) GetEntries() []*SLVRFGetStatsMsgRes {
	if m != nil {
		return m.Entries
	}
	return nil
}

// Common IPv4/IPv6 route attributes.
type SLRouteCommon struct {
	// Adminstrative distance of the route. [0-255].
	// RIB uses this field to break the tie when multiple
	// sources install the same route.
	// Lower distance is preferred over higher distance.
	// The per route object admin distance overrides the default's admin
	// distance set at VRF registration. see SLVrfReg
	AdminDistance uint32 `protobuf:"varint,1,opt,name=AdminDistance,json=adminDistance" json:"AdminDistance,omitempty"`
	// Local label associated with this route.
	// This is an optional field that can be used to simulatenously setup an
	// ILM entry (e.g. head end of an MPLS LSP) for the same route.
	// This is especially useful when setting up an MPLS LSP (a /32-route
	// towards the remote LSP peer e.g. 2.2.2.2/32) which is required for MPLS
	// VPN labeled routes.
	//
	// Note: MPLS VPN labeled routes can resolve only on:
	//    - /32-routes with valid LocalLabel and a valid egress MPLS path label
	LocalLabel uint32 `protobuf:"varint,2,opt,name=LocalLabel,json=localLabel" json:"LocalLabel,omitempty"`
	// Route Tag.
	// Routes are usually tagged to prevent loops during redistribution between
	// protocols.
	Tag uint32 `protobuf:"varint,3,opt,name=Tag,json=tag" json:"Tag,omitempty"`
}

func (m *SLRouteCommon) Reset()                    { *m = SLRouteCommon{} }
func (m *SLRouteCommon) String() string            { return proto.CompactTextString(m) }
func (*SLRouteCommon) ProtoMessage()               {}
func (*SLRouteCommon) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{12} }

// Route Path attributes.
//
// FRR Note:
// Failover from primary to backup is based on the trigger used (e.g.
// link down, BFD, etc.). Revertion is mainly timeout based. The timeout
// value is platform specific and is not configurable.
type SLRoutePath struct {
	// One of IPv4 or IPv6 address
	NexthopAddress *SLIpAddress `protobuf:"bytes,1,opt,name=NexthopAddress,json=nexthopAddress" json:"NexthopAddress,omitempty"`
	// Outgoing interface name for the path.
	NexthopInterface *SLInterface `protobuf:"bytes,2,opt,name=NexthopInterface,json=nexthopInterface" json:"NexthopInterface,omitempty"`
	// Load metric for this path.
	// Used for equal/unequal cost load balancing of traffic distribution.
	LoadMetric uint32 `protobuf:"varint,3,opt,name=LoadMetric,json=loadMetric" json:"LoadMetric,omitempty"`
	// Path VRF name. This field is used ONLY if the path is in a different
	// VRF than the route (e.g. VPN cases)
	VrfName string `protobuf:"bytes,4,opt,name=VrfName,json=vrfName" json:"VrfName,omitempty"`
	// Route Metric.
	// The metric is typically based on information like load, hop count,
	// MTU, reliability of the path, etc.
	Metric uint32 `protobuf:"varint,5,opt,name=Metric,json=metric" json:"Metric,omitempty"`
	// Path identifier.
	// Path-id is used to uniquely identify a path when it comes to
	// protection (Fast Re-Route - FRR). It is not used otherwise.
	//
	// In general, for FRR, There are 3 main path attributes:
	//    1. Primary. The path is the main path to carry traffic.
	//    2. Protected. A primary path with a configured backup path.
	//    3. Backup. The path is protecting a primary path.
	//
	// NOTE1: a primary path (A) can be simultaneously protected (by B), and
	// acting as a backup for another path (C).
	// In this example, the primary path C is protected by A (which happens to
	// be primary). So the primary path (A) is Primary, Protected, and Backup.
	//
	// The following are various path types based on combinations of attributes:
	//    1. Pure Primary i.e. Not Protected and is not a Backup.
	//       => PathId is optional. ProtectedPathBitmap = 0x0
	//    2. Primary and Protected. Path is not a Backup.
	//       => PathId is mandatory. ProtectedPathBitmap = 0x0
	//    3. Primary and Not Protected. Path is also a Backup.
	//       => PathId is mandatory. ProtectedPathBitmap = 0xYYYY
	//    4. Primary and Protected. Path is also a Backup.
	//       => PathId is mandatory. ProtectedPathBitmap = 0xYYYY
	//    5. Pure Backup. Protection is not allowed for backup paths.
	//       => PathId is mandatory. ProtectedPathBitmap = 0xYYYY
	//
	// NOTE2: Pure backup path-id uses a different range than primary path-ids.
	// The valid range of primary path IDs, and pure backup path IDS are
	// platform dependent and can be retrieved through the client init message.
	PathId uint32 `protobuf:"varint,6,opt,name=PathId,json=pathId" json:"PathId,omitempty"`
	// Path protection bitmap.
	// The bitmap of paths this Backup path is protecting.
	// Example: If this path is protecting paths with IDs 4, 5 and 6, then
	// set bitmap to:
	// 0x38 ==> 0011 1000
	//            || |-- path 4
	//            ||-- path 5
	//            |-- path 6
	// (1 << (pathId_1 - 1))  | (1 << (pathId_2 - 1)) | (1 << (pathId_3 - 1))
	ProtectedPathBitmap []uint64 `protobuf:"varint,7,rep,name=ProtectedPathBitmap,json=protectedPathBitmap" json:"ProtectedPathBitmap,omitempty"`
	// MPLS label stack.
	// Stack of labels that is pushed when the packet is switched out.
	// Label size is LSB 20 bits. Forwarding will set EXP, TTL and BOS.
	// For primary path, typically only 1 label is used.
	// For backup paths, more than 1 label can be used. If more than one label
	// is used, remote backup addresses must be specified.
	// The maximum number of labels pushed for primary and backup are
	// platform dependent.
	LabelStack []uint32 `protobuf:"varint,8,rep,name=LabelStack,json=labelStack" json:"LabelStack,omitempty"`
	// MPLS Remote router backup address.
	// This field is used only for backup MPLS path with more than one label
	// For N+1 backup labels, N remote backup addresses must be specified.
	RemoteAddress []*SLIpAddress `protobuf:"bytes,9,rep,name=RemoteAddress,json=remoteAddress" json:"RemoteAddress,omitempty"`
}

func (m *SLRoutePath) Reset()                    { *m = SLRoutePath{} }
func (m *SLRoutePath) String() string            { return proto.CompactTextString(m) }
func (*SLRoutePath) ProtoMessage()               {}
func (*SLRoutePath) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{13} }

func (m *SLRoutePath) GetNexthopAddress() *SLIpAddress {
	if m != nil {
		return m.NexthopAddress
	}
	return nil
}

func (m *SLRoutePath) GetNexthopInterface() *SLInterface {
	if m != nil {
		return m.NexthopInterface
	}
	return nil
}

func (m *SLRoutePath) GetRemoteAddress() []*SLIpAddress {
	if m != nil {
		return m.RemoteAddress
	}
	return nil
}

func init() {
	proto.RegisterType((*SLRouteGlobalsGetMsg)(nil), "service_layer.SLRouteGlobalsGetMsg")
	proto.RegisterType((*SLRouteGlobalsGetMsgRsp)(nil), "service_layer.SLRouteGlobalsGetMsgRsp")
	proto.RegisterType((*SLRouteGlobalStatsGetMsg)(nil), "service_layer.SLRouteGlobalStatsGetMsg")
	proto.RegisterType((*SLRouteGlobalStatsGetMsgRsp)(nil), "service_layer.SLRouteGlobalStatsGetMsgRsp")
	proto.RegisterType((*SLVrfReg)(nil), "service_layer.SLVrfReg")
	proto.RegisterType((*SLVrfRegMsg)(nil), "service_layer.SLVrfRegMsg")
	proto.RegisterType((*SLVrfRegMsgRes)(nil), "service_layer.SLVrfRegMsgRes")
	proto.RegisterType((*SLVrfRegMsgRsp)(nil), "service_layer.SLVrfRegMsgRsp")
	proto.RegisterType((*SLVrfRegGetMsg)(nil), "service_layer.SLVrfRegGetMsg")
	proto.RegisterType((*SLVrfRegGetMsgRsp)(nil), "service_layer.SLVrfRegGetMsgRsp")
	proto.RegisterType((*SLVRFGetStatsMsgRes)(nil), "service_layer.SLVRFGetStatsMsgRes")
	proto.RegisterType((*SLVRFGetStatsMsgRsp)(nil), "service_layer.SLVRFGetStatsMsgRsp")
	proto.RegisterType((*SLRouteCommon)(nil), "service_layer.SLRouteCommon")
	proto.RegisterType((*SLRoutePath)(nil), "service_layer.SLRoutePath")
}

var fileDescriptor7 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x93, 0x10, 0x27, 0xd3, 0x3a, 0x2a, 0x0e, 0x6d, 0xa2, 0xf0, 0x21, 0x64, 0x71, 0x40,
	0x1c, 0x2a, 0x28, 0x48, 0x20, 0xc4, 0x81, 0x16, 0xda, 0xaa, 0x52, 0xfa, 0xa1, 0x0d, 0xe2, 0x1a,
	0x6d, 0xed, 0x8d, 0x1b, 0xe1, 0x2f, 0xad, 0x37, 0x51, 0x7b, 0xe7, 0xca, 0x0d, 0xc4, 0x99, 0x1f,
	0xc2, 0x7f, 0x63, 0x76, 0xd7, 0x4e, 0xed, 0x3a, 0xad, 0x40, 0xe5, 0x94, 0xec, 0xcc, 0x7b, 0x3b,
	0xef, 0xcd, 0xce, 0xae, 0x61, 0x3d, 0x0d, 0xc6, 0x3c, 0x9e, 0x09, 0x36, 0x76, 0xe3, 0x30, 0x8c,
	0xa3, 0xcd, 0x84, 0xc7, 0x22, 0xb6, 0xad, 0x94, 0xf1, 0xf9, 0xd4, 0x65, 0xe3, 0x80, 0x5e, 0x30,
	0x3e, 0x90, 0x28, 0x9d, 0x1f, 0x8b, 0x8b, 0x84, 0xa5, 0x1a, 0xe5, 0x6c, 0xc0, 0xbd, 0xd1, 0x90,
	0x48, 0xf6, 0x7e, 0x10, 0x9f, 0xd2, 0x20, 0xdd, 0x67, 0xe2, 0x30, 0xf5, 0x9d, 0xdf, 0x06, 0xf4,
	0x96, 0x25, 0x48, 0x9a, 0xd8, 0x6f, 0xa1, 0xbd, 0xcb, 0xf9, 0x48, 0x50, 0x31, 0x4b, 0xfb, 0xc6,
	0x63, 0xe3, 0xe9, 0xca, 0xd6, 0x83, 0xcd, 0x52, 0xb5, 0xcd, 0xd1, 0x10, 0x11, 0x71, 0x86, 0x21,
	0x6d, 0x96, 0xc3, 0xed, 0x57, 0xb0, 0x7e, 0x48, 0xcf, 0x3f, 0xf3, 0x09, 0x67, 0xfe, 0x09, 0xe3,
	0xfa, 0x4f, 0x98, 0xfa, 0xfd, 0x1a, 0xee, 0x63, 0x91, 0xf5, 0x70, 0x59, 0xd2, 0x7e, 0x0e, 0x5d,
	0x64, 0x29, 0x35, 0x18, 0x57, 0xbf, 0x92, 0x53, 0x57, 0x9c, 0x6e, 0x58, 0x4d, 0x39, 0x03, 0xe8,
	0x97, 0xe4, 0xcb, 0xf2, 0xb9, 0xb7, 0x1f, 0x06, 0xdc, 0xbf, 0x2e, 0x79, 0x5b, 0x7f, 0x03, 0x68,
	0xa1, 0xec, 0x0f, 0xf1, 0x2c, 0x12, 0x99, 0xa5, 0xd6, 0x3c, 0x5b, 0xdb, 0x8f, 0x00, 0x54, 0x51,
	0x9d, 0xd5, 0xe2, 0x81, 0x2f, 0x22, 0xce, 0x57, 0x03, 0x5a, 0xa3, 0x21, 0xd2, 0x09, 0xf3, 0xed,
	0x3e, 0x98, 0xf8, 0xef, 0x88, 0x86, 0x4c, 0x49, 0x68, 0x13, 0x73, 0xae, 0x97, 0xf6, 0x13, 0xb0,
	0xb6, 0xbd, 0x70, 0x1a, 0x7d, 0x9c, 0xa6, 0x82, 0x46, 0x2e, 0xcb, 0xea, 0x58, 0xb4, 0x18, 0xb4,
	0xdf, 0x40, 0x0f, 0xf9, 0x27, 0x33, 0xee, 0xb3, 0x83, 0x48, 0xa0, 0x78, 0xb4, 0xc9, 0xdc, 0x38,
	0xf2, 0xd2, 0xac, 0x72, 0x6f, 0xbe, 0x3c, 0xed, 0x70, 0x58, 0xc9, 0x55, 0x60, 0x43, 0xec, 0x67,
	0xd0, 0x38, 0x4e, 0x18, 0x57, 0x2a, 0x3a, 0x5b, 0x1b, 0x95, 0x46, 0x20, 0xec, 0x38, 0x21, 0x8d,
	0x18, 0x31, 0xf6, 0x6b, 0x80, 0x05, 0x31, 0x45, 0x5d, 0x75, 0x6c, 0x5d, 0xaf, 0xc2, 0xd0, 0x10,
	0x02, 0xf3, 0x05, 0xd4, 0x99, 0x40, 0xa7, 0x50, 0x93, 0xb0, 0xf4, 0x56, 0x87, 0x50, 0xe8, 0x5d,
	0xad, 0xd4, 0x3b, 0xe7, 0x9b, 0x51, 0x2e, 0x84, 0xa7, 0xbd, 0x03, 0x96, 0xa6, 0x8d, 0x66, 0x61,
	0x48, 0xf9, 0xc5, 0x5f, 0x15, 0xb3, 0xd2, 0x22, 0x05, 0x7d, 0x9b, 0xa8, 0x79, 0x16, 0x88, 0xdc,
	0xf4, 0xc3, 0x6b, 0x4c, 0x6b, 0x73, 0xc4, 0xe4, 0x1a, 0xed, 0x9c, 0x5d, 0xca, 0xd1, 0xf3, 0x77,
	0xc3, 0xb9, 0x3b, 0xb0, 0xba, 0x1b, 0x09, 0x3e, 0x65, 0x69, 0x71, 0xbc, 0x56, 0x59, 0x21, 0x26,
	0xd9, 0xb8, 0xcf, 0x11, 0x3b, 0xd7, 0xf3, 0xd5, 0x22, 0xa6, 0xaf, 0x97, 0xce, 0x77, 0x03, 0xee,
	0x96, 0x4b, 0x49, 0xf3, 0x6b, 0x50, 0xdf, 0x8d, 0x27, 0xaa, 0x52, 0x8b, 0xd4, 0x59, 0x3c, 0x29,
	0xf7, 0xbd, 0xf6, 0x6f, 0x7d, 0x7f, 0x01, 0x66, 0xa6, 0x10, 0xab, 0xdf, 0x78, 0xf6, 0x66, 0xa6,
	0xda, 0x39, 0x86, 0x2e, 0x06, 0xc9, 0x1e, 0x4a, 0x52, 0xb7, 0x30, 0x3b, 0xfd, 0xeb, 0xbb, 0x50,
	0xbe, 0x44, 0xb5, 0xca, 0x25, 0xfa, 0x65, 0x2c, 0xd9, 0xf1, 0xbf, 0x3b, 0x7d, 0x77, 0xd5, 0xa9,
	0x53, 0x75, 0x7a, 0xd5, 0xd4, 0xa5, 0x69, 0x1f, 0x47, 0x6e, 0x98, 0xb9, 0x90, 0x2f, 0x72, 0xf5,
	0x4a, 0x1b, 0xcb, 0xae, 0x34, 0x5a, 0x1f, 0xc6, 0x2e, 0x0d, 0x86, 0xf4, 0x94, 0x05, 0xb9, 0xf5,
	0x60, 0x11, 0x91, 0x16, 0x3f, 0xd1, 0xfc, 0x55, 0xac, 0x0b, 0xea, 0x3b, 0x3f, 0xeb, 0xf2, 0x2e,
	0xeb, 0xc7, 0x91, 0x8a, 0x33, 0x9c, 0xf5, 0x8e, 0x1c, 0x86, 0xb3, 0x38, 0xd9, 0xf6, 0x3c, 0x9c,
	0xc1, 0xfc, 0x66, 0x0d, 0x2a, 0xea, 0x0f, 0x72, 0x04, 0xe9, 0x44, 0x25, 0x86, 0xbd, 0x07, 0x6b,
	0xd9, 0x1e, 0xea, 0xe1, 0x98, 0xd0, 0xec, 0x05, 0x5a, 0xba, 0x4b, 0x8e, 0x20, 0x6b, 0xd1, 0x15,
	0x8e, 0x76, 0x43, 0xbd, 0x43, 0x86, 0x3d, 0x71, 0xf3, 0xd7, 0x30, 0x58, 0x44, 0x8a, 0x23, 0xd0,
	0x28, 0x8f, 0xc0, 0x06, 0x34, 0x33, 0xd6, 0x1d, 0xc5, 0x6a, 0x86, 0x9a, 0x81, 0x71, 0xe9, 0xf2,
	0xc0, 0xeb, 0x37, 0x75, 0x3c, 0x51, 0x2b, 0xf9, 0xf5, 0x38, 0xc1, 0x8f, 0x1d, 0x73, 0x05, 0xf3,
	0x24, 0x60, 0x67, 0x2a, 0x42, 0x9a, 0xf4, 0x4d, 0x3c, 0xb8, 0x06, 0xe9, 0x26, 0xd5, 0x94, 0xd2,
	0x26, 0x5b, 0x8a, 0xa7, 0xe7, 0x7e, 0xe9, 0xb7, 0x10, 0x28, 0xb5, 0x2d, 0x22, 0xf6, 0x7b, 0xb0,
	0x08, 0x0b, 0x91, 0x97, 0xb7, 0xb1, 0xad, 0x86, 0xe0, 0xa6, 0x36, 0x5a, 0xbc, 0x48, 0x38, 0x6d,
	0xaa, 0xcf, 0xef, 0xcb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xc0, 0xf5, 0xee, 0xbd, 0x07,
	0x00, 0x00,
}
