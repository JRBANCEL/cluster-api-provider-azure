/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: ../agentpools.go
//
// Generated by this command:
//
//	mockgen -destination agentpools_mock.go -package mock_agentpools -source ../agentpools.go AgentPoolScope
//

// Package mock_agentpools is a generated GoMock package.
package mock_agentpools

import (
	reflect "reflect"
	time "time"

	v1api20231001 "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20231001"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockAgentPoolScope is a mock of AgentPoolScope interface.
type MockAgentPoolScope struct {
	ctrl     *gomock.Controller
	recorder *MockAgentPoolScopeMockRecorder
}

// MockAgentPoolScopeMockRecorder is the mock recorder for MockAgentPoolScope.
type MockAgentPoolScopeMockRecorder struct {
	mock *MockAgentPoolScope
}

// NewMockAgentPoolScope creates a new mock instance.
func NewMockAgentPoolScope(ctrl *gomock.Controller) *MockAgentPoolScope {
	mock := &MockAgentPoolScope{ctrl: ctrl}
	mock.recorder = &MockAgentPoolScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAgentPoolScope) EXPECT() *MockAgentPoolScopeMockRecorder {
	return m.recorder
}

// ASOOwner mocks base method.
func (m *MockAgentPoolScope) ASOOwner() client.Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ASOOwner")
	ret0, _ := ret[0].(client.Object)
	return ret0
}

// ASOOwner indicates an expected call of ASOOwner.
func (mr *MockAgentPoolScopeMockRecorder) ASOOwner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ASOOwner", reflect.TypeOf((*MockAgentPoolScope)(nil).ASOOwner))
}

// AgentPoolSpec mocks base method.
func (m *MockAgentPoolScope) AgentPoolSpec() azure.ASOResourceSpecGetter[*v1api20231001.ManagedClustersAgentPool] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AgentPoolSpec")
	ret0, _ := ret[0].(azure.ASOResourceSpecGetter[*v1api20231001.ManagedClustersAgentPool])
	return ret0
}

// AgentPoolSpec indicates an expected call of AgentPoolSpec.
func (mr *MockAgentPoolScopeMockRecorder) AgentPoolSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AgentPoolSpec", reflect.TypeOf((*MockAgentPoolScope)(nil).AgentPoolSpec))
}

// ClusterName mocks base method.
func (m *MockAgentPoolScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockAgentPoolScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockAgentPoolScope)(nil).ClusterName))
}

// DefaultedAzureCallTimeout mocks base method.
func (m *MockAgentPoolScope) DefaultedAzureCallTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureCallTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureCallTimeout indicates an expected call of DefaultedAzureCallTimeout.
func (mr *MockAgentPoolScopeMockRecorder) DefaultedAzureCallTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureCallTimeout", reflect.TypeOf((*MockAgentPoolScope)(nil).DefaultedAzureCallTimeout))
}

// DefaultedAzureServiceReconcileTimeout mocks base method.
func (m *MockAgentPoolScope) DefaultedAzureServiceReconcileTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureServiceReconcileTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureServiceReconcileTimeout indicates an expected call of DefaultedAzureServiceReconcileTimeout.
func (mr *MockAgentPoolScopeMockRecorder) DefaultedAzureServiceReconcileTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureServiceReconcileTimeout", reflect.TypeOf((*MockAgentPoolScope)(nil).DefaultedAzureServiceReconcileTimeout))
}

// DefaultedReconcilerRequeue mocks base method.
func (m *MockAgentPoolScope) DefaultedReconcilerRequeue() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedReconcilerRequeue")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedReconcilerRequeue indicates an expected call of DefaultedReconcilerRequeue.
func (mr *MockAgentPoolScopeMockRecorder) DefaultedReconcilerRequeue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedReconcilerRequeue", reflect.TypeOf((*MockAgentPoolScope)(nil).DefaultedReconcilerRequeue))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// GetClient mocks base method.
func (m *MockAgentPoolScope) GetClient() client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient")
	ret0, _ := ret[0].(client.Client)
	return ret0
}

// GetClient indicates an expected call of GetClient.
func (mr *MockAgentPoolScopeMockRecorder) GetClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockAgentPoolScope)(nil).GetClient))
}

// GetLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// Name mocks base method.
func (m *MockAgentPoolScope) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockAgentPoolScopeMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAgentPoolScope)(nil).Name))
}

// NodeResourceGroup mocks base method.
func (m *MockAgentPoolScope) NodeResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeResourceGroup indicates an expected call of NodeResourceGroup.
func (mr *MockAgentPoolScopeMockRecorder) NodeResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeResourceGroup", reflect.TypeOf((*MockAgentPoolScope)(nil).NodeResourceGroup))
}

// RemoveCAPIMachinePoolAnnotation mocks base method.
func (m *MockAgentPoolScope) RemoveCAPIMachinePoolAnnotation(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveCAPIMachinePoolAnnotation", key)
}

// RemoveCAPIMachinePoolAnnotation indicates an expected call of RemoveCAPIMachinePoolAnnotation.
func (mr *MockAgentPoolScopeMockRecorder) RemoveCAPIMachinePoolAnnotation(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCAPIMachinePoolAnnotation", reflect.TypeOf((*MockAgentPoolScope)(nil).RemoveCAPIMachinePoolAnnotation), key)
}

// SetAgentPoolProviderIDList mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolProviderIDList(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolProviderIDList", arg0)
}

// SetAgentPoolProviderIDList indicates an expected call of SetAgentPoolProviderIDList.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolProviderIDList(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolProviderIDList", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolProviderIDList), arg0)
}

// SetAgentPoolReady mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolReady(arg0 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolReady", arg0)
}

// SetAgentPoolReady indicates an expected call of SetAgentPoolReady.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolReady(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolReady", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolReady), arg0)
}

// SetAgentPoolReplicas mocks base method.
func (m *MockAgentPoolScope) SetAgentPoolReplicas(arg0 int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAgentPoolReplicas", arg0)
}

// SetAgentPoolReplicas indicates an expected call of SetAgentPoolReplicas.
func (mr *MockAgentPoolScopeMockRecorder) SetAgentPoolReplicas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAgentPoolReplicas", reflect.TypeOf((*MockAgentPoolScope)(nil).SetAgentPoolReplicas), arg0)
}

// SetCAPIMachinePoolAnnotation mocks base method.
func (m *MockAgentPoolScope) SetCAPIMachinePoolAnnotation(key, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCAPIMachinePoolAnnotation", key, value)
}

// SetCAPIMachinePoolAnnotation indicates an expected call of SetCAPIMachinePoolAnnotation.
func (mr *MockAgentPoolScopeMockRecorder) SetCAPIMachinePoolAnnotation(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCAPIMachinePoolAnnotation", reflect.TypeOf((*MockAgentPoolScope)(nil).SetCAPIMachinePoolAnnotation), key, value)
}

// SetCAPIMachinePoolReplicas mocks base method.
func (m *MockAgentPoolScope) SetCAPIMachinePoolReplicas(replicas *int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCAPIMachinePoolReplicas", replicas)
}

// SetCAPIMachinePoolReplicas indicates an expected call of SetCAPIMachinePoolReplicas.
func (mr *MockAgentPoolScopeMockRecorder) SetCAPIMachinePoolReplicas(replicas any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCAPIMachinePoolReplicas", reflect.TypeOf((*MockAgentPoolScope)(nil).SetCAPIMachinePoolReplicas), replicas)
}

// SetLongRunningOperationState mocks base method.
func (m *MockAgentPoolScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockAgentPoolScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockAgentPoolScope)(nil).SetLongRunningOperationState), arg0)
}

// SetSubnetName mocks base method.
func (m *MockAgentPoolScope) SetSubnetName() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSubnetName")
}

// SetSubnetName indicates an expected call of SetSubnetName.
func (mr *MockAgentPoolScopeMockRecorder) SetSubnetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSubnetName", reflect.TypeOf((*MockAgentPoolScope)(nil).SetSubnetName))
}

// UpdateDeleteStatus mocks base method.
func (m *MockAgentPoolScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockAgentPoolScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockAgentPoolScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockAgentPoolScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockAgentPoolScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}
