// Code generated by protoc-gen-go. DO NOT EDIT.
// source: premierleague.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TableRequest struct {
	TableName            string   `protobuf:"bytes,1,opt,name=tableName,proto3" json:"tableName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableRequest) Reset()         { *m = TableRequest{} }
func (m *TableRequest) String() string { return proto.CompactTextString(m) }
func (*TableRequest) ProtoMessage()    {}
func (*TableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{0}
}

func (m *TableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableRequest.Unmarshal(m, b)
}
func (m *TableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableRequest.Marshal(b, m, deterministic)
}
func (m *TableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableRequest.Merge(m, src)
}
func (m *TableRequest) XXX_Size() int {
	return xxx_messageInfo_TableRequest.Size(m)
}
func (m *TableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TableRequest proto.InternalMessageInfo

func (m *TableRequest) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

type Table struct {
	TeamName             string   `protobuf:"bytes,1,opt,name=teamName,proto3" json:"teamName,omitempty"`
	TeamPlayed           int32    `protobuf:"varint,2,opt,name=teamPlayed,proto3" json:"teamPlayed,omitempty"`
	TeamWon              int32    `protobuf:"varint,3,opt,name=teamWon,proto3" json:"teamWon,omitempty"`
	TeamDrawn            int32    `protobuf:"varint,4,opt,name=teamDrawn,proto3" json:"teamDrawn,omitempty"`
	TeamLost             int32    `protobuf:"varint,5,opt,name=teamLost,proto3" json:"teamLost,omitempty"`
	TeamGF               int32    `protobuf:"varint,6,opt,name=teamGF,proto3" json:"teamGF,omitempty"`
	TeamGA               int32    `protobuf:"varint,7,opt,name=teamGA,proto3" json:"teamGA,omitempty"`
	TeamGD               int32    `protobuf:"varint,8,opt,name=teamGD,proto3" json:"teamGD,omitempty"`
	TeamPoints           int32    `protobuf:"varint,9,opt,name=teamPoints,proto3" json:"teamPoints,omitempty"`
	TeamCapital          int32    `protobuf:"varint,10,opt,name=teamCapital,proto3" json:"teamCapital,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Table) Reset()         { *m = Table{} }
func (m *Table) String() string { return proto.CompactTextString(m) }
func (*Table) ProtoMessage()    {}
func (*Table) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{1}
}

func (m *Table) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Table.Unmarshal(m, b)
}
func (m *Table) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Table.Marshal(b, m, deterministic)
}
func (m *Table) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Table.Merge(m, src)
}
func (m *Table) XXX_Size() int {
	return xxx_messageInfo_Table.Size(m)
}
func (m *Table) XXX_DiscardUnknown() {
	xxx_messageInfo_Table.DiscardUnknown(m)
}

var xxx_messageInfo_Table proto.InternalMessageInfo

func (m *Table) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

func (m *Table) GetTeamPlayed() int32 {
	if m != nil {
		return m.TeamPlayed
	}
	return 0
}

func (m *Table) GetTeamWon() int32 {
	if m != nil {
		return m.TeamWon
	}
	return 0
}

func (m *Table) GetTeamDrawn() int32 {
	if m != nil {
		return m.TeamDrawn
	}
	return 0
}

func (m *Table) GetTeamLost() int32 {
	if m != nil {
		return m.TeamLost
	}
	return 0
}

func (m *Table) GetTeamGF() int32 {
	if m != nil {
		return m.TeamGF
	}
	return 0
}

func (m *Table) GetTeamGA() int32 {
	if m != nil {
		return m.TeamGA
	}
	return 0
}

func (m *Table) GetTeamGD() int32 {
	if m != nil {
		return m.TeamGD
	}
	return 0
}

func (m *Table) GetTeamPoints() int32 {
	if m != nil {
		return m.TeamPoints
	}
	return 0
}

func (m *Table) GetTeamCapital() int32 {
	if m != nil {
		return m.TeamCapital
	}
	return 0
}

type TableReply struct {
	Teams                []*Table `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableReply) Reset()         { *m = TableReply{} }
func (m *TableReply) String() string { return proto.CompactTextString(m) }
func (*TableReply) ProtoMessage()    {}
func (*TableReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{2}
}

func (m *TableReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableReply.Unmarshal(m, b)
}
func (m *TableReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableReply.Marshal(b, m, deterministic)
}
func (m *TableReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableReply.Merge(m, src)
}
func (m *TableReply) XXX_Size() int {
	return xxx_messageInfo_TableReply.Size(m)
}
func (m *TableReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TableReply.DiscardUnknown(m)
}

var xxx_messageInfo_TableReply proto.InternalMessageInfo

func (m *TableReply) GetTeams() []*Table {
	if m != nil {
		return m.Teams
	}
	return nil
}

func (m *TableReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type TeamRequest struct {
	TeamName             string   `protobuf:"bytes,1,opt,name=teamName,proto3" json:"teamName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TeamRequest) Reset()         { *m = TeamRequest{} }
func (m *TeamRequest) String() string { return proto.CompactTextString(m) }
func (*TeamRequest) ProtoMessage()    {}
func (*TeamRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{3}
}

func (m *TeamRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamRequest.Unmarshal(m, b)
}
func (m *TeamRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamRequest.Marshal(b, m, deterministic)
}
func (m *TeamRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamRequest.Merge(m, src)
}
func (m *TeamRequest) XXX_Size() int {
	return xxx_messageInfo_TeamRequest.Size(m)
}
func (m *TeamRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TeamRequest proto.InternalMessageInfo

func (m *TeamRequest) GetTeamName() string {
	if m != nil {
		return m.TeamName
	}
	return ""
}

type Player struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Team                 string   `protobuf:"bytes,2,opt,name=team,proto3" json:"team,omitempty"`
	Nationality          string   `protobuf:"bytes,3,opt,name=nationality,proto3" json:"nationality,omitempty"`
	Position             string   `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Appearences          int32    `protobuf:"varint,5,opt,name=appearences,proto3" json:"appearences,omitempty"`
	Goals                int32    `protobuf:"varint,6,opt,name=goals,proto3" json:"goals,omitempty"`
	Assists              int32    `protobuf:"varint,7,opt,name=assists,proto3" json:"assists,omitempty"`
	Passes               int32    `protobuf:"varint,8,opt,name=passes,proto3" json:"passes,omitempty"`
	Interceptions        int32    `protobuf:"varint,9,opt,name=interceptions,proto3" json:"interceptions,omitempty"`
	Tackles              int32    `protobuf:"varint,10,opt,name=tackles,proto3" json:"tackles,omitempty"`
	Fouls                int32    `protobuf:"varint,11,opt,name=fouls,proto3" json:"fouls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}
func (*Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{4}
}

func (m *Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Player.Unmarshal(m, b)
}
func (m *Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Player.Marshal(b, m, deterministic)
}
func (m *Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Player.Merge(m, src)
}
func (m *Player) XXX_Size() int {
	return xxx_messageInfo_Player.Size(m)
}
func (m *Player) XXX_DiscardUnknown() {
	xxx_messageInfo_Player.DiscardUnknown(m)
}

var xxx_messageInfo_Player proto.InternalMessageInfo

func (m *Player) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Player) GetTeam() string {
	if m != nil {
		return m.Team
	}
	return ""
}

func (m *Player) GetNationality() string {
	if m != nil {
		return m.Nationality
	}
	return ""
}

func (m *Player) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

func (m *Player) GetAppearences() int32 {
	if m != nil {
		return m.Appearences
	}
	return 0
}

func (m *Player) GetGoals() int32 {
	if m != nil {
		return m.Goals
	}
	return 0
}

func (m *Player) GetAssists() int32 {
	if m != nil {
		return m.Assists
	}
	return 0
}

func (m *Player) GetPasses() int32 {
	if m != nil {
		return m.Passes
	}
	return 0
}

func (m *Player) GetInterceptions() int32 {
	if m != nil {
		return m.Interceptions
	}
	return 0
}

func (m *Player) GetTackles() int32 {
	if m != nil {
		return m.Tackles
	}
	return 0
}

func (m *Player) GetFouls() int32 {
	if m != nil {
		return m.Fouls
	}
	return 0
}

type TeamReply struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Err                  string    `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TeamReply) Reset()         { *m = TeamReply{} }
func (m *TeamReply) String() string { return proto.CompactTextString(m) }
func (*TeamReply) ProtoMessage()    {}
func (*TeamReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{5}
}

func (m *TeamReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TeamReply.Unmarshal(m, b)
}
func (m *TeamReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TeamReply.Marshal(b, m, deterministic)
}
func (m *TeamReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TeamReply.Merge(m, src)
}
func (m *TeamReply) XXX_Size() int {
	return xxx_messageInfo_TeamReply.Size(m)
}
func (m *TeamReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TeamReply.DiscardUnknown(m)
}

var xxx_messageInfo_TeamReply proto.InternalMessageInfo

func (m *TeamReply) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *TeamReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type PositionRequest struct {
	Position             string   `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionRequest) Reset()         { *m = PositionRequest{} }
func (m *PositionRequest) String() string { return proto.CompactTextString(m) }
func (*PositionRequest) ProtoMessage()    {}
func (*PositionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{6}
}

func (m *PositionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionRequest.Unmarshal(m, b)
}
func (m *PositionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionRequest.Marshal(b, m, deterministic)
}
func (m *PositionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionRequest.Merge(m, src)
}
func (m *PositionRequest) XXX_Size() int {
	return xxx_messageInfo_PositionRequest.Size(m)
}
func (m *PositionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PositionRequest proto.InternalMessageInfo

func (m *PositionRequest) GetPosition() string {
	if m != nil {
		return m.Position
	}
	return ""
}

type PositionReply struct {
	Players              []*Player `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	Err                  string    `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PositionReply) Reset()         { *m = PositionReply{} }
func (m *PositionReply) String() string { return proto.CompactTextString(m) }
func (*PositionReply) ProtoMessage()    {}
func (*PositionReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4204cc63886395c8, []int{7}
}

func (m *PositionReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionReply.Unmarshal(m, b)
}
func (m *PositionReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionReply.Marshal(b, m, deterministic)
}
func (m *PositionReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionReply.Merge(m, src)
}
func (m *PositionReply) XXX_Size() int {
	return xxx_messageInfo_PositionReply.Size(m)
}
func (m *PositionReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionReply.DiscardUnknown(m)
}

var xxx_messageInfo_PositionReply proto.InternalMessageInfo

func (m *PositionReply) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *PositionReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*TableRequest)(nil), "pb.TableRequest")
	proto.RegisterType((*Table)(nil), "pb.Table")
	proto.RegisterType((*TableReply)(nil), "pb.TableReply")
	proto.RegisterType((*TeamRequest)(nil), "pb.TeamRequest")
	proto.RegisterType((*Player)(nil), "pb.Player")
	proto.RegisterType((*TeamReply)(nil), "pb.TeamReply")
	proto.RegisterType((*PositionRequest)(nil), "pb.PositionRequest")
	proto.RegisterType((*PositionReply)(nil), "pb.PositionReply")
}

func init() { proto.RegisterFile("premierleague.proto", fileDescriptor_4204cc63886395c8) }

var fileDescriptor_4204cc63886395c8 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x49, 0x93, 0xd4, 0x93, 0x86, 0x94, 0x0d, 0x42, 0xab, 0x08, 0x41, 0x64, 0xf5, 0x10,
	0x24, 0x08, 0x52, 0xfb, 0x00, 0x55, 0xd5, 0x88, 0x5e, 0x2a, 0x54, 0xb9, 0x95, 0x38, 0x6f, 0xc2,
	0x50, 0x59, 0xf8, 0x67, 0xd9, 0xdd, 0x80, 0xf2, 0x38, 0xbc, 0x00, 0x37, 0xde, 0x0f, 0xcd, 0xd8,
	0x6b, 0x6f, 0x24, 0xc4, 0x85, 0xdb, 0x7c, 0xdf, 0xcc, 0x78, 0x66, 0xbf, 0x6f, 0xd7, 0x30, 0xd3,
	0x06, 0x8b, 0x0c, 0x4d, 0x8e, 0xea, 0x71, 0x87, 0x2b, 0x6d, 0x2a, 0x57, 0x89, 0x9e, 0xde, 0x24,
	0x6f, 0xe1, 0xe4, 0x41, 0x6d, 0x72, 0x4c, 0xf1, 0xdb, 0x0e, 0xad, 0x13, 0x2f, 0x21, 0x76, 0x84,
	0x3f, 0xaa, 0x02, 0x65, 0xb4, 0x88, 0x96, 0x71, 0xda, 0x11, 0xc9, 0xcf, 0x1e, 0x0c, 0xb8, 0x5c,
	0xcc, 0xe1, 0xd8, 0xa1, 0x2a, 0x82, 0xb2, 0x16, 0x8b, 0x57, 0x00, 0x14, 0xdf, 0xe5, 0x6a, 0x8f,
	0x9f, 0x65, 0x6f, 0x11, 0x2d, 0x07, 0x69, 0xc0, 0x08, 0x09, 0x23, 0x42, 0x9f, 0xaa, 0x52, 0xf6,
	0x39, 0xe9, 0x21, 0x4f, 0x47, 0x55, 0xac, 0x8d, 0xfa, 0x51, 0xca, 0x23, 0xce, 0x75, 0x84, 0x9f,
	0x79, 0x5b, 0x59, 0x27, 0x07, 0x9c, 0x6c, 0xb1, 0x78, 0x01, 0x43, 0x8a, 0x6f, 0x3e, 0xc8, 0x21,
	0x67, 0x1a, 0xd4, 0xf2, 0x57, 0x72, 0x14, 0xf0, 0x57, 0x2d, 0xbf, 0x96, 0xc7, 0x01, 0xbf, 0x6e,
	0x77, 0xaf, 0xb2, 0xd2, 0x59, 0x19, 0x07, 0xbb, 0x33, 0x23, 0x16, 0x30, 0x26, 0x74, 0xad, 0x74,
	0xe6, 0x54, 0x2e, 0x81, 0x0b, 0x42, 0x2a, 0xb9, 0x04, 0x68, 0x14, 0xd5, 0xf9, 0x5e, 0xbc, 0x86,
	0x01, 0x25, 0xad, 0x8c, 0x16, 0xfd, 0xe5, 0xf8, 0x3c, 0x5e, 0xe9, 0xcd, 0xaa, 0x4e, 0xd7, 0xbc,
	0x38, 0x85, 0x3e, 0x1a, 0xc3, 0x2a, 0xc5, 0x29, 0x85, 0xc9, 0x1b, 0x18, 0x3f, 0xa0, 0x2a, 0xbc,
	0x23, 0xff, 0x50, 0x3a, 0xf9, 0xd5, 0x83, 0x21, 0x8b, 0x6a, 0x84, 0x80, 0xa3, 0xb2, 0x2b, 0xe1,
	0x98, 0x38, 0x2a, 0x6d, 0x3e, 0xce, 0x31, 0x1d, 0xa0, 0x54, 0x2e, 0xab, 0x4a, 0x95, 0x67, 0x6e,
	0xcf, 0x06, 0xc4, 0x69, 0x48, 0xd1, 0x40, 0x5d, 0xd9, 0x8c, 0x08, 0xf6, 0x20, 0x4e, 0x5b, 0x4c,
	0xdd, 0x4a, 0x6b, 0x54, 0x06, 0xcb, 0x2d, 0xda, 0xc6, 0x85, 0x90, 0x12, 0xcf, 0x61, 0xf0, 0x58,
	0xa9, 0xdc, 0x36, 0x3e, 0xd4, 0x80, 0x2c, 0x57, 0xd6, 0x66, 0xd6, 0xd9, 0xc6, 0x07, 0x0f, 0xc9,
	0x08, 0xad, 0xac, 0x45, 0xeb, 0x8d, 0xa8, 0x91, 0x38, 0x83, 0x49, 0x56, 0x3a, 0x34, 0x5b, 0xd4,
	0x34, 0xd9, 0x7b, 0x71, 0x48, 0xf2, 0x55, 0x52, 0xdb, 0xaf, 0x39, 0xda, 0xc6, 0x0a, 0x0f, 0x69,
	0x8f, 0x2f, 0xd5, 0x2e, 0xb7, 0x72, 0x5c, 0xef, 0xc1, 0x20, 0xb9, 0x86, 0xb8, 0xd6, 0x96, 0xbc,
	0x39, 0x83, 0x91, 0x66, 0xf1, 0xbc, 0x3b, 0x40, 0xee, 0xd4, 0x7a, 0xa6, 0x3e, 0xf5, 0x17, 0x83,
	0xde, 0xc1, 0xf4, 0xae, 0x11, 0x24, 0x30, 0xa9, 0xd5, 0x2c, 0x3a, 0xd4, 0x2c, 0xb9, 0x81, 0x49,
	0x57, 0xfe, 0x1f, 0x73, 0xcf, 0x7f, 0x47, 0x70, 0x72, 0xef, 0x94, 0xb3, 0xf7, 0x68, 0xbe, 0x67,
	0x5b, 0x14, 0xef, 0x21, 0xbe, 0xcd, 0xac, 0xab, 0x5f, 0xe4, 0x69, 0x77, 0xb5, 0xea, 0xa5, 0xe6,
	0x4f, 0x03, 0x46, 0xe7, 0xfb, 0xe4, 0x89, 0xb8, 0x80, 0x29, 0x37, 0xf8, 0xb7, 0x68, 0xac, 0x98,
	0x72, 0x51, 0x77, 0xdf, 0xe6, 0x93, 0x8e, 0xa8, 0x9b, 0x2e, 0x61, 0x46, 0x4d, 0xfe, 0x0c, 0xbe,
	0x71, 0xc6, 0x4b, 0x1f, 0xea, 0x30, 0x7f, 0x76, 0x48, 0xf2, 0x07, 0x36, 0x43, 0xfe, 0xdd, 0x5c,
	0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x39, 0x75, 0x26, 0x98, 0x85, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsServiceClient is the client API for StatsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsServiceClient interface {
	ListTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableReply, error)
	ListTeamPlayers(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*TeamReply, error)
	ListPositionPlayers(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReply, error)
}

type statsServiceClient struct {
	cc *grpc.ClientConn
}

func NewStatsServiceClient(cc *grpc.ClientConn) StatsServiceClient {
	return &statsServiceClient{cc}
}

func (c *statsServiceClient) ListTable(ctx context.Context, in *TableRequest, opts ...grpc.CallOption) (*TableReply, error) {
	out := new(TableReply)
	err := c.cc.Invoke(ctx, "/pb.StatsService/ListTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) ListTeamPlayers(ctx context.Context, in *TeamRequest, opts ...grpc.CallOption) (*TeamReply, error) {
	out := new(TeamReply)
	err := c.cc.Invoke(ctx, "/pb.StatsService/ListTeamPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsServiceClient) ListPositionPlayers(ctx context.Context, in *PositionRequest, opts ...grpc.CallOption) (*PositionReply, error) {
	out := new(PositionReply)
	err := c.cc.Invoke(ctx, "/pb.StatsService/ListPositionPlayers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServiceServer is the server API for StatsService service.
type StatsServiceServer interface {
	ListTable(context.Context, *TableRequest) (*TableReply, error)
	ListTeamPlayers(context.Context, *TeamRequest) (*TeamReply, error)
	ListPositionPlayers(context.Context, *PositionRequest) (*PositionReply, error)
}

// UnimplementedStatsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStatsServiceServer struct {
}

func (*UnimplementedStatsServiceServer) ListTable(ctx context.Context, req *TableRequest) (*TableReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTable not implemented")
}
func (*UnimplementedStatsServiceServer) ListTeamPlayers(ctx context.Context, req *TeamRequest) (*TeamReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTeamPlayers not implemented")
}
func (*UnimplementedStatsServiceServer) ListPositionPlayers(ctx context.Context, req *PositionRequest) (*PositionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPositionPlayers not implemented")
}

func RegisterStatsServiceServer(s *grpc.Server, srv StatsServiceServer) {
	s.RegisterService(&_StatsService_serviceDesc, srv)
}

func _StatsService_ListTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).ListTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatsService/ListTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).ListTable(ctx, req.(*TableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_ListTeamPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).ListTeamPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatsService/ListTeamPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).ListTeamPlayers(ctx, req.(*TeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatsService_ListPositionPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PositionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServiceServer).ListPositionPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatsService/ListPositionPlayers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServiceServer).ListPositionPlayers(ctx, req.(*PositionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StatsService",
	HandlerType: (*StatsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListTable",
			Handler:    _StatsService_ListTable_Handler,
		},
		{
			MethodName: "ListTeamPlayers",
			Handler:    _StatsService_ListTeamPlayers_Handler,
		},
		{
			MethodName: "ListPositionPlayers",
			Handler:    _StatsService_ListPositionPlayers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "premierleague.proto",
}
