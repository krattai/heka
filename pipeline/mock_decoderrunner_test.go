// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/mozilla-services/heka/pipeline (interfaces: DecoderRunner)

package pipeline

import (
	sync "sync"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of DecoderRunner interface
type MockDecoderRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockDecoderRunnerRecorder
}

// Recorder for MockDecoderRunner (not exported)
type _MockDecoderRunnerRecorder struct {
	mock *MockDecoderRunner
}

func NewMockDecoderRunner(ctrl *gomock.Controller) *MockDecoderRunner {
	mock := &MockDecoderRunner{ctrl: ctrl}
	mock.recorder = &_MockDecoderRunnerRecorder{mock}
	return mock
}

func (_m *MockDecoderRunner) EXPECT() *_MockDecoderRunnerRecorder {
	return _m.recorder
}

func (_m *MockDecoderRunner) Decoder() Decoder {
	ret := _m.ctrl.Call(_m, "Decoder")
	ret0, _ := ret[0].(Decoder)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) Decoder() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Decoder")
}

func (_m *MockDecoderRunner) InChan() chan *PipelinePack {
	ret := _m.ctrl.Call(_m, "InChan")
	ret0, _ := ret[0].(chan *PipelinePack)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) InChan() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InChan")
}

func (_m *MockDecoderRunner) LogError(_param0 error) {
	_m.ctrl.Call(_m, "LogError", _param0)
}

func (_mr *_MockDecoderRunnerRecorder) LogError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogError", arg0)
}

func (_m *MockDecoderRunner) LogMessage(_param0 string) {
	_m.ctrl.Call(_m, "LogMessage", _param0)
}

func (_mr *_MockDecoderRunnerRecorder) LogMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LogMessage", arg0)
}

func (_m *MockDecoderRunner) Name() string {
	ret := _m.ctrl.Call(_m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) Name() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Name")
}

func (_m *MockDecoderRunner) Plugin() Plugin {
	ret := _m.ctrl.Call(_m, "Plugin")
	ret0, _ := ret[0].(Plugin)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) Plugin() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Plugin")
}

func (_m *MockDecoderRunner) PluginGlobals() *PluginGlobals {
	ret := _m.ctrl.Call(_m, "PluginGlobals")
	ret0, _ := ret[0].(*PluginGlobals)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) PluginGlobals() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PluginGlobals")
}

func (_m *MockDecoderRunner) SetName(_param0 string) {
	_m.ctrl.Call(_m, "SetName", _param0)
}

func (_mr *_MockDecoderRunnerRecorder) SetName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetName", arg0)
}

func (_m *MockDecoderRunner) Start(_param0 PluginHelper, _param1 *sync.WaitGroup) {
	_m.ctrl.Call(_m, "Start", _param0, _param1)
}

func (_mr *_MockDecoderRunnerRecorder) Start(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0, arg1)
}

func (_m *MockDecoderRunner) UUID() string {
	ret := _m.ctrl.Call(_m, "UUID")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockDecoderRunnerRecorder) UUID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UUID")
}
