// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server_client.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//角色位置
type RolePos struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolePos) Reset()         { *m = RolePos{} }
func (m *RolePos) String() string { return proto.CompactTextString(m) }
func (*RolePos) ProtoMessage()    {}
func (*RolePos) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{0}
}

func (m *RolePos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolePos.Unmarshal(m, b)
}
func (m *RolePos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolePos.Marshal(b, m, deterministic)
}
func (m *RolePos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolePos.Merge(m, src)
}
func (m *RolePos) XXX_Size() int {
	return xxx_messageInfo_RolePos.Size(m)
}
func (m *RolePos) XXX_DiscardUnknown() {
	xxx_messageInfo_RolePos.DiscardUnknown(m)
}

var xxx_messageInfo_RolePos proto.InternalMessageInfo

func (m *RolePos) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RolePos) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *RolePos) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

//角色穿着
type RoleOutlook struct {
	Weapon               uint32   `protobuf:"varint,1,opt,name=weapon,proto3" json:"weapon,omitempty"`
	Helm                 uint32   `protobuf:"varint,2,opt,name=helm,proto3" json:"helm,omitempty"`
	Face                 uint32   `protobuf:"varint,3,opt,name=face,proto3" json:"face,omitempty"`
	Wing                 uint32   `protobuf:"varint,4,opt,name=wing,proto3" json:"wing,omitempty"`
	Bag                  uint32   `protobuf:"varint,5,opt,name=bag,proto3" json:"bag,omitempty"`
	Suit                 uint32   `protobuf:"varint,6,opt,name=suit,proto3" json:"suit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleOutlook) Reset()         { *m = RoleOutlook{} }
func (m *RoleOutlook) String() string { return proto.CompactTextString(m) }
func (*RoleOutlook) ProtoMessage()    {}
func (*RoleOutlook) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{1}
}

func (m *RoleOutlook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleOutlook.Unmarshal(m, b)
}
func (m *RoleOutlook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleOutlook.Marshal(b, m, deterministic)
}
func (m *RoleOutlook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleOutlook.Merge(m, src)
}
func (m *RoleOutlook) XXX_Size() int {
	return xxx_messageInfo_RoleOutlook.Size(m)
}
func (m *RoleOutlook) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleOutlook.DiscardUnknown(m)
}

var xxx_messageInfo_RoleOutlook proto.InternalMessageInfo

func (m *RoleOutlook) GetWeapon() uint32 {
	if m != nil {
		return m.Weapon
	}
	return 0
}

func (m *RoleOutlook) GetHelm() uint32 {
	if m != nil {
		return m.Helm
	}
	return 0
}

func (m *RoleOutlook) GetFace() uint32 {
	if m != nil {
		return m.Face
	}
	return 0
}

func (m *RoleOutlook) GetWing() uint32 {
	if m != nil {
		return m.Wing
	}
	return 0
}

func (m *RoleOutlook) GetBag() uint32 {
	if m != nil {
		return m.Bag
	}
	return 0
}

func (m *RoleOutlook) GetSuit() uint32 {
	if m != nil {
		return m.Suit
	}
	return 0
}

//角色基本信息
type RoleBaseInfo struct {
	Name                 string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level                uint32       `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	LifePoint            uint32       `protobuf:"varint,3,opt,name=life_point,json=lifePoint,proto3" json:"life_point,omitempty"`
	Pos                  *RolePos     `protobuf:"bytes,5,opt,name=pos,proto3" json:"pos,omitempty"`
	RoleAngle            uint32       `protobuf:"varint,6,opt,name=role_angle,json=roleAngle,proto3" json:"role_angle,omitempty"`
	Outlook              *RoleOutlook `protobuf:"bytes,7,opt,name=outlook,proto3" json:"outlook,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RoleBaseInfo) Reset()         { *m = RoleBaseInfo{} }
func (m *RoleBaseInfo) String() string { return proto.CompactTextString(m) }
func (*RoleBaseInfo) ProtoMessage()    {}
func (*RoleBaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{2}
}

func (m *RoleBaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleBaseInfo.Unmarshal(m, b)
}
func (m *RoleBaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleBaseInfo.Marshal(b, m, deterministic)
}
func (m *RoleBaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleBaseInfo.Merge(m, src)
}
func (m *RoleBaseInfo) XXX_Size() int {
	return xxx_messageInfo_RoleBaseInfo.Size(m)
}
func (m *RoleBaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleBaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleBaseInfo proto.InternalMessageInfo

func (m *RoleBaseInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleBaseInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *RoleBaseInfo) GetLifePoint() uint32 {
	if m != nil {
		return m.LifePoint
	}
	return 0
}

func (m *RoleBaseInfo) GetPos() *RolePos {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *RoleBaseInfo) GetRoleAngle() uint32 {
	if m != nil {
		return m.RoleAngle
	}
	return 0
}

func (m *RoleBaseInfo) GetOutlook() *RoleOutlook {
	if m != nil {
		return m.Outlook
	}
	return nil
}

//块信息
type BlockInfo struct {
	BlockType            uint32   `protobuf:"varint,1,opt,name=BlockType,proto3" json:"BlockType,omitempty"`
	Content              uint32   `protobuf:"varint,2,opt,name=Content,proto3" json:"Content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{3}
}

func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockInfo.Unmarshal(m, b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
}
func (m *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(m, src)
}
func (m *BlockInfo) XXX_Size() int {
	return xxx_messageInfo_BlockInfo.Size(m)
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetBlockType() uint32 {
	if m != nil {
		return m.BlockType
	}
	return 0
}

func (m *BlockInfo) GetContent() uint32 {
	if m != nil {
		return m.Content
	}
	return 0
}

//客户端请求让角色进入服务器地图
type ReqEnterGs struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEnterGs) Reset()         { *m = ReqEnterGs{} }
func (m *ReqEnterGs) String() string { return proto.CompactTextString(m) }
func (*ReqEnterGs) ProtoMessage()    {}
func (*ReqEnterGs) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{4}
}

func (m *ReqEnterGs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEnterGs.Unmarshal(m, b)
}
func (m *ReqEnterGs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEnterGs.Marshal(b, m, deterministic)
}
func (m *ReqEnterGs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEnterGs.Merge(m, src)
}
func (m *ReqEnterGs) XXX_Size() int {
	return xxx_messageInfo_ReqEnterGs.Size(m)
}
func (m *ReqEnterGs) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEnterGs.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEnterGs proto.InternalMessageInfo

func (m *ReqEnterGs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ReqEnterGs) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

//同步给客户端进入地图的信息
type RspEnterGs struct {
	RetCode              int32         `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	MainRoleIdx          int32         `protobuf:"varint,2,opt,name=mainRoleIdx,proto3" json:"mainRoleIdx,omitempty"`
	MainRole             *RoleBaseInfo `protobuf:"bytes,3,opt,name=mainRole,proto3" json:"mainRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RspEnterGs) Reset()         { *m = RspEnterGs{} }
func (m *RspEnterGs) String() string { return proto.CompactTextString(m) }
func (*RspEnterGs) ProtoMessage()    {}
func (*RspEnterGs) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{5}
}

func (m *RspEnterGs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspEnterGs.Unmarshal(m, b)
}
func (m *RspEnterGs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspEnterGs.Marshal(b, m, deterministic)
}
func (m *RspEnterGs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspEnterGs.Merge(m, src)
}
func (m *RspEnterGs) XXX_Size() int {
	return xxx_messageInfo_RspEnterGs.Size(m)
}
func (m *RspEnterGs) XXX_DiscardUnknown() {
	xxx_messageInfo_RspEnterGs.DiscardUnknown(m)
}

var xxx_messageInfo_RspEnterGs proto.InternalMessageInfo

func (m *RspEnterGs) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RspEnterGs) GetMainRoleIdx() int32 {
	if m != nil {
		return m.MainRoleIdx
	}
	return 0
}

func (m *RspEnterGs) GetMainRole() *RoleBaseInfo {
	if m != nil {
		return m.MainRole
	}
	return nil
}

//同步新的区块
type SyncChunkEnterRange struct {
	ChunkX               int32        `protobuf:"varint,1,opt,name=chunkX,proto3" json:"chunkX,omitempty"`
	ChunkZ               int32        `protobuf:"varint,2,opt,name=chunkZ,proto3" json:"chunkZ,omitempty"`
	Blocks               []*BlockInfo `protobuf:"bytes,3,rep,name=blocks,proto3" json:"blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SyncChunkEnterRange) Reset()         { *m = SyncChunkEnterRange{} }
func (m *SyncChunkEnterRange) String() string { return proto.CompactTextString(m) }
func (*SyncChunkEnterRange) ProtoMessage()    {}
func (*SyncChunkEnterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{6}
}

func (m *SyncChunkEnterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncChunkEnterRange.Unmarshal(m, b)
}
func (m *SyncChunkEnterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncChunkEnterRange.Marshal(b, m, deterministic)
}
func (m *SyncChunkEnterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncChunkEnterRange.Merge(m, src)
}
func (m *SyncChunkEnterRange) XXX_Size() int {
	return xxx_messageInfo_SyncChunkEnterRange.Size(m)
}
func (m *SyncChunkEnterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncChunkEnterRange.DiscardUnknown(m)
}

var xxx_messageInfo_SyncChunkEnterRange proto.InternalMessageInfo

func (m *SyncChunkEnterRange) GetChunkX() int32 {
	if m != nil {
		return m.ChunkX
	}
	return 0
}

func (m *SyncChunkEnterRange) GetChunkZ() int32 {
	if m != nil {
		return m.ChunkZ
	}
	return 0
}

func (m *SyncChunkEnterRange) GetBlocks() []*BlockInfo {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type SyncChunkLeaveRange struct {
	ChunkX               int32    `protobuf:"varint,1,opt,name=chunkX,proto3" json:"chunkX,omitempty"`
	ChunkZ               int32    `protobuf:"varint,2,opt,name=chunkZ,proto3" json:"chunkZ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncChunkLeaveRange) Reset()         { *m = SyncChunkLeaveRange{} }
func (m *SyncChunkLeaveRange) String() string { return proto.CompactTextString(m) }
func (*SyncChunkLeaveRange) ProtoMessage()    {}
func (*SyncChunkLeaveRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{7}
}

func (m *SyncChunkLeaveRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncChunkLeaveRange.Unmarshal(m, b)
}
func (m *SyncChunkLeaveRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncChunkLeaveRange.Marshal(b, m, deterministic)
}
func (m *SyncChunkLeaveRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncChunkLeaveRange.Merge(m, src)
}
func (m *SyncChunkLeaveRange) XXX_Size() int {
	return xxx_messageInfo_SyncChunkLeaveRange.Size(m)
}
func (m *SyncChunkLeaveRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncChunkLeaveRange.DiscardUnknown(m)
}

var xxx_messageInfo_SyncChunkLeaveRange proto.InternalMessageInfo

func (m *SyncChunkLeaveRange) GetChunkX() int32 {
	if m != nil {
		return m.ChunkX
	}
	return 0
}

func (m *SyncChunkLeaveRange) GetChunkZ() int32 {
	if m != nil {
		return m.ChunkZ
	}
	return 0
}

//同步某个角色进入同步范围
type SyncRoleEnterRange struct {
	RoleIdx              int32         `protobuf:"varint,1,opt,name=roleIdx,proto3" json:"roleIdx,omitempty"`
	RoleBaseInfo         *RoleBaseInfo `protobuf:"bytes,2,opt,name=roleBaseInfo,proto3" json:"roleBaseInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SyncRoleEnterRange) Reset()         { *m = SyncRoleEnterRange{} }
func (m *SyncRoleEnterRange) String() string { return proto.CompactTextString(m) }
func (*SyncRoleEnterRange) ProtoMessage()    {}
func (*SyncRoleEnterRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{8}
}

func (m *SyncRoleEnterRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRoleEnterRange.Unmarshal(m, b)
}
func (m *SyncRoleEnterRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRoleEnterRange.Marshal(b, m, deterministic)
}
func (m *SyncRoleEnterRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRoleEnterRange.Merge(m, src)
}
func (m *SyncRoleEnterRange) XXX_Size() int {
	return xxx_messageInfo_SyncRoleEnterRange.Size(m)
}
func (m *SyncRoleEnterRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRoleEnterRange.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRoleEnterRange proto.InternalMessageInfo

func (m *SyncRoleEnterRange) GetRoleIdx() int32 {
	if m != nil {
		return m.RoleIdx
	}
	return 0
}

func (m *SyncRoleEnterRange) GetRoleBaseInfo() *RoleBaseInfo {
	if m != nil {
		return m.RoleBaseInfo
	}
	return nil
}

//同步某个角色离开同步范围
type SyncRoleLeaveRange struct {
	RoleIdx              int32    `protobuf:"varint,1,opt,name=roleIdx,proto3" json:"roleIdx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRoleLeaveRange) Reset()         { *m = SyncRoleLeaveRange{} }
func (m *SyncRoleLeaveRange) String() string { return proto.CompactTextString(m) }
func (*SyncRoleLeaveRange) ProtoMessage()    {}
func (*SyncRoleLeaveRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{9}
}

func (m *SyncRoleLeaveRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRoleLeaveRange.Unmarshal(m, b)
}
func (m *SyncRoleLeaveRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRoleLeaveRange.Marshal(b, m, deterministic)
}
func (m *SyncRoleLeaveRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRoleLeaveRange.Merge(m, src)
}
func (m *SyncRoleLeaveRange) XXX_Size() int {
	return xxx_messageInfo_SyncRoleLeaveRange.Size(m)
}
func (m *SyncRoleLeaveRange) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRoleLeaveRange.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRoleLeaveRange proto.InternalMessageInfo

func (m *SyncRoleLeaveRange) GetRoleIdx() int32 {
	if m != nil {
		return m.RoleIdx
	}
	return 0
}

//客户端请求同步位置、方向、动作
type ReqRoleAction struct {
	Pos                  *RolePos `protobuf:"bytes,1,opt,name=pos,proto3" json:"pos,omitempty"`
	RoleAngle            uint32   `protobuf:"varint,2,opt,name=role_angle,json=roleAngle,proto3" json:"role_angle,omitempty"`
	Action               uint32   `protobuf:"varint,3,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRoleAction) Reset()         { *m = ReqRoleAction{} }
func (m *ReqRoleAction) String() string { return proto.CompactTextString(m) }
func (*ReqRoleAction) ProtoMessage()    {}
func (*ReqRoleAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{10}
}

func (m *ReqRoleAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRoleAction.Unmarshal(m, b)
}
func (m *ReqRoleAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRoleAction.Marshal(b, m, deterministic)
}
func (m *ReqRoleAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRoleAction.Merge(m, src)
}
func (m *ReqRoleAction) XXX_Size() int {
	return xxx_messageInfo_ReqRoleAction.Size(m)
}
func (m *ReqRoleAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRoleAction.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRoleAction proto.InternalMessageInfo

func (m *ReqRoleAction) GetPos() *RolePos {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *ReqRoleAction) GetRoleAngle() uint32 {
	if m != nil {
		return m.RoleAngle
	}
	return 0
}

func (m *ReqRoleAction) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

type SyncRoleAction struct {
	RoleIdx              int32    `protobuf:"varint,1,opt,name=roleIdx,proto3" json:"roleIdx,omitempty"`
	Pos                  *RolePos `protobuf:"bytes,2,opt,name=pos,proto3" json:"pos,omitempty"`
	RoleAngle            uint32   `protobuf:"varint,3,opt,name=role_angle,json=roleAngle,proto3" json:"role_angle,omitempty"`
	Action               uint32   `protobuf:"varint,4,opt,name=action,proto3" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncRoleAction) Reset()         { *m = SyncRoleAction{} }
func (m *SyncRoleAction) String() string { return proto.CompactTextString(m) }
func (*SyncRoleAction) ProtoMessage()    {}
func (*SyncRoleAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{11}
}

func (m *SyncRoleAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncRoleAction.Unmarshal(m, b)
}
func (m *SyncRoleAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncRoleAction.Marshal(b, m, deterministic)
}
func (m *SyncRoleAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncRoleAction.Merge(m, src)
}
func (m *SyncRoleAction) XXX_Size() int {
	return xxx_messageInfo_SyncRoleAction.Size(m)
}
func (m *SyncRoleAction) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncRoleAction.DiscardUnknown(m)
}

var xxx_messageInfo_SyncRoleAction proto.InternalMessageInfo

func (m *SyncRoleAction) GetRoleIdx() int32 {
	if m != nil {
		return m.RoleIdx
	}
	return 0
}

func (m *SyncRoleAction) GetPos() *RolePos {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *SyncRoleAction) GetRoleAngle() uint32 {
	if m != nil {
		return m.RoleAngle
	}
	return 0
}

func (m *SyncRoleAction) GetAction() uint32 {
	if m != nil {
		return m.Action
	}
	return 0
}

//请求操作块
type ReqOpBlock struct {
	OpCode               int32    `protobuf:"varint,1,opt,name=OpCode,proto3" json:"OpCode,omitempty"`
	BlockX               int32    `protobuf:"varint,2,opt,name=blockX,proto3" json:"blockX,omitempty"`
	BlockY               int32    `protobuf:"varint,3,opt,name=blockY,proto3" json:"blockY,omitempty"`
	BlockZ               int32    `protobuf:"varint,4,opt,name=blockZ,proto3" json:"blockZ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqOpBlock) Reset()         { *m = ReqOpBlock{} }
func (m *ReqOpBlock) String() string { return proto.CompactTextString(m) }
func (*ReqOpBlock) ProtoMessage()    {}
func (*ReqOpBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{12}
}

func (m *ReqOpBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqOpBlock.Unmarshal(m, b)
}
func (m *ReqOpBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqOpBlock.Marshal(b, m, deterministic)
}
func (m *ReqOpBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqOpBlock.Merge(m, src)
}
func (m *ReqOpBlock) XXX_Size() int {
	return xxx_messageInfo_ReqOpBlock.Size(m)
}
func (m *ReqOpBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqOpBlock.DiscardUnknown(m)
}

var xxx_messageInfo_ReqOpBlock proto.InternalMessageInfo

func (m *ReqOpBlock) GetOpCode() int32 {
	if m != nil {
		return m.OpCode
	}
	return 0
}

func (m *ReqOpBlock) GetBlockX() int32 {
	if m != nil {
		return m.BlockX
	}
	return 0
}

func (m *ReqOpBlock) GetBlockY() int32 {
	if m != nil {
		return m.BlockY
	}
	return 0
}

func (m *ReqOpBlock) GetBlockZ() int32 {
	if m != nil {
		return m.BlockZ
	}
	return 0
}

//同步块信息
type SyncBlock struct {
	BlockX               int32      `protobuf:"varint,1,opt,name=blockX,proto3" json:"blockX,omitempty"`
	BlockY               int32      `protobuf:"varint,2,opt,name=blockY,proto3" json:"blockY,omitempty"`
	BlockZ               int32      `protobuf:"varint,3,opt,name=blockZ,proto3" json:"blockZ,omitempty"`
	Info                 *BlockInfo `protobuf:"bytes,4,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SyncBlock) Reset()         { *m = SyncBlock{} }
func (m *SyncBlock) String() string { return proto.CompactTextString(m) }
func (*SyncBlock) ProtoMessage()    {}
func (*SyncBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{13}
}

func (m *SyncBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncBlock.Unmarshal(m, b)
}
func (m *SyncBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncBlock.Marshal(b, m, deterministic)
}
func (m *SyncBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncBlock.Merge(m, src)
}
func (m *SyncBlock) XXX_Size() int {
	return xxx_messageInfo_SyncBlock.Size(m)
}
func (m *SyncBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncBlock.DiscardUnknown(m)
}

var xxx_messageInfo_SyncBlock proto.InternalMessageInfo

func (m *SyncBlock) GetBlockX() int32 {
	if m != nil {
		return m.BlockX
	}
	return 0
}

func (m *SyncBlock) GetBlockY() int32 {
	if m != nil {
		return m.BlockY
	}
	return 0
}

func (m *SyncBlock) GetBlockZ() int32 {
	if m != nil {
		return m.BlockZ
	}
	return 0
}

func (m *SyncBlock) GetInfo() *BlockInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

//同步道具掉落
type SyncObjAdd struct {
	ObjId                uint32   `protobuf:"varint,1,opt,name=ObjId,proto3" json:"ObjId,omitempty"`
	ObjType              uint32   `protobuf:"varint,2,opt,name=ObjType,proto3" json:"ObjType,omitempty"`
	RemainTime           uint32   `protobuf:"varint,3,opt,name=remainTime,proto3" json:"remainTime,omitempty"`
	Count                uint32   `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
	OwnerRole            string   `protobuf:"bytes,5,opt,name=OwnerRole,proto3" json:"OwnerRole,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncObjAdd) Reset()         { *m = SyncObjAdd{} }
func (m *SyncObjAdd) String() string { return proto.CompactTextString(m) }
func (*SyncObjAdd) ProtoMessage()    {}
func (*SyncObjAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{14}
}

func (m *SyncObjAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncObjAdd.Unmarshal(m, b)
}
func (m *SyncObjAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncObjAdd.Marshal(b, m, deterministic)
}
func (m *SyncObjAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncObjAdd.Merge(m, src)
}
func (m *SyncObjAdd) XXX_Size() int {
	return xxx_messageInfo_SyncObjAdd.Size(m)
}
func (m *SyncObjAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncObjAdd.DiscardUnknown(m)
}

var xxx_messageInfo_SyncObjAdd proto.InternalMessageInfo

func (m *SyncObjAdd) GetObjId() uint32 {
	if m != nil {
		return m.ObjId
	}
	return 0
}

func (m *SyncObjAdd) GetObjType() uint32 {
	if m != nil {
		return m.ObjType
	}
	return 0
}

func (m *SyncObjAdd) GetRemainTime() uint32 {
	if m != nil {
		return m.RemainTime
	}
	return 0
}

func (m *SyncObjAdd) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *SyncObjAdd) GetOwnerRole() string {
	if m != nil {
		return m.OwnerRole
	}
	return ""
}

func init() {
	proto.RegisterType((*RolePos)(nil), "proto.role_pos")
	proto.RegisterType((*RoleOutlook)(nil), "proto.role_outlook")
	proto.RegisterType((*RoleBaseInfo)(nil), "proto.role_base_info")
	proto.RegisterType((*BlockInfo)(nil), "proto.block_info")
	proto.RegisterType((*ReqEnterGs)(nil), "proto.req_enter_gs")
	proto.RegisterType((*RspEnterGs)(nil), "proto.rsp_enter_gs")
	proto.RegisterType((*SyncChunkEnterRange)(nil), "proto.sync_chunk_enter_range")
	proto.RegisterType((*SyncChunkLeaveRange)(nil), "proto.sync_chunk_leave_range")
	proto.RegisterType((*SyncRoleEnterRange)(nil), "proto.sync_role_enter_range")
	proto.RegisterType((*SyncRoleLeaveRange)(nil), "proto.sync_role_leave_range")
	proto.RegisterType((*ReqRoleAction)(nil), "proto.req_role_action")
	proto.RegisterType((*SyncRoleAction)(nil), "proto.sync_role_action")
	proto.RegisterType((*ReqOpBlock)(nil), "proto.req_op_block")
	proto.RegisterType((*SyncBlock)(nil), "proto.sync_block")
	proto.RegisterType((*SyncObjAdd)(nil), "proto.sync_obj_add")
}

func init() { proto.RegisterFile("server_client.proto", fileDescriptor_55ccd995fd75bbbc) }

var fileDescriptor_55ccd995fd75bbbc = []byte{
	// 695 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0xf3, 0xb7, 0x99, 0xa4, 0xbf, 0xf6, 0xb7, 0xa5, 0x95, 0x0f, 0x80, 0xca, 0x4a, 0x48,
	0xe5, 0x40, 0xa5, 0x16, 0x2e, 0x48, 0x5c, 0x68, 0x39, 0xd0, 0x53, 0xd0, 0xaa, 0x87, 0x36, 0x17,
	0xcb, 0x4e, 0x36, 0xa9, 0x1b, 0x67, 0xd7, 0xd8, 0x4e, 0xdb, 0x14, 0x71, 0x45, 0x9c, 0xf9, 0x58,
	0x7c, 0x2a, 0x34, 0xb3, 0xe3, 0xc4, 0x45, 0x44, 0xaa, 0x38, 0x79, 0xdf, 0xf3, 0xec, 0xcc, 0x9b,
	0xb7, 0xb3, 0x0b, 0x3b, 0xb9, 0xce, 0x6e, 0x74, 0x16, 0x0c, 0x93, 0x58, 0x9b, 0xe2, 0x30, 0xcd,
	0x6c, 0x61, 0x45, 0x93, 0x3e, 0xf2, 0x2d, 0x6c, 0x64, 0x36, 0xd1, 0x41, 0x6a, 0x73, 0xd1, 0x03,
	0xef, 0xce, 0xf7, 0xf6, 0xbd, 0x83, 0x9a, 0xf2, 0xee, 0x10, 0x2d, 0xfc, 0x9a, 0x43, 0x0b, 0x44,
	0xf7, 0x7e, 0xdd, 0xa1, 0x7b, 0xf9, 0xc3, 0x83, 0x1e, 0x6d, 0xb3, 0xf3, 0x22, 0xb1, 0x76, 0x2a,
	0xf6, 0xa0, 0x75, 0xab, 0xc3, 0xd4, 0x1a, 0xda, 0xbf, 0xa9, 0x18, 0x09, 0x01, 0x8d, 0x2b, 0x9d,
	0xcc, 0x28, 0xcf, 0xa6, 0xa2, 0x35, 0x72, 0xe3, 0x70, 0xa8, 0x29, 0xdb, 0xa6, 0xa2, 0x35, 0x72,
	0xb7, 0xb1, 0x99, 0xf8, 0x0d, 0xc7, 0xe1, 0x5a, 0x6c, 0x43, 0x3d, 0x0a, 0x27, 0x7e, 0x93, 0x28,
	0x5c, 0x62, 0x54, 0x3e, 0x8f, 0x0b, 0xbf, 0xe5, 0xa2, 0x70, 0x2d, 0x7f, 0x79, 0xf0, 0x1f, 0x49,
	0x89, 0xc2, 0x5c, 0x07, 0xb1, 0x19, 0x5b, 0x0c, 0x33, 0xe1, 0x4c, 0x93, 0x94, 0x8e, 0xa2, 0xb5,
	0x78, 0x02, 0xcd, 0x44, 0xdf, 0xe8, 0x84, 0x95, 0x38, 0x20, 0x9e, 0x01, 0x24, 0xf1, 0x18, 0xbb,
	0x8f, 0x4d, 0xc1, 0x82, 0x3a, 0xc8, 0x7c, 0x46, 0x42, 0xbc, 0x80, 0x7a, 0x6a, 0x73, 0x52, 0xd0,
	0x3d, 0xde, 0x72, 0xc6, 0x1d, 0x96, 0x76, 0x29, 0xfc, 0x87, 0x19, 0x88, 0x08, 0xcd, 0x24, 0xd1,
	0x2c, 0xac, 0x83, 0xcc, 0x07, 0x24, 0xc4, 0x6b, 0x68, 0xb3, 0x45, 0x7e, 0x9b, 0xb2, 0xec, 0x54,
	0xb3, 0xf0, 0x2f, 0x55, 0xc6, 0xc8, 0x8f, 0x00, 0x51, 0x62, 0x87, 0x53, 0xd7, 0xc7, 0x53, 0xe8,
	0x9c, 0x20, 0x3a, 0x5f, 0xa4, 0x9a, 0x7d, 0x5d, 0x11, 0xc2, 0x87, 0xf6, 0xa9, 0x35, 0x85, 0x36,
	0x05, 0xf7, 0x54, 0x42, 0xf9, 0x1e, 0x7a, 0x99, 0xfe, 0x12, 0x68, 0x53, 0xe8, 0x2c, 0x98, 0xe4,
	0xd8, 0x7b, 0x61, 0xa7, 0xda, 0xb0, 0x21, 0x0e, 0xe0, 0x91, 0xa1, 0x88, 0xb3, 0x11, 0x6d, 0xef,
	0x28, 0x46, 0xf2, 0x1b, 0xf4, 0xb2, 0x3c, 0x5d, 0xed, 0xf6, 0xa1, 0x9d, 0xe9, 0xe2, 0xd4, 0x8e,
	0x9c, 0x86, 0xa6, 0x2a, 0xa1, 0xd8, 0x87, 0xee, 0x2c, 0x8c, 0x8d, 0xa2, 0x7d, 0x77, 0x94, 0xa6,
	0xa9, 0xaa, 0x94, 0x38, 0x82, 0x8d, 0x12, 0x92, 0xbb, 0xdd, 0xe3, 0xdd, 0x6a, 0xff, 0xcb, 0x23,
	0x53, 0xcb, 0x30, 0x99, 0xc3, 0x5e, 0xbe, 0x30, 0xc3, 0x60, 0x78, 0x35, 0x37, 0x53, 0x56, 0x91,
	0x85, 0x66, 0xa2, 0x51, 0x30, 0x91, 0x17, 0xac, 0x83, 0xd1, 0x92, 0x1f, 0xb0, 0x02, 0x46, 0xe2,
	0x15, 0xb4, 0xc8, 0xcc, 0xdc, 0xaf, 0xef, 0xd7, 0x0f, 0xba, 0xc7, 0xff, 0x73, 0xe9, 0x95, 0xc3,
	0x8a, 0x03, 0xe4, 0xa7, 0x07, 0x45, 0x13, 0x1d, 0xde, 0xe8, 0x7f, 0x2b, 0x2a, 0x13, 0xd8, 0xa5,
	0x4c, 0xd4, 0x5f, 0x55, 0x3d, 0xda, 0xc8, 0x46, 0x95, 0x36, 0xb2, 0x49, 0xef, 0xdc, 0x5d, 0x3a,
	0x09, 0x73, 0x7d, 0x66, 0xc6, 0x96, 0x12, 0xae, 0x35, 0xea, 0x41, 0xa8, 0x3c, 0xaa, 0x56, 0xab,
	0xca, 0x5e, 0x5b, 0x4d, 0x4e, 0x61, 0x0b, 0x87, 0xc3, 0x0d, 0xed, 0xb0, 0x88, 0xad, 0x29, 0xc7,
	0xdc, 0x7b, 0xf4, 0x98, 0xd7, 0xfe, 0x1c, 0xf3, 0x3d, 0x68, 0xb9, 0x5c, 0x7c, 0x87, 0x18, 0xc9,
	0xef, 0x1e, 0x6c, 0xaf, 0x04, 0x72, 0xb9, 0xf5, 0x4e, 0xb0, 0x90, 0xda, 0xa3, 0x85, 0xd4, 0xd7,
	0x0b, 0x69, 0x3c, 0x10, 0x62, 0xdc, 0x95, 0xb0, 0x69, 0x40, 0x27, 0x8e, 0x71, 0xfd, 0xb4, 0x32,
	0xd3, 0x8c, 0x90, 0xa7, 0x80, 0x8b, 0xf2, 0x58, 0x1d, 0x5a, 0xf2, 0x97, 0x54, 0xb2, 0xe4, 0x2f,
	0x97, 0xfc, 0x80, 0xea, 0x95, 0xfc, 0x40, 0x7e, 0x05, 0xa0, 0xbe, 0x97, 0xd5, 0x38, 0xab, 0xb7,
	0x26, 0x6b, 0x6d, 0x4d, 0xd6, 0x6a, 0xb5, 0x81, 0x78, 0x09, 0x0d, 0x1c, 0x02, 0xaa, 0xf5, 0xd7,
	0x79, 0xa6, 0xdf, 0xf2, 0xa7, 0x07, 0x3d, 0xaa, 0x6e, 0xa3, 0xeb, 0x20, 0x1c, 0x8d, 0xf0, 0x01,
	0xe8, 0x47, 0xd7, 0x67, 0x23, 0x7e, 0x44, 0x1c, 0xc0, 0x73, 0xe8, 0x47, 0xd7, 0xf4, 0xb8, 0xf0,
	0x03, 0xc2, 0x50, 0x3c, 0x07, 0xc8, 0x34, 0xde, 0xc8, 0xf3, 0x78, 0x56, 0x9a, 0x5c, 0x61, 0x30,
	0xdf, 0xa9, 0x9d, 0x9b, 0x82, 0x4d, 0x76, 0x00, 0x9f, 0xab, 0xfe, 0xad, 0xd1, 0x19, 0xdd, 0xf6,
	0x26, 0xbd, 0x29, 0x2b, 0x22, 0x6a, 0x91, 0xd8, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xab,
	0xe5, 0x9e, 0xfe, 0x8d, 0x06, 0x00, 0x00,
}
