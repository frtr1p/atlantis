// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events/runtime/policy (interfaces: SourceResolver)

package mocks

import (
	pegomock "github.com/petergtz/pegomock"
	valid "github.com/runatlantis/atlantis/server/events/yaml/valid"
	"reflect"
	"time"
)

type MockSourceResolver struct {
	fail func(message string, callerSkip ...int)
}

func NewMockSourceResolver(options ...pegomock.Option) *MockSourceResolver {
	mock := &MockSourceResolver{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockSourceResolver) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockSourceResolver) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockSourceResolver) Resolve(policySet valid.PolicySet) (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockSourceResolver().")
	}
	params := []pegomock.Param{policySet}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Resolve", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
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

func (mock *MockSourceResolver) VerifyWasCalledOnce() *VerifierMockSourceResolver {
	return &VerifierMockSourceResolver{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockSourceResolver) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockSourceResolver {
	return &VerifierMockSourceResolver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockSourceResolver) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockSourceResolver {
	return &VerifierMockSourceResolver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockSourceResolver) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockSourceResolver {
	return &VerifierMockSourceResolver{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockSourceResolver struct {
	mock                   *MockSourceResolver
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockSourceResolver) Resolve(policySet valid.PolicySet) *MockSourceResolver_Resolve_OngoingVerification {
	params := []pegomock.Param{policySet}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Resolve", params, verifier.timeout)
	return &MockSourceResolver_Resolve_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockSourceResolver_Resolve_OngoingVerification struct {
	mock              *MockSourceResolver
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockSourceResolver_Resolve_OngoingVerification) GetCapturedArguments() valid.PolicySet {
	policySet := c.GetAllCapturedArguments()
	return policySet[len(policySet)-1]
}

func (c *MockSourceResolver_Resolve_OngoingVerification) GetAllCapturedArguments() (_param0 []valid.PolicySet) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]valid.PolicySet, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(valid.PolicySet)
		}
	}
	return
}