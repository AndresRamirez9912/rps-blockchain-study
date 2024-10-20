// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lb/rps/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters of the module.
type Params struct {
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b582db7b6994176a, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// Game defines the Rock, Paper & Scissors game object
type Game struct {
	GameNumber   uint64   `protobuf:"varint,1,opt,name=game_number,json=gameNumber,proto3" json:"game_number,omitempty"`
	PlayerA      string   `protobuf:"bytes,2,opt,name=player_a,json=playerA,proto3" json:"player_a,omitempty"`
	PlayerB      string   `protobuf:"bytes,3,opt,name=player_b,json=playerB,proto3" json:"player_b,omitempty"`
	Status       string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Rounds       uint64   `protobuf:"varint,5,opt,name=rounds,proto3" json:"rounds,omitempty"`
	PlayerAMoves []string `protobuf:"bytes,6,rep,name=player_a_moves,json=playerAMoves,proto3" json:"player_a_moves,omitempty"`
	PlayerBMoves []string `protobuf:"bytes,7,rep,name=player_b_moves,json=playerBMoves,proto3" json:"player_b_moves,omitempty"`
	Score        []uint64 `protobuf:"varint,8,rep,packed,name=score,proto3" json:"score,omitempty"`
}

func (m *Game) Reset()         { *m = Game{} }
func (m *Game) String() string { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()    {}
func (*Game) Descriptor() ([]byte, []int) {
	return fileDescriptor_b582db7b6994176a, []int{1}
}
func (m *Game) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Game) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Game.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Game) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Game.Merge(m, src)
}
func (m *Game) XXX_Size() int {
	return m.Size()
}
func (m *Game) XXX_DiscardUnknown() {
	xxx_messageInfo_Game.DiscardUnknown(m)
}

var xxx_messageInfo_Game proto.InternalMessageInfo

func (m *Game) GetGameNumber() uint64 {
	if m != nil {
		return m.GameNumber
	}
	return 0
}

func (m *Game) GetPlayerA() string {
	if m != nil {
		return m.PlayerA
	}
	return ""
}

func (m *Game) GetPlayerB() string {
	if m != nil {
		return m.PlayerB
	}
	return ""
}

func (m *Game) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Game) GetRounds() uint64 {
	if m != nil {
		return m.Rounds
	}
	return 0
}

func (m *Game) GetPlayerAMoves() []string {
	if m != nil {
		return m.PlayerAMoves
	}
	return nil
}

func (m *Game) GetPlayerBMoves() []string {
	if m != nil {
		return m.PlayerBMoves
	}
	return nil
}

func (m *Game) GetScore() []uint64 {
	if m != nil {
		return m.Score
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "lb.rps.v1.Params")
	proto.RegisterType((*Game)(nil), "lb.rps.v1.Game")
}

func init() { proto.RegisterFile("lb/rps/v1/types.proto", fileDescriptor_b582db7b6994176a) }

var fileDescriptor_b582db7b6994176a = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4e, 0xc3, 0x30,
	0x18, 0x85, 0x9b, 0x36, 0x4d, 0x5b, 0x83, 0x18, 0xa2, 0x82, 0x0c, 0x43, 0x88, 0x2a, 0x90, 0xb2,
	0x34, 0xa1, 0xea, 0x01, 0x50, 0xb3, 0x30, 0x81, 0x50, 0xd8, 0x58, 0x22, 0x3b, 0xb5, 0xda, 0x4a,
	0x75, 0x1c, 0xf9, 0x77, 0xaa, 0xf6, 0x16, 0xdc, 0x81, 0x2b, 0x70, 0x08, 0xc6, 0x8a, 0x89, 0x11,
	0xb5, 0x17, 0x41, 0xb1, 0x83, 0x94, 0x91, 0xf1, 0x3d, 0x7f, 0xcf, 0x7a, 0xfa, 0x1f, 0x3a, 0x5f,
	0xd3, 0x48, 0x16, 0x10, 0x6d, 0x26, 0x91, 0xda, 0x15, 0x0c, 0xc2, 0x42, 0x0a, 0x25, 0xdc, 0xc1,
	0x9a, 0x86, 0xb2, 0x80, 0x70, 0x33, 0xb9, 0xba, 0xcc, 0x04, 0x70, 0x01, 0xa9, 0x7e, 0x88, 0x8c,
	0x30, 0xd4, 0xa8, 0x8f, 0x9c, 0x67, 0x22, 0x09, 0x87, 0xd1, 0x7b, 0x1b, 0xd9, 0x0f, 0x84, 0x33,
	0xf7, 0x1a, 0x9d, 0x2c, 0x08, 0x67, 0x69, 0x5e, 0x72, 0xca, 0x24, 0xb6, 0x7c, 0x2b, 0xb0, 0x13,
	0x54, 0x59, 0x4f, 0xda, 0x71, 0xa7, 0xa8, 0x5f, 0xac, 0xc9, 0x8e, 0xc9, 0x94, 0xe0, 0xb6, 0x6f,
	0x05, 0x83, 0x18, 0x7f, 0x7d, 0x8c, 0x87, 0xf5, 0xbf, 0xb3, 0xf9, 0x5c, 0x32, 0x80, 0x17, 0x25,
	0x57, 0xf9, 0x22, 0xe9, 0x19, 0x72, 0xd6, 0x08, 0x51, 0xdc, 0xf9, 0x5f, 0x28, 0x76, 0x2f, 0x90,
	0x03, 0x8a, 0xa8, 0x12, 0xb0, 0x5d, 0x45, 0x92, 0x5a, 0x55, 0xbe, 0x14, 0x65, 0x3e, 0x07, 0xdc,
	0xd5, 0xed, 0x6a, 0xe5, 0xde, 0xa0, 0xb3, 0xbf, 0x66, 0x29, 0x17, 0x1b, 0x06, 0xd8, 0xf1, 0x3b,
	0xc1, 0x20, 0x39, 0xad, 0x5b, 0x3c, 0x56, 0x5e, 0x83, 0xa2, 0x35, 0xd5, 0x6b, 0x52, 0xb1, 0xa1,
	0x86, 0xa8, 0x0b, 0x99, 0x90, 0x0c, 0xf7, 0xfd, 0x4e, 0x60, 0x27, 0x46, 0xc4, 0xf7, 0x9f, 0x07,
	0xcf, 0xda, 0x1f, 0x3c, 0xeb, 0xe7, 0xe0, 0x59, 0x6f, 0x47, 0xaf, 0xb5, 0x3f, 0x7a, 0xad, 0xef,
	0xa3, 0xd7, 0x7a, 0xbd, 0x5d, 0xac, 0xd4, 0xb2, 0xa4, 0x61, 0x26, 0x78, 0x74, 0xb7, 0x35, 0x9b,
	0x8c, 0xb3, 0x25, 0x59, 0xe5, 0xd1, 0x56, 0xef, 0xa3, 0xc7, 0xa1, 0x8e, 0xbe, 0xfb, 0xf4, 0x37,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0xa3, 0xc6, 0xd3, 0xb6, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *Game) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Game) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Game) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Score) > 0 {
		dAtA2 := make([]byte, len(m.Score)*10)
		var j1 int
		for _, num := range m.Score {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTypes(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x42
	}
	if len(m.PlayerBMoves) > 0 {
		for iNdEx := len(m.PlayerBMoves) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PlayerBMoves[iNdEx])
			copy(dAtA[i:], m.PlayerBMoves[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PlayerBMoves[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.PlayerAMoves) > 0 {
		for iNdEx := len(m.PlayerAMoves) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PlayerAMoves[iNdEx])
			copy(dAtA[i:], m.PlayerAMoves[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.PlayerAMoves[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.Rounds != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Rounds))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PlayerB) > 0 {
		i -= len(m.PlayerB)
		copy(dAtA[i:], m.PlayerB)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PlayerB)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PlayerA) > 0 {
		i -= len(m.PlayerA)
		copy(dAtA[i:], m.PlayerA)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.PlayerA)))
		i--
		dAtA[i] = 0x12
	}
	if m.GameNumber != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.GameNumber))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *Game) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GameNumber != 0 {
		n += 1 + sovTypes(uint64(m.GameNumber))
	}
	l = len(m.PlayerA)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.PlayerB)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Rounds != 0 {
		n += 1 + sovTypes(uint64(m.Rounds))
	}
	if len(m.PlayerAMoves) > 0 {
		for _, s := range m.PlayerAMoves {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.PlayerBMoves) > 0 {
		for _, s := range m.PlayerBMoves {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Score) > 0 {
		l = 0
		for _, e := range m.Score {
			l += sovTypes(uint64(e))
		}
		n += 1 + sovTypes(uint64(l)) + l
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Game) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Game: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Game: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GameNumber", wireType)
			}
			m.GameNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GameNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerB", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerB = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rounds", wireType)
			}
			m.Rounds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rounds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerAMoves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerAMoves = append(m.PlayerAMoves, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerBMoves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PlayerBMoves = append(m.PlayerBMoves, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Score = append(m.Score, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTypes
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTypes
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTypes
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Score) == 0 {
					m.Score = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTypes
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Score = append(m.Score, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
