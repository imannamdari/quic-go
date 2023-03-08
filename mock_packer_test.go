// Code generated by MockGen. DO NOT EDIT.
// Source: packet_packer.go

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	ackhandler "github.com/imannamdari/quic-go/internal/ackhandler"
	protocol "github.com/imannamdari/quic-go/internal/protocol"
	qerr "github.com/imannamdari/quic-go/internal/qerr"
	wire "github.com/imannamdari/quic-go/internal/wire"
)

// MockPacker is a mock of Packer interface.
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker.
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance.
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// HandleTransportParameters mocks base method.
func (m *MockPacker) HandleTransportParameters(arg0 *wire.TransportParameters) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HandleTransportParameters", arg0)
}

// HandleTransportParameters indicates an expected call of HandleTransportParameters.
func (mr *MockPackerMockRecorder) HandleTransportParameters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleTransportParameters", reflect.TypeOf((*MockPacker)(nil).HandleTransportParameters), arg0)
}

// MaybePackProbePacket mocks base method.
func (m *MockPacker) MaybePackProbePacket(arg0 protocol.EncryptionLevel, arg1 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaybePackProbePacket", arg0, arg1)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MaybePackProbePacket indicates an expected call of MaybePackProbePacket.
func (mr *MockPackerMockRecorder) MaybePackProbePacket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybePackProbePacket", reflect.TypeOf((*MockPacker)(nil).MaybePackProbePacket), arg0, arg1)
}

// PackApplicationClose mocks base method.
func (m *MockPacker) PackApplicationClose(arg0 *qerr.ApplicationError, arg1 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackApplicationClose", arg0, arg1)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackApplicationClose indicates an expected call of PackApplicationClose.
func (mr *MockPackerMockRecorder) PackApplicationClose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackApplicationClose", reflect.TypeOf((*MockPacker)(nil).PackApplicationClose), arg0, arg1)
}

// PackCoalescedPacket mocks base method.
func (m *MockPacker) PackCoalescedPacket(onlyAck bool, v protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackCoalescedPacket", onlyAck, v)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackCoalescedPacket indicates an expected call of PackCoalescedPacket.
func (mr *MockPackerMockRecorder) PackCoalescedPacket(onlyAck, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackCoalescedPacket", reflect.TypeOf((*MockPacker)(nil).PackCoalescedPacket), onlyAck, v)
}

// PackConnectionClose mocks base method.
func (m *MockPacker) PackConnectionClose(arg0 *qerr.TransportError, arg1 protocol.VersionNumber) (*coalescedPacket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackConnectionClose", arg0, arg1)
	ret0, _ := ret[0].(*coalescedPacket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PackConnectionClose indicates an expected call of PackConnectionClose.
func (mr *MockPackerMockRecorder) PackConnectionClose(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackConnectionClose", reflect.TypeOf((*MockPacker)(nil).PackConnectionClose), arg0, arg1)
}

// PackMTUProbePacket mocks base method.
func (m *MockPacker) PackMTUProbePacket(ping ackhandler.Frame, size protocol.ByteCount, now time.Time, v protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackMTUProbePacket", ping, size, now, v)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackMTUProbePacket indicates an expected call of PackMTUProbePacket.
func (mr *MockPackerMockRecorder) PackMTUProbePacket(ping, size, now, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackMTUProbePacket", reflect.TypeOf((*MockPacker)(nil).PackMTUProbePacket), ping, size, now, v)
}

// PackPacket mocks base method.
func (m *MockPacker) PackPacket(onlyAck bool, now time.Time, v protocol.VersionNumber) (shortHeaderPacket, *packetBuffer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PackPacket", onlyAck, now, v)
	ret0, _ := ret[0].(shortHeaderPacket)
	ret1, _ := ret[1].(*packetBuffer)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// PackPacket indicates an expected call of PackPacket.
func (mr *MockPackerMockRecorder) PackPacket(onlyAck, now, v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PackPacket", reflect.TypeOf((*MockPacker)(nil).PackPacket), onlyAck, now, v)
}

// SetMaxPacketSize mocks base method.
func (m *MockPacker) SetMaxPacketSize(arg0 protocol.ByteCount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxPacketSize", arg0)
}

// SetMaxPacketSize indicates an expected call of SetMaxPacketSize.
func (mr *MockPackerMockRecorder) SetMaxPacketSize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxPacketSize", reflect.TypeOf((*MockPacker)(nil).SetMaxPacketSize), arg0)
}

// SetToken mocks base method.
func (m *MockPacker) SetToken(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetToken", arg0)
}

// SetToken indicates an expected call of SetToken.
func (mr *MockPackerMockRecorder) SetToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetToken", reflect.TypeOf((*MockPacker)(nil).SetToken), arg0)
}
