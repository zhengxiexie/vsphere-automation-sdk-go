// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: State.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package transport_nodes

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for “mmState“ of method State#list.
const State_LIST_MM_STATE_ENTERING = "ENTERING"

// Possible value for “mmState“ of method State#list.
const State_LIST_MM_STATE_ENABLED = "ENABLED"

// Possible value for “mmState“ of method State#list.
const State_LIST_MM_STATE_EXITING = "EXITING"

// Possible value for “mmState“ of method State#list.
const State_LIST_MM_STATE_DISABLED = "DISABLED"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_PENDING = "PENDING"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_IN_PROGRESS = "IN_PROGRESS"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_SUCCESS = "SUCCESS"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_PARTIAL_SUCCESS = "PARTIAL_SUCCESS"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_FAILED = "FAILED"

// Possible value for “status“ of method State#list.
const State_LIST_STATUS_ORPHANED = "ORPHANED"

func stateGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StateGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.TransportNodeStateBindingType)
}

func stateGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["transport_node_id"] = vapiBindings_.NewStringType()
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["transportNodeId"] = vapiBindings_.NewStringType()
	pathParams["transport_node_id"] = "transportNodeId"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/transport-nodes/{transportNodeId}/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func stateListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["mm_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vtep_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mm_state"] = "MmState"
	fieldNameMap["status"] = "Status"
	fieldNameMap["vtep_ip"] = "VtepIp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StateListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.TransportNodeStateListResultBindingType)
}

func stateListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["mm_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["vtep_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["mm_state"] = "MmState"
	fieldNameMap["status"] = "Status"
	fieldNameMap["vtep_ip"] = "VtepIp"
	paramsTypeMap["vtep_ip"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["mm_state"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	queryParams["vtep_ip"] = "vtep_ip"
	queryParams["mm_state"] = "mm_state"
	queryParams["status"] = "status"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]map[string]string{}
	return vapiProtocol_.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		dispatchHeaderParams,
		bodyFieldsMap,
		"",
		"",
		"GET",
		"/api/v1/transport-nodes/state",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
