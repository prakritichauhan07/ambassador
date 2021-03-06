// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto

package envoy_config_cluster_dynamic_forward_proxy_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v2alpha "github.com/datawire/ambassador/pkg/api/envoy/config/common/dynamic_forward_proxy/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/gogo/protobuf/proto"
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

// Configuration for the dynamic forward proxy cluster. See the :ref:`architecture overview
// <arch_overview_http_dynamic_forward_proxy>` for more information.
// [#extension: envoy.clusters.dynamic_forward_proxy]
type ClusterConfig struct {
	// The DNS cache configuration that the cluster will attach to. Note this configuration must
	// match that of associated :ref:`dynamic forward proxy HTTP filter configuration
	// <envoy_api_field_config.filter.http.dynamic_forward_proxy.v2alpha.FilterConfig.dns_cache_config>`.
	DnsCacheConfig       *v2alpha.DnsCacheConfig `protobuf:"bytes,1,opt,name=dns_cache_config,json=dnsCacheConfig,proto3" json:"dns_cache_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClusterConfig) Reset()         { *m = ClusterConfig{} }
func (m *ClusterConfig) String() string { return proto.CompactTextString(m) }
func (*ClusterConfig) ProtoMessage()    {}
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_faeb9d327e132c11, []int{0}
}
func (m *ClusterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterConfig.Merge(m, src)
}
func (m *ClusterConfig) XXX_Size() int {
	return m.Size()
}
func (m *ClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterConfig proto.InternalMessageInfo

func (m *ClusterConfig) GetDnsCacheConfig() *v2alpha.DnsCacheConfig {
	if m != nil {
		return m.DnsCacheConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterConfig)(nil), "envoy.config.cluster.dynamic_forward_proxy.v2alpha.ClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/dynamic_forward_proxy/v2alpha/cluster.proto", fileDescriptor_faeb9d327e132c11)
}

var fileDescriptor_faeb9d327e132c11 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x4f, 0xa9, 0xcc, 0x4b, 0xcc, 0xcd, 0x4c, 0x8e, 0x4f, 0xcb, 0x2f, 0x2a, 0x4f, 0x2c, 0x4a,
	0x89, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x33, 0x4a, 0xcc, 0x29, 0xc8, 0x48, 0x84, 0xa9,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x32, 0x02, 0x9b, 0xa0, 0x07, 0x31, 0x41, 0x0f, 0x26,
	0x87, 0xd5, 0x04, 0x3d, 0xa8, 0x09, 0x52, 0x8e, 0xa8, 0xb6, 0xe6, 0xe7, 0xe6, 0xe6, 0xe7, 0x11,
	0xb0, 0x34, 0x25, 0xaf, 0x38, 0x3e, 0x39, 0x31, 0x39, 0x23, 0x15, 0x62, 0xad, 0x94, 0x5c, 0x69,
	0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x7e,
	0x6e, 0x66, 0x7a, 0x51, 0x62, 0x09, 0x4c, 0x5e, 0x16, 0x43, 0xbe, 0xb8, 0x24, 0xb1, 0xa4, 0xb4,
	0x18, 0x2a, 0x2d, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0xaa, 0x0f, 0x63, 0x40, 0x24,
	0x94, 0xda, 0x18, 0xb9, 0x78, 0x9d, 0x21, 0x9e, 0x70, 0x06, 0x3b, 0x4f, 0xa8, 0x94, 0x4b, 0x00,
	0x6e, 0x79, 0x3c, 0xc4, 0xc9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x8e, 0x7a, 0xa8, 0x7e,
	0x07, 0xfb, 0x03, 0xbf, 0xd7, 0xf5, 0x5c, 0xf2, 0x8a, 0x9d, 0x41, 0x26, 0x41, 0x0c, 0x77, 0xe2,
	0xf8, 0xe5, 0xc4, 0xda, 0xc5, 0xc8, 0x24, 0xc0, 0x18, 0xc4, 0x97, 0x82, 0x2a, 0xb3, 0x80, 0xf1,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0xfc, 0x34, 0xe3, 0x5f,
	0x3f, 0xab, 0x09, 0x2c, 0xb0, 0x53, 0x2b, 0x4a, 0x52, 0xf3, 0x8a, 0x41, 0xbe, 0x82, 0x05, 0x78,
	0x31, 0x2e, 0x6b, 0x8d, 0x77, 0x35, 0x9c, 0xb8, 0xc8, 0xc6, 0x24, 0xc0, 0xc0, 0xe5, 0x90, 0x99,
	0x0f, 0x71, 0x2f, 0x44, 0x86, 0xf4, 0x68, 0x73, 0xe2, 0x81, 0x06, 0x4c, 0x00, 0x28, 0xa4, 0x02,
	0x18, 0x93, 0xd8, 0xc0, 0x41, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xad, 0xaa, 0xc0, 0x68,
	0x45, 0x02, 0x00, 0x00,
}

func (m *ClusterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.DnsCacheConfig != nil {
		{
			size, err := m.DnsCacheConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCluster(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCluster(dAtA []byte, offset int, v uint64) int {
	offset -= sovCluster(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DnsCacheConfig != nil {
		l = m.DnsCacheConfig.Size()
		n += 1 + l + sovCluster(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCluster(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCluster(x uint64) (n int) {
	return sovCluster(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClusterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCluster
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
			return fmt.Errorf("proto: ClusterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DnsCacheConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCluster
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCluster
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCluster
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DnsCacheConfig == nil {
				m.DnsCacheConfig = &v2alpha.DnsCacheConfig{}
			}
			if err := m.DnsCacheConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCluster(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCluster
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCluster(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
					return 0, ErrIntOverflowCluster
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
				return 0, ErrInvalidLengthCluster
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCluster
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCluster
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCluster        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCluster          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCluster = fmt.Errorf("proto: unexpected end of group")
)
