// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Interfaces.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package network

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for ``source`` of method Interfaces#get.
const Interfaces_GET_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Interfaces#get.
const Interfaces_GET_SOURCE_CACHED = "cached"

// Possible value for ``source`` of method Interfaces#list.
const Interfaces_LIST_SOURCE_REALTIME = "realtime"

// Possible value for ``source`` of method Interfaces#list.
const Interfaces_LIST_SOURCE_CACHED = "cached"

func interfacesGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["interface_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InterfacesGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeInterfacePropertiesBindingType)
}

func interfacesGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["interface_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["interface_id"] = "InterfaceId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["interface_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	paramsTypeMap["interfaceId"] = vapiBindings_.NewStringType()
	pathParams["interface_id"] = "interfaceId"
	pathParams["node_id"] = "nodeId"
	queryParams["source"] = "source"
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
		"/api/v1/cluster/nodes/{nodeId}/network/interfaces/{interfaceId}",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}

func interfacesListInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func InterfacesListOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.NodeInterfacePropertiesListResultBindingType)
}

func interfacesListRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["node_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["node_id"] = "NodeId"
	fieldNameMap["source"] = "Source"
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["node_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["nodeId"] = vapiBindings_.NewStringType()
	pathParams["node_id"] = "nodeId"
	queryParams["source"] = "source"
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
		"/api/v1/cluster/nodes/{nodeId}/network/interfaces",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
