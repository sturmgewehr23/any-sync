// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/object/tree/objecttree (interfaces: ObjectTree,Storage)
//
// Generated by this command:
//
//	mockgen -destination mock_objecttree/mock_objecttree.go github.com/anyproto/any-sync/commonspace/object/tree/objecttree ObjectTree,Storage
//
// Package mock_objecttree is a generated GoMock package.
package mock_objecttree

import (
	context "context"
	reflect "reflect"
	time "time"

	list "github.com/anyproto/any-sync/commonspace/object/acl/list"
	objecttree "github.com/anyproto/any-sync/commonspace/object/tree/objecttree"
	treechangeproto "github.com/anyproto/any-sync/commonspace/object/tree/treechangeproto"
	gomock "go.uber.org/mock/gomock"
)

// MockObjectTree is a mock of ObjectTree interface.
type MockObjectTree struct {
	ctrl     *gomock.Controller
	recorder *MockObjectTreeMockRecorder
}

// MockObjectTreeMockRecorder is the mock recorder for MockObjectTree.
type MockObjectTreeMockRecorder struct {
	mock *MockObjectTree
}

// NewMockObjectTree creates a new mock instance.
func NewMockObjectTree(ctrl *gomock.Controller) *MockObjectTree {
	mock := &MockObjectTree{ctrl: ctrl}
	mock.recorder = &MockObjectTreeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockObjectTree) EXPECT() *MockObjectTreeMockRecorder {
	return m.recorder
}

// AclList mocks base method.
func (m *MockObjectTree) AclList() list.AclList {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AclList")
	ret0, _ := ret[0].(list.AclList)
	return ret0
}

// AclList indicates an expected call of AclList.
func (mr *MockObjectTreeMockRecorder) AclList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AclList", reflect.TypeOf((*MockObjectTree)(nil).AclList))
}

// AddContent mocks base method.
func (m *MockObjectTree) AddContent(arg0 context.Context, arg1 objecttree.SignableChangeContent) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddContent", arg0, arg1)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddContent indicates an expected call of AddContent.
func (mr *MockObjectTreeMockRecorder) AddContent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContent", reflect.TypeOf((*MockObjectTree)(nil).AddContent), arg0, arg1)
}

// AddContentWithValidator mocks base method.
func (m *MockObjectTree) AddContentWithValidator(arg0 context.Context, arg1 objecttree.SignableChangeContent, arg2 func(objecttree.StorageChange) error) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddContentWithValidator", arg0, arg1, arg2)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddContentWithValidator indicates an expected call of AddContentWithValidator.
func (mr *MockObjectTreeMockRecorder) AddContentWithValidator(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddContentWithValidator", reflect.TypeOf((*MockObjectTree)(nil).AddContentWithValidator), arg0, arg1, arg2)
}

// AddRawChanges mocks base method.
func (m *MockObjectTree) AddRawChanges(arg0 context.Context, arg1 objecttree.RawChangesPayload) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChanges", arg0, arg1)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRawChanges indicates an expected call of AddRawChanges.
func (mr *MockObjectTreeMockRecorder) AddRawChanges(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChanges", reflect.TypeOf((*MockObjectTree)(nil).AddRawChanges), arg0, arg1)
}

// AddRawChangesWithUpdater mocks base method.
func (m *MockObjectTree) AddRawChangesWithUpdater(arg0 context.Context, arg1 objecttree.RawChangesPayload, arg2 func(objecttree.ObjectTree, objecttree.Mode) error) (objecttree.AddResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawChangesWithUpdater", arg0, arg1, arg2)
	ret0, _ := ret[0].(objecttree.AddResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRawChangesWithUpdater indicates an expected call of AddRawChangesWithUpdater.
func (mr *MockObjectTreeMockRecorder) AddRawChangesWithUpdater(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawChangesWithUpdater", reflect.TypeOf((*MockObjectTree)(nil).AddRawChangesWithUpdater), arg0, arg1, arg2)
}

// ChangeInfo mocks base method.
func (m *MockObjectTree) ChangeInfo() *treechangeproto.TreeChangeInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeInfo")
	ret0, _ := ret[0].(*treechangeproto.TreeChangeInfo)
	return ret0
}

// ChangeInfo indicates an expected call of ChangeInfo.
func (mr *MockObjectTreeMockRecorder) ChangeInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeInfo", reflect.TypeOf((*MockObjectTree)(nil).ChangeInfo))
}

// ChangesAfterCommonSnapshotLoader mocks base method.
func (m *MockObjectTree) ChangesAfterCommonSnapshotLoader(arg0, arg1 []string) (objecttree.LoadIterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangesAfterCommonSnapshotLoader", arg0, arg1)
	ret0, _ := ret[0].(objecttree.LoadIterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChangesAfterCommonSnapshotLoader indicates an expected call of ChangesAfterCommonSnapshotLoader.
func (mr *MockObjectTreeMockRecorder) ChangesAfterCommonSnapshotLoader(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangesAfterCommonSnapshotLoader", reflect.TypeOf((*MockObjectTree)(nil).ChangesAfterCommonSnapshotLoader), arg0, arg1)
}

// Close mocks base method.
func (m *MockObjectTree) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockObjectTreeMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockObjectTree)(nil).Close))
}

// Debug mocks base method.
func (m *MockObjectTree) Debug(arg0 objecttree.DescriptionParser) (objecttree.DebugInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Debug", arg0)
	ret0, _ := ret[0].(objecttree.DebugInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Debug indicates an expected call of Debug.
func (mr *MockObjectTreeMockRecorder) Debug(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockObjectTree)(nil).Debug), arg0)
}

// Delete mocks base method.
func (m *MockObjectTree) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockObjectTreeMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockObjectTree)(nil).Delete))
}

// GetChange mocks base method.
func (m *MockObjectTree) GetChange(arg0 string) (*objecttree.Change, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChange", arg0)
	ret0, _ := ret[0].(*objecttree.Change)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChange indicates an expected call of GetChange.
func (mr *MockObjectTreeMockRecorder) GetChange(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChange", reflect.TypeOf((*MockObjectTree)(nil).GetChange), arg0)
}

// HasChanges mocks base method.
func (m *MockObjectTree) HasChanges(arg0 ...string) bool {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HasChanges", varargs...)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasChanges indicates an expected call of HasChanges.
func (mr *MockObjectTreeMockRecorder) HasChanges(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasChanges", reflect.TypeOf((*MockObjectTree)(nil).HasChanges), arg0...)
}

// Header mocks base method.
func (m *MockObjectTree) Header() *treechangeproto.RawTreeChangeWithId {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	return ret0
}

// Header indicates an expected call of Header.
func (mr *MockObjectTreeMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockObjectTree)(nil).Header))
}

// Heads mocks base method.
func (m *MockObjectTree) Heads() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Heads indicates an expected call of Heads.
func (mr *MockObjectTreeMockRecorder) Heads() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockObjectTree)(nil).Heads))
}

// Id mocks base method.
func (m *MockObjectTree) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockObjectTreeMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockObjectTree)(nil).Id))
}

// IsDerived mocks base method.
func (m *MockObjectTree) IsDerived() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsDerived")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsDerived indicates an expected call of IsDerived.
func (mr *MockObjectTreeMockRecorder) IsDerived() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsDerived", reflect.TypeOf((*MockObjectTree)(nil).IsDerived))
}

// IterateFrom mocks base method.
func (m *MockObjectTree) IterateFrom(arg0 string, arg1 func(*objecttree.Change, []byte) (any, error), arg2 func(*objecttree.Change) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateFrom", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateFrom indicates an expected call of IterateFrom.
func (mr *MockObjectTreeMockRecorder) IterateFrom(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateFrom", reflect.TypeOf((*MockObjectTree)(nil).IterateFrom), arg0, arg1, arg2)
}

// IterateRoot mocks base method.
func (m *MockObjectTree) IterateRoot(arg0 func(*objecttree.Change, []byte) (any, error), arg1 func(*objecttree.Change) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IterateRoot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// IterateRoot indicates an expected call of IterateRoot.
func (mr *MockObjectTreeMockRecorder) IterateRoot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateRoot", reflect.TypeOf((*MockObjectTree)(nil).IterateRoot), arg0, arg1)
}

// Len mocks base method.
func (m *MockObjectTree) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockObjectTreeMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockObjectTree)(nil).Len))
}

// Lock mocks base method.
func (m *MockObjectTree) Lock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock.
func (mr *MockObjectTreeMockRecorder) Lock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockObjectTree)(nil).Lock))
}

// PrepareChange mocks base method.
func (m *MockObjectTree) PrepareChange(arg0 objecttree.SignableChangeContent) (*treechangeproto.RawTreeChangeWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareChange", arg0)
	ret0, _ := ret[0].(*treechangeproto.RawTreeChangeWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrepareChange indicates an expected call of PrepareChange.
func (mr *MockObjectTreeMockRecorder) PrepareChange(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareChange", reflect.TypeOf((*MockObjectTree)(nil).PrepareChange), arg0)
}

// Root mocks base method.
func (m *MockObjectTree) Root() *objecttree.Change {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*objecttree.Change)
	return ret0
}

// Root indicates an expected call of Root.
func (mr *MockObjectTreeMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockObjectTree)(nil).Root))
}

// SetFlusher mocks base method.
func (m *MockObjectTree) SetFlusher(arg0 objecttree.Flusher) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetFlusher", arg0)
}

// SetFlusher indicates an expected call of SetFlusher.
func (mr *MockObjectTreeMockRecorder) SetFlusher(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFlusher", reflect.TypeOf((*MockObjectTree)(nil).SetFlusher), arg0)
}

// SnapshotPath mocks base method.
func (m *MockObjectTree) SnapshotPath() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SnapshotPath")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SnapshotPath indicates an expected call of SnapshotPath.
func (mr *MockObjectTreeMockRecorder) SnapshotPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SnapshotPath", reflect.TypeOf((*MockObjectTree)(nil).SnapshotPath))
}

// Storage mocks base method.
func (m *MockObjectTree) Storage() objecttree.Storage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(objecttree.Storage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockObjectTreeMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockObjectTree)(nil).Storage))
}

// TryClose mocks base method.
func (m *MockObjectTree) TryClose(arg0 time.Duration) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryClose", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TryClose indicates an expected call of TryClose.
func (mr *MockObjectTreeMockRecorder) TryClose(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryClose", reflect.TypeOf((*MockObjectTree)(nil).TryClose), arg0)
}

// TryLock mocks base method.
func (m *MockObjectTree) TryLock() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryLock")
	ret0, _ := ret[0].(bool)
	return ret0
}

// TryLock indicates an expected call of TryLock.
func (mr *MockObjectTreeMockRecorder) TryLock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryLock", reflect.TypeOf((*MockObjectTree)(nil).TryLock))
}

// Unlock mocks base method.
func (m *MockObjectTree) Unlock() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock.
func (mr *MockObjectTreeMockRecorder) Unlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockObjectTree)(nil).Unlock))
}

// UnmarshalledHeader mocks base method.
func (m *MockObjectTree) UnmarshalledHeader() *objecttree.Change {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnmarshalledHeader")
	ret0, _ := ret[0].(*objecttree.Change)
	return ret0
}

// UnmarshalledHeader indicates an expected call of UnmarshalledHeader.
func (mr *MockObjectTreeMockRecorder) UnmarshalledHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalledHeader", reflect.TypeOf((*MockObjectTree)(nil).UnmarshalledHeader))
}

// UnpackChange mocks base method.
func (m *MockObjectTree) UnpackChange(arg0 *treechangeproto.RawTreeChangeWithId) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnpackChange", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnpackChange indicates an expected call of UnpackChange.
func (mr *MockObjectTreeMockRecorder) UnpackChange(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnpackChange", reflect.TypeOf((*MockObjectTree)(nil).UnpackChange), arg0)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddAll mocks base method.
func (m *MockStorage) AddAll(arg0 context.Context, arg1 []objecttree.StorageChange, arg2 []string, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAll", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAll indicates an expected call of AddAll.
func (mr *MockStorageMockRecorder) AddAll(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAll", reflect.TypeOf((*MockStorage)(nil).AddAll), arg0, arg1, arg2, arg3)
}

// AddAllNoError mocks base method.
func (m *MockStorage) AddAllNoError(arg0 context.Context, arg1 []objecttree.StorageChange, arg2 []string, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAllNoError", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAllNoError indicates an expected call of AddAllNoError.
func (mr *MockStorageMockRecorder) AddAllNoError(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllNoError", reflect.TypeOf((*MockStorage)(nil).AddAllNoError), arg0, arg1, arg2, arg3)
}

// Close mocks base method.
func (m *MockStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorage)(nil).Close))
}

// CommonSnapshot mocks base method.
func (m *MockStorage) CommonSnapshot(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommonSnapshot", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommonSnapshot indicates an expected call of CommonSnapshot.
func (mr *MockStorageMockRecorder) CommonSnapshot(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonSnapshot", reflect.TypeOf((*MockStorage)(nil).CommonSnapshot), arg0)
}

// Delete mocks base method.
func (m *MockStorage) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorage)(nil).Delete), arg0)
}

// Get mocks base method.
func (m *MockStorage) Get(arg0 context.Context, arg1 string) (objecttree.StorageChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(objecttree.StorageChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStorageMockRecorder) Get(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStorage)(nil).Get), arg0, arg1)
}

// GetAfterOrder mocks base method.
func (m *MockStorage) GetAfterOrder(arg0 context.Context, arg1 string, arg2 func(context.Context, objecttree.StorageChange) (bool, error)) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAfterOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetAfterOrder indicates an expected call of GetAfterOrder.
func (mr *MockStorageMockRecorder) GetAfterOrder(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAfterOrder", reflect.TypeOf((*MockStorage)(nil).GetAfterOrder), arg0, arg1, arg2)
}

// Has mocks base method.
func (m *MockStorage) Has(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockStorageMockRecorder) Has(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockStorage)(nil).Has), arg0, arg1)
}

// Heads mocks base method.
func (m *MockStorage) Heads(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Heads", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Heads indicates an expected call of Heads.
func (mr *MockStorageMockRecorder) Heads(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Heads", reflect.TypeOf((*MockStorage)(nil).Heads), arg0)
}

// Id mocks base method.
func (m *MockStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockStorage)(nil).Id))
}

// Root mocks base method.
func (m *MockStorage) Root(arg0 context.Context) (objecttree.StorageChange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root", arg0)
	ret0, _ := ret[0].(objecttree.StorageChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockStorageMockRecorder) Root(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockStorage)(nil).Root), arg0)
}
