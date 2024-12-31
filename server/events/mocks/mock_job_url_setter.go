// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: JobURLSetter)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	command "github.com/runatlantis/atlantis/server/events/command"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockJobURLSetter struct {
	fail func(message string, callerSkip ...int)
}

func NewMockJobURLSetter(options ...pegomock.Option) *MockJobURLSetter {
	mock := &MockJobURLSetter{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockJobURLSetter) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockJobURLSetter) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockJobURLSetter) SetJobURLWithStatus(ctx command.ProjectContext, cmdName command.Name, status models.CommitStatus, res *command.ProjectResult) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockJobURLSetter().")
	}
	_params := []pegomock.Param{ctx, cmdName, status, res}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("SetJobURLWithStatus", _params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(error)
		}
	}
	return _ret0
}

func (mock *MockJobURLSetter) VerifyWasCalledOnce() *VerifierMockJobURLSetter {
	return &VerifierMockJobURLSetter{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockJobURLSetter) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockJobURLSetter {
	return &VerifierMockJobURLSetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockJobURLSetter) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockJobURLSetter {
	return &VerifierMockJobURLSetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockJobURLSetter) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockJobURLSetter {
	return &VerifierMockJobURLSetter{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockJobURLSetter struct {
	mock                   *MockJobURLSetter
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockJobURLSetter) SetJobURLWithStatus(ctx command.ProjectContext, cmdName command.Name, status models.CommitStatus, res *command.ProjectResult) *MockJobURLSetter_SetJobURLWithStatus_OngoingVerification {
	_params := []pegomock.Param{ctx, cmdName, status, res}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetJobURLWithStatus", _params, verifier.timeout)
	return &MockJobURLSetter_SetJobURLWithStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockJobURLSetter_SetJobURLWithStatus_OngoingVerification struct {
	mock              *MockJobURLSetter
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockJobURLSetter_SetJobURLWithStatus_OngoingVerification) GetCapturedArguments() (command.ProjectContext, command.Name, models.CommitStatus, *command.ProjectResult) {
	ctx, cmdName, status, res := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], cmdName[len(cmdName)-1], status[len(status)-1], res[len(res)-1]
}

func (c *MockJobURLSetter_SetJobURLWithStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []command.ProjectContext, _param1 []command.Name, _param2 []models.CommitStatus, _param3 []*command.ProjectResult) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]command.ProjectContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(command.ProjectContext)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]command.Name, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(command.Name)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]models.CommitStatus, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(models.CommitStatus)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]*command.ProjectResult, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(*command.ProjectResult)
			}
		}
	}
	return
}
