// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub/demo.proto

package hub

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	goshare "github.com/jamesshenjian/protos/src/go/goshare"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 模拟环境的订单记录
type DemoOrder struct {
	TaId                 string                  `protobuf:"bytes,1,opt,name=ta_id,json=taId,proto3" json:"taId"`
	BuId                 string                  `protobuf:"bytes,2,opt,name=bu_id,json=buId,proto3" json:"buId"`
	FrontId              int32                   `protobuf:"varint,3,opt,name=front_id,json=frontId,proto3" json:"frontId"`
	SessionId            int32                   `protobuf:"varint,4,opt,name=session_id,json=sessionId,proto3" json:"sessionId"`
	OrderRef             string                  `protobuf:"bytes,5,opt,name=order_ref,json=orderRef,proto3" json:"orderRef"`
	Exchange             string                  `protobuf:"bytes,6,opt,name=exchange,proto3" json:"exchange"`
	Symbol               string                  `protobuf:"bytes,7,opt,name=symbol,proto3" json:"symbol"`
	Product              string                  `protobuf:"bytes,8,opt,name=product,proto3" json:"product"`
	Direction            int32                   `protobuf:"varint,9,opt,name=direction,proto3" json:"direction"`
	Offset               int32                   `protobuf:"varint,10,opt,name=offset,proto3" json:"offset"`
	Price                float64                 `protobuf:"fixed64,11,opt,name=price,proto3" json:"price"`
	Volume               int32                   `protobuf:"varint,12,opt,name=volume,proto3" json:"volume"`
	PriceTick            float64                 `protobuf:"fixed64,13,opt,name=price_tick,json=priceTick,proto3" json:"priceTick"`
	Multiple             int32                   `protobuf:"varint,14,opt,name=multiple,proto3" json:"multiple"`
	TradingDay           int32                   `protobuf:"varint,15,opt,name=trading_day,json=tradingDay,proto3" json:"tradingDay"`
	Status               int32                   `protobuf:"varint,16,opt,name=status,proto3" json:"status"`
	DemoOrderId          string                  `protobuf:"bytes,17,opt,name=demo_order_id,json=demoOrderId,proto3" json:"demoOrderId"`
	UpperLimit           float64                 `protobuf:"fixed64,18,opt,name=upper_limit,json=upperLimit,proto3" json:"upperLimit"`
	LowerLimit           float64                 `protobuf:"fixed64,19,opt,name=lower_limit,json=lowerLimit,proto3" json:"lowerLimit"`
	MinLimitOrderVolume  int32                   `protobuf:"varint,20,opt,name=min_limit_order_volume,json=minLimitOrderVolume,proto3" json:"minLimitOrderVolume"`
	MinMarketOrderVolume int32                   `protobuf:"varint,21,opt,name=min_market_order_volume,json=minMarketOrderVolume,proto3" json:"minMarketOrderVolume"`
	VolumeTraded         int32                   `protobuf:"varint,22,opt,name=volume_traded,json=volumeTraded,proto3" json:"volumeTraded"`
	VolumeCanceled       int32                   `protobuf:"varint,23,opt,name=volume_canceled,json=volumeCanceled,proto3" json:"volumeCanceled"`
	SendTradingDay       int32                   `protobuf:"varint,24,opt,name=send_trading_day,json=sendTradingDay,proto3" json:"sendTradingDay"`
	ProductType          int32                   `protobuf:"varint,25,opt,name=product_type,json=productType,proto3" json:"productType"`
	OrderSourceType      int32                   `protobuf:"varint,26,opt,name=order_source_type,json=orderSourceType,proto3" json:"orderSourceType"`
	PriceType            int32                   `protobuf:"varint,27,opt,name=price_type,json=priceType,proto3" json:"priceType"`
	SendTime             int64                   `protobuf:"varint,28,opt,name=send_time,json=sendTime,proto3" json:"sendTime"`
	TimeRule             []*goshare.MarketStatus `protobuf:"bytes,29,rep,name=time_rule,json=timeRule,proto3" json:"timeRule"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *DemoOrder) Reset()         { *m = DemoOrder{} }
func (m *DemoOrder) String() string { return proto.CompactTextString(m) }
func (*DemoOrder) ProtoMessage()    {}
func (*DemoOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_983becccff96becb, []int{0}
}

func (m *DemoOrder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoOrder.Unmarshal(m, b)
}
func (m *DemoOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoOrder.Marshal(b, m, deterministic)
}
func (m *DemoOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoOrder.Merge(m, src)
}
func (m *DemoOrder) XXX_Size() int {
	return xxx_messageInfo_DemoOrder.Size(m)
}
func (m *DemoOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoOrder.DiscardUnknown(m)
}

var xxx_messageInfo_DemoOrder proto.InternalMessageInfo

func (m *DemoOrder) GetTaId() string {
	if m != nil {
		return m.TaId
	}
	return ""
}

func (m *DemoOrder) GetBuId() string {
	if m != nil {
		return m.BuId
	}
	return ""
}

func (m *DemoOrder) GetFrontId() int32 {
	if m != nil {
		return m.FrontId
	}
	return 0
}

func (m *DemoOrder) GetSessionId() int32 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *DemoOrder) GetOrderRef() string {
	if m != nil {
		return m.OrderRef
	}
	return ""
}

func (m *DemoOrder) GetExchange() string {
	if m != nil {
		return m.Exchange
	}
	return ""
}

func (m *DemoOrder) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *DemoOrder) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *DemoOrder) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

func (m *DemoOrder) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *DemoOrder) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *DemoOrder) GetVolume() int32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *DemoOrder) GetPriceTick() float64 {
	if m != nil {
		return m.PriceTick
	}
	return 0
}

func (m *DemoOrder) GetMultiple() int32 {
	if m != nil {
		return m.Multiple
	}
	return 0
}

func (m *DemoOrder) GetTradingDay() int32 {
	if m != nil {
		return m.TradingDay
	}
	return 0
}

func (m *DemoOrder) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *DemoOrder) GetDemoOrderId() string {
	if m != nil {
		return m.DemoOrderId
	}
	return ""
}

func (m *DemoOrder) GetUpperLimit() float64 {
	if m != nil {
		return m.UpperLimit
	}
	return 0
}

func (m *DemoOrder) GetLowerLimit() float64 {
	if m != nil {
		return m.LowerLimit
	}
	return 0
}

func (m *DemoOrder) GetMinLimitOrderVolume() int32 {
	if m != nil {
		return m.MinLimitOrderVolume
	}
	return 0
}

func (m *DemoOrder) GetMinMarketOrderVolume() int32 {
	if m != nil {
		return m.MinMarketOrderVolume
	}
	return 0
}

func (m *DemoOrder) GetVolumeTraded() int32 {
	if m != nil {
		return m.VolumeTraded
	}
	return 0
}

func (m *DemoOrder) GetVolumeCanceled() int32 {
	if m != nil {
		return m.VolumeCanceled
	}
	return 0
}

func (m *DemoOrder) GetSendTradingDay() int32 {
	if m != nil {
		return m.SendTradingDay
	}
	return 0
}

func (m *DemoOrder) GetProductType() int32 {
	if m != nil {
		return m.ProductType
	}
	return 0
}

func (m *DemoOrder) GetOrderSourceType() int32 {
	if m != nil {
		return m.OrderSourceType
	}
	return 0
}

func (m *DemoOrder) GetPriceType() int32 {
	if m != nil {
		return m.PriceType
	}
	return 0
}

func (m *DemoOrder) GetSendTime() int64 {
	if m != nil {
		return m.SendTime
	}
	return 0
}

func (m *DemoOrder) GetTimeRule() []*goshare.MarketStatus {
	if m != nil {
		return m.TimeRule
	}
	return nil
}

func init() {
	proto.RegisterType((*DemoOrder)(nil), "goshare.hub.DemoOrder")
}

func init() { proto.RegisterFile("hub/demo.proto", fileDescriptor_983becccff96becb) }

var fileDescriptor_983becccff96becb = []byte{
	// 603 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x93, 0xcb, 0x6e, 0xdb, 0x3a,
	0x10, 0x86, 0xe1, 0x93, 0x38, 0xb6, 0xc6, 0xb9, 0x32, 0x37, 0xe6, 0x86, 0xe3, 0x93, 0xb3, 0x88,
	0xdb, 0x85, 0x05, 0x24, 0xe8, 0x0b, 0xb4, 0xd9, 0x18, 0x68, 0x51, 0x40, 0x31, 0xba, 0xe8, 0x46,
	0x90, 0xc4, 0xb1, 0x4d, 0x44, 0x14, 0x05, 0x5e, 0xda, 0xfa, 0xd9, 0xfa, 0x72, 0x05, 0x87, 0x72,
	0x12, 0x74, 0xf9, 0x7f, 0xff, 0x3f, 0xe2, 0xcc, 0x90, 0x82, 0xfd, 0x95, 0x2f, 0x53, 0x81, 0x4a,
	0x4f, 0x5b, 0xa3, 0x9d, 0x66, 0xa3, 0xa5, 0xb6, 0xab, 0xc2, 0xe0, 0x74, 0xe5, 0xcb, 0x4b, 0xde,
	0x89, 0x54, 0x36, 0xd6, 0x19, 0xaf, 0xb0, 0x71, 0x31, 0x76, 0xfb, 0x7b, 0x00, 0xc9, 0x23, 0x2a,
	0xfd, 0xd5, 0x08, 0x34, 0xec, 0x18, 0xfa, 0xae, 0xc8, 0xa5, 0xe0, 0xbd, 0x71, 0x6f, 0x92, 0x64,
	0xdb, 0xae, 0x98, 0x89, 0x00, 0x4b, 0x1f, 0xe0, 0x3f, 0x11, 0x96, 0x7e, 0x26, 0xd8, 0x05, 0x0c,
	0x17, 0x46, 0x37, 0x2e, 0xf0, 0xad, 0x71, 0x6f, 0xd2, 0xcf, 0x06, 0xa4, 0x67, 0x82, 0xdd, 0x00,
	0x58, 0xb4, 0x56, 0xea, 0x26, 0x98, 0xdb, 0x64, 0x26, 0x1d, 0x99, 0x09, 0x76, 0x05, 0x89, 0x0e,
	0x87, 0xe5, 0x06, 0x17, 0xbc, 0x4f, 0x9f, 0x1c, 0x12, 0xc8, 0x70, 0xc1, 0x2e, 0x61, 0x88, 0xbf,
	0xaa, 0x55, 0xd1, 0x2c, 0x91, 0xef, 0x44, 0x6f, 0xa3, 0xd9, 0x19, 0xec, 0xd8, 0xb5, 0x2a, 0x75,
	0xcd, 0x07, 0xe4, 0x74, 0x8a, 0x71, 0x18, 0xb4, 0x46, 0x0b, 0x5f, 0x39, 0x3e, 0x24, 0x63, 0x23,
	0xd9, 0x35, 0x24, 0x42, 0x1a, 0xac, 0x9c, 0xd4, 0x0d, 0x4f, 0x62, 0x23, 0x2f, 0x20, 0x7c, 0x4f,
	0x2f, 0x16, 0x16, 0x1d, 0x07, 0xb2, 0x3a, 0xc5, 0x4e, 0xa0, 0xdf, 0x1a, 0x59, 0x21, 0x1f, 0x8d,
	0x7b, 0x93, 0x5e, 0x16, 0x45, 0x48, 0xff, 0xd0, 0xb5, 0x57, 0xc8, 0x77, 0x63, 0x3a, 0xaa, 0x30,
	0x2d, 0x05, 0x72, 0x27, 0xab, 0x67, 0xbe, 0x47, 0x25, 0x09, 0x91, 0xb9, 0xac, 0x9e, 0xc3, 0x40,
	0xca, 0xd7, 0x4e, 0xb6, 0x35, 0xf2, 0x7d, 0x2a, 0x7c, 0xd1, 0xec, 0x5f, 0x18, 0x39, 0x53, 0x08,
	0xd9, 0x2c, 0x73, 0x51, 0xac, 0xf9, 0x01, 0xd9, 0xd0, 0xa1, 0xc7, 0x62, 0x4d, 0x13, 0xbb, 0xc2,
	0x79, 0xcb, 0x0f, 0xe3, 0x99, 0x51, 0xb1, 0x5b, 0xd8, 0x0b, 0x37, 0x9d, 0xc7, 0x3d, 0x4a, 0xc1,
	0x8f, 0x68, 0xee, 0x91, 0xd8, 0x5c, 0xe4, 0x4c, 0x84, 0x8f, 0xfb, 0xb6, 0x45, 0x93, 0xd7, 0x52,
	0x49, 0xc7, 0x19, 0x35, 0x06, 0x84, 0x3e, 0x07, 0x12, 0x02, 0xb5, 0xfe, 0xf9, 0x12, 0x38, 0x8e,
	0x01, 0x42, 0x31, 0xf0, 0x00, 0x67, 0x4a, 0x36, 0xd1, 0xee, 0x8e, 0xea, 0x36, 0x70, 0x42, 0xdd,
	0x1c, 0x2b, 0xd9, 0x50, 0x92, 0x8e, 0xfc, 0x16, 0xd7, 0xf1, 0x01, 0xce, 0x43, 0x91, 0x2a, 0xcc,
	0x33, 0xfe, 0x55, 0x75, 0x4a, 0x55, 0x27, 0x4a, 0x36, 0x5f, 0xc8, 0x7d, 0x5b, 0xf6, 0x3f, 0xec,
	0xc5, 0x54, 0x1e, 0xc6, 0x47, 0xc1, 0xcf, 0x28, 0xbc, 0x1b, 0xe1, 0x9c, 0x18, 0xbb, 0x83, 0x83,
	0x2e, 0x54, 0x15, 0x4d, 0x85, 0x35, 0x0a, 0x7e, 0x4e, 0xb1, 0xfd, 0x88, 0x3f, 0x75, 0x94, 0x4d,
	0xe0, 0xd0, 0x62, 0x23, 0xf2, 0xb7, 0xdb, 0xe5, 0x31, 0x19, 0xf8, 0xfc, 0x75, 0xc3, 0xff, 0xc1,
	0x6e, 0xf7, 0x58, 0x72, 0xb7, 0x6e, 0x91, 0x5f, 0x50, 0x6a, 0xd4, 0xb1, 0xf9, 0xba, 0x45, 0xf6,
	0x1e, 0x8e, 0xe2, 0x18, 0x56, 0x7b, 0x13, 0xee, 0x39, 0xe4, 0x2e, 0x29, 0x77, 0x40, 0xc6, 0x13,
	0x71, 0xca, 0xbe, 0x3e, 0x86, 0x10, 0xba, 0x8a, 0x2f, 0x2e, 0x3e, 0x86, 0x60, 0x5f, 0x41, 0x12,
	0xfb, 0x92, 0x0a, 0xf9, 0xf5, 0xb8, 0x37, 0xd9, 0xca, 0x86, 0xd4, 0x90, 0x54, 0xc8, 0xee, 0x21,
	0x09, 0x3c, 0x37, 0xbe, 0x46, 0x7e, 0x33, 0xde, 0x9a, 0x8c, 0xee, 0x4f, 0xa7, 0x9b, 0x9f, 0x38,
	0x6e, 0xec, 0x89, 0xae, 0x3f, 0x1b, 0x86, 0x5c, 0xe6, 0x6b, 0xfc, 0xf8, 0xee, 0xfb, 0xdd, 0x52,
	0xba, 0x95, 0x2f, 0xa7, 0x95, 0x56, 0xa9, 0x92, 0x0d, 0x9a, 0xa2, 0x36, 0x68, 0x53, 0xfa, 0xb9,
	0x6d, 0x6a, 0x4d, 0x95, 0x2e, 0x75, 0xba, 0xf2, 0x65, 0xb9, 0x43, 0xe8, 0xe1, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x2f, 0x1f, 0x0e, 0x3a, 0x28, 0x04, 0x00, 0x00,
}
