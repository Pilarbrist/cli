// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakePushVersionActor struct {
	CloudControllerAPIVersionStub        func() string
	cloudControllerAPIVersionMutex       sync.RWMutex
	cloudControllerAPIVersionArgsForCall []struct {
	}
	cloudControllerAPIVersionReturns struct {
		result1 string
	}
	cloudControllerAPIVersionReturnsOnCall map[int]struct {
		result1 string
	}
	GetStreamingLogsForApplicationByNameAndSpaceStub        func(string, string, v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error)
	getStreamingLogsForApplicationByNameAndSpaceMutex       sync.RWMutex
	getStreamingLogsForApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v3action.NOAAClient
	}
	getStreamingLogsForApplicationByNameAndSpaceReturns struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}
	PollStartStub        func(string, chan<- v3action.Warnings) error
	pollStartMutex       sync.RWMutex
	pollStartArgsForCall []struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}
	pollStartReturns struct {
		result1 error
	}
	pollStartReturnsOnCall map[int]struct {
		result1 error
	}
	RestartApplicationStub        func(string) (v3action.Warnings, error)
	restartApplicationMutex       sync.RWMutex
	restartApplicationArgsForCall []struct {
		arg1 string
	}
	restartApplicationReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	restartApplicationReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePushVersionActor) CloudControllerAPIVersion() string {
	fake.cloudControllerAPIVersionMutex.Lock()
	ret, specificReturn := fake.cloudControllerAPIVersionReturnsOnCall[len(fake.cloudControllerAPIVersionArgsForCall)]
	fake.cloudControllerAPIVersionArgsForCall = append(fake.cloudControllerAPIVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("CloudControllerAPIVersion", []interface{}{})
	fake.cloudControllerAPIVersionMutex.Unlock()
	if fake.CloudControllerAPIVersionStub != nil {
		return fake.CloudControllerAPIVersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.cloudControllerAPIVersionReturns
	return fakeReturns.result1
}

func (fake *FakePushVersionActor) CloudControllerAPIVersionCallCount() int {
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	return len(fake.cloudControllerAPIVersionArgsForCall)
}

func (fake *FakePushVersionActor) CloudControllerAPIVersionCalls(stub func() string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = stub
}

func (fake *FakePushVersionActor) CloudControllerAPIVersionReturns(result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	fake.cloudControllerAPIVersionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePushVersionActor) CloudControllerAPIVersionReturnsOnCall(i int, result1 string) {
	fake.cloudControllerAPIVersionMutex.Lock()
	defer fake.cloudControllerAPIVersionMutex.Unlock()
	fake.CloudControllerAPIVersionStub = nil
	if fake.cloudControllerAPIVersionReturnsOnCall == nil {
		fake.cloudControllerAPIVersionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cloudControllerAPIVersionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpace(arg1 string, arg2 string, arg3 v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)]
	fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall = append(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v3action.NOAAClient
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetStreamingLogsForApplicationByNameAndSpace", []interface{}{arg1, arg2, arg3})
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetStreamingLogsForApplicationByNameAndSpaceStub != nil {
		return fake.GetStreamingLogsForApplicationByNameAndSpaceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getStreamingLogsForApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceCallCount() int {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceCalls(stub func(string, string, v3action.NOAAClient) (<-chan *v3action.LogMessage, <-chan error, v3action.Warnings, error)) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = stub
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceArgsForCall(i int) (string, string, v3action.NOAAClient) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getStreamingLogsForApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturns(result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	fake.getStreamingLogsForApplicationByNameAndSpaceReturns = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakePushVersionActor) GetStreamingLogsForApplicationByNameAndSpaceReturnsOnCall(i int, result1 <-chan *v3action.LogMessage, result2 <-chan error, result3 v3action.Warnings, result4 error) {
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Lock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.Unlock()
	fake.GetStreamingLogsForApplicationByNameAndSpaceStub = nil
	if fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 <-chan *v3action.LogMessage
			result2 <-chan error
			result3 v3action.Warnings
			result4 error
		})
	}
	fake.getStreamingLogsForApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 <-chan *v3action.LogMessage
		result2 <-chan error
		result3 v3action.Warnings
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakePushVersionActor) PollStart(arg1 string, arg2 chan<- v3action.Warnings) error {
	fake.pollStartMutex.Lock()
	ret, specificReturn := fake.pollStartReturnsOnCall[len(fake.pollStartArgsForCall)]
	fake.pollStartArgsForCall = append(fake.pollStartArgsForCall, struct {
		arg1 string
		arg2 chan<- v3action.Warnings
	}{arg1, arg2})
	fake.recordInvocation("PollStart", []interface{}{arg1, arg2})
	fake.pollStartMutex.Unlock()
	if fake.PollStartStub != nil {
		return fake.PollStartStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pollStartReturns
	return fakeReturns.result1
}

func (fake *FakePushVersionActor) PollStartCallCount() int {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	return len(fake.pollStartArgsForCall)
}

func (fake *FakePushVersionActor) PollStartCalls(stub func(string, chan<- v3action.Warnings) error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = stub
}

func (fake *FakePushVersionActor) PollStartArgsForCall(i int) (string, chan<- v3action.Warnings) {
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	argsForCall := fake.pollStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePushVersionActor) PollStartReturns(result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	fake.pollStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePushVersionActor) PollStartReturnsOnCall(i int, result1 error) {
	fake.pollStartMutex.Lock()
	defer fake.pollStartMutex.Unlock()
	fake.PollStartStub = nil
	if fake.pollStartReturnsOnCall == nil {
		fake.pollStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pollStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePushVersionActor) RestartApplication(arg1 string) (v3action.Warnings, error) {
	fake.restartApplicationMutex.Lock()
	ret, specificReturn := fake.restartApplicationReturnsOnCall[len(fake.restartApplicationArgsForCall)]
	fake.restartApplicationArgsForCall = append(fake.restartApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RestartApplication", []interface{}{arg1})
	fake.restartApplicationMutex.Unlock()
	if fake.RestartApplicationStub != nil {
		return fake.RestartApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.restartApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePushVersionActor) RestartApplicationCallCount() int {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	return len(fake.restartApplicationArgsForCall)
}

func (fake *FakePushVersionActor) RestartApplicationCalls(stub func(string) (v3action.Warnings, error)) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = stub
}

func (fake *FakePushVersionActor) RestartApplicationArgsForCall(i int) string {
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	argsForCall := fake.restartApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePushVersionActor) RestartApplicationReturns(result1 v3action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	fake.restartApplicationReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakePushVersionActor) RestartApplicationReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.restartApplicationMutex.Lock()
	defer fake.restartApplicationMutex.Unlock()
	fake.RestartApplicationStub = nil
	if fake.restartApplicationReturnsOnCall == nil {
		fake.restartApplicationReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.restartApplicationReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakePushVersionActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cloudControllerAPIVersionMutex.RLock()
	defer fake.cloudControllerAPIVersionMutex.RUnlock()
	fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RLock()
	defer fake.getStreamingLogsForApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollStartMutex.RLock()
	defer fake.pollStartMutex.RUnlock()
	fake.restartApplicationMutex.RLock()
	defer fake.restartApplicationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePushVersionActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.PushVersionActor = new(FakePushVersionActor)
