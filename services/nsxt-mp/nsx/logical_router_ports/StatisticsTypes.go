// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for service: Statistics.
// Includes binding types of a structures and enumerations defined in the service.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package logical_router_ports

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiProtocol_ "github.com/vmware/vsphere-automation-sdk-go/runtime/protocol"
	nsxModel "github.com/vmware/vsphere-automation-sdk-go/services/nsxt-mp/nsx/model"
	"reflect"
)

// Possible value for “source“ of method Statistics#get.
const Statistics_GET_SOURCE_REALTIME = "realtime"

// Possible value for “source“ of method Statistics#get.
const Statistics_GET_SOURCE_CACHED = "cached"

func statisticsGetInputType() vapiBindings_.StructType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["logical_router_port_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("operation-input", fields, reflect.TypeOf(vapiData_.StructValue{}), fieldNameMap, validators)
}

func StatisticsGetOutputType() vapiBindings_.BindingType {
	return vapiBindings_.NewReferenceType(nsxModel.LogicalRouterPortStatisticsBindingType)
}

func statisticsGetRestMetadata() vapiProtocol_.OperationRestMetadata {
	fields := map[string]vapiBindings_.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]vapiBindings_.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	dispatchHeaderParams := map[string]string{}
	bodyFieldsMap := map[string]string{}
	fields["logical_router_port_id"] = vapiBindings_.NewStringType()
	fields["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fields["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["logical_router_port_id"] = "LogicalRouterPortId"
	fieldNameMap["source"] = "Source"
	fieldNameMap["transport_node_id"] = "TransportNodeId"
	paramsTypeMap["logical_router_port_id"] = vapiBindings_.NewStringType()
	paramsTypeMap["source"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["transport_node_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	paramsTypeMap["logicalRouterPortId"] = vapiBindings_.NewStringType()
	pathParams["logical_router_port_id"] = "logicalRouterPortId"
	queryParams["source"] = "source"
	queryParams["transport_node_id"] = "transport_node_id"
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
		"/api/v1/logical-router-ports/{logicalRouterPortId}/statistics",
		"",
		resultHeaders,
		200,
		"",
		errorHeaders,
		map[string]int{"com.vmware.vapi.std.errors.invalid_request": 400, "com.vmware.vapi.std.errors.unauthorized": 403, "com.vmware.vapi.std.errors.service_unavailable": 503, "com.vmware.vapi.std.errors.internal_server_error": 500, "com.vmware.vapi.std.errors.not_found": 404})
}
