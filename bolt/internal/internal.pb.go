// Code generated by protoc-gen-gogo.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	Source
	Dashboard
	DashboardCell
	Axis
	Template
	TemplateValue
	TemplateQuery
	Server
	Layout
	Cell
	Query
	TimeShift
	Range
	AlertRule
	User
	Role
	Organization
*/
package internal

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Source struct {
	ID                 int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name               string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Type               string `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	Username           string `protobuf:"bytes,4,opt,name=Username,proto3" json:"Username,omitempty"`
	Password           string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	URL                string `protobuf:"bytes,6,opt,name=URL,proto3" json:"URL,omitempty"`
	Default            bool   `protobuf:"varint,7,opt,name=Default,proto3" json:"Default,omitempty"`
	Telegraf           string `protobuf:"bytes,8,opt,name=Telegraf,proto3" json:"Telegraf,omitempty"`
	InsecureSkipVerify bool   `protobuf:"varint,9,opt,name=InsecureSkipVerify,proto3" json:"InsecureSkipVerify,omitempty"`
	MetaURL            string `protobuf:"bytes,10,opt,name=MetaURL,proto3" json:"MetaURL,omitempty"`
	SharedSecret       string `protobuf:"bytes,11,opt,name=SharedSecret,proto3" json:"SharedSecret,omitempty"`
	Organization       string `protobuf:"bytes,12,opt,name=Organization,proto3" json:"Organization,omitempty"`
	Role               string `protobuf:"bytes,13,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

type Dashboard struct {
	ID           int64            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string           `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Cells        []*DashboardCell `protobuf:"bytes,3,rep,name=cells" json:"cells,omitempty"`
	Templates    []*Template      `protobuf:"bytes,4,rep,name=templates" json:"templates,omitempty"`
	Organization string           `protobuf:"bytes,5,opt,name=Organization,proto3" json:"Organization,omitempty"`
}

func (m *Dashboard) Reset()                    { *m = Dashboard{} }
func (m *Dashboard) String() string            { return proto.CompactTextString(m) }
func (*Dashboard) ProtoMessage()               {}
func (*Dashboard) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Dashboard) GetCells() []*DashboardCell {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (m *Dashboard) GetTemplates() []*Template {
	if m != nil {
		return m.Templates
	}
	return nil
}

type DashboardCell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	Name    string           `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Type    string           `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	ID      string           `protobuf:"bytes,8,opt,name=ID,proto3" json:"ID,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,9,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *DashboardCell) Reset()                    { *m = DashboardCell{} }
func (m *DashboardCell) String() string            { return proto.CompactTextString(m) }
func (*DashboardCell) ProtoMessage()               {}
func (*DashboardCell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *DashboardCell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *DashboardCell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

type Axis struct {
	LegacyBounds []int64  `protobuf:"varint,1,rep,name=legacyBounds" json:"legacyBounds,omitempty"`
	Bounds       []string `protobuf:"bytes,2,rep,name=bounds" json:"bounds,omitempty"`
	Label        string   `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Prefix       string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix       string   `protobuf:"bytes,5,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Base         string   `protobuf:"bytes,6,opt,name=base,proto3" json:"base,omitempty"`
	Scale        string   `protobuf:"bytes,7,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (m *Axis) Reset()                    { *m = Axis{} }
func (m *Axis) String() string            { return proto.CompactTextString(m) }
func (*Axis) ProtoMessage()               {}
func (*Axis) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

type Template struct {
	ID      string           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	TempVar string           `protobuf:"bytes,2,opt,name=temp_var,json=tempVar,proto3" json:"temp_var,omitempty"`
	Values  []*TemplateValue `protobuf:"bytes,3,rep,name=values" json:"values,omitempty"`
	Type    string           `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Label   string           `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	Query   *TemplateQuery   `protobuf:"bytes,6,opt,name=query" json:"query,omitempty"`
}

func (m *Template) Reset()                    { *m = Template{} }
func (m *Template) String() string            { return proto.CompactTextString(m) }
func (*Template) ProtoMessage()               {}
func (*Template) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Template) GetValues() []*TemplateValue {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Template) GetQuery() *TemplateQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

type TemplateValue struct {
	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value    string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Selected bool   `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
}

func (m *TemplateValue) Reset()                    { *m = TemplateValue{} }
func (m *TemplateValue) String() string            { return proto.CompactTextString(m) }
func (*TemplateValue) ProtoMessage()               {}
func (*TemplateValue) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

type TemplateQuery struct {
	Command     string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Db          string `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
	Rp          string `protobuf:"bytes,3,opt,name=rp,proto3" json:"rp,omitempty"`
	Measurement string `protobuf:"bytes,4,opt,name=measurement,proto3" json:"measurement,omitempty"`
	TagKey      string `protobuf:"bytes,5,opt,name=tag_key,json=tagKey,proto3" json:"tag_key,omitempty"`
	FieldKey    string `protobuf:"bytes,6,opt,name=field_key,json=fieldKey,proto3" json:"field_key,omitempty"`
}

func (m *TemplateQuery) Reset()                    { *m = TemplateQuery{} }
func (m *TemplateQuery) String() string            { return proto.CompactTextString(m) }
func (*TemplateQuery) ProtoMessage()               {}
func (*TemplateQuery) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

type Server struct {
	ID           int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Username     string `protobuf:"bytes,3,opt,name=Username,proto3" json:"Username,omitempty"`
	Password     string `protobuf:"bytes,4,opt,name=Password,proto3" json:"Password,omitempty"`
	URL          string `protobuf:"bytes,5,opt,name=URL,proto3" json:"URL,omitempty"`
	SrcID        int64  `protobuf:"varint,6,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	Active       bool   `protobuf:"varint,7,opt,name=Active,proto3" json:"Active,omitempty"`
	Organization string `protobuf:"bytes,8,opt,name=Organization,proto3" json:"Organization,omitempty"`
}

func (m *Server) Reset()                    { *m = Server{} }
func (m *Server) String() string            { return proto.CompactTextString(m) }
func (*Server) ProtoMessage()               {}
func (*Server) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

type Layout struct {
	ID           string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Application  string  `protobuf:"bytes,2,opt,name=Application,proto3" json:"Application,omitempty"`
	Measurement  string  `protobuf:"bytes,3,opt,name=Measurement,proto3" json:"Measurement,omitempty"`
	Cells        []*Cell `protobuf:"bytes,4,rep,name=Cells" json:"Cells,omitempty"`
	Autoflow     bool    `protobuf:"varint,5,opt,name=Autoflow,proto3" json:"Autoflow,omitempty"`
	Organization string  `protobuf:"bytes,6,opt,name=Organization,proto3" json:"Organization,omitempty"`
}

func (m *Layout) Reset()                    { *m = Layout{} }
func (m *Layout) String() string            { return proto.CompactTextString(m) }
func (*Layout) ProtoMessage()               {}
func (*Layout) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{8} }

func (m *Layout) GetCells() []*Cell {
	if m != nil {
		return m.Cells
	}
	return nil
}

type Cell struct {
	X       int32            `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y       int32            `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W       int32            `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H       int32            `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
	Queries []*Query         `protobuf:"bytes,5,rep,name=queries" json:"queries,omitempty"`
	I       string           `protobuf:"bytes,6,opt,name=i,proto3" json:"i,omitempty"`
	Name    string           `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	Yranges []int64          `protobuf:"varint,8,rep,name=yranges" json:"yranges,omitempty"`
	Ylabels []string         `protobuf:"bytes,9,rep,name=ylabels" json:"ylabels,omitempty"`
	Type    string           `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Axes    map[string]*Axis `protobuf:"bytes,11,rep,name=axes" json:"axes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Cell) Reset()                    { *m = Cell{} }
func (m *Cell) String() string            { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()               {}
func (*Cell) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{9} }

func (m *Cell) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

func (m *Cell) GetAxes() map[string]*Axis {
	if m != nil {
		return m.Axes
	}
	return nil
}

type Query struct {
	Command  string       `protobuf:"bytes,1,opt,name=Command,proto3" json:"Command,omitempty"`
	DB       string       `protobuf:"bytes,2,opt,name=DB,proto3" json:"DB,omitempty"`
	RP       string       `protobuf:"bytes,3,opt,name=RP,proto3" json:"RP,omitempty"`
	GroupBys []string     `protobuf:"bytes,4,rep,name=GroupBys" json:"GroupBys,omitempty"`
	Wheres   []string     `protobuf:"bytes,5,rep,name=Wheres" json:"Wheres,omitempty"`
	Label    string       `protobuf:"bytes,6,opt,name=Label,proto3" json:"Label,omitempty"`
	Range    *Range       `protobuf:"bytes,7,opt,name=Range" json:"Range,omitempty"`
	Source   string       `protobuf:"bytes,8,opt,name=Source,proto3" json:"Source,omitempty"`
	Shifts   []*TimeShift `protobuf:"bytes,9,rep,name=Shifts" json:"Shifts,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{10} }

func (m *Query) GetRange() *Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Query) GetShifts() []*TimeShift {
	if m != nil {
		return m.Shifts
	}
	return nil
}

type TimeShift struct {
	Label    string `protobuf:"bytes,1,opt,name=Label,proto3" json:"Label,omitempty"`
	Unit     string `protobuf:"bytes,2,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Quantity string `protobuf:"bytes,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (m *TimeShift) Reset()                    { *m = TimeShift{} }
func (m *TimeShift) String() string            { return proto.CompactTextString(m) }
func (*TimeShift) ProtoMessage()               {}
func (*TimeShift) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{11} }

type Range struct {
	Upper int64 `protobuf:"varint,1,opt,name=Upper,proto3" json:"Upper,omitempty"`
	Lower int64 `protobuf:"varint,2,opt,name=Lower,proto3" json:"Lower,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{12} }

type AlertRule struct {
	ID     string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	JSON   string `protobuf:"bytes,2,opt,name=JSON,proto3" json:"JSON,omitempty"`
	SrcID  int64  `protobuf:"varint,3,opt,name=SrcID,proto3" json:"SrcID,omitempty"`
	KapaID int64  `protobuf:"varint,4,opt,name=KapaID,proto3" json:"KapaID,omitempty"`
}

func (m *AlertRule) Reset()                    { *m = AlertRule{} }
func (m *AlertRule) String() string            { return proto.CompactTextString(m) }
func (*AlertRule) ProtoMessage()               {}
func (*AlertRule) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{13} }

type User struct {
	ID         uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Provider   string  `protobuf:"bytes,3,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Scheme     string  `protobuf:"bytes,4,opt,name=Scheme,proto3" json:"Scheme,omitempty"`
	Roles      []*Role `protobuf:"bytes,5,rep,name=Roles" json:"Roles,omitempty"`
	SuperAdmin bool    `protobuf:"varint,6,opt,name=SuperAdmin,proto3" json:"SuperAdmin,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{14} }

func (m *User) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Role struct {
	Organization string `protobuf:"bytes,1,opt,name=Organization,proto3" json:"Organization,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *Role) Reset()                    { *m = Role{} }
func (m *Role) String() string            { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()               {}
func (*Role) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{15} }

type Organization struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	DefaultRole string `protobuf:"bytes,3,opt,name=DefaultRole,proto3" json:"DefaultRole,omitempty"`
	Public      bool   `protobuf:"varint,4,opt,name=Public,proto3" json:"Public,omitempty"`
}

func (m *Organization) Reset()                    { *m = Organization{} }
func (m *Organization) String() string            { return proto.CompactTextString(m) }
func (*Organization) ProtoMessage()               {}
func (*Organization) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{16} }

func init() {
	proto.RegisterType((*Source)(nil), "internal.Source")
	proto.RegisterType((*Dashboard)(nil), "internal.Dashboard")
	proto.RegisterType((*DashboardCell)(nil), "internal.DashboardCell")
	proto.RegisterType((*Axis)(nil), "internal.Axis")
	proto.RegisterType((*Template)(nil), "internal.Template")
	proto.RegisterType((*TemplateValue)(nil), "internal.TemplateValue")
	proto.RegisterType((*TemplateQuery)(nil), "internal.TemplateQuery")
	proto.RegisterType((*Server)(nil), "internal.Server")
	proto.RegisterType((*Layout)(nil), "internal.Layout")
	proto.RegisterType((*Cell)(nil), "internal.Cell")
	proto.RegisterType((*Query)(nil), "internal.Query")
	proto.RegisterType((*TimeShift)(nil), "internal.TimeShift")
	proto.RegisterType((*Range)(nil), "internal.Range")
	proto.RegisterType((*AlertRule)(nil), "internal.AlertRule")
	proto.RegisterType((*User)(nil), "internal.User")
	proto.RegisterType((*Role)(nil), "internal.Role")
	proto.RegisterType((*Organization)(nil), "internal.Organization")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 1211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0x5f, 0x8f, 0xdb, 0x44,
	0x10, 0xd7, 0xc6, 0x71, 0x12, 0x4f, 0xae, 0x05, 0x2d, 0x15, 0x35, 0x45, 0x42, 0xc1, 0x02, 0xe9,
	0x10, 0xf4, 0x40, 0xad, 0x90, 0x10, 0x0f, 0x48, 0xb9, 0x0b, 0xaa, 0x8e, 0xfe, 0xbb, 0x6e, 0x7a,
	0xe5, 0x09, 0x55, 0x1b, 0x67, 0x72, 0xb1, 0xea, 0xd8, 0x66, 0x6d, 0xdf, 0xc5, 0x7c, 0x1b, 0x24,
	0x9e, 0x78, 0x44, 0xbc, 0x23, 0xf1, 0x84, 0xf8, 0x20, 0x7c, 0x05, 0x5e, 0xd1, 0xec, 0xae, 0x1d,
	0xa7, 0x09, 0x55, 0x5f, 0xe0, 0x6d, 0x7f, 0x33, 0xeb, 0xd9, 0x9d, 0x99, 0xdf, 0xfc, 0xbc, 0x70,
	0x3d, 0x4a, 0x0a, 0x54, 0x89, 0x8c, 0x8f, 0x32, 0x95, 0x16, 0x29, 0x1f, 0xd4, 0x38, 0xf8, 0xab,
	0x03, 0xbd, 0x69, 0x5a, 0xaa, 0x10, 0xf9, 0x75, 0xe8, 0x9c, 0x4e, 0x7c, 0x36, 0x62, 0x87, 0x8e,
	0xe8, 0x9c, 0x4e, 0x38, 0x87, 0xee, 0x23, 0xb9, 0x42, 0xbf, 0x33, 0x62, 0x87, 0x9e, 0xd0, 0x6b,
	0xb2, 0x3d, 0xad, 0x32, 0xf4, 0x1d, 0x63, 0xa3, 0x35, 0xbf, 0x05, 0x83, 0xf3, 0x9c, 0xa2, 0xad,
	0xd0, 0xef, 0x6a, 0x7b, 0x83, 0xc9, 0x77, 0x26, 0xf3, 0xfc, 0x2a, 0x55, 0x73, 0xdf, 0x35, 0xbe,
	0x1a, 0xf3, 0x37, 0xc1, 0x39, 0x17, 0x0f, 0xfc, 0x9e, 0x36, 0xd3, 0x92, 0xfb, 0xd0, 0x9f, 0xe0,
	0x42, 0x96, 0x71, 0xe1, 0xf7, 0x47, 0xec, 0x70, 0x20, 0x6a, 0x48, 0x71, 0x9e, 0x62, 0x8c, 0x17,
	0x4a, 0x2e, 0xfc, 0x81, 0x89, 0x53, 0x63, 0x7e, 0x04, 0xfc, 0x34, 0xc9, 0x31, 0x2c, 0x15, 0x4e,
	0x5f, 0x44, 0xd9, 0x33, 0x54, 0xd1, 0xa2, 0xf2, 0x3d, 0x1d, 0x60, 0x8f, 0x87, 0x4e, 0x79, 0x88,
	0x85, 0xa4, 0xb3, 0x41, 0x87, 0xaa, 0x21, 0x0f, 0xe0, 0x60, 0xba, 0x94, 0x0a, 0xe7, 0x53, 0x0c,
	0x15, 0x16, 0xfe, 0x50, 0xbb, 0xb7, 0x6c, 0xb4, 0xe7, 0xb1, 0xba, 0x90, 0x49, 0xf4, 0x83, 0x2c,
	0xa2, 0x34, 0xf1, 0x0f, 0xcc, 0x9e, 0xb6, 0x8d, 0xaa, 0x24, 0xd2, 0x18, 0xfd, 0x6b, 0xa6, 0x4a,
	0xb4, 0x0e, 0x7e, 0x65, 0xe0, 0x4d, 0x64, 0xbe, 0x9c, 0xa5, 0x52, 0xcd, 0x5f, 0xab, 0xd6, 0xb7,
	0xc1, 0x0d, 0x31, 0x8e, 0x73, 0xdf, 0x19, 0x39, 0x87, 0xc3, 0x3b, 0x37, 0x8f, 0x9a, 0x26, 0x36,
	0x71, 0x4e, 0x30, 0x8e, 0x85, 0xd9, 0xc5, 0x3f, 0x03, 0xaf, 0xc0, 0x55, 0x16, 0xcb, 0x02, 0x73,
	0xbf, 0xab, 0x3f, 0xe1, 0x9b, 0x4f, 0x9e, 0x5a, 0x97, 0xd8, 0x6c, 0xda, 0x49, 0xc5, 0xdd, 0x4d,
	0x25, 0xf8, 0xb9, 0x03, 0xd7, 0xb6, 0x8e, 0xe3, 0x07, 0xc0, 0xd6, 0xfa, 0xe6, 0xae, 0x60, 0x6b,
	0x42, 0x95, 0xbe, 0xb5, 0x2b, 0x58, 0x45, 0xe8, 0x4a, 0x73, 0xc3, 0x15, 0xec, 0x8a, 0xd0, 0x52,
	0x33, 0xc2, 0x15, 0x6c, 0xc9, 0x3f, 0x82, 0xfe, 0xf7, 0x25, 0xaa, 0x08, 0x73, 0xdf, 0xd5, 0xb7,
	0x7b, 0x63, 0x73, 0xbb, 0x27, 0x25, 0xaa, 0x4a, 0xd4, 0x7e, 0xaa, 0x86, 0x66, 0x93, 0xa1, 0x86,
	0x5e, 0x93, 0xad, 0x20, 0xe6, 0xf5, 0x8d, 0x8d, 0xd6, 0xb6, 0x8a, 0x86, 0x0f, 0x54, 0xc5, 0xcf,
	0xa1, 0x2b, 0xd7, 0x98, 0xfb, 0x9e, 0x8e, 0xff, 0xfe, 0xbf, 0x14, 0xec, 0x68, 0xbc, 0xc6, 0xfc,
	0xeb, 0xa4, 0x50, 0x95, 0xd0, 0xdb, 0x6f, 0xdd, 0x03, 0xaf, 0x31, 0x11, 0x2b, 0x5f, 0x60, 0xa5,
	0x13, 0xf4, 0x04, 0x2d, 0xf9, 0x07, 0xe0, 0x5e, 0xca, 0xb8, 0x34, 0xcd, 0x19, 0xde, 0xb9, 0xbe,
	0x09, 0x3b, 0x5e, 0x47, 0xb9, 0x30, 0xce, 0x2f, 0x3b, 0x5f, 0xb0, 0xe0, 0x17, 0x06, 0x5d, 0xb2,
	0x51, 0x65, 0x63, 0xbc, 0x90, 0x61, 0x75, 0x9c, 0x96, 0xc9, 0x3c, 0xf7, 0xd9, 0xc8, 0x39, 0x74,
	0xc4, 0x96, 0x8d, 0xbf, 0x0d, 0xbd, 0x99, 0xf1, 0x76, 0x46, 0xce, 0xa1, 0x27, 0x2c, 0xe2, 0x37,
	0xc0, 0x8d, 0xe5, 0x0c, 0x63, 0x3b, 0x63, 0x06, 0xd0, 0xee, 0x4c, 0xe1, 0x22, 0x5a, 0xdb, 0x11,
	0xb3, 0x88, 0xec, 0x79, 0xb9, 0x20, 0xbb, 0xe9, 0x9e, 0x45, 0x54, 0xae, 0x99, 0xcc, 0x9b, 0x12,
	0xd2, 0x9a, 0x22, 0xe7, 0xa1, 0x8c, 0xeb, 0x1a, 0x1a, 0x10, 0xfc, 0xc6, 0x68, 0xb6, 0x0c, 0x27,
	0x5a, 0xbc, 0x34, 0x15, 0x7d, 0x07, 0x06, 0xc4, 0x97, 0xe7, 0x97, 0x52, 0x59, 0x6e, 0xf6, 0x09,
	0x3f, 0x93, 0x8a, 0x7f, 0x0a, 0x3d, 0x9d, 0xf9, 0x1e, 0x7e, 0xd6, 0xe1, 0x9e, 0x91, 0x5f, 0xd8,
	0x6d, 0x4d, 0x07, 0xbb, 0xad, 0x0e, 0x36, 0xc9, 0xba, 0xed, 0x64, 0x6f, 0x83, 0x4b, 0x54, 0xa8,
	0xf4, 0xed, 0xf7, 0x46, 0x36, 0x84, 0x31, 0xbb, 0x82, 0x73, 0xb8, 0xb6, 0x75, 0x62, 0x73, 0x12,
	0xdb, 0x3e, 0x69, 0xd3, 0x45, 0xcf, 0x76, 0x8d, 0x74, 0x25, 0xc7, 0x18, 0xc3, 0x02, 0xe7, 0xba,
	0xde, 0x03, 0xd1, 0xe0, 0xe0, 0x47, 0xb6, 0x89, 0xab, 0xcf, 0x23, 0xe5, 0x08, 0xd3, 0xd5, 0x4a,
	0x26, 0x73, 0x1b, 0xba, 0x86, 0x54, 0xb7, 0xf9, 0xcc, 0x86, 0xee, 0xcc, 0x67, 0x84, 0x55, 0x66,
	0x3b, 0xd8, 0x51, 0x19, 0x1f, 0xc1, 0x70, 0x85, 0x32, 0x2f, 0x15, 0xae, 0x30, 0x29, 0x6c, 0x09,
	0xda, 0x26, 0x7e, 0x13, 0xfa, 0x85, 0xbc, 0x78, 0x4e, 0xdc, 0xb3, 0x9d, 0x2c, 0xe4, 0xc5, 0x7d,
	0xac, 0xf8, 0xbb, 0xe0, 0x2d, 0x22, 0x8c, 0xe7, 0xda, 0x65, 0xda, 0x39, 0xd0, 0x86, 0xfb, 0x58,
	0x05, 0x7f, 0x30, 0xe8, 0x4d, 0x51, 0x5d, 0xa2, 0x7a, 0x2d, 0x49, 0x69, 0x4b, 0xb5, 0xf3, 0x0a,
	0xa9, 0xee, 0xee, 0x97, 0x6a, 0x77, 0x23, 0xd5, 0x37, 0xc0, 0x9d, 0xaa, 0xf0, 0x74, 0xa2, 0x6f,
	0xe4, 0x08, 0x03, 0x88, 0x8d, 0xe3, 0xb0, 0x88, 0x2e, 0xd1, 0xea, 0xb7, 0x45, 0x3b, 0x4a, 0x33,
	0xd8, 0xa3, 0x34, 0xbf, 0x33, 0xe8, 0x3d, 0x90, 0x55, 0x5a, 0x16, 0x3b, 0x2c, 0x1c, 0xc1, 0x70,
	0x9c, 0x65, 0x71, 0x14, 0x9a, 0xaf, 0x4d, 0x46, 0x6d, 0x13, 0xed, 0x78, 0xd8, 0xaa, 0xaf, 0xc9,
	0xad, 0x6d, 0xa2, 0x29, 0x3e, 0xd1, 0x6a, 0x6a, 0xa4, 0xb1, 0x35, 0xc5, 0x46, 0x44, 0xb5, 0x93,
	0x8a, 0x30, 0x2e, 0x8b, 0x74, 0x11, 0xa7, 0x57, 0x3a, 0xdb, 0x81, 0x68, 0xf0, 0x4e, 0x12, 0xbd,
	0x3d, 0x49, 0xfc, 0xd9, 0x81, 0xee, 0xff, 0xa5, 0x92, 0x07, 0xc0, 0x22, 0x7b, 0x09, 0x16, 0x35,
	0x9a, 0xd9, 0x6f, 0x69, 0xa6, 0x0f, 0xfd, 0x4a, 0xc9, 0xe4, 0x02, 0x73, 0x7f, 0xa0, 0x15, 0xa8,
	0x86, 0xda, 0xa3, 0x67, 0xcd, 0x88, 0xa5, 0x27, 0x6a, 0xd8, 0xcc, 0x0e, 0xb4, 0x66, 0xe7, 0x13,
	0xab, 0xab, 0x43, 0x7d, 0x23, 0x7f, 0xbb, 0x74, 0xff, 0x9d, 0x9c, 0xfe, 0xcd, 0xc0, 0x6d, 0x06,
	0xef, 0x64, 0x7b, 0xf0, 0x4e, 0x36, 0x83, 0x37, 0x39, 0xae, 0x07, 0x6f, 0x72, 0x4c, 0x58, 0x9c,
	0xd5, 0x83, 0x27, 0xce, 0xa8, 0xa1, 0xf7, 0x54, 0x5a, 0x66, 0xc7, 0x95, 0xe9, 0xbc, 0x27, 0x1a,
	0x4c, 0x6c, 0xfd, 0x76, 0x89, 0xca, 0x96, 0xda, 0x13, 0x16, 0x11, 0xb7, 0x1f, 0x68, 0x51, 0x32,
	0xc5, 0x35, 0x80, 0x7f, 0x08, 0xae, 0xa0, 0xe2, 0xe9, 0x0a, 0x6f, 0xf5, 0x45, 0x9b, 0x85, 0xf1,
	0x52, 0x50, 0xf3, 0x9e, 0xb2, 0x24, 0xaf, 0x5f, 0x57, 0x1f, 0x43, 0x6f, 0xba, 0x8c, 0x16, 0x45,
	0xfd, 0x77, 0x7a, 0xab, 0x25, 0x6a, 0xd1, 0x0a, 0xb5, 0x4f, 0xd8, 0x2d, 0xc1, 0x13, 0xf0, 0x1a,
	0xe3, 0xe6, 0x3a, 0xac, 0x7d, 0x1d, 0x0e, 0xdd, 0xf3, 0x24, 0x2a, 0xea, 0xf1, 0xa6, 0x35, 0x25,
	0xfb, 0xa4, 0x94, 0x49, 0x11, 0x15, 0x55, 0x3d, 0xde, 0x35, 0x0e, 0xee, 0xda, 0xeb, 0x53, 0xb8,
	0xf3, 0x2c, 0x43, 0x65, 0xa5, 0xc2, 0x00, 0x7d, 0x48, 0x7a, 0x85, 0x46, 0xe5, 0x1d, 0x61, 0x40,
	0xf0, 0x1d, 0x78, 0xe3, 0x18, 0x55, 0x21, 0xca, 0x78, 0xf7, 0xdf, 0xc0, 0xa1, 0xfb, 0xcd, 0xf4,
	0xf1, 0xa3, 0xfa, 0x06, 0xb4, 0xde, 0xc8, 0x82, 0xf3, 0x92, 0x2c, 0xdc, 0x97, 0x99, 0x3c, 0x9d,
	0x68, 0x9e, 0x3b, 0xc2, 0xa2, 0xe0, 0x27, 0x06, 0x5d, 0xd2, 0x9f, 0x56, 0xe8, 0xee, 0xab, 0xb4,
	0xeb, 0x4c, 0xa5, 0x97, 0xd1, 0x1c, 0x55, 0x9d, 0x5c, 0x8d, 0x75, 0xd1, 0xc3, 0x25, 0x36, 0x0f,
	0x50, 0x8b, 0x88, 0x6b, 0xf4, 0xf8, 0xaa, 0x67, 0xa9, 0xc5, 0x35, 0x32, 0x0b, 0xe3, 0xe4, 0xef,
	0x01, 0x4c, 0xcb, 0x0c, 0xd5, 0x78, 0xbe, 0x8a, 0xcc, 0x58, 0x0f, 0x44, 0xcb, 0x12, 0x7c, 0x65,
	0x9e, 0x73, 0x3b, 0x02, 0xc0, 0xf6, 0x3f, 0xfd, 0x5e, 0xbe, 0x79, 0x10, 0x6f, 0x7f, 0xf7, 0x5a,
	0xd9, 0x8e, 0x60, 0x68, 0xdf, 0xbe, 0xfa, 0x25, 0x69, 0x05, 0xad, 0x65, 0xa2, 0x9c, 0xcf, 0xca,
	0x59, 0x1c, 0x85, 0x3a, 0xe7, 0x81, 0xb0, 0x68, 0xd6, 0xd3, 0x4f, 0xfc, 0xbb, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe3, 0xb7, 0x65, 0x8a, 0xf4, 0x0b, 0x00, 0x00,
}
