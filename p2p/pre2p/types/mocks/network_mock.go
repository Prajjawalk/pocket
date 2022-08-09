// Code generated by MockGen. DO NOT EDIT.
// Source: p2p/pre2p/types/network.go

// Package mock_types is a generated GoMock package.
package mock_types

import (
	"github.com/pokt-network/pocket/p2p/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	crypto "github.com/pokt-network/pocket/shared/crypto"
)

// MockNetwork is a mock of Network interface.
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork.
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance.
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// AddPeerToAddrBook mocks base method.
func (m *MockNetwork) AddPeerToAddrBook(peer *types.NetworkPeer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPeerToAddrBook", peer)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPeerToAddrBook indicates an expected call of AddPeerToAddrBook.
func (mr *MockNetworkMockRecorder) AddPeerToAddrBook(peer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeerToAddrBook", reflect.TypeOf((*MockNetwork)(nil).AddPeerToAddrBook), peer)
}

// GetAddrBook mocks base method.
func (m *MockNetwork) GetAddrBook() types.AddrBook {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddrBook")
	ret0, _ := ret[0].(types.AddrBook)
	return ret0
}

// GetAddrBook indicates an expected call of GetAddrBook.
func (mr *MockNetworkMockRecorder) GetAddrBook() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddrBook", reflect.TypeOf((*MockNetwork)(nil).GetAddrBook))
}

// HandleNetworkData mocks base method.
func (m *MockNetwork) HandleNetworkData(data []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleNetworkData", data)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleNetworkData indicates an expected call of HandleNetworkData.
func (mr *MockNetworkMockRecorder) HandleNetworkData(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleNetworkData", reflect.TypeOf((*MockNetwork)(nil).HandleNetworkData), data)
}

// NetworkBroadcast mocks base method.
func (m *MockNetwork) NetworkBroadcast(data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkBroadcast", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetworkBroadcast indicates an expected call of NetworkBroadcast.
func (mr *MockNetworkMockRecorder) NetworkBroadcast(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkBroadcast", reflect.TypeOf((*MockNetwork)(nil).NetworkBroadcast), data)
}

// NetworkSend mocks base method.
func (m *MockNetwork) NetworkSend(data []byte, address crypto.Address) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkSend", data, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// NetworkSend indicates an expected call of NetworkSend.
func (mr *MockNetworkMockRecorder) NetworkSend(data, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkSend", reflect.TypeOf((*MockNetwork)(nil).NetworkSend), data, address)
}

// RemovePeerToAddrBook mocks base method.
func (m *MockNetwork) RemovePeerToAddrBook(peer *types.NetworkPeer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePeerToAddrBook", peer)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePeerToAddrBook indicates an expected call of RemovePeerToAddrBook.
func (mr *MockNetworkMockRecorder) RemovePeerToAddrBook(peer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePeerToAddrBook", reflect.TypeOf((*MockNetwork)(nil).RemovePeerToAddrBook), peer)
}

// MockTransport is a mock of Transport interface.
type MockTransport struct {
	ctrl     *gomock.Controller
	recorder *MockTransportMockRecorder
}

// MockTransportMockRecorder is the mock recorder for MockTransport.
type MockTransportMockRecorder struct {
	mock *MockTransport
}

// NewMockTransport creates a new mock instance.
func NewMockTransport(ctrl *gomock.Controller) *MockTransport {
	mock := &MockTransport{ctrl: ctrl}
	mock.recorder = &MockTransportMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransport) EXPECT() *MockTransportMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockTransport) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockTransportMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTransport)(nil).Close))
}

// IsListener mocks base method.
func (m *MockTransport) IsListener() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsListener")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsListener indicates an expected call of IsListener.
func (mr *MockTransportMockRecorder) IsListener() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsListener", reflect.TypeOf((*MockTransport)(nil).IsListener))
}

// Read mocks base method.
func (m *MockTransport) Read() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockTransportMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockTransport)(nil).Read))
}

// Write mocks base method.
func (m *MockTransport) Write(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockTransportMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockTransport)(nil).Write), arg0)
}