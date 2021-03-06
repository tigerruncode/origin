// Code generated by protoc-gen-gogo.
// source: github.com/openshift/api/config/v1/generated.proto
// DO NOT EDIT!

/*
	Package v1 is a generated protocol buffer package.

	It is generated from these files:
		github.com/openshift/api/config/v1/generated.proto

	It has these top-level messages:
		CertInfo
		HTTPServingInfo
		NamedCertificate
		ServingInfo
*/
package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *CertInfo) Reset()                    { *m = CertInfo{} }
func (*CertInfo) ProtoMessage()               {}
func (*CertInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *HTTPServingInfo) Reset()                    { *m = HTTPServingInfo{} }
func (*HTTPServingInfo) ProtoMessage()               {}
func (*HTTPServingInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *NamedCertificate) Reset()                    { *m = NamedCertificate{} }
func (*NamedCertificate) ProtoMessage()               {}
func (*NamedCertificate) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *ServingInfo) Reset()                    { *m = ServingInfo{} }
func (*ServingInfo) ProtoMessage()               {}
func (*ServingInfo) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*CertInfo)(nil), "github.com.openshift.api.config.v1.CertInfo")
	proto.RegisterType((*HTTPServingInfo)(nil), "github.com.openshift.api.config.v1.HTTPServingInfo")
	proto.RegisterType((*NamedCertificate)(nil), "github.com.openshift.api.config.v1.NamedCertificate")
	proto.RegisterType((*ServingInfo)(nil), "github.com.openshift.api.config.v1.ServingInfo")
}
func (m *CertInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CertInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.CertFile)))
	i += copy(dAtA[i:], m.CertFile)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.KeyFile)))
	i += copy(dAtA[i:], m.KeyFile)
	return i, nil
}

func (m *HTTPServingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPServingInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.ServingInfo.Size()))
	n1, err := m.ServingInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.MaxRequestsInFlight))
	dAtA[i] = 0x18
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.RequestTimeoutSeconds))
	return i, nil
}

func (m *NamedCertificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamedCertificate) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.CertInfo.Size()))
	n2, err := m.CertInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *ServingInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServingInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.BindAddress)))
	i += copy(dAtA[i:], m.BindAddress)
	dAtA[i] = 0x12
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.BindNetwork)))
	i += copy(dAtA[i:], m.BindNetwork)
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.CertInfo.Size()))
	n3, err := m.CertInfo.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.ClientCA)))
	i += copy(dAtA[i:], m.ClientCA)
	if len(m.NamedCertificates) > 0 {
		for _, msg := range m.NamedCertificates {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintGenerated(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.MinTLSVersion)))
	i += copy(dAtA[i:], m.MinTLSVersion)
	if len(m.CipherSuites) > 0 {
		for _, s := range m.CipherSuites {
			dAtA[i] = 0x3a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Generated(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CertInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.CertFile)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.KeyFile)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *HTTPServingInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ServingInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	n += 1 + sovGenerated(uint64(m.MaxRequestsInFlight))
	n += 1 + sovGenerated(uint64(m.RequestTimeoutSeconds))
	return n
}

func (m *NamedCertificate) Size() (n int) {
	var l int
	_ = l
	if len(m.Names) > 0 {
		for _, s := range m.Names {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = m.CertInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *ServingInfo) Size() (n int) {
	var l int
	_ = l
	l = len(m.BindAddress)
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.BindNetwork)
	n += 1 + l + sovGenerated(uint64(l))
	l = m.CertInfo.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = len(m.ClientCA)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.NamedCertificates) > 0 {
		for _, e := range m.NamedCertificates {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.MinTLSVersion)
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.CipherSuites) > 0 {
		for _, s := range m.CipherSuites {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CertInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CertInfo{`,
		`CertFile:` + fmt.Sprintf("%v", this.CertFile) + `,`,
		`KeyFile:` + fmt.Sprintf("%v", this.KeyFile) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPServingInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPServingInfo{`,
		`ServingInfo:` + strings.Replace(strings.Replace(this.ServingInfo.String(), "ServingInfo", "ServingInfo", 1), `&`, ``, 1) + `,`,
		`MaxRequestsInFlight:` + fmt.Sprintf("%v", this.MaxRequestsInFlight) + `,`,
		`RequestTimeoutSeconds:` + fmt.Sprintf("%v", this.RequestTimeoutSeconds) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NamedCertificate) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&NamedCertificate{`,
		`Names:` + fmt.Sprintf("%v", this.Names) + `,`,
		`CertInfo:` + strings.Replace(strings.Replace(this.CertInfo.String(), "CertInfo", "CertInfo", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ServingInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ServingInfo{`,
		`BindAddress:` + fmt.Sprintf("%v", this.BindAddress) + `,`,
		`BindNetwork:` + fmt.Sprintf("%v", this.BindNetwork) + `,`,
		`CertInfo:` + strings.Replace(strings.Replace(this.CertInfo.String(), "CertInfo", "CertInfo", 1), `&`, ``, 1) + `,`,
		`ClientCA:` + fmt.Sprintf("%v", this.ClientCA) + `,`,
		`NamedCertificates:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.NamedCertificates), "NamedCertificate", "NamedCertificate", 1), `&`, ``, 1) + `,`,
		`MinTLSVersion:` + fmt.Sprintf("%v", this.MinTLSVersion) + `,`,
		`CipherSuites:` + fmt.Sprintf("%v", this.CipherSuites) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CertInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CertInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CertInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CertFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyFile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyFile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *HTTPServingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HTTPServingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPServingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServingInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ServingInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxRequestsInFlight", wireType)
			}
			m.MaxRequestsInFlight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxRequestsInFlight |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestTimeoutSeconds", wireType)
			}
			m.RequestTimeoutSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestTimeoutSeconds |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *NamedCertificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamedCertificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamedCertificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Names", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Names = append(m.Names, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CertInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *ServingInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ServingInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServingInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BindAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BindNetwork", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BindNetwork = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CertInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CertInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientCA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientCA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NamedCertificates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NamedCertificates = append(m.NamedCertificates, NamedCertificate{})
			if err := m.NamedCertificates[len(m.NamedCertificates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinTLSVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinTLSVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CipherSuites", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CipherSuites = append(m.CipherSuites, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenerated
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/openshift/api/config/v1/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xcf, 0x4e, 0xdb, 0x40,
	0x10, 0xc6, 0xe3, 0xa4, 0x14, 0xd8, 0x80, 0x00, 0x23, 0xa4, 0x94, 0xaa, 0x0e, 0xca, 0x29, 0x95,
	0xa8, 0x2d, 0x28, 0xad, 0x2a, 0xf5, 0x84, 0x91, 0x50, 0x51, 0x0b, 0xaa, 0x36, 0x51, 0x0f, 0xdc,
	0x8c, 0x3d, 0xb6, 0x57, 0x89, 0x77, 0xdd, 0xdd, 0x75, 0x5a, 0x0e, 0x95, 0xfa, 0x06, 0xed, 0x0b,
	0xf4, 0x7d, 0x38, 0x72, 0xe4, 0x14, 0x15, 0xf7, 0x21, 0x7a, 0xad, 0xbc, 0x36, 0xc1, 0xf9, 0x83,
	0xe0, 0xd0, 0x9b, 0x77, 0xe6, 0xfb, 0x7d, 0x33, 0x9e, 0x19, 0xb4, 0x1b, 0x10, 0x19, 0x26, 0x67,
	0xa6, 0xcb, 0x22, 0x8b, 0xc5, 0x40, 0x45, 0x48, 0x7c, 0x69, 0x39, 0x31, 0xb1, 0x5c, 0x46, 0x7d,
	0x12, 0x58, 0x83, 0x1d, 0x2b, 0x00, 0x0a, 0xdc, 0x91, 0xe0, 0x99, 0x31, 0x67, 0x92, 0xe9, 0xad,
	0x5b, 0xc6, 0x1c, 0x31, 0xa6, 0x13, 0x13, 0x33, 0x67, 0xcc, 0xc1, 0xce, 0xe6, 0x8b, 0x92, 0x6f,
	0xc0, 0x02, 0x66, 0x29, 0xf4, 0x2c, 0xf1, 0xd5, 0x4b, 0x3d, 0xd4, 0x57, 0x6e, 0xb9, 0xb9, 0xd7,
	0x7b, 0x23, 0x4c, 0xc2, 0xb2, 0xc2, 0x91, 0xe3, 0x86, 0x84, 0x02, 0x3f, 0xb7, 0xe2, 0x5e, 0x90,
	0x05, 0x84, 0x15, 0x81, 0x74, 0x66, 0x34, 0xb2, 0x69, 0xdd, 0x45, 0xf1, 0x84, 0x4a, 0x12, 0xc1,
	0x14, 0xf0, 0xfa, 0x3e, 0x40, 0xb8, 0x21, 0x44, 0xce, 0x14, 0xf7, 0xf2, 0x2e, 0x2e, 0x91, 0xa4,
	0x6f, 0x11, 0x2a, 0x85, 0xe4, 0x93, 0x50, 0xcb, 0x45, 0x0b, 0x07, 0xc0, 0xe5, 0x11, 0xf5, 0x99,
	0xbe, 0x8d, 0x16, 0x5c, 0xe0, 0xf2, 0x90, 0xf4, 0xa1, 0xa1, 0x6d, 0x69, 0xed, 0x45, 0x7b, 0xf5,
	0x62, 0xd8, 0xac, 0xa4, 0xc3, 0xa6, 0xd2, 0x64, 0x71, 0x3c, 0x52, 0xe8, 0xcf, 0xd1, 0x7c, 0x0f,
	0xce, 0x95, 0xb8, 0xaa, 0xc4, 0x2b, 0x85, 0x78, 0xfe, 0x7d, 0x1e, 0xc6, 0x37, 0xf9, 0xd6, 0xaf,
	0x2a, 0x5a, 0x79, 0xd7, 0xed, 0x7e, 0xec, 0x00, 0x1f, 0x10, 0x1a, 0xa8, 0x62, 0x3e, 0xaa, 0x8b,
	0xdb, 0xa7, 0xaa, 0x57, 0xdf, 0xb5, 0xcc, 0xfb, 0xb7, 0x66, 0x96, 0x5c, 0xec, 0xf5, 0xa2, 0x66,
	0xbd, 0x14, 0xc4, 0x65, 0x63, 0xfd, 0x18, 0xad, 0x47, 0xce, 0x57, 0x0c, 0x9f, 0x13, 0x10, 0x52,
	0x1c, 0xd1, 0xc3, 0x3e, 0x09, 0x42, 0xa9, 0x5a, 0xae, 0xd9, 0x4f, 0x0b, 0x7c, 0xfd, 0x78, 0x5a,
	0x82, 0x67, 0x71, 0x7a, 0x07, 0x6d, 0xf0, 0x3c, 0xd6, 0x25, 0x11, 0xb0, 0x44, 0x76, 0xc0, 0x65,
	0xd4, 0x13, 0x8d, 0x9a, 0x32, 0x7c, 0x56, 0x18, 0x6e, 0xe0, 0x59, 0x22, 0x3c, 0x9b, 0x6d, 0xfd,
	0xd0, 0xd0, 0xea, 0x89, 0x13, 0x81, 0x97, 0x8d, 0x99, 0xf8, 0xc4, 0x75, 0x24, 0xe8, 0x4d, 0x34,
	0x47, 0x9d, 0x08, 0x44, 0x43, 0xdb, 0xaa, 0xb5, 0x17, 0xed, 0xc5, 0x74, 0xd8, 0x9c, 0xcb, 0x44,
	0x02, 0xe7, 0x71, 0xfd, 0x34, 0x5f, 0x97, 0x1a, 0x5f, 0x55, 0x8d, 0x6f, 0xfb, 0x21, 0xe3, 0xbb,
	0x59, 0xf7, 0xf8, 0x72, 0xd5, 0xe0, 0x46, 0x7e, 0xad, 0xbf, 0x35, 0x54, 0x1e, 0xa9, 0xfe, 0x0a,
	0xd5, 0xcf, 0x08, 0xf5, 0xf6, 0x3d, 0x8f, 0x83, 0x10, 0xc5, 0x75, 0x8c, 0x86, 0x6f, 0xdf, 0xa6,
	0x70, 0x59, 0x77, 0x83, 0x9d, 0x80, 0xfc, 0xc2, 0x78, 0xaf, 0xb8, 0x93, 0x31, 0xac, 0x48, 0xe1,
	0xb2, 0x6e, 0xec, 0xcf, 0x6a, 0xff, 0xf7, 0xcf, 0xd4, 0x91, 0xf7, 0x09, 0x50, 0x79, 0xb0, 0xdf,
	0x78, 0x34, 0x71, 0xe4, 0x45, 0x1c, 0x8f, 0x14, 0xfa, 0x37, 0xb4, 0x46, 0x27, 0x16, 0x23, 0x1a,
	0x73, 0x5b, 0xb5, 0x76, 0x7d, 0x77, 0xef, 0x21, 0x2d, 0x4d, 0x6e, 0xd5, 0x7e, 0x52, 0x14, 0x5b,
	0x9b, 0xcc, 0x08, 0x3c, 0x5d, 0x49, 0x7f, 0x8b, 0x96, 0x23, 0x42, 0xbb, 0x1f, 0x3a, 0x9f, 0x80,
	0x0b, 0xc2, 0x68, 0xe3, 0xb1, 0xea, 0x78, 0xa3, 0x30, 0x59, 0x3e, 0x2e, 0x27, 0xf1, 0xb8, 0x56,
	0xdf, 0x43, 0x4b, 0x2e, 0x89, 0x43, 0xe0, 0x9d, 0x84, 0x64, 0x6d, 0xcf, 0xab, 0x3b, 0x5a, 0x4d,
	0x87, 0xcd, 0xa5, 0x83, 0x52, 0x1c, 0x8f, 0xa9, 0xec, 0xf6, 0xc5, 0xb5, 0x51, 0xb9, 0xbc, 0x36,
	0x2a, 0x57, 0xd7, 0x46, 0xe5, 0x7b, 0x6a, 0x68, 0x17, 0xa9, 0xa1, 0x5d, 0xa6, 0x86, 0x76, 0x95,
	0x1a, 0xda, 0xef, 0xd4, 0xd0, 0x7e, 0xfe, 0x31, 0x2a, 0xa7, 0xd5, 0xc1, 0xce, 0xbf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x8c, 0xff, 0x55, 0xc4, 0x96, 0x05, 0x00, 0x00,
}
