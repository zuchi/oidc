// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/caos/oidc/pkg/op (interfaces: Configuration)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	op "github.com/caos/oidc/pkg/op"
	gomock "github.com/golang/mock/gomock"
	language "golang.org/x/text/language"
)

// MockConfiguration is a mock of Configuration interface.
type MockConfiguration struct {
	ctrl     *gomock.Controller
	recorder *MockConfigurationMockRecorder
}

// MockConfigurationMockRecorder is the mock recorder for MockConfiguration.
type MockConfigurationMockRecorder struct {
	mock *MockConfiguration
}

// NewMockConfiguration creates a new mock instance.
func NewMockConfiguration(ctrl *gomock.Controller) *MockConfiguration {
	mock := &MockConfiguration{ctrl: ctrl}
	mock.recorder = &MockConfigurationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfiguration) EXPECT() *MockConfigurationMockRecorder {
	return m.recorder
}

// AuthMethodPostSupported mocks base method.
func (m *MockConfiguration) AuthMethodPostSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMethodPostSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMethodPostSupported indicates an expected call of AuthMethodPostSupported.
func (mr *MockConfigurationMockRecorder) AuthMethodPostSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMethodPostSupported", reflect.TypeOf((*MockConfiguration)(nil).AuthMethodPostSupported))
}

// AuthMethodPrivateKeyJWTSupported mocks base method.
func (m *MockConfiguration) AuthMethodPrivateKeyJWTSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMethodPrivateKeyJWTSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMethodPrivateKeyJWTSupported indicates an expected call of AuthMethodPrivateKeyJWTSupported.
func (mr *MockConfigurationMockRecorder) AuthMethodPrivateKeyJWTSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMethodPrivateKeyJWTSupported", reflect.TypeOf((*MockConfiguration)(nil).AuthMethodPrivateKeyJWTSupported))
}

// AuthorizationEndpoint mocks base method.
func (m *MockConfiguration) AuthorizationEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthorizationEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// AuthorizationEndpoint indicates an expected call of AuthorizationEndpoint.
func (mr *MockConfigurationMockRecorder) AuthorizationEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthorizationEndpoint", reflect.TypeOf((*MockConfiguration)(nil).AuthorizationEndpoint))
}

// CodeMethodS256Supported mocks base method.
func (m *MockConfiguration) CodeMethodS256Supported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CodeMethodS256Supported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// CodeMethodS256Supported indicates an expected call of CodeMethodS256Supported.
func (mr *MockConfigurationMockRecorder) CodeMethodS256Supported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CodeMethodS256Supported", reflect.TypeOf((*MockConfiguration)(nil).CodeMethodS256Supported))
}

// EndSessionEndpoint mocks base method.
func (m *MockConfiguration) EndSessionEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndSessionEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// EndSessionEndpoint indicates an expected call of EndSessionEndpoint.
func (mr *MockConfigurationMockRecorder) EndSessionEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndSessionEndpoint", reflect.TypeOf((*MockConfiguration)(nil).EndSessionEndpoint))
}

// GrantTypeJWTAuthorizationSupported mocks base method.
func (m *MockConfiguration) GrantTypeJWTAuthorizationSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantTypeJWTAuthorizationSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GrantTypeJWTAuthorizationSupported indicates an expected call of GrantTypeJWTAuthorizationSupported.
func (mr *MockConfigurationMockRecorder) GrantTypeJWTAuthorizationSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantTypeJWTAuthorizationSupported", reflect.TypeOf((*MockConfiguration)(nil).GrantTypeJWTAuthorizationSupported))
}

// GrantTypeRefreshTokenSupported mocks base method.
func (m *MockConfiguration) GrantTypeRefreshTokenSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantTypeRefreshTokenSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GrantTypeRefreshTokenSupported indicates an expected call of GrantTypeRefreshTokenSupported.
func (mr *MockConfigurationMockRecorder) GrantTypeRefreshTokenSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantTypeRefreshTokenSupported", reflect.TypeOf((*MockConfiguration)(nil).GrantTypeRefreshTokenSupported))
}

// GrantTypeTokenExchangeSupported mocks base method.
func (m *MockConfiguration) GrantTypeTokenExchangeSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantTypeTokenExchangeSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// GrantTypeTokenExchangeSupported indicates an expected call of GrantTypeTokenExchangeSupported.
func (mr *MockConfigurationMockRecorder) GrantTypeTokenExchangeSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantTypeTokenExchangeSupported", reflect.TypeOf((*MockConfiguration)(nil).GrantTypeTokenExchangeSupported))
}

// IntrospectionAuthMethodPrivateKeyJWTSupported mocks base method.
func (m *MockConfiguration) IntrospectionAuthMethodPrivateKeyJWTSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectionAuthMethodPrivateKeyJWTSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IntrospectionAuthMethodPrivateKeyJWTSupported indicates an expected call of IntrospectionAuthMethodPrivateKeyJWTSupported.
func (mr *MockConfigurationMockRecorder) IntrospectionAuthMethodPrivateKeyJWTSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectionAuthMethodPrivateKeyJWTSupported", reflect.TypeOf((*MockConfiguration)(nil).IntrospectionAuthMethodPrivateKeyJWTSupported))
}

// IntrospectionEndpoint mocks base method.
func (m *MockConfiguration) IntrospectionEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectionEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// IntrospectionEndpoint indicates an expected call of IntrospectionEndpoint.
func (mr *MockConfigurationMockRecorder) IntrospectionEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectionEndpoint", reflect.TypeOf((*MockConfiguration)(nil).IntrospectionEndpoint))
}

// IntrospectionEndpointSigningAlgorithmsSupported mocks base method.
func (m *MockConfiguration) IntrospectionEndpointSigningAlgorithmsSupported() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntrospectionEndpointSigningAlgorithmsSupported")
	ret0, _ := ret[0].([]string)
	return ret0
}

// IntrospectionEndpointSigningAlgorithmsSupported indicates an expected call of IntrospectionEndpointSigningAlgorithmsSupported.
func (mr *MockConfigurationMockRecorder) IntrospectionEndpointSigningAlgorithmsSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntrospectionEndpointSigningAlgorithmsSupported", reflect.TypeOf((*MockConfiguration)(nil).IntrospectionEndpointSigningAlgorithmsSupported))
}

// Issuer mocks base method.
func (m *MockConfiguration) Issuer() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Issuer")
	ret0, _ := ret[0].(string)
	return ret0
}

// Issuer indicates an expected call of Issuer.
func (mr *MockConfigurationMockRecorder) Issuer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Issuer", reflect.TypeOf((*MockConfiguration)(nil).Issuer))
}

// KeysEndpoint mocks base method.
func (m *MockConfiguration) KeysEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "KeysEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// KeysEndpoint indicates an expected call of KeysEndpoint.
func (mr *MockConfigurationMockRecorder) KeysEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KeysEndpoint", reflect.TypeOf((*MockConfiguration)(nil).KeysEndpoint))
}

// RequestObjectSigningAlgorithmsSupported mocks base method.
func (m *MockConfiguration) RequestObjectSigningAlgorithmsSupported() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestObjectSigningAlgorithmsSupported")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RequestObjectSigningAlgorithmsSupported indicates an expected call of RequestObjectSigningAlgorithmsSupported.
func (mr *MockConfigurationMockRecorder) RequestObjectSigningAlgorithmsSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestObjectSigningAlgorithmsSupported", reflect.TypeOf((*MockConfiguration)(nil).RequestObjectSigningAlgorithmsSupported))
}

// RequestObjectSupported mocks base method.
func (m *MockConfiguration) RequestObjectSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestObjectSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RequestObjectSupported indicates an expected call of RequestObjectSupported.
func (mr *MockConfigurationMockRecorder) RequestObjectSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestObjectSupported", reflect.TypeOf((*MockConfiguration)(nil).RequestObjectSupported))
}

// RevocationAuthMethodPrivateKeyJWTSupported mocks base method.
func (m *MockConfiguration) RevocationAuthMethodPrivateKeyJWTSupported() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevocationAuthMethodPrivateKeyJWTSupported")
	ret0, _ := ret[0].(bool)
	return ret0
}

// RevocationAuthMethodPrivateKeyJWTSupported indicates an expected call of RevocationAuthMethodPrivateKeyJWTSupported.
func (mr *MockConfigurationMockRecorder) RevocationAuthMethodPrivateKeyJWTSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevocationAuthMethodPrivateKeyJWTSupported", reflect.TypeOf((*MockConfiguration)(nil).RevocationAuthMethodPrivateKeyJWTSupported))
}

// RevocationEndpoint mocks base method.
func (m *MockConfiguration) RevocationEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevocationEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// RevocationEndpoint indicates an expected call of RevocationEndpoint.
func (mr *MockConfigurationMockRecorder) RevocationEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevocationEndpoint", reflect.TypeOf((*MockConfiguration)(nil).RevocationEndpoint))
}

// RevocationEndpointSigningAlgorithmsSupported mocks base method.
func (m *MockConfiguration) RevocationEndpointSigningAlgorithmsSupported() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevocationEndpointSigningAlgorithmsSupported")
	ret0, _ := ret[0].([]string)
	return ret0
}

// RevocationEndpointSigningAlgorithmsSupported indicates an expected call of RevocationEndpointSigningAlgorithmsSupported.
func (mr *MockConfigurationMockRecorder) RevocationEndpointSigningAlgorithmsSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevocationEndpointSigningAlgorithmsSupported", reflect.TypeOf((*MockConfiguration)(nil).RevocationEndpointSigningAlgorithmsSupported))
}

// SupportedUILocales mocks base method.
func (m *MockConfiguration) SupportedUILocales() []language.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportedUILocales")
	ret0, _ := ret[0].([]language.Tag)
	return ret0
}

// SupportedUILocales indicates an expected call of SupportedUILocales.
func (mr *MockConfigurationMockRecorder) SupportedUILocales() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedUILocales", reflect.TypeOf((*MockConfiguration)(nil).SupportedUILocales))
}

// TokenEndpoint mocks base method.
func (m *MockConfiguration) TokenEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// TokenEndpoint indicates an expected call of TokenEndpoint.
func (mr *MockConfigurationMockRecorder) TokenEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenEndpoint", reflect.TypeOf((*MockConfiguration)(nil).TokenEndpoint))
}

// TokenEndpointSigningAlgorithmsSupported mocks base method.
func (m *MockConfiguration) TokenEndpointSigningAlgorithmsSupported() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenEndpointSigningAlgorithmsSupported")
	ret0, _ := ret[0].([]string)
	return ret0
}

// TokenEndpointSigningAlgorithmsSupported indicates an expected call of TokenEndpointSigningAlgorithmsSupported.
func (mr *MockConfigurationMockRecorder) TokenEndpointSigningAlgorithmsSupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenEndpointSigningAlgorithmsSupported", reflect.TypeOf((*MockConfiguration)(nil).TokenEndpointSigningAlgorithmsSupported))
}

// UserinfoEndpoint mocks base method.
func (m *MockConfiguration) UserinfoEndpoint() op.Endpoint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserinfoEndpoint")
	ret0, _ := ret[0].(op.Endpoint)
	return ret0
}

// UserinfoEndpoint indicates an expected call of UserinfoEndpoint.
func (mr *MockConfigurationMockRecorder) UserinfoEndpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserinfoEndpoint", reflect.TypeOf((*MockConfiguration)(nil).UserinfoEndpoint))
}
