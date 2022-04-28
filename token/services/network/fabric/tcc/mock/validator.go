// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-token-sdk/token"
	"github.com/hyperledger-labs/fabric-token-sdk/token/services/network/fabric/tcc"
)

type Validator struct {
	UnmarshallAndVerifyStub        func(token.Ledger, string, []byte) ([]interface{}, error)
	unmarshallAndVerifyMutex       sync.RWMutex
	unmarshallAndVerifyArgsForCall []struct {
		arg1 token.Ledger
		arg2 string
		arg3 []byte
	}
	unmarshallAndVerifyReturns struct {
		result1 []interface{}
		result2 error
	}
	unmarshallAndVerifyReturnsOnCall map[int]struct {
		result1 []interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Validator) UnmarshallAndVerify(arg1 token.Ledger, arg2 string, arg3 []byte) ([]interface{}, error) {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.unmarshallAndVerifyMutex.Lock()
	ret, specificReturn := fake.unmarshallAndVerifyReturnsOnCall[len(fake.unmarshallAndVerifyArgsForCall)]
	fake.unmarshallAndVerifyArgsForCall = append(fake.unmarshallAndVerifyArgsForCall, struct {
		arg1 token.Ledger
		arg2 string
		arg3 []byte
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("UnmarshallAndVerify", []interface{}{arg1, arg2, arg3Copy})
	fake.unmarshallAndVerifyMutex.Unlock()
	if fake.UnmarshallAndVerifyStub != nil {
		return fake.UnmarshallAndVerifyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.unmarshallAndVerifyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *Validator) UnmarshallAndVerifyCallCount() int {
	fake.unmarshallAndVerifyMutex.RLock()
	defer fake.unmarshallAndVerifyMutex.RUnlock()
	return len(fake.unmarshallAndVerifyArgsForCall)
}

func (fake *Validator) UnmarshallAndVerifyCalls(stub func(token.Ledger, string, []byte) ([]interface{}, error)) {
	fake.unmarshallAndVerifyMutex.Lock()
	defer fake.unmarshallAndVerifyMutex.Unlock()
	fake.UnmarshallAndVerifyStub = stub
}

func (fake *Validator) UnmarshallAndVerifyArgsForCall(i int) (token.Ledger, string, []byte) {
	fake.unmarshallAndVerifyMutex.RLock()
	defer fake.unmarshallAndVerifyMutex.RUnlock()
	argsForCall := fake.unmarshallAndVerifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *Validator) UnmarshallAndVerifyReturns(result1 []interface{}, result2 error) {
	fake.unmarshallAndVerifyMutex.Lock()
	defer fake.unmarshallAndVerifyMutex.Unlock()
	fake.UnmarshallAndVerifyStub = nil
	fake.unmarshallAndVerifyReturns = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *Validator) UnmarshallAndVerifyReturnsOnCall(i int, result1 []interface{}, result2 error) {
	fake.unmarshallAndVerifyMutex.Lock()
	defer fake.unmarshallAndVerifyMutex.Unlock()
	fake.UnmarshallAndVerifyStub = nil
	if fake.unmarshallAndVerifyReturnsOnCall == nil {
		fake.unmarshallAndVerifyReturnsOnCall = make(map[int]struct {
			result1 []interface{}
			result2 error
		})
	}
	fake.unmarshallAndVerifyReturnsOnCall[i] = struct {
		result1 []interface{}
		result2 error
	}{result1, result2}
}

func (fake *Validator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.unmarshallAndVerifyMutex.RLock()
	defer fake.unmarshallAndVerifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Validator) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ tcc.Validator = new(Validator)