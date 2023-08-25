// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ProtonMail/proton-bridge/v3/internal/services/syncservice (interfaces: ApplyStageInput,BuildStageInput,BuildStageOutput,DownloadStageInput,DownloadStageOutput,MetadataStageInput,MetadataStageOutput,StateProvider,Regulator,UpdateApplier,MessageBuilder,APIClient,Reporter,DownloadRateModifier)

// Package syncservice is a generated GoMock package.
package syncservice

import (
	bytes "bytes"
	context "context"
	io "io"
	reflect "reflect"

	proton "github.com/ProtonMail/go-proton-api"
	crypto "github.com/ProtonMail/gopenpgp/v2/crypto"
	gomock "github.com/golang/mock/gomock"
)

// MockApplyStageInput is a mock of ApplyStageInput interface.
type MockApplyStageInput struct {
	ctrl     *gomock.Controller
	recorder *MockApplyStageInputMockRecorder
}

// MockApplyStageInputMockRecorder is the mock recorder for MockApplyStageInput.
type MockApplyStageInputMockRecorder struct {
	mock *MockApplyStageInput
}

// NewMockApplyStageInput creates a new mock instance.
func NewMockApplyStageInput(ctrl *gomock.Controller) *MockApplyStageInput {
	mock := &MockApplyStageInput{ctrl: ctrl}
	mock.recorder = &MockApplyStageInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplyStageInput) EXPECT() *MockApplyStageInputMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockApplyStageInput) Consume(arg0 context.Context) (ApplyRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(ApplyRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockApplyStageInputMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockApplyStageInput)(nil).Consume), arg0)
}

// MockBuildStageInput is a mock of BuildStageInput interface.
type MockBuildStageInput struct {
	ctrl     *gomock.Controller
	recorder *MockBuildStageInputMockRecorder
}

// MockBuildStageInputMockRecorder is the mock recorder for MockBuildStageInput.
type MockBuildStageInputMockRecorder struct {
	mock *MockBuildStageInput
}

// NewMockBuildStageInput creates a new mock instance.
func NewMockBuildStageInput(ctrl *gomock.Controller) *MockBuildStageInput {
	mock := &MockBuildStageInput{ctrl: ctrl}
	mock.recorder = &MockBuildStageInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildStageInput) EXPECT() *MockBuildStageInputMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockBuildStageInput) Consume(arg0 context.Context) (BuildRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(BuildRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockBuildStageInputMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockBuildStageInput)(nil).Consume), arg0)
}

// MockBuildStageOutput is a mock of BuildStageOutput interface.
type MockBuildStageOutput struct {
	ctrl     *gomock.Controller
	recorder *MockBuildStageOutputMockRecorder
}

// MockBuildStageOutputMockRecorder is the mock recorder for MockBuildStageOutput.
type MockBuildStageOutputMockRecorder struct {
	mock *MockBuildStageOutput
}

// NewMockBuildStageOutput creates a new mock instance.
func NewMockBuildStageOutput(ctrl *gomock.Controller) *MockBuildStageOutput {
	mock := &MockBuildStageOutput{ctrl: ctrl}
	mock.recorder = &MockBuildStageOutputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuildStageOutput) EXPECT() *MockBuildStageOutputMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBuildStageOutput) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockBuildStageOutputMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBuildStageOutput)(nil).Close))
}

// Produce mocks base method.
func (m *MockBuildStageOutput) Produce(arg0 context.Context, arg1 ApplyRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Produce", arg0, arg1)
}

// Produce indicates an expected call of Produce.
func (mr *MockBuildStageOutputMockRecorder) Produce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockBuildStageOutput)(nil).Produce), arg0, arg1)
}

// MockDownloadStageInput is a mock of DownloadStageInput interface.
type MockDownloadStageInput struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadStageInputMockRecorder
}

// MockDownloadStageInputMockRecorder is the mock recorder for MockDownloadStageInput.
type MockDownloadStageInputMockRecorder struct {
	mock *MockDownloadStageInput
}

// NewMockDownloadStageInput creates a new mock instance.
func NewMockDownloadStageInput(ctrl *gomock.Controller) *MockDownloadStageInput {
	mock := &MockDownloadStageInput{ctrl: ctrl}
	mock.recorder = &MockDownloadStageInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadStageInput) EXPECT() *MockDownloadStageInputMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockDownloadStageInput) Consume(arg0 context.Context) (DownloadRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(DownloadRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockDownloadStageInputMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockDownloadStageInput)(nil).Consume), arg0)
}

// MockDownloadStageOutput is a mock of DownloadStageOutput interface.
type MockDownloadStageOutput struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadStageOutputMockRecorder
}

// MockDownloadStageOutputMockRecorder is the mock recorder for MockDownloadStageOutput.
type MockDownloadStageOutputMockRecorder struct {
	mock *MockDownloadStageOutput
}

// NewMockDownloadStageOutput creates a new mock instance.
func NewMockDownloadStageOutput(ctrl *gomock.Controller) *MockDownloadStageOutput {
	mock := &MockDownloadStageOutput{ctrl: ctrl}
	mock.recorder = &MockDownloadStageOutputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadStageOutput) EXPECT() *MockDownloadStageOutputMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDownloadStageOutput) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDownloadStageOutputMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDownloadStageOutput)(nil).Close))
}

// Produce mocks base method.
func (m *MockDownloadStageOutput) Produce(arg0 context.Context, arg1 BuildRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Produce", arg0, arg1)
}

// Produce indicates an expected call of Produce.
func (mr *MockDownloadStageOutputMockRecorder) Produce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockDownloadStageOutput)(nil).Produce), arg0, arg1)
}

// MockMetadataStageInput is a mock of MetadataStageInput interface.
type MockMetadataStageInput struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataStageInputMockRecorder
}

// MockMetadataStageInputMockRecorder is the mock recorder for MockMetadataStageInput.
type MockMetadataStageInputMockRecorder struct {
	mock *MockMetadataStageInput
}

// NewMockMetadataStageInput creates a new mock instance.
func NewMockMetadataStageInput(ctrl *gomock.Controller) *MockMetadataStageInput {
	mock := &MockMetadataStageInput{ctrl: ctrl}
	mock.recorder = &MockMetadataStageInputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataStageInput) EXPECT() *MockMetadataStageInputMockRecorder {
	return m.recorder
}

// Consume mocks base method.
func (m *MockMetadataStageInput) Consume(arg0 context.Context) (*Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Consume", arg0)
	ret0, _ := ret[0].(*Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Consume indicates an expected call of Consume.
func (mr *MockMetadataStageInputMockRecorder) Consume(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Consume", reflect.TypeOf((*MockMetadataStageInput)(nil).Consume), arg0)
}

// MockMetadataStageOutput is a mock of MetadataStageOutput interface.
type MockMetadataStageOutput struct {
	ctrl     *gomock.Controller
	recorder *MockMetadataStageOutputMockRecorder
}

// MockMetadataStageOutputMockRecorder is the mock recorder for MockMetadataStageOutput.
type MockMetadataStageOutputMockRecorder struct {
	mock *MockMetadataStageOutput
}

// NewMockMetadataStageOutput creates a new mock instance.
func NewMockMetadataStageOutput(ctrl *gomock.Controller) *MockMetadataStageOutput {
	mock := &MockMetadataStageOutput{ctrl: ctrl}
	mock.recorder = &MockMetadataStageOutputMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetadataStageOutput) EXPECT() *MockMetadataStageOutputMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMetadataStageOutput) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockMetadataStageOutputMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetadataStageOutput)(nil).Close))
}

// Produce mocks base method.
func (m *MockMetadataStageOutput) Produce(arg0 context.Context, arg1 DownloadRequest) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Produce", arg0, arg1)
}

// Produce indicates an expected call of Produce.
func (mr *MockMetadataStageOutputMockRecorder) Produce(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Produce", reflect.TypeOf((*MockMetadataStageOutput)(nil).Produce), arg0, arg1)
}

// MockStateProvider is a mock of StateProvider interface.
type MockStateProvider struct {
	ctrl     *gomock.Controller
	recorder *MockStateProviderMockRecorder
}

// MockStateProviderMockRecorder is the mock recorder for MockStateProvider.
type MockStateProviderMockRecorder struct {
	mock *MockStateProvider
}

// NewMockStateProvider creates a new mock instance.
func NewMockStateProvider(ctrl *gomock.Controller) *MockStateProvider {
	mock := &MockStateProvider{ctrl: ctrl}
	mock.recorder = &MockStateProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateProvider) EXPECT() *MockStateProviderMockRecorder {
	return m.recorder
}

// AddFailedMessageID mocks base method.
func (m *MockStateProvider) AddFailedMessageID(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFailedMessageID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFailedMessageID indicates an expected call of AddFailedMessageID.
func (mr *MockStateProviderMockRecorder) AddFailedMessageID(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFailedMessageID", reflect.TypeOf((*MockStateProvider)(nil).AddFailedMessageID), varargs...)
}

// ClearSyncStatus mocks base method.
func (m *MockStateProvider) ClearSyncStatus(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearSyncStatus", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearSyncStatus indicates an expected call of ClearSyncStatus.
func (mr *MockStateProviderMockRecorder) ClearSyncStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearSyncStatus", reflect.TypeOf((*MockStateProvider)(nil).ClearSyncStatus), arg0)
}

// GetSyncStatus mocks base method.
func (m *MockStateProvider) GetSyncStatus(arg0 context.Context) (Status, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncStatus", arg0)
	ret0, _ := ret[0].(Status)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncStatus indicates an expected call of GetSyncStatus.
func (mr *MockStateProviderMockRecorder) GetSyncStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncStatus", reflect.TypeOf((*MockStateProvider)(nil).GetSyncStatus), arg0)
}

// RemFailedMessageID mocks base method.
func (m *MockStateProvider) RemFailedMessageID(arg0 context.Context, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemFailedMessageID", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemFailedMessageID indicates an expected call of RemFailedMessageID.
func (mr *MockStateProviderMockRecorder) RemFailedMessageID(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemFailedMessageID", reflect.TypeOf((*MockStateProvider)(nil).RemFailedMessageID), varargs...)
}

// SetHasLabels mocks base method.
func (m *MockStateProvider) SetHasLabels(arg0 context.Context, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHasLabels", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHasLabels indicates an expected call of SetHasLabels.
func (mr *MockStateProviderMockRecorder) SetHasLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHasLabels", reflect.TypeOf((*MockStateProvider)(nil).SetHasLabels), arg0, arg1)
}

// SetHasMessages mocks base method.
func (m *MockStateProvider) SetHasMessages(arg0 context.Context, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHasMessages", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHasMessages indicates an expected call of SetHasMessages.
func (mr *MockStateProviderMockRecorder) SetHasMessages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHasMessages", reflect.TypeOf((*MockStateProvider)(nil).SetHasMessages), arg0, arg1)
}

// SetLastMessageID mocks base method.
func (m *MockStateProvider) SetLastMessageID(arg0 context.Context, arg1 string, arg2 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLastMessageID", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLastMessageID indicates an expected call of SetLastMessageID.
func (mr *MockStateProviderMockRecorder) SetLastMessageID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLastMessageID", reflect.TypeOf((*MockStateProvider)(nil).SetLastMessageID), arg0, arg1, arg2)
}

// SetMessageCount mocks base method.
func (m *MockStateProvider) SetMessageCount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMessageCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMessageCount indicates an expected call of SetMessageCount.
func (mr *MockStateProviderMockRecorder) SetMessageCount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessageCount", reflect.TypeOf((*MockStateProvider)(nil).SetMessageCount), arg0, arg1)
}

// MockRegulator is a mock of Regulator interface.
type MockRegulator struct {
	ctrl     *gomock.Controller
	recorder *MockRegulatorMockRecorder
}

// MockRegulatorMockRecorder is the mock recorder for MockRegulator.
type MockRegulatorMockRecorder struct {
	mock *MockRegulator
}

// NewMockRegulator creates a new mock instance.
func NewMockRegulator(ctrl *gomock.Controller) *MockRegulator {
	mock := &MockRegulator{ctrl: ctrl}
	mock.recorder = &MockRegulatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegulator) EXPECT() *MockRegulatorMockRecorder {
	return m.recorder
}

// Sync mocks base method.
func (m *MockRegulator) Sync(arg0 context.Context, arg1 *Job) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Sync", arg0, arg1)
}

// Sync indicates an expected call of Sync.
func (mr *MockRegulatorMockRecorder) Sync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockRegulator)(nil).Sync), arg0, arg1)
}

// MockUpdateApplier is a mock of UpdateApplier interface.
type MockUpdateApplier struct {
	ctrl     *gomock.Controller
	recorder *MockUpdateApplierMockRecorder
}

// MockUpdateApplierMockRecorder is the mock recorder for MockUpdateApplier.
type MockUpdateApplierMockRecorder struct {
	mock *MockUpdateApplier
}

// NewMockUpdateApplier creates a new mock instance.
func NewMockUpdateApplier(ctrl *gomock.Controller) *MockUpdateApplier {
	mock := &MockUpdateApplier{ctrl: ctrl}
	mock.recorder = &MockUpdateApplierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUpdateApplier) EXPECT() *MockUpdateApplierMockRecorder {
	return m.recorder
}

// ApplySyncUpdates mocks base method.
func (m *MockUpdateApplier) ApplySyncUpdates(arg0 context.Context, arg1 []BuildResult) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplySyncUpdates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplySyncUpdates indicates an expected call of ApplySyncUpdates.
func (mr *MockUpdateApplierMockRecorder) ApplySyncUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplySyncUpdates", reflect.TypeOf((*MockUpdateApplier)(nil).ApplySyncUpdates), arg0, arg1)
}

// SyncLabels mocks base method.
func (m *MockUpdateApplier) SyncLabels(arg0 context.Context, arg1 map[string]proton.Label) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncLabels", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncLabels indicates an expected call of SyncLabels.
func (mr *MockUpdateApplierMockRecorder) SyncLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncLabels", reflect.TypeOf((*MockUpdateApplier)(nil).SyncLabels), arg0, arg1)
}

// SyncSystemLabelsOnly mocks base method.
func (m *MockUpdateApplier) SyncSystemLabelsOnly(arg0 context.Context, arg1 map[string]proton.Label) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncSystemLabelsOnly", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncSystemLabelsOnly indicates an expected call of SyncSystemLabelsOnly.
func (mr *MockUpdateApplierMockRecorder) SyncSystemLabelsOnly(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncSystemLabelsOnly", reflect.TypeOf((*MockUpdateApplier)(nil).SyncSystemLabelsOnly), arg0, arg1)
}

// MockMessageBuilder is a mock of MessageBuilder interface.
type MockMessageBuilder struct {
	ctrl     *gomock.Controller
	recorder *MockMessageBuilderMockRecorder
}

// MockMessageBuilderMockRecorder is the mock recorder for MockMessageBuilder.
type MockMessageBuilderMockRecorder struct {
	mock *MockMessageBuilder
}

// NewMockMessageBuilder creates a new mock instance.
func NewMockMessageBuilder(ctrl *gomock.Controller) *MockMessageBuilder {
	mock := &MockMessageBuilder{ctrl: ctrl}
	mock.recorder = &MockMessageBuilderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageBuilder) EXPECT() *MockMessageBuilderMockRecorder {
	return m.recorder
}

// BuildMessage mocks base method.
func (m *MockMessageBuilder) BuildMessage(arg0 map[string]proton.Label, arg1 proton.FullMessage, arg2 *crypto.KeyRing, arg3 *bytes.Buffer) (BuildResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildMessage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(BuildResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BuildMessage indicates an expected call of BuildMessage.
func (mr *MockMessageBuilderMockRecorder) BuildMessage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildMessage", reflect.TypeOf((*MockMessageBuilder)(nil).BuildMessage), arg0, arg1, arg2, arg3)
}

// WithKeys mocks base method.
func (m *MockMessageBuilder) WithKeys(arg0 func(*crypto.KeyRing, map[string]*crypto.KeyRing) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithKeys", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithKeys indicates an expected call of WithKeys.
func (mr *MockMessageBuilderMockRecorder) WithKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithKeys", reflect.TypeOf((*MockMessageBuilder)(nil).WithKeys), arg0)
}

// MockAPIClient is a mock of APIClient interface.
type MockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIClientMockRecorder
}

// MockAPIClientMockRecorder is the mock recorder for MockAPIClient.
type MockAPIClientMockRecorder struct {
	mock *MockAPIClient
}

// NewMockAPIClient creates a new mock instance.
func NewMockAPIClient(ctrl *gomock.Controller) *MockAPIClient {
	mock := &MockAPIClient{ctrl: ctrl}
	mock.recorder = &MockAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIClient) EXPECT() *MockAPIClientMockRecorder {
	return m.recorder
}

// GetAttachment mocks base method.
func (m *MockAPIClient) GetAttachment(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachment", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttachment indicates an expected call of GetAttachment.
func (mr *MockAPIClientMockRecorder) GetAttachment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachment", reflect.TypeOf((*MockAPIClient)(nil).GetAttachment), arg0, arg1)
}

// GetAttachmentInto mocks base method.
func (m *MockAPIClient) GetAttachmentInto(arg0 context.Context, arg1 string, arg2 io.ReaderFrom) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttachmentInto", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAttachmentInto indicates an expected call of GetAttachmentInto.
func (mr *MockAPIClientMockRecorder) GetAttachmentInto(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttachmentInto", reflect.TypeOf((*MockAPIClient)(nil).GetAttachmentInto), arg0, arg1, arg2)
}

// GetFullMessage mocks base method.
func (m *MockAPIClient) GetFullMessage(arg0 context.Context, arg1 string, arg2 proton.Scheduler, arg3 proton.AttachmentAllocator) (proton.FullMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFullMessage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(proton.FullMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFullMessage indicates an expected call of GetFullMessage.
func (mr *MockAPIClientMockRecorder) GetFullMessage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFullMessage", reflect.TypeOf((*MockAPIClient)(nil).GetFullMessage), arg0, arg1, arg2, arg3)
}

// GetGroupedMessageCount mocks base method.
func (m *MockAPIClient) GetGroupedMessageCount(arg0 context.Context) ([]proton.MessageGroupCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupedMessageCount", arg0)
	ret0, _ := ret[0].([]proton.MessageGroupCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupedMessageCount indicates an expected call of GetGroupedMessageCount.
func (mr *MockAPIClientMockRecorder) GetGroupedMessageCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupedMessageCount", reflect.TypeOf((*MockAPIClient)(nil).GetGroupedMessageCount), arg0)
}

// GetLabels mocks base method.
func (m *MockAPIClient) GetLabels(arg0 context.Context, arg1 ...proton.LabelType) ([]proton.Label, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetLabels", varargs...)
	ret0, _ := ret[0].([]proton.Label)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLabels indicates an expected call of GetLabels.
func (mr *MockAPIClientMockRecorder) GetLabels(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLabels", reflect.TypeOf((*MockAPIClient)(nil).GetLabels), varargs...)
}

// GetMessage mocks base method.
func (m *MockAPIClient) GetMessage(arg0 context.Context, arg1 string) (proton.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", arg0, arg1)
	ret0, _ := ret[0].(proton.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockAPIClientMockRecorder) GetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockAPIClient)(nil).GetMessage), arg0, arg1)
}

// GetMessageMetadataPage mocks base method.
func (m *MockAPIClient) GetMessageMetadataPage(arg0 context.Context, arg1, arg2 int, arg3 proton.MessageFilter) ([]proton.MessageMetadata, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageMetadataPage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]proton.MessageMetadata)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageMetadataPage indicates an expected call of GetMessageMetadataPage.
func (mr *MockAPIClientMockRecorder) GetMessageMetadataPage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageMetadataPage", reflect.TypeOf((*MockAPIClient)(nil).GetMessageMetadataPage), arg0, arg1, arg2, arg3)
}

// MockReporter is a mock of Reporter interface.
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter.
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance.
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// InitializeProgressCounter mocks base method.
func (m *MockReporter) InitializeProgressCounter(arg0 context.Context, arg1, arg2 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InitializeProgressCounter", arg0, arg1, arg2)
}

// InitializeProgressCounter indicates an expected call of InitializeProgressCounter.
func (mr *MockReporterMockRecorder) InitializeProgressCounter(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeProgressCounter", reflect.TypeOf((*MockReporter)(nil).InitializeProgressCounter), arg0, arg1, arg2)
}

// OnError mocks base method.
func (m *MockReporter) OnError(arg0 context.Context, arg1 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnError", arg0, arg1)
}

// OnError indicates an expected call of OnError.
func (mr *MockReporterMockRecorder) OnError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnError", reflect.TypeOf((*MockReporter)(nil).OnError), arg0, arg1)
}

// OnFinished mocks base method.
func (m *MockReporter) OnFinished(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnFinished", arg0)
}

// OnFinished indicates an expected call of OnFinished.
func (mr *MockReporterMockRecorder) OnFinished(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFinished", reflect.TypeOf((*MockReporter)(nil).OnFinished), arg0)
}

// OnProgress mocks base method.
func (m *MockReporter) OnProgress(arg0 context.Context, arg1 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnProgress", arg0, arg1)
}

// OnProgress indicates an expected call of OnProgress.
func (mr *MockReporterMockRecorder) OnProgress(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnProgress", reflect.TypeOf((*MockReporter)(nil).OnProgress), arg0, arg1)
}

// OnStart mocks base method.
func (m *MockReporter) OnStart(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnStart", arg0)
}

// OnStart indicates an expected call of OnStart.
func (mr *MockReporterMockRecorder) OnStart(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnStart", reflect.TypeOf((*MockReporter)(nil).OnStart), arg0)
}

// MockDownloadRateModifier is a mock of DownloadRateModifier interface.
type MockDownloadRateModifier struct {
	ctrl     *gomock.Controller
	recorder *MockDownloadRateModifierMockRecorder
}

// MockDownloadRateModifierMockRecorder is the mock recorder for MockDownloadRateModifier.
type MockDownloadRateModifierMockRecorder struct {
	mock *MockDownloadRateModifier
}

// NewMockDownloadRateModifier creates a new mock instance.
func NewMockDownloadRateModifier(ctrl *gomock.Controller) *MockDownloadRateModifier {
	mock := &MockDownloadRateModifier{ctrl: ctrl}
	mock.recorder = &MockDownloadRateModifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDownloadRateModifier) EXPECT() *MockDownloadRateModifierMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockDownloadRateModifier) Apply(arg0 bool, arg1, arg2 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	return ret0
}

// Apply indicates an expected call of Apply.
func (mr *MockDownloadRateModifierMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockDownloadRateModifier)(nil).Apply), arg0, arg1, arg2)
}
