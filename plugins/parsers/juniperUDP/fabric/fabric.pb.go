// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fabric.proto

/*
Package fabric is a generated protocol buffer package.

It is generated from these files:
	fabric.proto

It has these top-level messages:
	FabricMessage
	EdgeStats
	ClassStats
	Counters
*/
package fabric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import telemetry_top "github.com/influxdata/telegraf/plugins/parsers/juniperUDP/telemetry_top"


// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FabricMessageSensorLocation int32

const (
	FabricMessage_Linecard      FabricMessageSensorLocation = 1
	FabricMessage_Switch_Fabric FabricMessageSensorLocation = 2
)

var FabricMessageSensorLocation_name = map[int32]string{
	1: "Linecard",
	2: "Switch_Fabric",
}
var FabricMessageSensorLocation_value = map[string]int32{
	"Linecard":      1,
	"Switch_Fabric": 2,
}

func (x FabricMessageSensorLocation) Enum() *FabricMessageSensorLocation {
	p := new(FabricMessageSensorLocation)
	*p = x
	return p
}
func (x FabricMessageSensorLocation) String() string {
	return proto.EnumName(FabricMessageSensorLocation_name, int32(x))
}
func (x *FabricMessageSensorLocation) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FabricMessageSensorLocation_value, data, "FabricMessageSensorLocation")
	if err != nil {
		return err
	}
	*x = FabricMessageSensorLocation(value)
	return nil
}
func (FabricMessageSensorLocation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0, 0}
}

type EdgeStatsIdentifierType int32

const (
	EdgeStats_Switch_Fabric EdgeStatsIdentifierType = 1
	EdgeStats_Linecard      EdgeStatsIdentifierType = 2
)

var EdgeStatsIdentifierType_name = map[int32]string{
	1: "Switch_Fabric",
	2: "Linecard",
}
var EdgeStatsIdentifierType_value = map[string]int32{
	"Switch_Fabric": 1,
	"Linecard":      2,
}

func (x EdgeStatsIdentifierType) Enum() *EdgeStatsIdentifierType {
	p := new(EdgeStatsIdentifierType)
	*p = x
	return p
}
func (x EdgeStatsIdentifierType) String() string {
	return proto.EnumName(EdgeStatsIdentifierType_name, int32(x))
}
func (x *EdgeStatsIdentifierType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EdgeStatsIdentifierType_value, data, "EdgeStatsIdentifierType")
	if err != nil {
		return err
	}
	*x = EdgeStatsIdentifierType(value)
	return nil
}
func (EdgeStatsIdentifierType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type FabricMessage struct {
	Edges            []*EdgeStats                 `protobuf:"bytes,1,rep,name=edges" json:"edges,omitempty"`
	Location         *FabricMessageSensorLocation `protobuf:"varint,2,opt,name=location,enum=FabricMessageSensorLocation" json:"location,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *FabricMessage) Reset()                    { *m = FabricMessage{} }
func (m *FabricMessage) String() string            { return proto.CompactTextString(m) }
func (*FabricMessage) ProtoMessage()               {}
func (*FabricMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FabricMessage) GetEdges() []*EdgeStats {
	if m != nil {
		return m.Edges
	}
	return nil
}

func (m *FabricMessage) GetLocation() FabricMessageSensorLocation {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return FabricMessage_Linecard
}

type EdgeStats struct {
	SourceType       *EdgeStatsIdentifierType `protobuf:"varint,1,opt,name=source_type,enum=EdgeStatsIdentifierType" json:"source_type,omitempty"`
	SourceSlot       *uint32                  `protobuf:"varint,2,opt,name=source_slot" json:"source_slot,omitempty"`
	SourcePfe        *uint32                  `protobuf:"varint,3,opt,name=source_pfe" json:"source_pfe,omitempty"`
	DestinationType  *EdgeStatsIdentifierType `protobuf:"varint,4,opt,name=destination_type,enum=EdgeStatsIdentifierType" json:"destination_type,omitempty"`
	DestinationSlot  *uint32                  `protobuf:"varint,5,opt,name=destination_slot" json:"destination_slot,omitempty"`
	DestinationPfe   *uint32                  `protobuf:"varint,6,opt,name=destination_pfe" json:"destination_pfe,omitempty"`
	ClassStats       []*ClassStats            `protobuf:"bytes,7,rep,name=class_stats" json:"class_stats,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *EdgeStats) Reset()                    { *m = EdgeStats{} }
func (m *EdgeStats) String() string            { return proto.CompactTextString(m) }
func (*EdgeStats) ProtoMessage()               {}
func (*EdgeStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *EdgeStats) GetSourceType() EdgeStatsIdentifierType {
	if m != nil && m.SourceType != nil {
		return *m.SourceType
	}
	return EdgeStats_Switch_Fabric
}

func (m *EdgeStats) GetSourceSlot() uint32 {
	if m != nil && m.SourceSlot != nil {
		return *m.SourceSlot
	}
	return 0
}

func (m *EdgeStats) GetSourcePfe() uint32 {
	if m != nil && m.SourcePfe != nil {
		return *m.SourcePfe
	}
	return 0
}

func (m *EdgeStats) GetDestinationType() EdgeStatsIdentifierType {
	if m != nil && m.DestinationType != nil {
		return *m.DestinationType
	}
	return EdgeStats_Switch_Fabric
}

func (m *EdgeStats) GetDestinationSlot() uint32 {
	if m != nil && m.DestinationSlot != nil {
		return *m.DestinationSlot
	}
	return 0
}

func (m *EdgeStats) GetDestinationPfe() uint32 {
	if m != nil && m.DestinationPfe != nil {
		return *m.DestinationPfe
	}
	return 0
}

func (m *EdgeStats) GetClassStats() []*ClassStats {
	if m != nil {
		return m.ClassStats
	}
	return nil
}

type ClassStats struct {
	Priority         *string   `protobuf:"bytes,1,opt,name=priority" json:"priority,omitempty"`
	TransmitCounts   *Counters `protobuf:"bytes,2,opt,name=transmit_counts" json:"transmit_counts,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ClassStats) Reset()                    { *m = ClassStats{} }
func (m *ClassStats) String() string            { return proto.CompactTextString(m) }
func (*ClassStats) ProtoMessage()               {}
func (*ClassStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClassStats) GetPriority() string {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return ""
}

func (m *ClassStats) GetTransmitCounts() *Counters {
	if m != nil {
		return m.TransmitCounts
	}
	return nil
}

type Counters struct {
	Packets               *uint64 `protobuf:"varint,1,opt,name=packets" json:"packets,omitempty"`
	Bytes                 *uint64 `protobuf:"varint,2,opt,name=bytes" json:"bytes,omitempty"`
	PacketsPerSecond      *uint64 `protobuf:"varint,3,opt,name=packets_per_second" json:"packets_per_second,omitempty"`
	BytesPerSecond        *uint64 `protobuf:"varint,4,opt,name=bytes_per_second" json:"bytes_per_second,omitempty"`
	DropPackets           *uint64 `protobuf:"varint,5,opt,name=drop_packets" json:"drop_packets,omitempty"`
	DropBytes             *uint64 `protobuf:"varint,6,opt,name=drop_bytes" json:"drop_bytes,omitempty"`
	DropPacketsPerSecond  *uint64 `protobuf:"varint,7,opt,name=drop_packets_per_second" json:"drop_packets_per_second,omitempty"`
	DropBytesPerSecond    *uint64 `protobuf:"varint,8,opt,name=drop_bytes_per_second" json:"drop_bytes_per_second,omitempty"`
	QueueDepthAverage     *uint64 `protobuf:"varint,9,opt,name=queue_depth_average" json:"queue_depth_average,omitempty"`
	QueueDepthCurrent     *uint64 `protobuf:"varint,10,opt,name=queue_depth_current" json:"queue_depth_current,omitempty"`
	QueueDepthPeak        *uint64 `protobuf:"varint,11,opt,name=queue_depth_peak" json:"queue_depth_peak,omitempty"`
	QueueDepthMaximum     *uint64 `protobuf:"varint,12,opt,name=queue_depth_maximum" json:"queue_depth_maximum,omitempty"`
	ErrorPackets          *uint64 `protobuf:"varint,13,opt,name=error_packets" json:"error_packets,omitempty"`
	ErrorPacketsPerSecond *uint64 `protobuf:"varint,14,opt,name=error_packets_per_second" json:"error_packets_per_second,omitempty"`
	XXX_unrecognized      []byte  `json:"-"`
}

func (m *Counters) Reset()                    { *m = Counters{} }
func (m *Counters) String() string            { return proto.CompactTextString(m) }
func (*Counters) ProtoMessage()               {}
func (*Counters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Counters) GetPackets() uint64 {
	if m != nil && m.Packets != nil {
		return *m.Packets
	}
	return 0
}

func (m *Counters) GetBytes() uint64 {
	if m != nil && m.Bytes != nil {
		return *m.Bytes
	}
	return 0
}

func (m *Counters) GetPacketsPerSecond() uint64 {
	if m != nil && m.PacketsPerSecond != nil {
		return *m.PacketsPerSecond
	}
	return 0
}

func (m *Counters) GetBytesPerSecond() uint64 {
	if m != nil && m.BytesPerSecond != nil {
		return *m.BytesPerSecond
	}
	return 0
}

func (m *Counters) GetDropPackets() uint64 {
	if m != nil && m.DropPackets != nil {
		return *m.DropPackets
	}
	return 0
}

func (m *Counters) GetDropBytes() uint64 {
	if m != nil && m.DropBytes != nil {
		return *m.DropBytes
	}
	return 0
}

func (m *Counters) GetDropPacketsPerSecond() uint64 {
	if m != nil && m.DropPacketsPerSecond != nil {
		return *m.DropPacketsPerSecond
	}
	return 0
}

func (m *Counters) GetDropBytesPerSecond() uint64 {
	if m != nil && m.DropBytesPerSecond != nil {
		return *m.DropBytesPerSecond
	}
	return 0
}

func (m *Counters) GetQueueDepthAverage() uint64 {
	if m != nil && m.QueueDepthAverage != nil {
		return *m.QueueDepthAverage
	}
	return 0
}

func (m *Counters) GetQueueDepthCurrent() uint64 {
	if m != nil && m.QueueDepthCurrent != nil {
		return *m.QueueDepthCurrent
	}
	return 0
}

func (m *Counters) GetQueueDepthPeak() uint64 {
	if m != nil && m.QueueDepthPeak != nil {
		return *m.QueueDepthPeak
	}
	return 0
}

func (m *Counters) GetQueueDepthMaximum() uint64 {
	if m != nil && m.QueueDepthMaximum != nil {
		return *m.QueueDepthMaximum
	}
	return 0
}

func (m *Counters) GetErrorPackets() uint64 {
	if m != nil && m.ErrorPackets != nil {
		return *m.ErrorPackets
	}
	return 0
}

func (m *Counters) GetErrorPacketsPerSecond() uint64 {
	if m != nil && m.ErrorPacketsPerSecond != nil {
		return *m.ErrorPacketsPerSecond
	}
	return 0
}

var E_FabricMessageExt = &proto.ExtensionDesc{
	ExtendedType:  (*telemetry_top.JuniperNetworksSensors)(nil),
	ExtensionType: (*FabricMessage)(nil),
	Field:         2,
	Name:          "fabricMessageExt",
	Tag:           "bytes,2,opt,name=fabricMessageExt",
	Filename:      "fabric.proto",
}

func init() {
	proto.RegisterType((*FabricMessage)(nil), "fabric_message")
	proto.RegisterType((*EdgeStats)(nil), "edge_stats")
	proto.RegisterType((*ClassStats)(nil), "class_stats")
	proto.RegisterType((*Counters)(nil), "counters")
	proto.RegisterEnum("FabricMessageSensorLocation", FabricMessageSensorLocation_name, FabricMessageSensorLocation_value)
	proto.RegisterEnum("EdgeStatsIdentifierType", EdgeStatsIdentifierType_name, EdgeStatsIdentifierType_value)
	proto.RegisterExtension(E_FabricMessageExt)
}

func init() { proto.RegisterFile("fabric.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x14, 0x85, 0xe5, 0xfc, 0x34, 0xce, 0x75, 0xfe, 0x70, 0x81, 0xb8, 0x29, 0xa2, 0xa9, 0x85, 0xa0,
	0x2b, 0x2f, 0xb2, 0x01, 0x21, 0x16, 0xa8, 0x12, 0x2c, 0x2a, 0x60, 0xd3, 0x07, 0x18, 0x4d, 0xed,
	0x9b, 0x74, 0x94, 0xd8, 0x33, 0xcc, 0x5c, 0xd3, 0x66, 0xcb, 0x63, 0xf0, 0x58, 0x2c, 0x78, 0x1e,
	0x94, 0x71, 0x1c, 0x26, 0x41, 0xaa, 0xba, 0x9c, 0x39, 0xe7, 0x9e, 0x6f, 0xe6, 0x24, 0x63, 0xe8,
	0xcd, 0xf9, 0x8d, 0x16, 0x69, 0xa2, 0xb4, 0x24, 0x39, 0x39, 0x26, 0x5c, 0x61, 0x8e, 0xa4, 0xd7,
	0x8c, 0xa4, 0xaa, 0x36, 0xe3, 0x5f, 0x1e, 0x0c, 0x2a, 0x17, 0xcb, 0xd1, 0x18, 0xbe, 0xc0, 0x70,
	0x02, 0x6d, 0xcc, 0x16, 0x68, 0x22, 0x6f, 0xda, 0xbc, 0x08, 0x66, 0x41, 0xb2, 0x59, 0x31, 0x43,
	0x9c, 0x4c, 0xf8, 0x0e, 0xfc, 0x95, 0x4c, 0x39, 0x09, 0x59, 0x44, 0x8d, 0xa9, 0x77, 0x31, 0x98,
	0x9d, 0x25, 0xfb, 0xe3, 0x89, 0xc1, 0xc2, 0x48, 0xcd, 0x6a, 0xdb, 0x65, 0xfb, 0xe7, 0xc7, 0x86,
	0xef, 0xc5, 0x33, 0x18, 0x1e, 0x28, 0x61, 0x0f, 0xfc, 0x2f, 0xa2, 0xc0, 0x94, 0xeb, 0x6c, 0xe4,
	0x85, 0x4f, 0xa0, 0x7f, 0x7d, 0x27, 0x28, 0xbd, 0x65, 0x9f, 0x6d, 0xe0, 0xa8, 0x11, 0xff, 0x6e,
	0x00, 0x38, 0xf0, 0xb7, 0x10, 0x18, 0x59, 0xea, 0x14, 0x19, 0xad, 0x15, 0x46, 0x9e, 0xe5, 0x9f,
	0x3a, 0xc7, 0x4b, 0x44, 0x86, 0x05, 0x89, 0xb9, 0x40, 0x6d, 0x2d, 0x5b, 0x76, 0x38, 0xd9, 0x0d,
	0x9a, 0x95, 0x24, 0x7b, 0xf0, 0x7e, 0xad, 0x9d, 0x00, 0x6c, 0x35, 0x35, 0xc7, 0xa8, 0xe9, 0x4a,
	0x1f, 0x60, 0x94, 0xa1, 0x21, 0x51, 0xd8, 0xe3, 0x56, 0xd0, 0xd6, 0xa3, 0xa1, 0x67, 0xfb, 0xd3,
	0x96, 0xdc, 0x76, 0xe3, 0x5f, 0xc2, 0xd0, 0x35, 0x6c, 0xf0, 0x47, 0xae, 0x7e, 0x0e, 0x41, 0xba,
	0xe2, 0xc6, 0x54, 0x98, 0xa8, 0x63, 0x7f, 0x8d, 0x5e, 0xe2, 0xec, 0x6d, 0x4a, 0x3d, 0xa0, 0xff,
	0x5f, 0xa3, 0xb7, 0xd7, 0x73, 0x23, 0xbe, 0xda, 0x8b, 0x0d, 0xc7, 0xe0, 0x2b, 0x2d, 0xa4, 0x16,
	0xb4, 0xb6, 0x8d, 0x76, 0x6b, 0x7c, 0x0c, 0x43, 0xd2, 0xbc, 0x30, 0xb9, 0x20, 0x96, 0xca, 0xb2,
	0x20, 0x63, 0x8b, 0x0b, 0x66, 0xdd, 0xc4, 0x2e, 0x51, 0x9b, 0xf8, 0x4f, 0x13, 0xfc, 0x7a, 0x11,
	0x3e, 0x87, 0x8e, 0xe2, 0xe9, 0x12, 0xc9, 0xd8, 0xa0, 0x96, 0x0d, 0x8a, 0xbc, 0xf0, 0x29, 0xb4,
	0x6f, 0xd6, 0x84, 0xd5, 0xf8, 0x6e, 0xf7, 0x1c, 0xc2, 0xad, 0x9b, 0x29, 0xd4, 0xcc, 0x60, 0x2a,
	0x8b, 0xcc, 0xf6, 0x5f, 0x59, 0xa6, 0xb6, 0x41, 0x3b, 0xe8, 0x1a, 0x5a, 0xae, 0xe1, 0x14, 0x7a,
	0x99, 0x96, 0x8a, 0xd5, 0xd8, 0xb6, 0x0b, 0x38, 0x01, 0xb0, 0x62, 0xc5, 0x3e, 0x72, 0xa5, 0xd7,
	0x30, 0x76, 0xe7, 0xdc, 0xfc, 0x8e, 0x9b, 0xff, 0x0a, 0x9e, 0xfd, 0x8b, 0x70, 0x5d, 0xbe, 0xeb,
	0x8a, 0xe1, 0xf8, 0x7b, 0x89, 0x25, 0xb2, 0x0c, 0x15, 0xdd, 0x32, 0xfe, 0x03, 0x35, 0x5f, 0x60,
	0xd4, 0x7d, 0xc0, 0x93, 0x96, 0x5a, 0x63, 0x41, 0x11, 0x1c, 0x5c, 0xd7, 0xf5, 0x28, 0xe4, 0xcb,
	0x28, 0x78, 0x20, 0x24, 0xe7, 0xf7, 0x22, 0x2f, 0xf3, 0xa8, 0xe7, 0x7a, 0x5e, 0x40, 0x1f, 0xb5,
	0x96, 0x7a, 0xd7, 0x49, 0xdf, 0xbd, 0xf8, 0x1b, 0x88, 0xf6, 0x54, 0xf7, 0x4e, 0x03, 0x27, 0xe6,
	0xfd, 0x25, 0x8c, 0xaa, 0x67, 0xfd, 0xb5, 0x7a, 0xd5, 0x9f, 0xee, 0x29, 0x1c, 0x27, 0x57, 0x65,
	0x21, 0x14, 0xea, 0x6f, 0x48, 0x77, 0x52, 0x2f, 0xcd, 0xb5, 0x7d, 0xcf, 0xf5, 0xdf, 0x62, 0x78,
	0xf0, 0x21, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x4d, 0xc8, 0xbc, 0x7e, 0x04, 0x00, 0x00,
}
