// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/locking (interfaces: Backend)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v3"
	command "github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockBackend struct {
	fail func(message string, callerSkip ...int)
}

func NewMockBackend(options ...pegomock.Option) *MockBackend {
	mock := &MockBackend{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockBackend) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockBackend) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockBackend) CheckCommandLock(_param0 command.Name) (*command.Lock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CheckCommandLock", params, []reflect.Type{reflect.TypeOf((**command.Lock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *command.Lock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*command.Lock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) DeletePullStatus(_param0 models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeletePullStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockBackend) GetLock(_param0 models.Project, _param1 string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetLock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) GetPullStatus(_param0 models.PullRequest) (*models.PullStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullStatus", params, []reflect.Type{reflect.TypeOf((**models.PullStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.PullStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.PullStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) List() ([]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("List", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) LockCommand(_param0 command.Name, _param1 time.Time) (*command.Lock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("LockCommand", params, []reflect.Type{reflect.TypeOf((**command.Lock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *command.Lock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*command.Lock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) TryLock(_param0 models.ProjectLock) (bool, models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("TryLock", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 models.ProjectLock
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(models.ProjectLock)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockBackend) Unlock(_param0 models.Project, _param1 string) (*models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Unlock", params, []reflect.Type{reflect.TypeOf((**models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) UnlockByPull(_param0 string, _param1 int) ([]models.ProjectLock, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockByPull", params, []reflect.Type{reflect.TypeOf((*[]models.ProjectLock)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []models.ProjectLock
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]models.ProjectLock)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) UnlockCommand(_param0 command.Name) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UnlockCommand", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockBackend) UpdateProjectStatus(_param0 models.PullRequest, _param1 string, _param2 string, _param3 models.ProjectPlanStatus) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateProjectStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockBackend) UpdatePullWithResults(_param0 models.PullRequest, _param1 []command.ProjectResult) (models.PullStatus, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockBackend().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdatePullWithResults", params, []reflect.Type{reflect.TypeOf((*models.PullStatus)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 models.PullStatus
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(models.PullStatus)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockBackend) VerifyWasCalledOnce() *VerifierMockBackend {
	return &VerifierMockBackend{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockBackend) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockBackend {
	return &VerifierMockBackend{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockBackend) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockBackend {
	return &VerifierMockBackend{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockBackend) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockBackend {
	return &VerifierMockBackend{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockBackend struct {
	mock                   *MockBackend
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockBackend) CheckCommandLock(_param0 command.Name) *MockBackend_CheckCommandLock_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CheckCommandLock", params, verifier.timeout)
	return &MockBackend_CheckCommandLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_CheckCommandLock_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_CheckCommandLock_OngoingVerification) GetCapturedArguments() command.Name {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockBackend_CheckCommandLock_OngoingVerification) GetAllCapturedArguments() (_param0 []command.Name) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.Name)
		}
	}
	return
}

func (verifier *VerifierMockBackend) DeletePullStatus(_param0 models.PullRequest) *MockBackend_DeletePullStatus_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeletePullStatus", params, verifier.timeout)
	return &MockBackend_DeletePullStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_DeletePullStatus_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_DeletePullStatus_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockBackend_DeletePullStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockBackend) GetLock(_param0 models.Project, _param1 string) *MockBackend_GetLock_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetLock", params, verifier.timeout)
	return &MockBackend_GetLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_GetLock_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_GetLock_OngoingVerification) GetCapturedArguments() (models.Project, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockBackend_GetLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockBackend) GetPullStatus(_param0 models.PullRequest) *MockBackend_GetPullStatus_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullStatus", params, verifier.timeout)
	return &MockBackend_GetPullStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_GetPullStatus_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_GetPullStatus_OngoingVerification) GetCapturedArguments() models.PullRequest {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockBackend_GetPullStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockBackend) List() *MockBackend_List_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "List", params, verifier.timeout)
	return &MockBackend_List_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_List_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_List_OngoingVerification) GetCapturedArguments() {
}

func (c *MockBackend_List_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockBackend) LockCommand(_param0 command.Name, _param1 time.Time) *MockBackend_LockCommand_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "LockCommand", params, verifier.timeout)
	return &MockBackend_LockCommand_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_LockCommand_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_LockCommand_OngoingVerification) GetCapturedArguments() (command.Name, time.Time) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockBackend_LockCommand_OngoingVerification) GetAllCapturedArguments() (_param0 []command.Name, _param1 []time.Time) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.Name)
		}
		_param1 = make([]time.Time, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(time.Time)
		}
	}
	return
}

func (verifier *VerifierMockBackend) TryLock(_param0 models.ProjectLock) *MockBackend_TryLock_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "TryLock", params, verifier.timeout)
	return &MockBackend_TryLock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_TryLock_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_TryLock_OngoingVerification) GetCapturedArguments() models.ProjectLock {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockBackend_TryLock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.ProjectLock) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.ProjectLock, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.ProjectLock)
		}
	}
	return
}

func (verifier *VerifierMockBackend) Unlock(_param0 models.Project, _param1 string) *MockBackend_Unlock_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Unlock", params, verifier.timeout)
	return &MockBackend_Unlock_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_Unlock_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_Unlock_OngoingVerification) GetCapturedArguments() (models.Project, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockBackend_Unlock_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Project, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Project, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Project)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockBackend) UnlockByPull(_param0 string, _param1 int) *MockBackend_UnlockByPull_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockByPull", params, verifier.timeout)
	return &MockBackend_UnlockByPull_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_UnlockByPull_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_UnlockByPull_OngoingVerification) GetCapturedArguments() (string, int) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockBackend_UnlockByPull_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []int) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]int, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(int)
		}
	}
	return
}

func (verifier *VerifierMockBackend) UnlockCommand(_param0 command.Name) *MockBackend_UnlockCommand_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UnlockCommand", params, verifier.timeout)
	return &MockBackend_UnlockCommand_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_UnlockCommand_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_UnlockCommand_OngoingVerification) GetCapturedArguments() command.Name {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockBackend_UnlockCommand_OngoingVerification) GetAllCapturedArguments() (_param0 []command.Name) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]command.Name, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(command.Name)
		}
	}
	return
}

func (verifier *VerifierMockBackend) UpdateProjectStatus(_param0 models.PullRequest, _param1 string, _param2 string, _param3 models.ProjectPlanStatus) *MockBackend_UpdateProjectStatus_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateProjectStatus", params, verifier.timeout)
	return &MockBackend_UpdateProjectStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_UpdateProjectStatus_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_UpdateProjectStatus_OngoingVerification) GetCapturedArguments() (models.PullRequest, string, string, models.ProjectPlanStatus) {
	_param0, _param1, _param2, _param3 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1]
}

func (c *MockBackend_UpdateProjectStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 []string, _param2 []string, _param3 []models.ProjectPlanStatus) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]models.ProjectPlanStatus, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(models.ProjectPlanStatus)
		}
	}
	return
}

func (verifier *VerifierMockBackend) UpdatePullWithResults(_param0 models.PullRequest, _param1 []command.ProjectResult) *MockBackend_UpdatePullWithResults_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdatePullWithResults", params, verifier.timeout)
	return &MockBackend_UpdatePullWithResults_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockBackend_UpdatePullWithResults_OngoingVerification struct {
	mock              *MockBackend
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockBackend_UpdatePullWithResults_OngoingVerification) GetCapturedArguments() (models.PullRequest, []command.ProjectResult) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockBackend_UpdatePullWithResults_OngoingVerification) GetAllCapturedArguments() (_param0 []models.PullRequest, _param1 [][]command.ProjectResult) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.PullRequest)
		}
		_param1 = make([][]command.ProjectResult, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.([]command.ProjectResult)
		}
	}
	return
}
