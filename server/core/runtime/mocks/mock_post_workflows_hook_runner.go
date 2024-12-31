// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/core/runtime (interfaces: PostWorkflowHookRunner)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockPostWorkflowHookRunner struct {
	fail func(message string, callerSkip ...int)
}

func NewMockPostWorkflowHookRunner(options ...pegomock.Option) *MockPostWorkflowHookRunner {
	mock := &MockPostWorkflowHookRunner{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockPostWorkflowHookRunner) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockPostWorkflowHookRunner) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockPostWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, shell string, shellArgs string, path string) (string, string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockPostWorkflowHookRunner().")
	}
	_params := []pegomock.Param{ctx, command, shell, shellArgs, path}
	_result := pegomock.GetGenericMockFrom(mock).Invoke("Run", _params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var _ret0 string
	var _ret1 string
	var _ret2 error
	if len(_result) != 0 {
		if _result[0] != nil {
			_ret0 = _result[0].(string)
		}
		if _result[1] != nil {
			_ret1 = _result[1].(string)
		}
		if _result[2] != nil {
			_ret2 = _result[2].(error)
		}
	}
	return _ret0, _ret1, _ret2
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledOnce() *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockPostWorkflowHookRunner) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockPostWorkflowHookRunner {
	return &VerifierMockPostWorkflowHookRunner{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockPostWorkflowHookRunner struct {
	mock                   *MockPostWorkflowHookRunner
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockPostWorkflowHookRunner) Run(ctx models.WorkflowHookCommandContext, command string, shell string, shellArgs string, path string) *MockPostWorkflowHookRunner_Run_OngoingVerification {
	_params := []pegomock.Param{ctx, command, shell, shellArgs, path}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", _params, verifier.timeout)
	return &MockPostWorkflowHookRunner_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockPostWorkflowHookRunner_Run_OngoingVerification struct {
	mock              *MockPostWorkflowHookRunner
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockPostWorkflowHookRunner_Run_OngoingVerification) GetCapturedArguments() (models.WorkflowHookCommandContext, string, string, string, string) {
	ctx, command, shell, shellArgs, path := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1], command[len(command)-1], shell[len(shell)-1], shellArgs[len(shellArgs)-1], path[len(path)-1]
}

func (c *MockPostWorkflowHookRunner_Run_OngoingVerification) GetAllCapturedArguments() (_param0 []models.WorkflowHookCommandContext, _param1 []string, _param2 []string, _param3 []string, _param4 []string) {
	_params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(_params) > 0 {
		if len(_params) > 0 {
			_param0 = make([]models.WorkflowHookCommandContext, len(c.methodInvocations))
			for u, param := range _params[0] {
				_param0[u] = param.(models.WorkflowHookCommandContext)
			}
		}
		if len(_params) > 1 {
			_param1 = make([]string, len(c.methodInvocations))
			for u, param := range _params[1] {
				_param1[u] = param.(string)
			}
		}
		if len(_params) > 2 {
			_param2 = make([]string, len(c.methodInvocations))
			for u, param := range _params[2] {
				_param2[u] = param.(string)
			}
		}
		if len(_params) > 3 {
			_param3 = make([]string, len(c.methodInvocations))
			for u, param := range _params[3] {
				_param3[u] = param.(string)
			}
		}
		if len(_params) > 4 {
			_param4 = make([]string, len(c.methodInvocations))
			for u, param := range _params[4] {
				_param4[u] = param.(string)
			}
		}
	}
	return
}
