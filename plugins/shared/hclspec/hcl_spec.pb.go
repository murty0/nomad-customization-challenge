// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugins/shared/hclspec/hcl_spec.proto

// Spec allows exposing the specification for an HCL body, allowing for parsing and
//validation.
//
//Certain expressions within a specification may use the following functions.
//The documentation for each spec type above specifies where functions may
//be used.
//
// `abs(number)` returns the absolute (positive) value of the given number.
// `coalesce(vals...)` returns the first non-null value given.
// `concat(lists...)` concatenates together all of the given lists to produce a new list.
// `hasindex(val, idx)` returns true if the expression `val[idx]` could succeed.
// `int(number)` returns the integer portion of the given number, rounding towards zero.
// `jsondecode(str)` interprets the given string as JSON and returns the resulting data structure.
// `jsonencode(val)` returns a JSON-serialized version of the given value.
// `length(collection)` returns the number of elements in the given collection (list, set, map, object, or tuple).
// `lower(string)` returns the given string with all uppercase letters converted to lowercase.
// `max(numbers...)` returns the greatest of the given numbers.
// `min(numbers...)` returns the smallest of the given numbers.
// `reverse(string)` returns the given string with all of the characters in reverse order.
// `strlen(string)` returns the number of characters in the given string.
// `substr(string, offset, length)` returns the requested substring of the given string.
// `upper(string)` returns the given string with all lowercase letters converted to uppercase.
//
//## Type Expressions
//
//Type expressions are used to describe the expected type of an attribute, as
//an additional validation constraint.
//
//A type expression uses primitive type names and compound type constructors.
//A type constructor builds a new type based on one or more type expression
//arguments.
//
//The following type names and type constructors are supported:
//
// `any` is a wildcard that accepts a value of any type. (In HCL terms, this
// is the _dynamic pseudo-type_.)
// `string` is a Unicode string.
// `number` is an arbitrary-precision floating point number.
// `bool` is a boolean value (`true` or `false`)
// `list(element_type)` constructs a list type with the given element type
// `set(element_type)` constructs a set type with the given element type
// `map(element_type)` constructs a map type with the given element type
// `object({name1 = element_type, name2 = element_type, ...})` constructs
// an object type with the given attribute types.
// `tuple([element_type, element_type, ...])` constructs a tuple type with
// the given element types. This can be used, for example, to require an
// array with a particular number of elements, or with elements of different
// types.
//
//`null` is a valid value of any type, and not a type itself.

package hclspec

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

// Spec defines the available specification types.
type Spec struct {
	// Types that are valid to be assigned to Block:
	//	*Spec_Object
	//	*Spec_Array
	//	*Spec_Attr
	//	*Spec_BlockValue
	//	*Spec_BlockAttrs
	//	*Spec_BlockList
	//	*Spec_BlockSet
	//	*Spec_BlockMap
	//	*Spec_Default
	//	*Spec_Literal
	Block                isSpec_Block `protobuf_oneof:"block"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Spec) Reset()         { *m = Spec{} }
func (m *Spec) String() string { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()    {}
func (*Spec) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{0}
}

func (m *Spec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Spec.Unmarshal(m, b)
}
func (m *Spec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Spec.Marshal(b, m, deterministic)
}
func (m *Spec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spec.Merge(m, src)
}
func (m *Spec) XXX_Size() int {
	return xxx_messageInfo_Spec.Size(m)
}
func (m *Spec) XXX_DiscardUnknown() {
	xxx_messageInfo_Spec.DiscardUnknown(m)
}

var xxx_messageInfo_Spec proto.InternalMessageInfo

type isSpec_Block interface {
	isSpec_Block()
}

type Spec_Object struct {
	Object *Object `protobuf:"bytes,1,opt,name=object,proto3,oneof"`
}

type Spec_Array struct {
	Array *Array `protobuf:"bytes,2,opt,name=array,proto3,oneof"`
}

type Spec_Attr struct {
	Attr *Attr `protobuf:"bytes,3,opt,name=Attr,proto3,oneof"`
}

type Spec_BlockValue struct {
	BlockValue *Block `protobuf:"bytes,4,opt,name=block_value,json=blockValue,proto3,oneof"`
}

type Spec_BlockAttrs struct {
	BlockAttrs *BlockAttrs `protobuf:"bytes,5,opt,name=block_attrs,json=blockAttrs,proto3,oneof"`
}

type Spec_BlockList struct {
	BlockList *BlockList `protobuf:"bytes,6,opt,name=block_list,json=blockList,proto3,oneof"`
}

type Spec_BlockSet struct {
	BlockSet *BlockSet `protobuf:"bytes,7,opt,name=block_set,json=blockSet,proto3,oneof"`
}

type Spec_BlockMap struct {
	BlockMap *BlockMap `protobuf:"bytes,8,opt,name=block_map,json=blockMap,proto3,oneof"`
}

type Spec_Default struct {
	Default *Default `protobuf:"bytes,9,opt,name=default,proto3,oneof"`
}

type Spec_Literal struct {
	Literal *Literal `protobuf:"bytes,10,opt,name=literal,proto3,oneof"`
}

func (*Spec_Object) isSpec_Block() {}

func (*Spec_Array) isSpec_Block() {}

func (*Spec_Attr) isSpec_Block() {}

func (*Spec_BlockValue) isSpec_Block() {}

func (*Spec_BlockAttrs) isSpec_Block() {}

func (*Spec_BlockList) isSpec_Block() {}

func (*Spec_BlockSet) isSpec_Block() {}

func (*Spec_BlockMap) isSpec_Block() {}

func (*Spec_Default) isSpec_Block() {}

func (*Spec_Literal) isSpec_Block() {}

func (m *Spec) GetBlock() isSpec_Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *Spec) GetObject() *Object {
	if x, ok := m.GetBlock().(*Spec_Object); ok {
		return x.Object
	}
	return nil
}

func (m *Spec) GetArray() *Array {
	if x, ok := m.GetBlock().(*Spec_Array); ok {
		return x.Array
	}
	return nil
}

func (m *Spec) GetAttr() *Attr {
	if x, ok := m.GetBlock().(*Spec_Attr); ok {
		return x.Attr
	}
	return nil
}

func (m *Spec) GetBlockValue() *Block {
	if x, ok := m.GetBlock().(*Spec_BlockValue); ok {
		return x.BlockValue
	}
	return nil
}

func (m *Spec) GetBlockAttrs() *BlockAttrs {
	if x, ok := m.GetBlock().(*Spec_BlockAttrs); ok {
		return x.BlockAttrs
	}
	return nil
}

func (m *Spec) GetBlockList() *BlockList {
	if x, ok := m.GetBlock().(*Spec_BlockList); ok {
		return x.BlockList
	}
	return nil
}

func (m *Spec) GetBlockSet() *BlockSet {
	if x, ok := m.GetBlock().(*Spec_BlockSet); ok {
		return x.BlockSet
	}
	return nil
}

func (m *Spec) GetBlockMap() *BlockMap {
	if x, ok := m.GetBlock().(*Spec_BlockMap); ok {
		return x.BlockMap
	}
	return nil
}

func (m *Spec) GetDefault() *Default {
	if x, ok := m.GetBlock().(*Spec_Default); ok {
		return x.Default
	}
	return nil
}

func (m *Spec) GetLiteral() *Literal {
	if x, ok := m.GetBlock().(*Spec_Literal); ok {
		return x.Literal
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Spec) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Spec_Object)(nil),
		(*Spec_Array)(nil),
		(*Spec_Attr)(nil),
		(*Spec_BlockValue)(nil),
		(*Spec_BlockAttrs)(nil),
		(*Spec_BlockList)(nil),
		(*Spec_BlockSet)(nil),
		(*Spec_BlockMap)(nil),
		(*Spec_Default)(nil),
		(*Spec_Literal)(nil),
	}
}

// Attr spec type reads the value of an attribute in the current body
// and returns that value as its result. It also creates validation constraints
// for the given attribute name and its value.
//
// ```hcl
// Attr {
// name     = "document_root"
// type     = string
// required = true
// }
// ```
//
// `Attr` spec blocks accept the following arguments:
//
// `name` (required) - The attribute name to expect within the HCL input file.
// This may be omitted when a default name selector is created by a parent
// `Object` spec, if the input attribute name should match the output JSON
// object property name.
//
// `type` (optional) - A [type expression](#type-expressions) that the given
// attribute value must conform to. If this argument is set, `hcldec` will
// automatically convert the given input value to this type or produce an
// error if that is not possible.
//
// `required` (optional) - If set to `true`, `hcldec` will produce an error
// if a value is not provided for the source attribute.
//
// `Attr` is a leaf spec type, so no nested spec blocks are permitted.
type Attr struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Required             bool     `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attr) Reset()         { *m = Attr{} }
func (m *Attr) String() string { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()    {}
func (*Attr) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{1}
}

func (m *Attr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attr.Unmarshal(m, b)
}
func (m *Attr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attr.Marshal(b, m, deterministic)
}
func (m *Attr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attr.Merge(m, src)
}
func (m *Attr) XXX_Size() int {
	return xxx_messageInfo_Attr.Size(m)
}
func (m *Attr) XXX_DiscardUnknown() {
	xxx_messageInfo_Attr.DiscardUnknown(m)
}

var xxx_messageInfo_Attr proto.InternalMessageInfo

func (m *Attr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Attr) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Attr) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

// Block spec type applies one nested spec block to the contents of a
// block within the current body and returns the result of that spec. It also
// creates validation constraints for the given block type name.
//
// ```hcl
// Block {
// name = "logging"
//
// Object {
// Attr "level" {
// type = string
// }
// Attr "file" {
// type = string
// }
// }
// }
// ```
//
// `Block` spec blocks accept the following arguments:
//
// `name` (required) - The block type name to expect within the HCL
// input file. This may be omitted when a default name selector is created
// by a parent `Object` spec, if the input block type name should match the
// output JSON object property name.
//
// `required` (optional) - If set to `true`, `hcldec` will produce an error
// if a block of the specified type is not present in the current body.
//
// `Block` creates a validation constraint that there must be zero or one blocks
// of the given type name, or exactly one if `required` is set.
//
// `Block` expects a single nested spec block, which is applied to the body of
// the block of the given type when it is present.
type Block struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Required             bool     `protobuf:"varint,2,opt,name=required,proto3" json:"required,omitempty"`
	Nested               *Spec    `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{2}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Block) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *Block) GetNested() *Spec {
	if m != nil {
		return m.Nested
	}
	return nil
}

// The BlockAttrs spec type is similar to an Attr spec block of a map type,
// but it produces a map from the attributes of a block rather than from an
// attribute's expression.
//
// ```hcl
// BlockAttrs {
// name     = "variables"
// type     = string
// required = false
// }
// ```
//
// This allows a map with user-defined keys to be produced within block syntax,
// but due to the constraints of that syntax it also means that the user will
// be unable to dynamically-generate either individual key names using key
// expressions or the entire map value using a `for` expression.
//
// `BlockAttrs` spec blocks accept the following arguments:
//
// `name` (required) - The block type name to expect within the HCL
// input file. This may be omitted when a default name selector is created
// by a parent `object` spec, if the input block type name should match the
// output JSON object property name.
//
// `type` (required) - The value type to require for each of the
// attributes within a matched block. The resulting value will be a JSON
// object whose property values are of this type.
//
// `required` (optional) - If `true`, an error will be produced if a block
// of the given type is not present. If `false` -- the default -- an absent
// block will be indicated by producing `null`.
type BlockAttrs struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Required             bool     `protobuf:"varint,3,opt,name=required,proto3" json:"required,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockAttrs) Reset()         { *m = BlockAttrs{} }
func (m *BlockAttrs) String() string { return proto.CompactTextString(m) }
func (*BlockAttrs) ProtoMessage()    {}
func (*BlockAttrs) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{3}
}

func (m *BlockAttrs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockAttrs.Unmarshal(m, b)
}
func (m *BlockAttrs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockAttrs.Marshal(b, m, deterministic)
}
func (m *BlockAttrs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockAttrs.Merge(m, src)
}
func (m *BlockAttrs) XXX_Size() int {
	return xxx_messageInfo_BlockAttrs.Size(m)
}
func (m *BlockAttrs) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockAttrs.DiscardUnknown(m)
}

var xxx_messageInfo_BlockAttrs proto.InternalMessageInfo

func (m *BlockAttrs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockAttrs) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BlockAttrs) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

// BlockList spec type is similar to `Block`, but it accepts zero or
// more blocks of a specified type rather than requiring zero or one. The
// result is a JSON array with one entry per block of the given type.
//
// ```hcl
// BlockList {
// name = "log_file"
//
// Object {
// Attr "level" {
// type = string
// }
// Attr "filename" {
// type     = string
// required = true
// }
// }
// }
// ```
//
// `BlockList` spec blocks accept the following arguments:
//
// `name` (required) - The block type name to expect within the HCL
// input file. This may be omitted when a default name selector is created
// by a parent `Object` spec, if the input block type name should match the
// output JSON object property name.
//
// `min_items` (optional) - If set to a number greater than zero, `hcldec` will
// produce an error if fewer than the given number of blocks are present.
//
// `max_items` (optional) - If set to a number greater than zero, `hcldec` will
// produce an error if more than the given number of blocks are present. This
// attribute must be greater than or equal to `min_items` if both are set.
//
// `Block` creates a validation constraint on the number of blocks of the given
// type that must be present.
//
// `Block` expects a single nested spec block, which is applied to the body of
// each matching block to produce the resulting list items.
type BlockList struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MinItems             uint64   `protobuf:"varint,2,opt,name=min_items,json=minItems,proto3" json:"min_items,omitempty"`
	MaxItems             uint64   `protobuf:"varint,3,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	Nested               *Spec    `protobuf:"bytes,4,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockList) Reset()         { *m = BlockList{} }
func (m *BlockList) String() string { return proto.CompactTextString(m) }
func (*BlockList) ProtoMessage()    {}
func (*BlockList) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{4}
}

func (m *BlockList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockList.Unmarshal(m, b)
}
func (m *BlockList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockList.Marshal(b, m, deterministic)
}
func (m *BlockList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockList.Merge(m, src)
}
func (m *BlockList) XXX_Size() int {
	return xxx_messageInfo_BlockList.Size(m)
}
func (m *BlockList) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockList.DiscardUnknown(m)
}

var xxx_messageInfo_BlockList proto.InternalMessageInfo

func (m *BlockList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockList) GetMinItems() uint64 {
	if m != nil {
		return m.MinItems
	}
	return 0
}

func (m *BlockList) GetMaxItems() uint64 {
	if m != nil {
		return m.MaxItems
	}
	return 0
}

func (m *BlockList) GetNested() *Spec {
	if m != nil {
		return m.Nested
	}
	return nil
}

// BlockSet spec type behaves the same as BlockList except that
// the result is in no specific order and any duplicate items are removed.
//
// ```hcl
// BlockSet {
// name = "log_file"
//
// Object {
// Attr "level" {
// type = string
// }
// Attr "filename" {
// type     = string
// required = true
// }
// }
// }
// ```
//
// The contents of `BlockSet` are the same as for `BlockList`.
type BlockSet struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MinItems             uint64   `protobuf:"varint,2,opt,name=min_items,json=minItems,proto3" json:"min_items,omitempty"`
	MaxItems             uint64   `protobuf:"varint,3,opt,name=max_items,json=maxItems,proto3" json:"max_items,omitempty"`
	Nested               *Spec    `protobuf:"bytes,4,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockSet) Reset()         { *m = BlockSet{} }
func (m *BlockSet) String() string { return proto.CompactTextString(m) }
func (*BlockSet) ProtoMessage()    {}
func (*BlockSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{5}
}

func (m *BlockSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockSet.Unmarshal(m, b)
}
func (m *BlockSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockSet.Marshal(b, m, deterministic)
}
func (m *BlockSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockSet.Merge(m, src)
}
func (m *BlockSet) XXX_Size() int {
	return xxx_messageInfo_BlockSet.Size(m)
}
func (m *BlockSet) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockSet.DiscardUnknown(m)
}

var xxx_messageInfo_BlockSet proto.InternalMessageInfo

func (m *BlockSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockSet) GetMinItems() uint64 {
	if m != nil {
		return m.MinItems
	}
	return 0
}

func (m *BlockSet) GetMaxItems() uint64 {
	if m != nil {
		return m.MaxItems
	}
	return 0
}

func (m *BlockSet) GetNested() *Spec {
	if m != nil {
		return m.Nested
	}
	return nil
}

// BlockMap spec type is similar to `Block`, but it accepts zero or
// more blocks of a specified type rather than requiring zero or one. The
// result is a JSON object, or possibly multiple nested JSON objects, whose
// properties are derived from the labels set on each matching block.
//
// ```hcl
// BlockMap {
// name = "log_file"
// labels = ["filename"]
//
// Object {
// Attr "level" {
// type     = string
// required = true
// }
// }
// }
// ```
//
// `BlockMap` spec blocks accept the following arguments:
//
// `name` (required) - The block type name to expect within the HCL
// input file. This may be omitted when a default name selector is created
// by a parent `Object` spec, if the input block type name should match the
// output JSON object property name.
//
// `labels` (required) - A list of user-oriented block label names. Each entry
// in this list creates one level of object within the output value, and
// requires one additional block header label on any child block of this type.
// Block header labels are the quoted strings that appear after the block type
// name but before the opening `{`.
//
// `Block` creates a validation constraint on the number of labels that blocks
// of the given type must have.
//
// `Block` expects a single nested spec block, which is applied to the body of
// each matching block to produce the resulting map items.
type BlockMap struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Labels               []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Nested               *Spec    `protobuf:"bytes,3,opt,name=nested,proto3" json:"nested,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockMap) Reset()         { *m = BlockMap{} }
func (m *BlockMap) String() string { return proto.CompactTextString(m) }
func (*BlockMap) ProtoMessage()    {}
func (*BlockMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{6}
}

func (m *BlockMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMap.Unmarshal(m, b)
}
func (m *BlockMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMap.Marshal(b, m, deterministic)
}
func (m *BlockMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMap.Merge(m, src)
}
func (m *BlockMap) XXX_Size() int {
	return xxx_messageInfo_BlockMap.Size(m)
}
func (m *BlockMap) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMap.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMap proto.InternalMessageInfo

func (m *BlockMap) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BlockMap) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *BlockMap) GetNested() *Spec {
	if m != nil {
		return m.Nested
	}
	return nil
}

// Literal spec type returns a given literal value, and creates no
// validation constraints. It is most commonly used with the `Default` spec
// type to create a fallback value, but can also be used e.g. to fill out
// required properties in an `Object` spec that do not correspond to any
// construct in the input configuration.
//
// ```hcl
// Literal {
// value = "hello world"
// }
// ```
//
// `Literal` spec blocks accept the following argument:
//
// `value` (required) - The value to return. This attribute may be an expression
// that uses [functions](#spec-definition-functions).
//
// `Literal` is a leaf spec type, so no nested spec blocks are permitted.
type Literal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Literal) Reset()         { *m = Literal{} }
func (m *Literal) String() string { return proto.CompactTextString(m) }
func (*Literal) ProtoMessage()    {}
func (*Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{7}
}

func (m *Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Literal.Unmarshal(m, b)
}
func (m *Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Literal.Marshal(b, m, deterministic)
}
func (m *Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Literal.Merge(m, src)
}
func (m *Literal) XXX_Size() int {
	return xxx_messageInfo_Literal.Size(m)
}
func (m *Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_Literal proto.InternalMessageInfo

func (m *Literal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Default spec type evaluates a sequence of nested specs in turn and
// returns the result of the first one that produces a non-null value.
// It creates no validation constraints of its own, but passes on the validation
// constraints from its first nested block.
//
// ```hcl
// Default {
// Attr {
// name = "private"
// type = bool
// }
// Literal {
// value = false
// }
// }
// ```
//
// A `Default` spec block must have at least one nested spec block, and should
// generally have at least two since otherwise the `Default` wrapper is a no-op.
//
// The second and any subsequent spec blocks are _fallback_ specs. These exhibit
// their usual behavior but are not able to impose validation constraints on the
// current body since they are not evaluated unless all prior specs produce
// `null` as their result.
type Default struct {
	Primary              *Spec    `protobuf:"bytes,1,opt,name=primary,proto3" json:"primary,omitempty"`
	Default              *Spec    `protobuf:"bytes,2,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Default) Reset()         { *m = Default{} }
func (m *Default) String() string { return proto.CompactTextString(m) }
func (*Default) ProtoMessage()    {}
func (*Default) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{8}
}

func (m *Default) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Default.Unmarshal(m, b)
}
func (m *Default) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Default.Marshal(b, m, deterministic)
}
func (m *Default) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Default.Merge(m, src)
}
func (m *Default) XXX_Size() int {
	return xxx_messageInfo_Default.Size(m)
}
func (m *Default) XXX_DiscardUnknown() {
	xxx_messageInfo_Default.DiscardUnknown(m)
}

var xxx_messageInfo_Default proto.InternalMessageInfo

func (m *Default) GetPrimary() *Spec {
	if m != nil {
		return m.Primary
	}
	return nil
}

func (m *Default) GetDefault() *Spec {
	if m != nil {
		return m.Default
	}
	return nil
}

// Object spec type is the most commonly used at the root of a spec file.
// Its result is a JSON object whose properties are set based on any nested
// spec blocks:
//
// ```hcl
// Object {
// Attr "name" {
// type = "string"
// }
// Block "address" {
// Object {
// Attr "street" {
// type = "string"
// }
// # ...
// }
// }
// }
// ```
//
// Nested spec blocks inside `Object` must always have an extra block label
// `"name"`, `"address"` and `"street"` in the above example) that specifies
// the name of the property that should be created in the JSON object result.
// This label also acts as a default name selector for the nested spec, allowing
// the `Attr` blocks in the above example to omit the usually-required `name`
// argument in cases where the HCL input name and JSON output name are the same.
//
// An `Object` spec block creates no validation constraints, but it passes on
// any validation constraints created by the nested specs.
type Object struct {
	Attributes           map[string]*Spec `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{9}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetAttributes() map[string]*Spec {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Array spec type produces a JSON array whose elements are set based on
// any nested spec blocks:
//
// ```hcl
// Array {
// Attr {
// name = "first_element"
// type = "string"
// }
// Attr {
// name = "second_element"
// type = "string"
// }
// }
// ```
//
// An `Array` spec block creates no validation constraints, but it passes on
// any validation constraints created by the nested specs.
type Array struct {
	Values               []*Spec  `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Array) Reset()         { *m = Array{} }
func (m *Array) String() string { return proto.CompactTextString(m) }
func (*Array) ProtoMessage()    {}
func (*Array) Descriptor() ([]byte, []int) {
	return fileDescriptor_28863966909039be, []int{10}
}

func (m *Array) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Array.Unmarshal(m, b)
}
func (m *Array) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Array.Marshal(b, m, deterministic)
}
func (m *Array) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Array.Merge(m, src)
}
func (m *Array) XXX_Size() int {
	return xxx_messageInfo_Array.Size(m)
}
func (m *Array) XXX_DiscardUnknown() {
	xxx_messageInfo_Array.DiscardUnknown(m)
}

var xxx_messageInfo_Array proto.InternalMessageInfo

func (m *Array) GetValues() []*Spec {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*Spec)(nil), "hashicorp.nomad.plugins.shared.hclspec.Spec")
	proto.RegisterType((*Attr)(nil), "hashicorp.nomad.plugins.shared.hclspec.Attr")
	proto.RegisterType((*Block)(nil), "hashicorp.nomad.plugins.shared.hclspec.Block")
	proto.RegisterType((*BlockAttrs)(nil), "hashicorp.nomad.plugins.shared.hclspec.BlockAttrs")
	proto.RegisterType((*BlockList)(nil), "hashicorp.nomad.plugins.shared.hclspec.BlockList")
	proto.RegisterType((*BlockSet)(nil), "hashicorp.nomad.plugins.shared.hclspec.BlockSet")
	proto.RegisterType((*BlockMap)(nil), "hashicorp.nomad.plugins.shared.hclspec.BlockMap")
	proto.RegisterType((*Literal)(nil), "hashicorp.nomad.plugins.shared.hclspec.Literal")
	proto.RegisterType((*Default)(nil), "hashicorp.nomad.plugins.shared.hclspec.Default")
	proto.RegisterType((*Object)(nil), "hashicorp.nomad.plugins.shared.hclspec.Object")
	proto.RegisterMapType((map[string]*Spec)(nil), "hashicorp.nomad.plugins.shared.hclspec.Object.AttributesEntry")
	proto.RegisterType((*Array)(nil), "hashicorp.nomad.plugins.shared.hclspec.Array")
}

func init() {
	proto.RegisterFile("plugins/shared/hclspec/hcl_spec.proto", fileDescriptor_28863966909039be)
}

var fileDescriptor_28863966909039be = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x96, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0xc4, 0xaf, 0xd3, 0xc3, 0xf3, 0x68, 0x85, 0x90, 0x55, 0x0e, 0x54, 0x96, 0x40,
	0x3d, 0x80, 0x0b, 0xe5, 0x82, 0x38, 0x20, 0x35, 0x6a, 0x91, 0x81, 0x46, 0xad, 0xb6, 0x82, 0x03,
	0x07, 0xa2, 0xb5, 0xb3, 0x10, 0x13, 0xbf, 0xb1, 0xbb, 0x41, 0x8d, 0x04, 0x1f, 0x84, 0x03, 0xf0,
	0xa9, 0xf8, 0x3e, 0x68, 0x5f, 0x9c, 0x14, 0x94, 0x43, 0x1c, 0x7a, 0xe0, 0x94, 0x19, 0x8f, 0xfe,
	0x3f, 0xcf, 0xec, 0xce, 0x78, 0x02, 0x77, 0x9a, 0x62, 0xfe, 0x3e, 0xaf, 0xf8, 0x01, 0x9f, 0x12,
	0x46, 0x27, 0x07, 0xd3, 0xac, 0xe0, 0x0d, 0xcd, 0xe4, 0xef, 0x58, 0x1a, 0x71, 0xc3, 0x6a, 0x51,
	0xa3, 0xbb, 0x53, 0xc2, 0xa7, 0x79, 0x56, 0xb3, 0x26, 0xae, 0xea, 0x92, 0x4c, 0x62, 0x23, 0x8b,
	0xb5, 0x2c, 0x36, 0xb2, 0xe8, 0x9b, 0x0b, 0xf6, 0x45, 0x43, 0x33, 0x94, 0x80, 0x5b, 0xa7, 0x1f,
	0x68, 0x26, 0x42, 0x6b, 0xcf, 0xda, 0xdf, 0x39, 0x8c, 0xe3, 0xcd, 0x08, 0xf1, 0x99, 0x52, 0x25,
	0x3d, 0x6c, 0xf4, 0xe8, 0x04, 0x1c, 0xc2, 0x18, 0x59, 0x84, 0x7d, 0x05, 0xba, 0xbf, 0x29, 0xe8,
	0x48, 0x8a, 0x92, 0x1e, 0xd6, 0x6a, 0x34, 0x04, 0xfb, 0x48, 0x08, 0x16, 0x0e, 0x14, 0xe5, 0xde,
	0xc6, 0x14, 0x21, 0x58, 0xd2, 0xc3, 0x4a, 0x8b, 0xce, 0x61, 0x27, 0x2d, 0xea, 0x6c, 0x36, 0xfe,
	0x44, 0x8a, 0x39, 0x0d, 0xed, 0x6e, 0x09, 0x0d, 0xa5, 0x34, 0xe9, 0x61, 0x50, 0x8c, 0xd7, 0x12,
	0x81, 0x5e, 0xb5, 0x44, 0x22, 0x04, 0xe3, 0xa1, 0xa3, 0x88, 0x87, 0x9d, 0x88, 0x32, 0x33, 0xbe,
	0xc4, 0x2a, 0x0f, 0x61, 0xd0, 0xde, 0xb8, 0xc8, 0xb9, 0x08, 0x5d, 0x45, 0x7d, 0xd8, 0x89, 0x7a,
	0x9a, 0x73, 0x79, 0x09, 0x41, 0xda, 0x3a, 0xe8, 0x0c, 0xb4, 0x33, 0xe6, 0x54, 0x84, 0x9e, 0x42,
	0x3e, 0xe8, 0x84, 0xbc, 0xa0, 0x92, 0xe8, 0xa7, 0xc6, 0x5e, 0x01, 0x4b, 0xd2, 0x84, 0xfe, 0x16,
	0xc0, 0x11, 0x69, 0x96, 0xc0, 0x11, 0x69, 0xd0, 0x4b, 0xf0, 0x26, 0xf4, 0x1d, 0x99, 0x17, 0x22,
	0x0c, 0x14, 0xee, 0x60, 0x53, 0xdc, 0xb1, 0x96, 0x25, 0x3d, 0xdc, 0x12, 0x24, 0xac, 0xc8, 0x05,
	0x65, 0xa4, 0x08, 0xa1, 0x1b, 0xec, 0x54, 0xcb, 0x24, 0xcc, 0x10, 0x86, 0x1e, 0x38, 0x2a, 0xcb,
	0xe8, 0x85, 0xee, 0x42, 0x84, 0xc0, 0xae, 0x48, 0x49, 0xd5, 0x70, 0x04, 0x58, 0xd9, 0xf2, 0x99,
	0x58, 0x34, 0x54, 0xf5, 0x79, 0x80, 0x95, 0x8d, 0x76, 0xc1, 0x67, 0xf4, 0xe3, 0x3c, 0x67, 0x74,
	0xa2, 0x3a, 0xd7, 0xc7, 0x4b, 0x3f, 0xfa, 0x02, 0x8e, 0x3a, 0x86, 0xb5, 0xb0, 0xab, 0xc2, 0xfe,
	0xef, 0x42, 0x74, 0x0c, 0x6e, 0x45, 0xb9, 0x30, 0xc8, 0x0e, 0xc3, 0x20, 0x27, 0x1b, 0x1b, 0x6d,
	0x74, 0x0e, 0xb0, 0xea, 0xbf, 0x6b, 0x29, 0xe8, 0x87, 0x05, 0xc1, 0xb2, 0xf9, 0xd6, 0x12, 0x6f,
	0x41, 0x50, 0xe6, 0xd5, 0x38, 0x17, 0xb4, 0xe4, 0x0a, 0x6b, 0x63, 0xbf, 0xcc, 0xab, 0xe7, 0xd2,
	0x57, 0x41, 0x72, 0x69, 0x82, 0x03, 0x13, 0x24, 0x97, 0x3a, 0xb8, 0xaa, 0xd9, 0xfe, 0x8b, 0x9a,
	0xbf, 0x5b, 0xe0, 0xb7, 0xbd, 0xfc, 0x4f, 0x26, 0xf8, 0xd9, 0xe4, 0x27, 0xc7, 0x61, 0x5d, 0x7e,
	0x37, 0xc1, 0x2d, 0x48, 0x4a, 0x0b, 0x99, 0xdc, 0x60, 0x3f, 0xc0, 0xc6, 0xbb, 0xa6, 0x96, 0xb8,
	0x0d, 0x9e, 0x69, 0x7e, 0x74, 0x03, 0x1c, 0xfd, 0x91, 0xd4, 0x6f, 0xd7, 0x4e, 0xf4, 0xd5, 0x02,
	0xcf, 0xcc, 0x1a, 0x7a, 0x06, 0x5e, 0xc3, 0xf2, 0x92, 0xb0, 0x85, 0x59, 0x11, 0xdd, 0xde, 0xd9,
	0x8a, 0x25, 0xa7, 0x9d, 0xfa, 0xfe, 0x36, 0x1c, 0x23, 0x8e, 0x7e, 0x5a, 0xe0, 0xea, 0xe5, 0x83,
	0xde, 0x02, 0xc8, 0xef, 0x71, 0x9e, 0xce, 0x05, 0xe5, 0xa1, 0xb5, 0x37, 0xd8, 0xdf, 0x39, 0x7c,
	0xda, 0x6d, 0x81, 0xa9, 0xc5, 0xa1, 0x01, 0x27, 0x95, 0x60, 0x0b, 0x7c, 0x85, 0xb8, 0x3b, 0x83,
	0xff, 0xfe, 0x08, 0xa3, 0xff, 0x61, 0x30, 0xa3, 0x0b, 0x73, 0x5a, 0xd2, 0x44, 0xc3, 0xf6, 0x04,
	0xb7, 0xa9, 0x4a, 0x4b, 0x9f, 0xf4, 0x1f, 0x5b, 0xd1, 0x08, 0x1c, 0xb5, 0x0a, 0xe5, 0x1d, 0xab,
	0xa7, 0x6d, 0x45, 0x1d, 0xef, 0x58, 0x6b, 0x87, 0xc1, 0x1b, 0xcf, 0x3c, 0x4f, 0x5d, 0xf5, 0xdf,
	0xe0, 0xd1, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x20, 0x9f, 0xf2, 0x44, 0x08, 0x00, 0x00,
}
