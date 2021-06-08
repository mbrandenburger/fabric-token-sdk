// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-smart-client/platform/view/api"
	"github.com/hyperledger-labs/fabric-token-sdk/token/core/zkatdlog/crypto/issue/nonanonym"
)

type SigningIdentity struct {
	SerializeStub        func() ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct{}
	serializeReturns     struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	VerifyStub        func(message []byte, signature []byte) error
	verifyMutex       sync.RWMutex
	verifyArgsForCall []struct {
		message   []byte
		signature []byte
	}
	verifyReturns struct {
		result1 error
	}
	verifyReturnsOnCall map[int]struct {
		result1 error
	}
	SignStub        func(raw []byte) ([]byte, error)
	signMutex       sync.RWMutex
	signArgsForCall []struct {
		raw []byte
	}
	signReturns struct {
		result1 []byte
		result2 error
	}
	signReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	GetPublicVersionStub        func() api.Identity
	getPublicVersionMutex       sync.RWMutex
	getPublicVersionArgsForCall []struct{}
	getPublicVersionReturns     struct {
		result1 api.Identity
	}
	getPublicVersionReturnsOnCall map[int]struct {
		result1 api.Identity
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SigningIdentity) Serialize() ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct{}{})
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if fake.SerializeStub != nil {
		return fake.SerializeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.serializeReturns.result1, fake.serializeReturns.result2
}

func (fake *SigningIdentity) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *SigningIdentity) SerializeReturns(result1 []byte, result2 error) {
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) Verify(message []byte, signature []byte) error {
	var messageCopy []byte
	if message != nil {
		messageCopy = make([]byte, len(message))
		copy(messageCopy, message)
	}
	var signatureCopy []byte
	if signature != nil {
		signatureCopy = make([]byte, len(signature))
		copy(signatureCopy, signature)
	}
	fake.verifyMutex.Lock()
	ret, specificReturn := fake.verifyReturnsOnCall[len(fake.verifyArgsForCall)]
	fake.verifyArgsForCall = append(fake.verifyArgsForCall, struct {
		message   []byte
		signature []byte
	}{messageCopy, signatureCopy})
	fake.recordInvocation("Verify", []interface{}{messageCopy, signatureCopy})
	fake.verifyMutex.Unlock()
	if fake.VerifyStub != nil {
		return fake.VerifyStub(message, signature)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyReturns.result1
}

func (fake *SigningIdentity) VerifyCallCount() int {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return len(fake.verifyArgsForCall)
}

func (fake *SigningIdentity) VerifyArgsForCall(i int) ([]byte, []byte) {
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	return fake.verifyArgsForCall[i].message, fake.verifyArgsForCall[i].signature
}

func (fake *SigningIdentity) VerifyReturns(result1 error) {
	fake.VerifyStub = nil
	fake.verifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) VerifyReturnsOnCall(i int, result1 error) {
	fake.VerifyStub = nil
	if fake.verifyReturnsOnCall == nil {
		fake.verifyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SigningIdentity) Sign(raw []byte) ([]byte, error) {
	var rawCopy []byte
	if raw != nil {
		rawCopy = make([]byte, len(raw))
		copy(rawCopy, raw)
	}
	fake.signMutex.Lock()
	ret, specificReturn := fake.signReturnsOnCall[len(fake.signArgsForCall)]
	fake.signArgsForCall = append(fake.signArgsForCall, struct {
		raw []byte
	}{rawCopy})
	fake.recordInvocation("Sign", []interface{}{rawCopy})
	fake.signMutex.Unlock()
	if fake.SignStub != nil {
		return fake.SignStub(raw)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.signReturns.result1, fake.signReturns.result2
}

func (fake *SigningIdentity) SignCallCount() int {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return len(fake.signArgsForCall)
}

func (fake *SigningIdentity) SignArgsForCall(i int) []byte {
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	return fake.signArgsForCall[i].raw
}

func (fake *SigningIdentity) SignReturns(result1 []byte, result2 error) {
	fake.SignStub = nil
	fake.signReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) SignReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.SignStub = nil
	if fake.signReturnsOnCall == nil {
		fake.signReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.signReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *SigningIdentity) GetPublicVersion() api.Identity {
	fake.getPublicVersionMutex.Lock()
	ret, specificReturn := fake.getPublicVersionReturnsOnCall[len(fake.getPublicVersionArgsForCall)]
	fake.getPublicVersionArgsForCall = append(fake.getPublicVersionArgsForCall, struct{}{})
	fake.recordInvocation("GetPublicVersion", []interface{}{})
	fake.getPublicVersionMutex.Unlock()
	if fake.GetPublicVersionStub != nil {
		return fake.GetPublicVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.getPublicVersionReturns.result1
}

func (fake *SigningIdentity) GetPublicVersionCallCount() int {
	fake.getPublicVersionMutex.RLock()
	defer fake.getPublicVersionMutex.RUnlock()
	return len(fake.getPublicVersionArgsForCall)
}

func (fake *SigningIdentity) GetPublicVersionReturns(result1 api.Identity) {
	fake.GetPublicVersionStub = nil
	fake.getPublicVersionReturns = struct {
		result1 api.Identity
	}{result1}
}

func (fake *SigningIdentity) GetPublicVersionReturnsOnCall(i int, result1 api.Identity) {
	fake.GetPublicVersionStub = nil
	if fake.getPublicVersionReturnsOnCall == nil {
		fake.getPublicVersionReturnsOnCall = make(map[int]struct {
			result1 api.Identity
		})
	}
	fake.getPublicVersionReturnsOnCall[i] = struct {
		result1 api.Identity
	}{result1}
}

func (fake *SigningIdentity) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	fake.verifyMutex.RLock()
	defer fake.verifyMutex.RUnlock()
	fake.signMutex.RLock()
	defer fake.signMutex.RUnlock()
	fake.getPublicVersionMutex.RLock()
	defer fake.getPublicVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SigningIdentity) recordInvocation(key string, args []interface{}) {
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

var _ nonanonym.SigningIdentity = new(SigningIdentity)
