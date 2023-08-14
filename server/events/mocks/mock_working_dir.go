// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: WorkingDir)

package mocks

import (
	pegomock "github.com/petergtz/pegomock/v4"
	models "github.com/runatlantis/atlantis/server/events/models"
	"reflect"
	"time"
)

type MockWorkingDir struct {
	fail func(message string, callerSkip ...int)
}

func NewMockWorkingDir(options ...pegomock.Option) *MockWorkingDir {
	mock := &MockWorkingDir{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockWorkingDir) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockWorkingDir) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockWorkingDir) Clone(headRepo models.Repo, p models.PullRequest, workspace string, shouldRecheckUpstream bool) (string, bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{headRepo, p, workspace, shouldRecheckUpstream}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Clone", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 bool
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(bool)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockWorkingDir) Delete(r models.Repo, p models.PullRequest) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{r, p}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Delete", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockWorkingDir) DeleteForWorkspace(r models.Repo, p models.PullRequest, workspace string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{r, p, workspace}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeleteForWorkspace", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockWorkingDir) DeletePlan(r models.Repo, p models.PullRequest, workspace string, path string, projectName string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{r, p, workspace, path, projectName}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DeletePlan", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockWorkingDir) GetPullDir(r models.Repo, p models.PullRequest) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{r, p}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetPullDir", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockWorkingDir) GetWorkingDir(r models.Repo, p models.PullRequest, workspace string) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{r, p, workspace}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetWorkingDir", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockWorkingDir) HasDiverged(cloneDir string) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockWorkingDir().")
	}
	params := []pegomock.Param{cloneDir}
	result := pegomock.GetGenericMockFrom(mock).Invoke("HasDiverged", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockWorkingDir) VerifyWasCalledOnce() *VerifierMockWorkingDir {
	return &VerifierMockWorkingDir{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockWorkingDir) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockWorkingDir {
	return &VerifierMockWorkingDir{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockWorkingDir) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockWorkingDir {
	return &VerifierMockWorkingDir{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockWorkingDir) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockWorkingDir {
	return &VerifierMockWorkingDir{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockWorkingDir struct {
	mock                   *MockWorkingDir
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockWorkingDir) Clone(headRepo models.Repo, p models.PullRequest, workspace string, shouldRecheckUpstream bool) *MockWorkingDir_Clone_OngoingVerification {
	params := []pegomock.Param{headRepo, p, workspace, shouldRecheckUpstream}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Clone", params, verifier.timeout)
	return &MockWorkingDir_Clone_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_Clone_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_Clone_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string, bool) {
	headRepo, p, workspace, shouldRecheckUpstream := c.GetAllCapturedArguments()
	return headRepo[len(headRepo)-1], p[len(p)-1], workspace[len(workspace)-1], shouldRecheckUpstream[len(shouldRecheckUpstream)-1]
}

func (c *MockWorkingDir_Clone_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string, _param3 []bool) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]bool, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(bool)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) Delete(r models.Repo, p models.PullRequest) *MockWorkingDir_Delete_OngoingVerification {
	params := []pegomock.Param{r, p}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Delete", params, verifier.timeout)
	return &MockWorkingDir_Delete_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_Delete_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_Delete_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	r, p := c.GetAllCapturedArguments()
	return r[len(r)-1], p[len(p)-1]
}

func (c *MockWorkingDir_Delete_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) DeleteForWorkspace(r models.Repo, p models.PullRequest, workspace string) *MockWorkingDir_DeleteForWorkspace_OngoingVerification {
	params := []pegomock.Param{r, p, workspace}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeleteForWorkspace", params, verifier.timeout)
	return &MockWorkingDir_DeleteForWorkspace_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_DeleteForWorkspace_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_DeleteForWorkspace_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string) {
	r, p, workspace := c.GetAllCapturedArguments()
	return r[len(r)-1], p[len(p)-1], workspace[len(workspace)-1]
}

func (c *MockWorkingDir_DeleteForWorkspace_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) DeletePlan(r models.Repo, p models.PullRequest, workspace string, path string, projectName string) *MockWorkingDir_DeletePlan_OngoingVerification {
	params := []pegomock.Param{r, p, workspace, path, projectName}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DeletePlan", params, verifier.timeout)
	return &MockWorkingDir_DeletePlan_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_DeletePlan_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_DeletePlan_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string, string, string) {
	r, p, workspace, path, projectName := c.GetAllCapturedArguments()
	return r[len(r)-1], p[len(p)-1], workspace[len(workspace)-1], path[len(path)-1], projectName[len(projectName)-1]
}

func (c *MockWorkingDir_DeletePlan_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string, _param3 []string, _param4 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(c.methodInvocations))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(c.methodInvocations))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) GetPullDir(r models.Repo, p models.PullRequest) *MockWorkingDir_GetPullDir_OngoingVerification {
	params := []pegomock.Param{r, p}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetPullDir", params, verifier.timeout)
	return &MockWorkingDir_GetPullDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_GetPullDir_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_GetPullDir_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest) {
	r, p := c.GetAllCapturedArguments()
	return r[len(r)-1], p[len(p)-1]
}

func (c *MockWorkingDir_GetPullDir_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) GetWorkingDir(r models.Repo, p models.PullRequest, workspace string) *MockWorkingDir_GetWorkingDir_OngoingVerification {
	params := []pegomock.Param{r, p, workspace}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetWorkingDir", params, verifier.timeout)
	return &MockWorkingDir_GetWorkingDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_GetWorkingDir_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_GetWorkingDir_OngoingVerification) GetCapturedArguments() (models.Repo, models.PullRequest, string) {
	r, p, workspace := c.GetAllCapturedArguments()
	return r[len(r)-1], p[len(p)-1], workspace[len(workspace)-1]
}

func (c *MockWorkingDir_GetWorkingDir_OngoingVerification) GetAllCapturedArguments() (_param0 []models.Repo, _param1 []models.PullRequest, _param2 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]models.Repo, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(models.Repo)
		}
		_param1 = make([]models.PullRequest, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(models.PullRequest)
		}
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockWorkingDir) HasDiverged(cloneDir string) *MockWorkingDir_HasDiverged_OngoingVerification {
	params := []pegomock.Param{cloneDir}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "HasDiverged", params, verifier.timeout)
	return &MockWorkingDir_HasDiverged_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockWorkingDir_HasDiverged_OngoingVerification struct {
	mock              *MockWorkingDir
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockWorkingDir_HasDiverged_OngoingVerification) GetCapturedArguments() string {
	cloneDir := c.GetAllCapturedArguments()
	return cloneDir[len(cloneDir)-1]
}

func (c *MockWorkingDir_HasDiverged_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}
