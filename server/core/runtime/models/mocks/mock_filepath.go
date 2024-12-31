// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime/models (interfaces: FilePath)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/core/runtime/models"
	"reflect"
	"time"
)

type MockFilePath struct {
	fail func(message string, callerSkip ...int)
}

func NewMockFilePath(options ...pegomock.Option) *MockFilePath {
	mock := &MockFilePath{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockFilePath) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockFilePath) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockFilePath) Join(elem ...string) models.FilePath {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockFilePath().")
	}
	_params := []pegomock.Param{}
	for _, param := range elem {
		_params = append(_params, param)
	}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Join", _params, []reflect.Type{reflect.TypeOf((*models.FilePath)(nil)).Elem()})
	var _ret0 models.FilePath
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(models.FilePath)
		}
	}
	return _ret0
}

func (mock *MockFilePath) NotExists() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockFilePath().")
	}
	_params := []pegomock.Param{}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("NotExists", _params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var _ret0 bool
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(bool)
		}
	}
	return _ret0
}

func (mock *MockFilePath) Resolve() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockFilePath().")
	}
	_params := []pegomock.Param{}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Resolve", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var _ret0 string
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
	}
	return _ret0
}

func (mock *MockFilePath) Symlink(newname string) (models.FilePath, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockFilePath().")
	}
	_params := []pegomock.Param{newname}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Symlink", _params, []reflect.Type{reflect.TypeOf((*models.FilePath)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 models.FilePath
	var _ret1 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(models.FilePath)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(error)
		}
	}
	return _ret0, _ret1
}

func (mock *MockFilePath) VerifyWasCalledOnce() *VerifierMockFilePath {
	return &VerifierMockFilePath{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockFilePath) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockFilePath {
	return &VerifierMockFilePath{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockFilePath) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockFilePath {
	return &VerifierMockFilePath{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockFilePath) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockFilePath {
	return &VerifierMockFilePath{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockFilePath struct {
	mock                   *MockFilePath
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockFilePath) Join(elem ...string) *MockFilePath_Join_OngoingVerification {
	_params := []pegomock.Param{}
	for _, param := range elem {
		_params = append(_params, param)
	}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Join", _params, verifier.timeout)
	return &MockFilePath_Join_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockFilePath_Join_OngoingVerification struct {
	mock              *MockFilePath
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockFilePath_Join_OngoingVerification) GetCapturedArguments() []string {
	elem := c.GetAllCapturedArguments()
	return elem[len(elem)-1]
}

func (c *MockFilePath_Join_OngoingVerification) GetAllCapturedArguments() (_param0 [][]string) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		_param0 = make([][]string, len(c.methodInvocations))
		for u := 0; u < len(c.methodInvocations); u++ {
			_param0[u] = make([]string, len(_params)-0)
			for x := 0; x < len(_params); x++ {
				if _params[x][u] != nil {
					_param0[u][x-0] = _params[x][u].(string)
				}
			}
		}
	}
	return
}

func (verifier *VerifierMockFilePath) NotExists() *MockFilePath_NotExists_OngoingVerification {
	_params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "NotExists", _params, verifier.timeout)
	return &MockFilePath_NotExists_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockFilePath_NotExists_OngoingVerification struct {
	mock              *MockFilePath
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockFilePath_NotExists_OngoingVerification) GetCapturedArguments() {
}

func (c *MockFilePath_NotExists_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockFilePath) Resolve() *MockFilePath_Resolve_OngoingVerification {
	_params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Resolve", _params, verifier.timeout)
	return &MockFilePath_Resolve_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockFilePath_Resolve_OngoingVerification struct {
	mock              *MockFilePath
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockFilePath_Resolve_OngoingVerification) GetCapturedArguments() {
}

func (c *MockFilePath_Resolve_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockFilePath) Symlink(newname string) *MockFilePath_Symlink_OngoingVerification {
	_params := []pegomock.Param{newname}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Symlink", _params, verifier.timeout)
	return &MockFilePath_Symlink_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockFilePath_Symlink_OngoingVerification struct {
	mock              *MockFilePath
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockFilePath_Symlink_OngoingVerification) GetCapturedArguments() string {
	newname := c.GetAllCapturedArguments()
	return newname[len(newname)-1]
}

func (c *MockFilePath_Symlink_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]string, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(string)
			}
		}
	}
	return
}
