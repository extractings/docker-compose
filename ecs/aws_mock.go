// Code generated by MockGen. DO NOT EDIT.
// Container: github.com/docker/compose-cli/ecs (interfaces: API)

// Package ecs is a generated GoMock package.
package ecs

import (
	context "context"
	cloudformation "github.com/aws/aws-sdk-go/service/cloudformation"
	ecs "github.com/aws/aws-sdk-go/service/ecs"
	secrets "github.com/docker/compose-cli/api/secrets"
	compose "github.com/docker/compose-cli/pkg/api"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// CheckRequirements mocks base method
func (m *MockAPI) CheckRequirements(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckRequirements", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckRequirements indicates an expected call of CheckRequirements
func (mr *MockAPIMockRecorder) CheckRequirements(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckRequirements", reflect.TypeOf((*MockAPI)(nil).CheckRequirements), arg0, arg1)
}

// CheckVPC mocks base method
func (m *MockAPI) CheckVPC(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckVPC", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckVPC indicates an expected call of CheckVPC
func (mr *MockAPIMockRecorder) CheckVPC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckVPC", reflect.TypeOf((*MockAPI)(nil).CheckVPC), arg0, arg1)
}

// CreateChangeSet mocks base method
func (m *MockAPI) CreateChangeSet(arg0 context.Context, arg1, arg2 string, arg3 []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateChangeSet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateChangeSet indicates an expected call of CreateChangeSet
func (mr *MockAPIMockRecorder) CreateChangeSet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateChangeSet", reflect.TypeOf((*MockAPI)(nil).CreateChangeSet), arg0, arg1, arg2, arg3)
}

// CreateCluster mocks base method
func (m *MockAPI) CreateCluster(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCluster", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCluster indicates an expected call of CreateCluster
func (mr *MockAPIMockRecorder) CreateCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCluster", reflect.TypeOf((*MockAPI)(nil).CreateCluster), arg0, arg1)
}

// CreateFileSystem mocks base method
func (m *MockAPI) CreateFileSystem(arg0 context.Context, arg1 map[string]string, arg2 VolumeCreateOptions) (awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFileSystem", arg0, arg1, arg2)
	ret0, _ := ret[0].(awsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFileSystem indicates an expected call of CreateFileSystem
func (mr *MockAPIMockRecorder) CreateFileSystem(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFileSystem", reflect.TypeOf((*MockAPI)(nil).CreateFileSystem), arg0, arg1, arg2)
}

// CreateSecret mocks base method
func (m *MockAPI) CreateSecret(arg0 context.Context, arg1 secrets.Secret) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecret", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecret indicates an expected call of CreateSecret
func (mr *MockAPIMockRecorder) CreateSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecret", reflect.TypeOf((*MockAPI)(nil).CreateSecret), arg0, arg1)
}

// CreateStack mocks base method
func (m *MockAPI) CreateStack(arg0 context.Context, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStack", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateStack indicates an expected call of CreateStack
func (mr *MockAPIMockRecorder) CreateStack(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStack", reflect.TypeOf((*MockAPI)(nil).CreateStack), arg0, arg1, arg2, arg3)
}

// DeleteAutoscalingGroup mocks base method
func (m *MockAPI) DeleteAutoscalingGroup(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAutoscalingGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAutoscalingGroup indicates an expected call of DeleteAutoscalingGroup
func (mr *MockAPIMockRecorder) DeleteAutoscalingGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAutoscalingGroup", reflect.TypeOf((*MockAPI)(nil).DeleteAutoscalingGroup), arg0, arg1)
}

// DeleteCapacityProvider mocks base method
func (m *MockAPI) DeleteCapacityProvider(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCapacityProvider", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCapacityProvider indicates an expected call of DeleteCapacityProvider
func (mr *MockAPIMockRecorder) DeleteCapacityProvider(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCapacityProvider", reflect.TypeOf((*MockAPI)(nil).DeleteCapacityProvider), arg0, arg1)
}

// DeleteFileSystem mocks base method
func (m *MockAPI) DeleteFileSystem(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFileSystem", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFileSystem indicates an expected call of DeleteFileSystem
func (mr *MockAPIMockRecorder) DeleteFileSystem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFileSystem", reflect.TypeOf((*MockAPI)(nil).DeleteFileSystem), arg0, arg1)
}

// DeleteSecret mocks base method
func (m *MockAPI) DeleteSecret(arg0 context.Context, arg1 string, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret
func (mr *MockAPIMockRecorder) DeleteSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockAPI)(nil).DeleteSecret), arg0, arg1, arg2)
}

// DeleteStack mocks base method
func (m *MockAPI) DeleteStack(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStack indicates an expected call of DeleteStack
func (mr *MockAPIMockRecorder) DeleteStack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStack", reflect.TypeOf((*MockAPI)(nil).DeleteStack), arg0, arg1)
}

// DescribeService mocks base method
func (m *MockAPI) DescribeService(arg0 context.Context, arg1, arg2 string) (compose.ServiceStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeService", arg0, arg1, arg2)
	ret0, _ := ret[0].(compose.ServiceStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeService indicates an expected call of DescribeService
func (mr *MockAPIMockRecorder) DescribeService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeService", reflect.TypeOf((*MockAPI)(nil).DescribeService), arg0, arg1, arg2)
}

// DescribeServiceTasks mocks base method
func (m *MockAPI) DescribeServiceTasks(arg0 context.Context, arg1, arg2, arg3 string) ([]compose.ContainerSummary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeServiceTasks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]compose.ContainerSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeServiceTasks indicates an expected call of DescribeServiceTasks
func (mr *MockAPIMockRecorder) DescribeServiceTasks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeServiceTasks", reflect.TypeOf((*MockAPI)(nil).DescribeServiceTasks), arg0, arg1, arg2, arg3)
}

// DescribeStackEvents mocks base method
func (m *MockAPI) DescribeStackEvents(arg0 context.Context, arg1 string) ([]*cloudformation.StackEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStackEvents", arg0, arg1)
	ret0, _ := ret[0].([]*cloudformation.StackEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStackEvents indicates an expected call of DescribeStackEvents
func (mr *MockAPIMockRecorder) DescribeStackEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStackEvents", reflect.TypeOf((*MockAPI)(nil).DescribeStackEvents), arg0, arg1)
}

// GetDefaultVPC mocks base method
func (m *MockAPI) GetDefaultVPC(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDefaultVPC", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultVPC indicates an expected call of GetDefaultVPC
func (mr *MockAPIMockRecorder) GetDefaultVPC(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultVPC", reflect.TypeOf((*MockAPI)(nil).GetDefaultVPC), arg0)
}

// GetLoadBalancerURL mocks base method
func (m *MockAPI) GetLoadBalancerURL(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLoadBalancerURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLoadBalancerURL indicates an expected call of GetLoadBalancerURL
func (mr *MockAPIMockRecorder) GetLoadBalancerURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLoadBalancerURL", reflect.TypeOf((*MockAPI)(nil).GetLoadBalancerURL), arg0, arg1)
}

// GetLogs mocks base method
func (m *MockAPI) GetLogs(arg0 context.Context, arg1 string, arg2 func(string, string, string), arg3 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetLogs indicates an expected call of GetLogs
func (mr *MockAPIMockRecorder) GetLogs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockAPI)(nil).GetLogs), arg0, arg1, arg2, arg3)
}

// GetParameter mocks base method
func (m *MockAPI) GetParameter(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParameter", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetParameter indicates an expected call of GetParameter
func (mr *MockAPIMockRecorder) GetParameter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParameter", reflect.TypeOf((*MockAPI)(nil).GetParameter), arg0, arg1)
}

// GetPublicIPs mocks base method
func (m *MockAPI) GetPublicIPs(arg0 context.Context, arg1 ...string) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublicIPs", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicIPs indicates an expected call of GetPublicIPs
func (mr *MockAPIMockRecorder) GetPublicIPs(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicIPs", reflect.TypeOf((*MockAPI)(nil).GetPublicIPs), varargs...)
}

// GetRoleArn mocks base method
func (m *MockAPI) GetRoleArn(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleArn", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleArn indicates an expected call of GetRoleArn
func (mr *MockAPIMockRecorder) GetRoleArn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleArn", reflect.TypeOf((*MockAPI)(nil).GetRoleArn), arg0, arg1)
}

// GetServiceTaskDefinition mocks base method
func (m *MockAPI) GetServiceTaskDefinition(arg0 context.Context, arg1 string, arg2 []string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceTaskDefinition", arg0, arg1, arg2)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceTaskDefinition indicates an expected call of GetServiceTaskDefinition
func (mr *MockAPIMockRecorder) GetServiceTaskDefinition(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceTaskDefinition", reflect.TypeOf((*MockAPI)(nil).GetServiceTaskDefinition), arg0, arg1, arg2)
}

// GetServiceTasks mocks base method
func (m *MockAPI) GetServiceTasks(arg0 context.Context, arg1, arg2 string, arg3 bool) ([]*ecs.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceTasks", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*ecs.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetServiceTasks indicates an expected call of GetServiceTasks
func (mr *MockAPIMockRecorder) GetServiceTasks(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceTasks", reflect.TypeOf((*MockAPI)(nil).GetServiceTasks), arg0, arg1, arg2, arg3)
}

// GetStackClusterID mocks base method
func (m *MockAPI) GetStackClusterID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStackClusterID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStackClusterID indicates an expected call of GetStackClusterID
func (mr *MockAPIMockRecorder) GetStackClusterID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStackClusterID", reflect.TypeOf((*MockAPI)(nil).GetStackClusterID), arg0, arg1)
}

// GetStackID mocks base method
func (m *MockAPI) GetStackID(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStackID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStackID indicates an expected call of GetStackID
func (mr *MockAPIMockRecorder) GetStackID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStackID", reflect.TypeOf((*MockAPI)(nil).GetStackID), arg0, arg1)
}

// GetSubNets mocks base method
func (m *MockAPI) GetSubNets(arg0 context.Context, arg1 string) ([]awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubNets", arg0, arg1)
	ret0, _ := ret[0].([]awsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubNets indicates an expected call of GetSubNets
func (mr *MockAPIMockRecorder) GetSubNets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubNets", reflect.TypeOf((*MockAPI)(nil).GetSubNets), arg0, arg1)
}

// GetTaskStoppedReason mocks base method
func (m *MockAPI) GetTaskStoppedReason(arg0 context.Context, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskStoppedReason", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskStoppedReason indicates an expected call of GetTaskStoppedReason
func (mr *MockAPIMockRecorder) GetTaskStoppedReason(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskStoppedReason", reflect.TypeOf((*MockAPI)(nil).GetTaskStoppedReason), arg0, arg1, arg2)
}

// InspectSecret mocks base method
func (m *MockAPI) InspectSecret(arg0 context.Context, arg1 string) (secrets.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InspectSecret", arg0, arg1)
	ret0, _ := ret[0].(secrets.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectSecret indicates an expected call of InspectSecret
func (mr *MockAPIMockRecorder) InspectSecret(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectSecret", reflect.TypeOf((*MockAPI)(nil).InspectSecret), arg0, arg1)
}

// IsPublicSubnet mocks base method
func (m *MockAPI) IsPublicSubnet(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPublicSubnet", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPublicSubnet indicates an expected call of IsPublicSubnet
func (mr *MockAPIMockRecorder) IsPublicSubnet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPublicSubnet", reflect.TypeOf((*MockAPI)(nil).IsPublicSubnet), arg0, arg1)
}

// ListFileSystems mocks base method
func (m *MockAPI) ListFileSystems(arg0 context.Context, arg1 map[string]string) ([]awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFileSystems", arg0, arg1)
	ret0, _ := ret[0].([]awsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFileSystems indicates an expected call of ListFileSystems
func (mr *MockAPIMockRecorder) ListFileSystems(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFileSystems", reflect.TypeOf((*MockAPI)(nil).ListFileSystems), arg0, arg1)
}

// ListSecrets mocks base method
func (m *MockAPI) ListSecrets(arg0 context.Context) ([]secrets.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecrets", arg0)
	ret0, _ := ret[0].([]secrets.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets
func (mr *MockAPIMockRecorder) ListSecrets(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockAPI)(nil).ListSecrets), arg0)
}

// ListStackParameters mocks base method
func (m *MockAPI) ListStackParameters(arg0 context.Context, arg1 string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStackParameters", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStackParameters indicates an expected call of ListStackParameters
func (mr *MockAPIMockRecorder) ListStackParameters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackParameters", reflect.TypeOf((*MockAPI)(nil).ListStackParameters), arg0, arg1)
}

// ListStackResources mocks base method
func (m *MockAPI) ListStackResources(arg0 context.Context, arg1 string) (stackResources, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStackResources", arg0, arg1)
	ret0, _ := ret[0].(stackResources)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStackResources indicates an expected call of ListStackResources
func (mr *MockAPIMockRecorder) ListStackResources(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackResources", reflect.TypeOf((*MockAPI)(nil).ListStackResources), arg0, arg1)
}

// ListStackServices mocks base method
func (m *MockAPI) ListStackServices(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStackServices", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStackServices indicates an expected call of ListStackServices
func (mr *MockAPIMockRecorder) ListStackServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStackServices", reflect.TypeOf((*MockAPI)(nil).ListStackServices), arg0, arg1)
}

// ListStacks mocks base method
func (m *MockAPI) ListStacks(arg0 context.Context) ([]compose.Stack, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStacks", arg0)
	ret0, _ := ret[0].([]compose.Stack)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStacks indicates an expected call of ListStacks
func (mr *MockAPIMockRecorder) ListStacks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStacks", reflect.TypeOf((*MockAPI)(nil).ListStacks), arg0)
}

// ListTasks mocks base method
func (m *MockAPI) ListTasks(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockAPIMockRecorder) ListTasks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockAPI)(nil).ListTasks), arg0, arg1, arg2)
}

// ResolveCluster mocks base method
func (m *MockAPI) ResolveCluster(arg0 context.Context, arg1 string) (awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveCluster", arg0, arg1)
	ret0, _ := ret[0].(awsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveCluster indicates an expected call of ResolveCluster
func (mr *MockAPIMockRecorder) ResolveCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveCluster", reflect.TypeOf((*MockAPI)(nil).ResolveCluster), arg0, arg1)
}

// ResolveFileSystem mocks base method
func (m *MockAPI) ResolveFileSystem(arg0 context.Context, arg1 string) (awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveFileSystem", arg0, arg1)
	ret0, _ := ret[0].(awsResource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResolveFileSystem indicates an expected call of ResolveFileSystem
func (mr *MockAPIMockRecorder) ResolveFileSystem(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveFileSystem", reflect.TypeOf((*MockAPI)(nil).ResolveFileSystem), arg0, arg1)
}

// ResolveLoadBalancer mocks base method
func (m *MockAPI) ResolveLoadBalancer(arg0 context.Context, arg1 string) (awsResource, string, string, []awsResource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResolveLoadBalancer", arg0, arg1)
	ret0, _ := ret[0].(awsResource)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].([]awsResource)
	ret4, _ := ret[4].(error)
	return ret0, ret1, ret2, ret3, ret4
}

// ResolveLoadBalancer indicates an expected call of ResolveLoadBalancer
func (mr *MockAPIMockRecorder) ResolveLoadBalancer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResolveLoadBalancer", reflect.TypeOf((*MockAPI)(nil).ResolveLoadBalancer), arg0, arg1)
}

// SecurityGroupExists mocks base method
func (m *MockAPI) SecurityGroupExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecurityGroupExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupExists indicates an expected call of SecurityGroupExists
func (mr *MockAPIMockRecorder) SecurityGroupExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupExists", reflect.TypeOf((*MockAPI)(nil).SecurityGroupExists), arg0, arg1)
}

// StackExists mocks base method
func (m *MockAPI) StackExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StackExists indicates an expected call of StackExists
func (mr *MockAPIMockRecorder) StackExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackExists", reflect.TypeOf((*MockAPI)(nil).StackExists), arg0, arg1)
}

// UpdateStack mocks base method
func (m *MockAPI) UpdateStack(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStack", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStack indicates an expected call of UpdateStack
func (mr *MockAPIMockRecorder) UpdateStack(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStack", reflect.TypeOf((*MockAPI)(nil).UpdateStack), arg0, arg1)
}

// WaitStackComplete mocks base method
func (m *MockAPI) WaitStackComplete(arg0 context.Context, arg1 string, arg2 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitStackComplete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitStackComplete indicates an expected call of WaitStackComplete
func (mr *MockAPIMockRecorder) WaitStackComplete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitStackComplete", reflect.TypeOf((*MockAPI)(nil).WaitStackComplete), arg0, arg1, arg2)
}

// getURLWithPortMapping mocks base method
func (m *MockAPI) getURLWithPortMapping(arg0 context.Context, arg1 []string) ([]compose.PortPublisher, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getURLWithPortMapping", arg0, arg1)
	ret0, _ := ret[0].([]compose.PortPublisher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getURLWithPortMapping indicates an expected call of getURLWithPortMapping
func (mr *MockAPIMockRecorder) getURLWithPortMapping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getURLWithPortMapping", reflect.TypeOf((*MockAPI)(nil).getURLWithPortMapping), arg0, arg1)
}
