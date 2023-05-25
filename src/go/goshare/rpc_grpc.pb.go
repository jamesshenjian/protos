// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: goshare/rpc.proto

package goshare

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	KlineCache_SaveKlineSeries_FullMethodName   = "/goshare.KlineCache/SaveKlineSeries"
	KlineCache_GetKlineSeries_FullMethodName    = "/goshare.KlineCache/GetKlineSeries"
	KlineCache_SaveInstrument_FullMethodName    = "/goshare.KlineCache/SaveInstrument"
	KlineCache_GetInstrument_FullMethodName     = "/goshare.KlineCache/GetInstrument"
	KlineCache_GetInstrumentList_FullMethodName = "/goshare.KlineCache/GetInstrumentList"
	KlineCache_SaveBonusHistory_FullMethodName  = "/goshare.KlineCache/SaveBonusHistory"
	KlineCache_GetBonusHistory_FullMethodName   = "/goshare.KlineCache/GetBonusHistory"
	KlineCache_BatchDelete_FullMethodName       = "/goshare.KlineCache/BatchDelete"
	KlineCache_BatchGet_FullMethodName          = "/goshare.KlineCache/BatchGet"
	KlineCache_BatchPut_FullMethodName          = "/goshare.KlineCache/BatchPut"
	KlineCache_Delete_FullMethodName            = "/goshare.KlineCache/Delete"
	KlineCache_Get_FullMethodName               = "/goshare.KlineCache/Get"
	KlineCache_Put_FullMethodName               = "/goshare.KlineCache/Put"
	KlineCache_ReverseScan_FullMethodName       = "/goshare.KlineCache/ReverseScan"
	KlineCache_Scan_FullMethodName              = "/goshare.KlineCache/Scan"
)

// KlineCacheClient is the client API for KlineCache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KlineCacheClient interface {
	// 保存K线
	SaveKlineSeries(ctx context.Context, in *ReqSaveKlineSeries, opts ...grpc.CallOption) (*RspSaveKlineSeries, error)
	// 读取K线
	GetKlineSeries(ctx context.Context, in *ReqGetKlineSeries, opts ...grpc.CallOption) (*RspGetKlineSeries, error)
	// 保存合约
	SaveInstrument(ctx context.Context, in *ReqSaveInstrument, opts ...grpc.CallOption) (*RspSaveInstrument, error)
	// 读取合约
	GetInstrument(ctx context.Context, in *ReqGetInstrument, opts ...grpc.CallOption) (*RspGetInstrument, error)
	// 批量读取
	GetInstrumentList(ctx context.Context, in *ReqGetInstrumentList, opts ...grpc.CallOption) (*RspGetInstrumentList, error)
	// 保存分红历史
	SaveBonusHistory(ctx context.Context, in *ReqSaveBonusHistory, opts ...grpc.CallOption) (*RspSaveBonusHistory, error)
	// 读取分红历史
	GetBonusHistory(ctx context.Context, in *ReqGetBonusHistory, opts ...grpc.CallOption) (*RspGetBonusHistory, error)
	// 批量删除
	BatchDelete(ctx context.Context, in *ReqBatchDelete, opts ...grpc.CallOption) (*RspBatchDelete, error)
	// 批量获取
	BatchGet(ctx context.Context, in *ReqBatchGet, opts ...grpc.CallOption) (*RspBatchGet, error)
	// 批量保存
	BatchPut(ctx context.Context, in *ReqBatchPut, opts ...grpc.CallOption) (*RspBatchPut, error)
	// 删除
	Delete(ctx context.Context, in *ReqDelete, opts ...grpc.CallOption) (*RspDelete, error)
	// 取得
	Get(ctx context.Context, in *ReqGet, opts ...grpc.CallOption) (*RspGet, error)
	// 保存
	Put(ctx context.Context, in *ReqPut, opts ...grpc.CallOption) (*RspPut, error)
	// 反向扫描
	ReverseScan(ctx context.Context, in *ReqReverseScan, opts ...grpc.CallOption) (*RspReverseScan, error)
	// 扫描
	Scan(ctx context.Context, in *ReqScan, opts ...grpc.CallOption) (*RspScan, error)
}

type klineCacheClient struct {
	cc grpc.ClientConnInterface
}

func NewKlineCacheClient(cc grpc.ClientConnInterface) KlineCacheClient {
	return &klineCacheClient{cc}
}

func (c *klineCacheClient) SaveKlineSeries(ctx context.Context, in *ReqSaveKlineSeries, opts ...grpc.CallOption) (*RspSaveKlineSeries, error) {
	out := new(RspSaveKlineSeries)
	err := c.cc.Invoke(ctx, KlineCache_SaveKlineSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) GetKlineSeries(ctx context.Context, in *ReqGetKlineSeries, opts ...grpc.CallOption) (*RspGetKlineSeries, error) {
	out := new(RspGetKlineSeries)
	err := c.cc.Invoke(ctx, KlineCache_GetKlineSeries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) SaveInstrument(ctx context.Context, in *ReqSaveInstrument, opts ...grpc.CallOption) (*RspSaveInstrument, error) {
	out := new(RspSaveInstrument)
	err := c.cc.Invoke(ctx, KlineCache_SaveInstrument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) GetInstrument(ctx context.Context, in *ReqGetInstrument, opts ...grpc.CallOption) (*RspGetInstrument, error) {
	out := new(RspGetInstrument)
	err := c.cc.Invoke(ctx, KlineCache_GetInstrument_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) GetInstrumentList(ctx context.Context, in *ReqGetInstrumentList, opts ...grpc.CallOption) (*RspGetInstrumentList, error) {
	out := new(RspGetInstrumentList)
	err := c.cc.Invoke(ctx, KlineCache_GetInstrumentList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) SaveBonusHistory(ctx context.Context, in *ReqSaveBonusHistory, opts ...grpc.CallOption) (*RspSaveBonusHistory, error) {
	out := new(RspSaveBonusHistory)
	err := c.cc.Invoke(ctx, KlineCache_SaveBonusHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) GetBonusHistory(ctx context.Context, in *ReqGetBonusHistory, opts ...grpc.CallOption) (*RspGetBonusHistory, error) {
	out := new(RspGetBonusHistory)
	err := c.cc.Invoke(ctx, KlineCache_GetBonusHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) BatchDelete(ctx context.Context, in *ReqBatchDelete, opts ...grpc.CallOption) (*RspBatchDelete, error) {
	out := new(RspBatchDelete)
	err := c.cc.Invoke(ctx, KlineCache_BatchDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) BatchGet(ctx context.Context, in *ReqBatchGet, opts ...grpc.CallOption) (*RspBatchGet, error) {
	out := new(RspBatchGet)
	err := c.cc.Invoke(ctx, KlineCache_BatchGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) BatchPut(ctx context.Context, in *ReqBatchPut, opts ...grpc.CallOption) (*RspBatchPut, error) {
	out := new(RspBatchPut)
	err := c.cc.Invoke(ctx, KlineCache_BatchPut_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) Delete(ctx context.Context, in *ReqDelete, opts ...grpc.CallOption) (*RspDelete, error) {
	out := new(RspDelete)
	err := c.cc.Invoke(ctx, KlineCache_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) Get(ctx context.Context, in *ReqGet, opts ...grpc.CallOption) (*RspGet, error) {
	out := new(RspGet)
	err := c.cc.Invoke(ctx, KlineCache_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) Put(ctx context.Context, in *ReqPut, opts ...grpc.CallOption) (*RspPut, error) {
	out := new(RspPut)
	err := c.cc.Invoke(ctx, KlineCache_Put_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) ReverseScan(ctx context.Context, in *ReqReverseScan, opts ...grpc.CallOption) (*RspReverseScan, error) {
	out := new(RspReverseScan)
	err := c.cc.Invoke(ctx, KlineCache_ReverseScan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *klineCacheClient) Scan(ctx context.Context, in *ReqScan, opts ...grpc.CallOption) (*RspScan, error) {
	out := new(RspScan)
	err := c.cc.Invoke(ctx, KlineCache_Scan_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KlineCacheServer is the server API for KlineCache service.
// All implementations must embed UnimplementedKlineCacheServer
// for forward compatibility
type KlineCacheServer interface {
	// 保存K线
	SaveKlineSeries(context.Context, *ReqSaveKlineSeries) (*RspSaveKlineSeries, error)
	// 读取K线
	GetKlineSeries(context.Context, *ReqGetKlineSeries) (*RspGetKlineSeries, error)
	// 保存合约
	SaveInstrument(context.Context, *ReqSaveInstrument) (*RspSaveInstrument, error)
	// 读取合约
	GetInstrument(context.Context, *ReqGetInstrument) (*RspGetInstrument, error)
	// 批量读取
	GetInstrumentList(context.Context, *ReqGetInstrumentList) (*RspGetInstrumentList, error)
	// 保存分红历史
	SaveBonusHistory(context.Context, *ReqSaveBonusHistory) (*RspSaveBonusHistory, error)
	// 读取分红历史
	GetBonusHistory(context.Context, *ReqGetBonusHistory) (*RspGetBonusHistory, error)
	// 批量删除
	BatchDelete(context.Context, *ReqBatchDelete) (*RspBatchDelete, error)
	// 批量获取
	BatchGet(context.Context, *ReqBatchGet) (*RspBatchGet, error)
	// 批量保存
	BatchPut(context.Context, *ReqBatchPut) (*RspBatchPut, error)
	// 删除
	Delete(context.Context, *ReqDelete) (*RspDelete, error)
	// 取得
	Get(context.Context, *ReqGet) (*RspGet, error)
	// 保存
	Put(context.Context, *ReqPut) (*RspPut, error)
	// 反向扫描
	ReverseScan(context.Context, *ReqReverseScan) (*RspReverseScan, error)
	// 扫描
	Scan(context.Context, *ReqScan) (*RspScan, error)
	mustEmbedUnimplementedKlineCacheServer()
}

// UnimplementedKlineCacheServer must be embedded to have forward compatible implementations.
type UnimplementedKlineCacheServer struct {
}

func (UnimplementedKlineCacheServer) SaveKlineSeries(context.Context, *ReqSaveKlineSeries) (*RspSaveKlineSeries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveKlineSeries not implemented")
}
func (UnimplementedKlineCacheServer) GetKlineSeries(context.Context, *ReqGetKlineSeries) (*RspGetKlineSeries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKlineSeries not implemented")
}
func (UnimplementedKlineCacheServer) SaveInstrument(context.Context, *ReqSaveInstrument) (*RspSaveInstrument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveInstrument not implemented")
}
func (UnimplementedKlineCacheServer) GetInstrument(context.Context, *ReqGetInstrument) (*RspGetInstrument, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrument not implemented")
}
func (UnimplementedKlineCacheServer) GetInstrumentList(context.Context, *ReqGetInstrumentList) (*RspGetInstrumentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInstrumentList not implemented")
}
func (UnimplementedKlineCacheServer) SaveBonusHistory(context.Context, *ReqSaveBonusHistory) (*RspSaveBonusHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveBonusHistory not implemented")
}
func (UnimplementedKlineCacheServer) GetBonusHistory(context.Context, *ReqGetBonusHistory) (*RspGetBonusHistory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBonusHistory not implemented")
}
func (UnimplementedKlineCacheServer) BatchDelete(context.Context, *ReqBatchDelete) (*RspBatchDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}
func (UnimplementedKlineCacheServer) BatchGet(context.Context, *ReqBatchGet) (*RspBatchGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedKlineCacheServer) BatchPut(context.Context, *ReqBatchPut) (*RspBatchPut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPut not implemented")
}
func (UnimplementedKlineCacheServer) Delete(context.Context, *ReqDelete) (*RspDelete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKlineCacheServer) Get(context.Context, *ReqGet) (*RspGet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKlineCacheServer) Put(context.Context, *ReqPut) (*RspPut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedKlineCacheServer) ReverseScan(context.Context, *ReqReverseScan) (*RspReverseScan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReverseScan not implemented")
}
func (UnimplementedKlineCacheServer) Scan(context.Context, *ReqScan) (*RspScan, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (UnimplementedKlineCacheServer) mustEmbedUnimplementedKlineCacheServer() {}

// UnsafeKlineCacheServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KlineCacheServer will
// result in compilation errors.
type UnsafeKlineCacheServer interface {
	mustEmbedUnimplementedKlineCacheServer()
}

func RegisterKlineCacheServer(s grpc.ServiceRegistrar, srv KlineCacheServer) {
	s.RegisterService(&KlineCache_ServiceDesc, srv)
}

func _KlineCache_SaveKlineSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveKlineSeries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).SaveKlineSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_SaveKlineSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).SaveKlineSeries(ctx, req.(*ReqSaveKlineSeries))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_GetKlineSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetKlineSeries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).GetKlineSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_GetKlineSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).GetKlineSeries(ctx, req.(*ReqGetKlineSeries))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_SaveInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveInstrument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).SaveInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_SaveInstrument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).SaveInstrument(ctx, req.(*ReqSaveInstrument))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_GetInstrument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetInstrument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).GetInstrument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_GetInstrument_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).GetInstrument(ctx, req.(*ReqGetInstrument))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_GetInstrumentList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetInstrumentList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).GetInstrumentList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_GetInstrumentList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).GetInstrumentList(ctx, req.(*ReqGetInstrumentList))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_SaveBonusHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqSaveBonusHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).SaveBonusHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_SaveBonusHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).SaveBonusHistory(ctx, req.(*ReqSaveBonusHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_GetBonusHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGetBonusHistory)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).GetBonusHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_GetBonusHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).GetBonusHistory(ctx, req.(*ReqGetBonusHistory))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBatchDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_BatchDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).BatchDelete(ctx, req.(*ReqBatchDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBatchGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_BatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).BatchGet(ctx, req.(*ReqBatchGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_BatchPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqBatchPut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).BatchPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_BatchPut_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).BatchPut(ctx, req.(*ReqBatchPut))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).Delete(ctx, req.(*ReqDelete))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqGet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).Get(ctx, req.(*ReqGet))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqPut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_Put_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).Put(ctx, req.(*ReqPut))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_ReverseScan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqReverseScan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).ReverseScan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_ReverseScan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).ReverseScan(ctx, req.(*ReqReverseScan))
	}
	return interceptor(ctx, in, info, handler)
}

func _KlineCache_Scan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqScan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KlineCacheServer).Scan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KlineCache_Scan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KlineCacheServer).Scan(ctx, req.(*ReqScan))
	}
	return interceptor(ctx, in, info, handler)
}

// KlineCache_ServiceDesc is the grpc.ServiceDesc for KlineCache service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KlineCache_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goshare.KlineCache",
	HandlerType: (*KlineCacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveKlineSeries",
			Handler:    _KlineCache_SaveKlineSeries_Handler,
		},
		{
			MethodName: "GetKlineSeries",
			Handler:    _KlineCache_GetKlineSeries_Handler,
		},
		{
			MethodName: "SaveInstrument",
			Handler:    _KlineCache_SaveInstrument_Handler,
		},
		{
			MethodName: "GetInstrument",
			Handler:    _KlineCache_GetInstrument_Handler,
		},
		{
			MethodName: "GetInstrumentList",
			Handler:    _KlineCache_GetInstrumentList_Handler,
		},
		{
			MethodName: "SaveBonusHistory",
			Handler:    _KlineCache_SaveBonusHistory_Handler,
		},
		{
			MethodName: "GetBonusHistory",
			Handler:    _KlineCache_GetBonusHistory_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _KlineCache_BatchDelete_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _KlineCache_BatchGet_Handler,
		},
		{
			MethodName: "BatchPut",
			Handler:    _KlineCache_BatchPut_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KlineCache_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KlineCache_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KlineCache_Put_Handler,
		},
		{
			MethodName: "ReverseScan",
			Handler:    _KlineCache_ReverseScan_Handler,
		},
		{
			MethodName: "Scan",
			Handler:    _KlineCache_Scan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goshare/rpc.proto",
}
