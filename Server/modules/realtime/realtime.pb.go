// Code generated by protoc-gen-go. DO NOT EDIT.
// source: realtime.proto

package realtime

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Packet struct {
	SenderId   string   `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	IsServer   bool     `protobuf:"varint,2,opt,name=is_server,json=isServer,proto3" json:"is_server,omitempty"`
	Recipients []string `protobuf:"bytes,3,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Types that are valid to be assigned to Type:
	//	*Packet_UpdateTransform
	//	*Packet_Spawn
	//	*Packet_RequestSpawn
	//	*Packet_Destroy
	//	*Packet_RequestDestroy
	//	*Packet_Meme
	//	*Packet_MatchInformation
	//	*Packet_Initialized
	Type                 isPacket_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *Packet) GetIsServer() bool {
	if m != nil {
		return m.IsServer
	}
	return false
}

func (m *Packet) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

type isPacket_Type interface {
	isPacket_Type()
}

type Packet_UpdateTransform struct {
	UpdateTransform *UpdateTransform `protobuf:"bytes,4,opt,name=update_transform,json=updateTransform,proto3,oneof"`
}

type Packet_Spawn struct {
	Spawn *Spawn `protobuf:"bytes,5,opt,name=spawn,proto3,oneof"`
}

type Packet_RequestSpawn struct {
	RequestSpawn *Spawn `protobuf:"bytes,6,opt,name=request_spawn,json=requestSpawn,proto3,oneof"`
}

type Packet_Destroy struct {
	Destroy *Destroy `protobuf:"bytes,7,opt,name=destroy,proto3,oneof"`
}

type Packet_RequestDestroy struct {
	RequestDestroy *Destroy `protobuf:"bytes,8,opt,name=request_destroy,json=requestDestroy,proto3,oneof"`
}

type Packet_Meme struct {
	Meme *Meme `protobuf:"bytes,10,opt,name=meme,proto3,oneof"`
}

type Packet_MatchInformation struct {
	MatchInformation *MatchInformation `protobuf:"bytes,11,opt,name=match_information,json=matchInformation,proto3,oneof"`
}

type Packet_Initialized struct {
	Initialized *Initialized `protobuf:"bytes,12,opt,name=initialized,proto3,oneof"`
}

func (*Packet_UpdateTransform) isPacket_Type() {}

func (*Packet_Spawn) isPacket_Type() {}

func (*Packet_RequestSpawn) isPacket_Type() {}

func (*Packet_Destroy) isPacket_Type() {}

func (*Packet_RequestDestroy) isPacket_Type() {}

func (*Packet_Meme) isPacket_Type() {}

func (*Packet_MatchInformation) isPacket_Type() {}

func (*Packet_Initialized) isPacket_Type() {}

func (m *Packet) GetType() isPacket_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Packet) GetUpdateTransform() *UpdateTransform {
	if x, ok := m.GetType().(*Packet_UpdateTransform); ok {
		return x.UpdateTransform
	}
	return nil
}

func (m *Packet) GetSpawn() *Spawn {
	if x, ok := m.GetType().(*Packet_Spawn); ok {
		return x.Spawn
	}
	return nil
}

func (m *Packet) GetRequestSpawn() *Spawn {
	if x, ok := m.GetType().(*Packet_RequestSpawn); ok {
		return x.RequestSpawn
	}
	return nil
}

func (m *Packet) GetDestroy() *Destroy {
	if x, ok := m.GetType().(*Packet_Destroy); ok {
		return x.Destroy
	}
	return nil
}

func (m *Packet) GetRequestDestroy() *Destroy {
	if x, ok := m.GetType().(*Packet_RequestDestroy); ok {
		return x.RequestDestroy
	}
	return nil
}

func (m *Packet) GetMeme() *Meme {
	if x, ok := m.GetType().(*Packet_Meme); ok {
		return x.Meme
	}
	return nil
}

func (m *Packet) GetMatchInformation() *MatchInformation {
	if x, ok := m.GetType().(*Packet_MatchInformation); ok {
		return x.MatchInformation
	}
	return nil
}

func (m *Packet) GetInitialized() *Initialized {
	if x, ok := m.GetType().(*Packet_Initialized); ok {
		return x.Initialized
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Packet) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Packet_UpdateTransform)(nil),
		(*Packet_Spawn)(nil),
		(*Packet_RequestSpawn)(nil),
		(*Packet_Destroy)(nil),
		(*Packet_RequestDestroy)(nil),
		(*Packet_Meme)(nil),
		(*Packet_MatchInformation)(nil),
		(*Packet_Initialized)(nil),
	}
}

// General purpose transform update typically shared by several different types of objects
type UpdateTransform struct {
	Transform            *Transform `protobuf:"bytes,1,opt,name=transform,proto3" json:"transform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateTransform) Reset()         { *m = UpdateTransform{} }
func (m *UpdateTransform) String() string { return proto.CompactTextString(m) }
func (*UpdateTransform) ProtoMessage()    {}
func (*UpdateTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{1}
}

func (m *UpdateTransform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTransform.Unmarshal(m, b)
}
func (m *UpdateTransform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTransform.Marshal(b, m, deterministic)
}
func (m *UpdateTransform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTransform.Merge(m, src)
}
func (m *UpdateTransform) XXX_Size() int {
	return xxx_messageInfo_UpdateTransform.Size(m)
}
func (m *UpdateTransform) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTransform.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTransform proto.InternalMessageInfo

func (m *UpdateTransform) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

type Spawn struct {
	// Types that are valid to be assigned to Type:
	//	*Spawn_Any
	//	*Spawn_Tree
	//	*Spawn_Animal
	Type                 isSpawn_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Spawn) Reset()         { *m = Spawn{} }
func (m *Spawn) String() string { return proto.CompactTextString(m) }
func (*Spawn) ProtoMessage()    {}
func (*Spawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{2}
}

func (m *Spawn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spawn.Unmarshal(m, b)
}
func (m *Spawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spawn.Marshal(b, m, deterministic)
}
func (m *Spawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spawn.Merge(m, src)
}
func (m *Spawn) XXX_Size() int {
	return xxx_messageInfo_Spawn.Size(m)
}
func (m *Spawn) XXX_DiscardUnknown() {
	xxx_messageInfo_Spawn.DiscardUnknown(m)
}

var xxx_messageInfo_Spawn proto.InternalMessageInfo

type isSpawn_Type interface {
	isSpawn_Type()
}

type Spawn_Any struct {
	Any *Transform `protobuf:"bytes,1,opt,name=any,proto3,oneof"`
}

type Spawn_Tree struct {
	Tree *Tree `protobuf:"bytes,2,opt,name=tree,proto3,oneof"`
}

type Spawn_Animal struct {
	Animal *Animal `protobuf:"bytes,3,opt,name=animal,proto3,oneof"`
}

func (*Spawn_Any) isSpawn_Type() {}

func (*Spawn_Tree) isSpawn_Type() {}

func (*Spawn_Animal) isSpawn_Type() {}

func (m *Spawn) GetType() isSpawn_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Spawn) GetAny() *Transform {
	if x, ok := m.GetType().(*Spawn_Any); ok {
		return x.Any
	}
	return nil
}

func (m *Spawn) GetTree() *Tree {
	if x, ok := m.GetType().(*Spawn_Tree); ok {
		return x.Tree
	}
	return nil
}

func (m *Spawn) GetAnimal() *Animal {
	if x, ok := m.GetType().(*Spawn_Animal); ok {
		return x.Animal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Spawn) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Spawn_Any)(nil),
		(*Spawn_Tree)(nil),
		(*Spawn_Animal)(nil),
	}
}

type Destroy struct {
	// Types that are valid to be assigned to Type:
	//	*Destroy_Any
	//	*Destroy_Tree
	//	*Destroy_Animal
	Type                 isDestroy_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Destroy) Reset()         { *m = Destroy{} }
func (m *Destroy) String() string { return proto.CompactTextString(m) }
func (*Destroy) ProtoMessage()    {}
func (*Destroy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{3}
}

func (m *Destroy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destroy.Unmarshal(m, b)
}
func (m *Destroy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destroy.Marshal(b, m, deterministic)
}
func (m *Destroy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destroy.Merge(m, src)
}
func (m *Destroy) XXX_Size() int {
	return xxx_messageInfo_Destroy.Size(m)
}
func (m *Destroy) XXX_DiscardUnknown() {
	xxx_messageInfo_Destroy.DiscardUnknown(m)
}

var xxx_messageInfo_Destroy proto.InternalMessageInfo

type isDestroy_Type interface {
	isDestroy_Type()
}

type Destroy_Any struct {
	Any *Transform `protobuf:"bytes,1,opt,name=any,proto3,oneof"`
}

type Destroy_Tree struct {
	Tree *Tree `protobuf:"bytes,2,opt,name=tree,proto3,oneof"`
}

type Destroy_Animal struct {
	Animal *Animal `protobuf:"bytes,3,opt,name=animal,proto3,oneof"`
}

func (*Destroy_Any) isDestroy_Type() {}

func (*Destroy_Tree) isDestroy_Type() {}

func (*Destroy_Animal) isDestroy_Type() {}

func (m *Destroy) GetType() isDestroy_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Destroy) GetAny() *Transform {
	if x, ok := m.GetType().(*Destroy_Any); ok {
		return x.Any
	}
	return nil
}

func (m *Destroy) GetTree() *Tree {
	if x, ok := m.GetType().(*Destroy_Tree); ok {
		return x.Tree
	}
	return nil
}

func (m *Destroy) GetAnimal() *Animal {
	if x, ok := m.GetType().(*Destroy_Animal); ok {
		return x.Animal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Destroy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Destroy_Any)(nil),
		(*Destroy_Tree)(nil),
		(*Destroy_Animal)(nil),
	}
}

type Meme struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MemeName             string   `protobuf:"bytes,2,opt,name=meme_name,json=memeName,proto3" json:"meme_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Meme) Reset()         { *m = Meme{} }
func (m *Meme) String() string { return proto.CompactTextString(m) }
func (*Meme) ProtoMessage()    {}
func (*Meme) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{4}
}

func (m *Meme) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Meme.Unmarshal(m, b)
}
func (m *Meme) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Meme.Marshal(b, m, deterministic)
}
func (m *Meme) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Meme.Merge(m, src)
}
func (m *Meme) XXX_Size() int {
	return xxx_messageInfo_Meme.Size(m)
}
func (m *Meme) XXX_DiscardUnknown() {
	xxx_messageInfo_Meme.DiscardUnknown(m)
}

var xxx_messageInfo_Meme proto.InternalMessageInfo

func (m *Meme) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Meme) GetMemeName() string {
	if m != nil {
		return m.MemeName
	}
	return ""
}

// Server sending some information on the current state before starting
type MatchInformation struct {
	// TODO: some state config https://heroiclabs.com/docs/tutorial-remote-configuration/
	// like map size, initial evolution parameters ...
	Seed                 int32    `protobuf:"varint,1,opt,name=seed,proto3" json:"seed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MatchInformation) Reset()         { *m = MatchInformation{} }
func (m *MatchInformation) String() string { return proto.CompactTextString(m) }
func (*MatchInformation) ProtoMessage()    {}
func (*MatchInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{5}
}

func (m *MatchInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MatchInformation.Unmarshal(m, b)
}
func (m *MatchInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MatchInformation.Marshal(b, m, deterministic)
}
func (m *MatchInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MatchInformation.Merge(m, src)
}
func (m *MatchInformation) XXX_Size() int {
	return xxx_messageInfo_MatchInformation.Size(m)
}
func (m *MatchInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_MatchInformation.DiscardUnknown(m)
}

var xxx_messageInfo_MatchInformation proto.InternalMessageInfo

func (m *MatchInformation) GetSeed() int32 {
	if m != nil {
		return m.Seed
	}
	return 0
}

// Client notifying being ready to handle gameplay
type Initialized struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Initialized) Reset()         { *m = Initialized{} }
func (m *Initialized) String() string { return proto.CompactTextString(m) }
func (*Initialized) ProtoMessage()    {}
func (*Initialized) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{6}
}

func (m *Initialized) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Initialized.Unmarshal(m, b)
}
func (m *Initialized) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Initialized.Marshal(b, m, deterministic)
}
func (m *Initialized) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Initialized.Merge(m, src)
}
func (m *Initialized) XXX_Size() int {
	return xxx_messageInfo_Initialized.Size(m)
}
func (m *Initialized) XXX_DiscardUnknown() {
	xxx_messageInfo_Initialized.DiscardUnknown(m)
}

var xxx_messageInfo_Initialized proto.InternalMessageInfo

type Transform struct {
	Id                   uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Position             *Vector3    `protobuf:"bytes,2,opt,name=position,proto3" json:"position,omitempty"`
	Rotation             *Quaternion `protobuf:"bytes,3,opt,name=rotation,proto3" json:"rotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Transform) Reset()         { *m = Transform{} }
func (m *Transform) String() string { return proto.CompactTextString(m) }
func (*Transform) ProtoMessage()    {}
func (*Transform) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{7}
}

func (m *Transform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transform.Unmarshal(m, b)
}
func (m *Transform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transform.Marshal(b, m, deterministic)
}
func (m *Transform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transform.Merge(m, src)
}
func (m *Transform) XXX_Size() int {
	return xxx_messageInfo_Transform.Size(m)
}
func (m *Transform) XXX_DiscardUnknown() {
	xxx_messageInfo_Transform.DiscardUnknown(m)
}

var xxx_messageInfo_Transform proto.InternalMessageInfo

func (m *Transform) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Transform) GetPosition() *Vector3 {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Transform) GetRotation() *Quaternion {
	if m != nil {
		return m.Rotation
	}
	return nil
}

type Tree struct {
	Transform            *Transform `protobuf:"bytes,1,opt,name=transform,proto3" json:"transform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Tree) Reset()         { *m = Tree{} }
func (m *Tree) String() string { return proto.CompactTextString(m) }
func (*Tree) ProtoMessage()    {}
func (*Tree) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{8}
}

func (m *Tree) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tree.Unmarshal(m, b)
}
func (m *Tree) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tree.Marshal(b, m, deterministic)
}
func (m *Tree) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tree.Merge(m, src)
}
func (m *Tree) XXX_Size() int {
	return xxx_messageInfo_Tree.Size(m)
}
func (m *Tree) XXX_DiscardUnknown() {
	xxx_messageInfo_Tree.DiscardUnknown(m)
}

var xxx_messageInfo_Tree proto.InternalMessageInfo

func (m *Tree) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

type Animal struct {
	Transform            *Transform `protobuf:"bytes,1,opt,name=transform,proto3" json:"transform,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Animal) Reset()         { *m = Animal{} }
func (m *Animal) String() string { return proto.CompactTextString(m) }
func (*Animal) ProtoMessage()    {}
func (*Animal) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{9}
}

func (m *Animal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Animal.Unmarshal(m, b)
}
func (m *Animal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Animal.Marshal(b, m, deterministic)
}
func (m *Animal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Animal.Merge(m, src)
}
func (m *Animal) XXX_Size() int {
	return xxx_messageInfo_Animal.Size(m)
}
func (m *Animal) XXX_DiscardUnknown() {
	xxx_messageInfo_Animal.DiscardUnknown(m)
}

var xxx_messageInfo_Animal proto.InternalMessageInfo

func (m *Animal) GetTransform() *Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

type Vector3 struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector3) Reset()         { *m = Vector3{} }
func (m *Vector3) String() string { return proto.CompactTextString(m) }
func (*Vector3) ProtoMessage()    {}
func (*Vector3) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{10}
}

func (m *Vector3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector3.Unmarshal(m, b)
}
func (m *Vector3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector3.Marshal(b, m, deterministic)
}
func (m *Vector3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector3.Merge(m, src)
}
func (m *Vector3) XXX_Size() int {
	return xxx_messageInfo_Vector3.Size(m)
}
func (m *Vector3) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector3.DiscardUnknown(m)
}

var xxx_messageInfo_Vector3 proto.InternalMessageInfo

func (m *Vector3) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector3) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector3) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Quaternion struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	W                    float32  `protobuf:"fixed32,4,opt,name=w,proto3" json:"w,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Quaternion) Reset()         { *m = Quaternion{} }
func (m *Quaternion) String() string { return proto.CompactTextString(m) }
func (*Quaternion) ProtoMessage()    {}
func (*Quaternion) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcbdca058206953b, []int{11}
}

func (m *Quaternion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Quaternion.Unmarshal(m, b)
}
func (m *Quaternion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Quaternion.Marshal(b, m, deterministic)
}
func (m *Quaternion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Quaternion.Merge(m, src)
}
func (m *Quaternion) XXX_Size() int {
	return xxx_messageInfo_Quaternion.Size(m)
}
func (m *Quaternion) XXX_DiscardUnknown() {
	xxx_messageInfo_Quaternion.DiscardUnknown(m)
}

var xxx_messageInfo_Quaternion proto.InternalMessageInfo

func (m *Quaternion) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Quaternion) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Quaternion) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Quaternion) GetW() float32 {
	if m != nil {
		return m.W
	}
	return 0
}

func init() {
	proto.RegisterType((*Packet)(nil), "niwrad.realtime.Packet")
	proto.RegisterType((*UpdateTransform)(nil), "niwrad.realtime.UpdateTransform")
	proto.RegisterType((*Spawn)(nil), "niwrad.realtime.Spawn")
	proto.RegisterType((*Destroy)(nil), "niwrad.realtime.Destroy")
	proto.RegisterType((*Meme)(nil), "niwrad.realtime.Meme")
	proto.RegisterType((*MatchInformation)(nil), "niwrad.realtime.MatchInformation")
	proto.RegisterType((*Initialized)(nil), "niwrad.realtime.Initialized")
	proto.RegisterType((*Transform)(nil), "niwrad.realtime.Transform")
	proto.RegisterType((*Tree)(nil), "niwrad.realtime.Tree")
	proto.RegisterType((*Animal)(nil), "niwrad.realtime.Animal")
	proto.RegisterType((*Vector3)(nil), "niwrad.realtime.Vector3")
	proto.RegisterType((*Quaternion)(nil), "niwrad.realtime.Quaternion")
}

func init() {
	proto.RegisterFile("realtime.proto", fileDescriptor_dcbdca058206953b)
}

var fileDescriptor_dcbdca058206953b = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xe3, 0x34, 0x6d, 0xd3, 0xd3, 0x6e, 0xdd, 0x2c, 0x7d, 0x1f, 0x11, 0x43, 0xa8, 0xe4,
	0x02, 0x55, 0x42, 0x8a, 0x04, 0x9d, 0x04, 0x12, 0x42, 0x1a, 0x63, 0x17, 0x99, 0xd0, 0xa6, 0xe1,
	0x0d, 0x2e, 0xb8, 0xa9, 0x4c, 0x73, 0x10, 0x16, 0x8b, 0x53, 0x1c, 0x97, 0xae, 0x7b, 0x04, 0x5e,
	0x02, 0x71, 0xcb, 0x73, 0xf1, 0x20, 0xc8, 0x4e, 0xd2, 0x8e, 0x06, 0x34, 0xa4, 0xdd, 0x70, 0x97,
	0x73, 0xfc, 0xfb, 0x1f, 0xfb, 0x1c, 0xff, 0xdd, 0xc2, 0xa6, 0x42, 0x7e, 0xae, 0x45, 0x8a, 0xd1,
	0x54, 0x65, 0x3a, 0xa3, 0x7d, 0x29, 0xe6, 0x8a, 0x27, 0x51, 0x95, 0x0e, 0x7f, 0x78, 0xd0, 0x3a,
	0xe1, 0x93, 0x8f, 0xa8, 0xe9, 0x0e, 0x74, 0x72, 0x94, 0x09, 0xaa, 0xb1, 0x48, 0x02, 0x32, 0x20,
	0xc3, 0x0e, 0xf3, 0x8b, 0xc4, 0x61, 0x62, 0x16, 0x45, 0x3e, 0xce, 0x51, 0x7d, 0x46, 0x15, 0xb8,
	0x03, 0x32, 0xf4, 0x99, 0x2f, 0xf2, 0x53, 0x1b, 0xd3, 0xbb, 0x00, 0x0a, 0x27, 0x62, 0x2a, 0x50,
	0xea, 0x3c, 0x68, 0x0c, 0x1a, 0xc3, 0x0e, 0xbb, 0x92, 0xa1, 0x47, 0xb0, 0x35, 0x9b, 0x26, 0x5c,
	0xe3, 0x58, 0x2b, 0x2e, 0xf3, 0xf7, 0x99, 0x4a, 0x03, 0x6f, 0x40, 0x86, 0xdd, 0x47, 0x83, 0x68,
	0xed, 0x40, 0xd1, 0x6b, 0x0b, 0x9e, 0x55, 0x5c, 0xec, 0xb0, 0xfe, 0xec, 0xd7, 0x14, 0x8d, 0xa0,
	0x99, 0x4f, 0xf9, 0x5c, 0x06, 0x4d, 0x5b, 0xe3, 0xff, 0x5a, 0x8d, 0x53, 0xb3, 0x1a, 0x3b, 0xac,
	0xc0, 0xe8, 0x33, 0xd8, 0x50, 0xf8, 0x69, 0x86, 0xb9, 0x1e, 0x17, 0xba, 0xd6, 0x35, 0xba, 0x5e,
	0x89, 0xdb, 0x98, 0xee, 0x42, 0x3b, 0xc1, 0x5c, 0xab, 0x6c, 0x11, 0xb4, 0xad, 0x30, 0xa8, 0x09,
	0x0f, 0x8a, 0xf5, 0xd8, 0x61, 0x15, 0x4a, 0x5f, 0x40, 0xbf, 0xda, 0xb4, 0x52, 0xfb, 0xd7, 0xaa,
	0x37, 0x4b, 0x49, 0x99, 0xa1, 0x0f, 0xc0, 0x4b, 0x31, 0xc5, 0x00, 0xac, 0xf2, 0xbf, 0x9a, 0xf2,
	0x08, 0x53, 0x8c, 0x1d, 0x66, 0x21, 0x7a, 0x02, 0xdb, 0x29, 0xd7, 0x93, 0x0f, 0x63, 0x21, 0xcd,
	0x98, 0xb8, 0x16, 0x99, 0x0c, 0xba, 0x56, 0x79, 0xaf, 0xae, 0x34, 0xe4, 0xe1, 0x0a, 0x8c, 0x1d,
	0xb6, 0x95, 0xae, 0xe5, 0xe8, 0x1e, 0x74, 0x85, 0x14, 0x5a, 0xf0, 0x73, 0x71, 0x89, 0x49, 0xd0,
	0xb3, 0xb5, 0xee, 0xd4, 0x6a, 0x1d, 0xae, 0x98, 0xd8, 0x61, 0x57, 0x25, 0xfb, 0x2d, 0xf0, 0xf4,
	0x62, 0x8a, 0xe1, 0x4b, 0xe8, 0xaf, 0x5d, 0x2c, 0x7d, 0x02, 0x9d, 0x95, 0x1b, 0x88, 0x2d, 0x7d,
	0xbb, 0x56, 0x7a, 0x89, 0xb3, 0x15, 0x1c, 0x7e, 0x25, 0xd0, 0x2c, 0xae, 0x26, 0x82, 0x06, 0x97,
	0x8b, 0xeb, 0xd5, 0xb1, 0xc3, 0x0c, 0x68, 0xe6, 0xa9, 0x15, 0xa2, 0x35, 0xf0, 0xef, 0xe6, 0x79,
	0xa6, 0xd0, 0xce, 0xd3, 0x40, 0xf4, 0x21, 0xb4, 0xb8, 0x14, 0x29, 0x3f, 0x0f, 0x1a, 0x16, 0xbf,
	0x55, 0xc3, 0x9f, 0xdb, 0xe5, 0xd8, 0x61, 0x25, 0xb8, 0x6c, 0xf7, 0x1b, 0x81, 0x76, 0x75, 0x87,
	0xff, 0xea, 0x19, 0x47, 0xe0, 0x19, 0xfb, 0xd0, 0x4d, 0x70, 0xcb, 0xf7, 0xee, 0x31, 0x57, 0xd8,
	0x97, 0x6e, 0xec, 0x34, 0x96, 0x3c, 0x2d, 0x0e, 0xd1, 0x61, 0xbe, 0x49, 0x1c, 0xf3, 0x14, 0xc3,
	0xfb, 0xb0, 0xb5, 0xee, 0x1c, 0x4a, 0xc1, 0xcb, 0x11, 0x8b, 0x12, 0x4d, 0x66, 0xbf, 0xc3, 0x0d,
	0xe8, 0x5e, 0x71, 0x45, 0xf8, 0x85, 0x40, 0x67, 0x75, 0xf3, 0xeb, 0x3b, 0xee, 0x82, 0x3f, 0xcd,
	0x72, 0x61, 0xfd, 0xea, 0xfe, 0xe1, 0x8d, 0xbc, 0xc1, 0x89, 0xce, 0xd4, 0x88, 0x2d, 0x49, 0xfa,
	0x18, 0x7c, 0x95, 0xe9, 0xc2, 0xe5, 0x45, 0xf3, 0x3b, 0x35, 0xd5, 0xab, 0x19, 0xd7, 0xa8, 0xa4,
	0xc8, 0x24, 0x5b, 0xc2, 0xe1, 0x1e, 0x78, 0x66, 0x86, 0x37, 0x30, 0xe0, 0x3e, 0xb4, 0x8a, 0xb1,
	0xde, 0xa0, 0xc6, 0x08, 0xda, 0x65, 0x4f, 0xb4, 0x07, 0xe4, 0xc2, 0x8a, 0x5d, 0x46, 0x2e, 0x4c,
	0xb4, 0xb0, 0x63, 0x70, 0x19, 0x59, 0x98, 0xe8, 0xd2, 0xb6, 0xe7, 0x32, 0x72, 0x19, 0x1e, 0x00,
	0xac, 0x5a, 0xfa, 0x7b, 0x9d, 0x89, 0xe6, 0xf6, 0x17, 0xd7, 0x65, 0x64, 0xbe, 0x1f, 0xbe, 0xdd,
	0xae, 0xce, 0xf6, 0xb4, 0xfa, 0xf8, 0xee, 0xf6, 0x8e, 0x51, 0x47, 0xac, 0x0c, 0xdf, 0xb5, 0xec,
	0xff, 0xc5, 0xe8, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0x5a, 0x49, 0x42, 0x41, 0x06, 0x00,
	0x00,
}